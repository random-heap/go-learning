package main

import "fmt"

type Member struct {
	id     int
	name   string
	email  string
	gender int
	age    int
}

func (m Member) setName(name string) {
	m.name = name
}

func (m *Member) setName1(name string) {
	m.name = name
}

func change(m1 Member, m2 *Member) {
	m1.age = 18
	m2.age = 18
}

type Animal struct {
	name   string
	color  string
	height float32
	weight float32
	age    int
}

//奔跑
func (a Animal)Run() {
	fmt.Println(a.name + "is running")
}
//吃东西
func (a Animal)Eat() {
	fmt.Println(a.name + "is eating")
}

type Cat struct {
	a Animal
}

func main() {

	var m1 Member
	m1 = Member{1, "yao", "yao@mail", 1, 21}
	fmt.Println("m1 = ", m1)

	m2 := Member{2, "hong", "hong@mail", 2, 20}
	fmt.Println("m2 = ", m2)

	var m3 *Member = &Member{}
	m3.id = 3
	m3.name = "xiao"
	fmt.Println("m3 = ", m3)

	change(m1, &m2)

	fmt.Println("m1 = ", m1)
	fmt.Println("m2 = ", m2)

	//值传递
	m1.setName("yaomaoze")
	fmt.Println("m1 = ", m1)

	//指针传递
	m1.setName1("yaomaoze")
	fmt.Println("m1 = ", m1)

	var c = Cat {
		a: Animal {
			name:   "猫猫",
			color:  "橙色",
			weight: 10,
			height: 30,
			age:    5,
		},
	}
	fmt.Println(c.a.name)
	c.a.Run()

}


