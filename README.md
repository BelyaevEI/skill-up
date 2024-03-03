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
```
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
