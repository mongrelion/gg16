package main

import (
	"fmt"
)

type Car struct {
	on bool
}

func (c *Car) Start() {
	c.on = true
}

func (c *Car) Stop() {
	c.on = false
}

func main() {
	c := new(Car)
	c.Start()
	fmt.Println("Car is on! Brrrrmmmmmrmmrmrmm")

	c.Stop()
	fmt.Println("Car is off!")
}
