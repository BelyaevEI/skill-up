# Оглавление:

* [SOLID в Golang](#SOLID-в-Golang)

## SOLID в Golang
SOLID - это сокращенная аббревиатура от 5 принципов начинающихся на каждую букву из SOLID.

1. Single Responsibility Principle(SRP) - принцип единственной ответственности.

Структура (в оригинале класс) должна иметь только одну причину для изменения, другими словами, структура должна иметь только одну ответственность.
На самом верхнем уровне мы декомпозируем систему на пакеты. В соответствии с этим принципом каждый пакет должен заниматься какой-то отдельной задачей.Дальше пакет мы делим на структуры с набором методов. Каждая структура и связанные с ней методы несут отвественность за какую-то более специфическую задачу внутри пакета. Каждый метод структуры выполняет какую-то одну единственную задачу.

2. Open/Closed Principle (принцип открытости/закрытости)(OCP). 

Структуры должны быть открыты для расширения, но закрыты для модификации. Это значит, что поведение структуры может быть расширено без изменения ее кода.
Наглядный пример в коде:
```Golang
type Animal interface {
	MakeSound() string
}

type Lion struct {}
func (lion *Lion) MakeSound() string {
	return "roar"
}

type Squirrel struct {}
func (squirrel *Squirrel) MakeSound() string {
	return "squeak"
}

type Snake struct {}
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
```
3. Liskov substitution principle (принцип подстановки Барбары Лисков)(LSP).
   
Объекты в программе должны быть заменяемыми на экземпляры их подтипов без изменения правильности работы программы. Если у вас есть класс-родитель и класс-потомок, то любой код, который использует родительский класс, должен работать так же хорошо и с объектами дочернего класса.

Рассмотрим пример, который нарушает принцип LSP:

```Golang
package main

import "fmt"

// Bird базовый тип
type Bird struct {}

func (b *Bird) Fly() {
    fmt.Println("Птица летит")
}

// Penguin - подтип Bird, но не может летать
type Penguin struct {
    Bird
}

func main() {
    var bird = &Bird{}
    bird.Fly()

    var penguin = &Penguin{}
    penguin.Fly() // Нарушение LSP, т.к. пингвины не летают
}

```

Penguin наследуется от Bird, но не соответствует поведению, ожидаемому от Bird, что нарушает LSP. В данном случае, так как пингвины не умеют летать, нам следует отделить способность летать от базового класса Bird:

```Golang
package main

import "fmt"

// Bird базовый тип
type Bird struct {}

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
```

