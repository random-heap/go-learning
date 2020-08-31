package main

import (
	"fmt"
)

type NOTFoundError struct {
	name string
}

func (e *NOTFoundError) Error() string {
	return fmt.Sprintf("%s  is not found, please new again", e.name)
}

func NewNotFoundError(name string) error {
	return &NOTFoundError{name}
}

func runDIYError() {
	err := NewNotFoundError("your object")

	// 根据switch，确定是哪种error
	switch err.(type) {
	case *NOTFoundError:
		fmt.Printf("error : %v \n",err)
	default: // 其他类型的错误
		fmt.Println("other error")
	}
}


func _main() {
	//sources := []string{"hello", "world", "souyunku", "gostack"}
	//fmt.Println(getN(0, sources))//直接调用，会打印两项内容，字符串元素以及error空对象
	//fmt.Println(getN(1, sources))
	//fmt.Println(getN(2, sources))
	//fmt.Println(getN(3, sources))
	//
	//target, err := getN(4, sources)//将返回结果赋值
	//if err != nil {//常见的错误处理，如果error不为nil，则进行错误处理
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(target)

	//notFound := errors.New("404 not found")
	//fmt.Println(notFound)
	//
	//fmt.Println(fmt.Errorf("404: page %v is not found","index.html"))

	//runDIYError()

	//runDefer()

	//deferCall()

	runSimplePanic()
}

func runSimplePanic(){
	defer func() {
		fmt.Println("before panic")
	}()
	panic("simple panic")
}

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}


func runDefer(){
	defer func() {
		fmt.Println("3")
	}()//括号表示定义function之后直接执行

	fmt.Println("1")

	defer func(index string) {
		fmt.Println(index)
	}("2")
}


//定义函数获取第N个元素，正常返回元素以及为nil的error，异常返回空元素以及error
func getN(n int, sources []string) (string, error) {
	if n > len(sources)-1 {
		return "", fmt.Errorf("%d, out of index range %d", n, len(sources) - 1)
	}
	return sources[n], nil
}

