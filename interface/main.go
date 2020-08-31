package main

import "fmt"

type I interface {
	Get() int
	Set(int)
}

type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s *S) Set(age int) {
	s.Age = age
}

func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

func printAll(values []interface{}) {
	for _, val := range values {
		fmt.Println(val)
	}
}

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {

	//s := S{}
	//var i I
	//i = &s
	//f(i)


	//names := []string{"stanley", "david", "oscar"}
	//var interfaceSlice []interface{} = make([]interface{}, len(names))
	//for i, d := range names {
	//	interfaceSlice[i] = d
	//}
	//
	//printAll(interfaceSlice)

	animals := []Animal{Dog{}, &Cat{}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
