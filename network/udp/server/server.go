package main

import (
	"fmt"
	"net"
)

func main() {
	//0.本应从步骤1开始,但是在写步骤1的时候发现,步骤1还需要*UDPAddr类型的参数,所以需要先创建一个*DUPAddr
	//组织一个udp地址结构,指定服务器的IP+port
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("net.ResolveUDPAddr()函数执行出错,错误为:%v\n", err)
		return
	}
	fmt.Printf("UDP服务器地址结构创建完成!!!\n")

	//1.创建用户通信的socket
	//由于ListenUDP需要一个*UDPAddr类型的参数,所以我们还需要先创建一个监听地址
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Printf("net.ListenUDP()函数执行出错,错误为:%v\n", err)
		return
	}
	defer udpConn.Close()
	fmt.Printf("UDP服务器通信socket创建完成!!!\n")

	//2.读取客户端发送的数据(阻塞发生在ReadFromUDP()方法中)
	buf := make([]byte, 4096)
	//ReadFromUDP()方法返回三个值,分别是读取到的字节数,客户端的地址,error
	n, clientUDPAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {

		fmt.Printf("*UDPAddr.ReadFromUDP()方法执行出错,错误为:%v\n", err)
		return
	}
	//3.模拟处理数据
	fmt.Printf("服务器读到%v的数据:%s", clientUDPAddr, buf[:n])

	//4.回写数据给客户端
	_, err = udpConn.WriteToUDP([]byte("I am OK!"), clientUDPAddr)
	if err != nil {
		fmt.Printf("*UDPAddr.WriteToUDP()方法执行出错,错误为:%v\n", err)
		return
	}
}
