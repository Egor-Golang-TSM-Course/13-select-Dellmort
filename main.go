package main

import (
	"lesson13/task1"
	"time"
)

func main() {
	names := []string{"Vasya", "Alex", "Fedya"}
	seconds := []time.Duration{2, 5, 8}

	for i, name := range names {
		go task1.NewChef(name, seconds[i]).Cooking()
	}

	time.Sleep(10 * time.Second)
}
