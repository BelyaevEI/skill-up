package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Lion struct{}

func (lion *Lion) MakeSound() string {
	return "roar"
}

type Squirrel struct{}

func (squirrel *Squirrel) MakeSound() string {
	return "squeak"
}

type Snake struct{}

func (snake *Snake) MakeSound() string {
	return "hiss"
}

func AnimalSounds() {
	animals := []Animal{
		&Lion{},
		&Squirrel{},
		&Snake{},
	}

	for _, animal := range animals {
		fmt.Println(animal.MakeSound())
	}
}

func main() {
	AnimalSounds()
}
