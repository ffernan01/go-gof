package main

func main() {
	f := &ConcreteFactory{}
	c := NewClient(f)
	c.Act()
}
