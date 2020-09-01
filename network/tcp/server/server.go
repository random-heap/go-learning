package main

import (
	"fmt"
	"net"
)

func main() {

	//指定服务器 通信协议,IP地址,port. 创建一个用户监听的socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("服务器设置监听失败,原因是:%v\n", err)
		return
	}

	defer listener.Close()

	fmt.Println("服务器等待客户端建立连接...")

	//阻塞监听客户端连接请求,成功建立连接,返回用于通信的socket
	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("服务器监听失败,原因是:%v\n", err)
	}

	defer conn.Close()

	fmt.Println("服务器与客户端成功建立连接!!!")

	buf := make([]byte, 1024)
	//读取客户端发送的数据
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("Conn Read()错误,原因是:%v\n", err)
	}

	fmt.Println("服务器读到数据:", string(buf[:n]))

	//处理数据 -- 回写给客户端
	if string(buf[:n]) == "Are you ready?" {
		conn.Write([]byte("I am ready!!!"))
	} else {
		conn.Write([]byte("I don`t know what you say!!!"))
	}
}
