package main

import "fmt"

// Bird базовый тип
type Bird struct{}

func (b *Bird) MakeSound() {
	fmt.Println("Птица издает звук")
}

// FlyingBird интерфейс для летающих птиц
type FlyingBird interface {
	Fly()
}

// Sparrow подтип Bird, который умеет летать
type Sparrow struct {
	Bird
}

func (s *Sparrow) Fly() {
	fmt.Println("Воробей летит")
}

// Penguin подтип Bird, но не реализует интерфейс FlyingBird
type Penguin struct {
	Bird
}

func main() {
	var sparrow FlyingBird = &Sparrow{}
	sparrow.Fly()

	var penguin = &Penguin{}
	penguin.MakeSound() // Penguin может издавать звук, но не летать
}
