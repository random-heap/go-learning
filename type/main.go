package main

import (
	"fmt"
	"time"
	"unicode/utf8"
	"unsafe"
)

func main() {

	_bool()
	//_int()
	//_float()
	//_complex()
	//_string()
	//_rune()
	//_pointer()
}

func _pointer() {
	var p *int
	var i int = 10
	p = &i
	*p++

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)
}

func _rune() {
	var chinese = "我是中国人， I am Chinese"
	fmt.Println("chinese length", len(chinese))
	fmt.Println("chinese word length", len([]rune(chinese)))
	fmt.Println("chinese word length", utf8.RuneCountInString(chinese))
}

func _string() {
	var s1 string = ""
	var s2 string = "hello world."
	s3 := "hello world!!!"

	fmt.Println("the len of s1", len(s1))
	fmt.Println("the size of s1", unsafe.Sizeof(s1))
	fmt.Println("the len of s2", len(s2))
	fmt.Println("the size of s1", unsafe.Sizeof(s2))
	fmt.Println("the len of s3", len(s3))
	fmt.Println("the size of s1", unsafe.Sizeof(s3))

}

func _complex() {
	var c64 complex64 = 5 + 10i
	var c128 complex128 = 5 + 10i

	fmt.Println("the size of complex64", unsafe.Sizeof(c64))
	fmt.Println("the size of complex128", unsafe.Sizeof(c128))
}

func _float() {

	var f32 float32 = 0
	var f64 float64 = 0

	fmt.Println("the size of float32", unsafe.Sizeof(f32))
	fmt.Println("the size of float64", unsafe.Sizeof(f64))

}

func _int() {

	var b byte = 0

	var i int = 0
	var i8 int8 = 0
	var i16 int16 = 0
	var i32 int32 = 0
	var i64 int64 = 0

	var  ui uint = 0
	var ui8 uint8 = 0
	var ui16 uint16 = 0
	var ui32 uint32 = 0
	var ui64 uint64 = 0

	fmt.Println("the size of byte", unsafe.Sizeof(b))

	fmt.Println("the size of int", unsafe.Sizeof(i))
	fmt.Println("the size of int8", unsafe.Sizeof(i8))
	fmt.Println("the size of int16", unsafe.Sizeof(i16))
	fmt.Println("the size of int32", unsafe.Sizeof(i32))
	fmt.Println("the size of int64", unsafe.Sizeof(i64))

	fmt.Println("the size of unsigned int", unsafe.Sizeof(ui))
	fmt.Println("the size of unsigned int8", unsafe.Sizeof(ui8))
	fmt.Println("the size of unsigned int16", unsafe.Sizeof(ui16))
	fmt.Println("the size of unsigned int32", unsafe.Sizeof(ui32))
	fmt.Println("the size of unsigned int64", unsafe.Sizeof(ui64))

}

func _bool() {

	count := 0
	var flag bool = true

	fmt.Println("the size of bool", unsafe.Sizeof(flag))

	for flag {

		fmt.Printf("%d : bool type test\n", count)
		count++
		time.Sleep(1000)
		if count == 10 {
			flag = false
		}
	}
}
