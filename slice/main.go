package main

import "fmt"

func main() {

	s := make([]string, 3, 10)
	fmt.Println("raw slice =", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("1.slice =", s)
	fmt.Println("1.len:", len(s))
	fmt.Println("1.cap:", cap(s))

	s = append(s, "d", "e", "f")
	s = append(s, "i", "j", "k", "o", "p", "q")

	fmt.Println("2.slice =", s)
	fmt.Println("2.len:", len(s))
	fmt.Println("2.cap:", cap(s))

	l := s[:5]
	fmt.Println("l.slice =", l)
	fmt.Println("l.len:", len(l))
	fmt.Println("l.cap:", cap(l))

	t := s[3:]
	fmt.Println("t.slice =", t)
	fmt.Println("t.len:", len(t))
	fmt.Println("t.cap:", cap(t))

	m := s[1:5]
	fmt.Println("m.slice =", m)
	fmt.Println("m.len:", len(m))
	fmt.Println("m.cap:", cap(m))

	for i, v := range m {
		fmt.Printf("index %d, value %s\n", i, v)
	}

	arr := make([]int, 5, 10)
	arr = append(arr, 1, 2, 3, 4, 5)
	fmt.Printf("%p\n", &arr)

	f1(arr)
	fmt.Println("main slice is", arr)

}

func f1(slice []int) {
	fmt.Printf("%p\n", &slice)
	fmt.Println("f1 raw slice is", slice)
	slice[0] = 10
	fmt.Println("f1 after slice is,", slice)
}
