package task3

import (
	"fmt"
	"time"
)

type Client struct {
	uid     int
	name    string
	doneCh  chan bool
	Message chan *Message
}

func NewClient(uid int, name string) *Client {
	return &Client{
		uid:     uid,
		name:    name,
		doneCh:  make(chan bool),
		Message: make(chan *Message),
	}
}

func (c *Client) SendMessageToCh(message string, ch chan<- *Message) {
	fmt.Println(c.name, "Отправляю сообщение", message)
	newMessage := NewMessage(c.uid, message)
	ch <- newMessage
}

func (c *Client) Close() {
	c.doneCh <- true
}

func (c *Client) Listen() {
	for {
		select {
		case <-c.doneCh:
			fmt.Println(c.name, "Останавливаю свою работу")
			return
		case message := <-c.Message:
			fmt.Printf("%s Принял сообщение: %v\n", c.name, message.Message)
		}
	}
}

type Server struct {
	clients     []*Client
	messageChan chan *Message
	doneCh      chan bool
}

func NewServer(clients []*Client, ch chan *Message) *Server {
	return &Server{
		clients:     clients,
		doneCh:      make(chan bool),
		messageChan: ch,
	}
}

func (s *Server) Listen() {
	for {
		select {
		case <-s.doneCh:
			fmt.Println("Server:", "Останавливаю свою работу")
			return

		case message := <-s.messageChan:
			fmt.Println("Server:", "Получил сообщение", message.Message)
			for _, client := range s.clients {
				if message.Uid != client.uid {
					client.Message <- message
					fmt.Println("Успешно отправил сообщение клиенту", client.name)
				}
			}
		}
	}
}

func (s *Server) Close() {
	s.doneCh <- true
}

func Start() {
	ch := make(chan *Message)

	cl := NewClient(1, "Alex")
	go cl.Listen()
	go cl.SendMessageToCh("message1", ch)

	cl2 := NewClient(2, "Fedya")
	go cl2.Listen()
	go cl2.SendMessageToCh("message2", ch)

	cl3 := NewClient(3, "Vitya")
	go cl3.Listen()
	go cl3.SendMessageToCh("message3", ch)

	server := NewServer([]*Client{cl, cl2, cl3}, ch)
	go server.Listen()

	time.Sleep(3 * time.Second)
	cl.Close()
	cl2.Close()
	cl3.Close()
	server.Close()
	time.Sleep(1 * time.Second)
}
