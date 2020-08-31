package main

import "fmt"

type Animal struct {
	Name string
	Real bool
}

func (c Animal) Dump() {
	fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

type FlyingAnimal struct {
	Animal
	WingSpan int
}

type Footer interface {
	Foo1()
	Foo2()
	Foo3()
}

type Foo struct {
}

func (f Foo) Foo1() {
	fmt.Println("Foo1() here")
}

func (f Foo) Foo2() {
	fmt.Println("Foo2() here")
}

func (f Foo) Foo3() {
	fmt.Println("Foo3() here")
}

func main() {
	dragon := &FlyingAnimal{
		Animal{"Dragon", false, },
		15,
	}

	fmt.Println(dragon.Name)
	fmt.Println(dragon.Real)
	fmt.Println(dragon.WingSpan)

	var footer = Foo{}
	footer.Foo1()
	footer.Foo2()
	footer.Foo3()
}