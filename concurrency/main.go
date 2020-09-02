package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	dataChan := make(chan int, 100)

	go func() {
		for {
			select {
			case data := <- dataChan:
				fmt.Println("data -> ", data, time.Now())
				time.Sleep(1 * time.Second)
			}
		}
	}()

	//填充数据
	for i := 0; i < 20; i++ {
		dataChan <- i
	}

	fmt.Println("发送完毕")

	for {
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}
