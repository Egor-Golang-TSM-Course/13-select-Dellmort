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
