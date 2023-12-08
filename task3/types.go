package task3

type Message struct {
	Uid     int
	Message string
}

func NewMessage(uid int, message string) *Message {
	return &Message{
		Uid:     uid,
		Message: message,
	}
}
