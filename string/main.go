package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//数值、字符串互相转换
	var s1 string = "100"
	num1, err1 := strconv.Atoi(s1)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("num1 = ", num1)
	}

	var num2 int = 100
	s2 := strconv.Itoa(num2)
	fmt.Println("s2 = ", s2)

	//字符串比较
	var str1 string = "hello, world"
	var str2 string = "你好，世界"
	var comp1 = strings.Compare(str1, str2)
	fmt.Println("comp1 = ", comp1)

	//包含
	var exist bool = strings.Contains("hello world", "1hello")
	fmt.Println("exist = ", exist)

	//查找
	var index int = strings.Index("what a amazing!!!", "amazing")
	fmt.Println("index = ", index)

	//统计
	var count int = strings.Count("chinese", "e")
	fmt.Println("count = ", count)

	//替换
	var newStr1 string = strings.Replace("I have a dream.", "dream", "idea", 1)
	fmt.Println("newStr1 = ", newStr1)

	//删除
	var newStr2 string = strings.Trim("/hello, world/", "/")
	fmt.Println("newStr2 = ", newStr2)

	//大小写转换
	fmt.Println(strings.ToLower("Hello, World..."))
	fmt.Println(strings.ToUpper("Hello, World..."))

	//前缀、后缀
	fmt.Println(strings.HasPrefix("Golang, Java, C++", "Golang"))
	fmt.Println(strings.HasSuffix("Golang, Java, C++", "C++"))

	//字符串分割
	var s3 string = "Yesterday    once more"
	wordSlice1 := strings.Fields(s3)
	for i, v := range wordSlice1 {
		fmt.Printf("wordSlice1 下标 %d 对应值 = %s \n", i, v)
	}

	var s4 string = "111,222,333"
	wordSlice2 := strings.Split(s4, ",")
	for i, v := range wordSlice2 {
		fmt.Printf("wordSlice2 下标 %d 对应值 = %s \n", i, v)
	}

	//拼接
	var s5 string = "a" + "b" + "c"
	fmt.Println("s5 = ", s5)
	var s6 string = strings.Join(wordSlice2, ";")
	fmt.Println("s6 = ", s6)

	var buffer bytes.Buffer
	for i := 0; i < 10; i++ {
		buffer.WriteString("let's learn golang.\n")
	}
	fmt.Println("buffer is " + buffer.String())

}