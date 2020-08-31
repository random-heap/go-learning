package main

import "fmt"

func main() {

	var arr1 [3]int
	arr1[0] = 1

	arr2 := [5]int{1, 2, 3, 4, 5}   //指定长度为5，并赋5个初始值
	arr3 := [5]int{1, 2, 3}         //指定长度为5，对前3个元素进行赋值，其他元素为零值
	arr4 := [5]int{4: 1}            //指定长度为5，对第5个元素赋值
	arr5 := [...]int{1, 2, 3, 4, 5} //不指定长度，对数组赋以5个值
	arr6 := [...]int{8: 1}          //不指定长度，对第9个元素（下标为8）赋值1

	fmt.Println("arr1 =", arr1, " len is ", len(arr1))
	fmt.Println("arr2 =", arr2, " len is ", len(arr2))
	fmt.Println("arr3 =", arr3, " len is ", len(arr3))
	fmt.Println("arr4 =", arr4, " len is ", len(arr4))
	fmt.Println("arr5 =", arr5, " len is ", len(arr5))
	fmt.Println("arr6 =", arr6, " len is ", len(arr6))

	for i, v := range arr5 {
		fmt.Println(i, v)
	}

	f1(arr5)

	fmt.Println("main :", arr5)
}

func f1(arr [5]int) {
	fmt.Println("f1 update before", arr)

	arr[0] = 0

	fmt.Println("f1 update after", arr)
}