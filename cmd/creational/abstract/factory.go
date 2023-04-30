package main

type (
	LivingThingFactory interface {
		CreatePerson() LivingThing
		CreateAnimal() LivingThing
	}

	ConcreteFactory struct{}
)

func (f *ConcreteFactory) CreatePerson() LivingThing {
	return &Person{}
}

func (f *ConcreteFactory) CreateAnimal() LivingThing {
	return &Animal{}
}
