package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//isExist, err := nativeFileExist("D:/data0/go/test")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(isExist)
	//if !isExist {
	//	os.Mkdir("D:/data0/go/test", 0777)
	//}

	//nativeCreate()

	//nativeRead()

	//ioutilRead()

	bufioRead()
}

func nativeFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func nativeCreate() {

	file, err := os.Create("D:/data0/go/test/test.txt")
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString("hello world!!!")

	defer file.Close()
}

func nativeRead() {

	file, err := os.OpenFile("D:/data0/go/test/test.txt", os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var chunk []byte
	buf := make([]byte, 1024)

	for {
		//从file读取到buf中
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}

	fmt.Println(string(chunk))
}

func ioutilRead() {
	f, err := ioutil.ReadFile("D:/data0/go/test/test.txt")
	if err != nil {
		fmt.Println("read fail", err)
	}

	fmt.Println(string(f))
}

func bufioRead() {
	fi, err := os.Open("D:/data0/go/test/test.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	r := bufio.NewReader(fi)
	var chunks []byte

	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		//fmt.Println(string(buf))
		chunks = append(chunks, buf...)
	}
	fmt.Println(string(chunks))
}
