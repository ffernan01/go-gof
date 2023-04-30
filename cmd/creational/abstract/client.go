package main

import "fmt"

type Client struct {
	factory LivingThingFactory
}

func NewClient(f LivingThingFactory) *Client {
	return &Client{factory: f}
}

func (c *Client) Act() {
	p := c.factory.CreatePerson()
	a := c.factory.CreateAnimal()

	fmt.Printf("created person: %s\n", p.Identify())
	fmt.Printf("created animal: %s\n", a.Identify())
}
