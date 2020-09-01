package main

import (
	"fmt"
	"net"
)

func main() {
	//指定服务器IP+port创建 通信套接字
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("net.Dial err:%v\n", err)
		return
	}
	defer conn.Close()

	//主动写数据给服务器
	_, err = conn.Write([]byte("Are you ready?"))
	if err != nil {
		fmt.Printf("conn.Write err:%v\n", err)
		return
	}

	buf := make([]byte, 1024)
	//接受服务器回发的数据
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("conn.Read err:%v\n", err)
		return
	}

	fmt.Printf("服务器回发的数据为:%v\n", string(buf[:n]))
}
