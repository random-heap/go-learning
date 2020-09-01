package main

import "fmt"

func add(x int , y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

type Greeting func(name string) string

func (g Greeting) say(n string)  {
	fmt.Println(g(n))
}

func english(name string) string {
	return "Hello, " + name
}

func french(name string) string {
	return "Bonjour, " + name
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	fmt.Println(add(1, 2))

	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)

	g := Greeting(english)
	g.say("World")
	g = Greeting(french)
	g.say("World")

	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())


}
