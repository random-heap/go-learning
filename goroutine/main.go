package main

import (
	"fmt"
	"time"
)

func say(s string) {

	fmt.Println(s)
	go func() {
		fmt.Println("============")
	}()

}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func main() {
	go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	a := <- ch
	b := <- ch
	fmt.Println("a = ", a, "b = ", b)

	//主线程退出后其他goruntine都退出
	time.Sleep(2000)
}