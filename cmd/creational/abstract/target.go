package main

type (
	LivingThing interface {
		Identify() string
	}

	Person struct{}

	Animal struct{}
)

func (p *Person) Identify() string {
	return "I'm a person"
}

func (a *Animal) Identify() string {
	return "I'm an animal"
}
