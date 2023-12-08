package task2

import (
	"fmt"
	"math/rand"
	"time"
)

type Restaurant struct {
	Name  string
	Count int
}

func NewRestaurant(name string) *Restaurant {
	return &Restaurant{
		Name: name,
	}
}

func (r *Restaurant) GenerateSales(ch chan<- *Restaurant) {
	rand.NewSource(time.Now().UnixNano())
	r.Count = rand.Intn(500)
	ch <- r
}

func Start() {
	restaurantNames := [3]string{"Eleon", "Mishka", "Hallal"}
	reportCh := make(chan *Restaurant)
	go func() {
		var countSales int
		for restaurant := range reportCh {
			fmt.Printf("Ресторн %s, сделал %d продаж за день\n", restaurant.Name, restaurant.Count)
			countSales += restaurant.Count
		}

		fmt.Println("Всего продаж за день:", countSales)
	}()

	for _, name := range restaurantNames {
		go NewRestaurant(name).GenerateSales(reportCh)
	}

	time.Sleep(5 * time.Second)
	close(reportCh)
	time.Sleep(1 * time.Second)
}
