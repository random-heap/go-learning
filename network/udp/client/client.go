package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("net.Dial()函数执行出错,错误为:%v\n", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("hello, I`m a client in UDP!"))

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("Conn.Read()方法执行出错,错误为:%v\n", err)
		return
	}
	fmt.Printf("服务器发来数据:%s\n", buf[:n])
}
