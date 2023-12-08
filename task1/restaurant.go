package task1

import (
	"fmt"
	"time"
)

type Chef struct {
	name        string
	cookingTime time.Duration
}

func NewChef(name string, cookTime time.Duration) *Chef {
	return &Chef{
		name:        name,
		cookingTime: cookTime,
	}
}

func (c *Chef) Cooking() {
	fmt.Printf("%s начал готовку блюда, она займет %d секунд\n", c.name, c.cookingTime)
	time.Sleep(c.cookingTime * time.Second)
	fmt.Printf("Шеф-повар %s завершил готовить блюдо.\n", c.name)
}

func Start() {
	names := [3]string{"Vasya", "Alex", "Fedya"}
	seconds := [3]time.Duration{2, 3, 1}

	for i, name := range names {
		go NewChef(name, seconds[i]).Cooking()
	}
}
