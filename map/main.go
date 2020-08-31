package main

import (
	"fmt"
	"sort"
)

func main()  {

	var m1 map[string]int
	m1 = make(map[string]int)

	values := []string{"a", "b", "c", "d", "e"}
	for i, v := range values {
		m1[v] = i
	}
	fmt.Println(m1)

	delete(m1, "a")

	fmt.Println(m1)

	value := m1["b"]
	fmt.Println("value = ", value)

	m1["b"] = 5
	value = m1["b"]
	fmt.Println("value = ", value)

	for k, v := range m1 {
		fmt.Printf("key = %s, value = %d\n", k, v)
	}

	//判断key是否存在
	v, ok := m1["b"]
	if ok {
		fmt.Println("存在key=", v)
	} else {
		fmt.Println("不存在key=", v)
	}

	m1["f"] = 10
	v, ok = m1["f"]
	if ok {
		fmt.Println("存在key=", v)
	} else {
		fmt.Println("不存在key=", v)
	}

	languages := map[string]int {
		"Java":0,
		"C":1,
		"C++":2,
		"Python":3,
		"C#":4,
		"PHP":5,
		"JavaScript":6,
		"Visual Basic.NET":7,
		"Perl":8,
		"Assembly language":9,
		"Ruby":10,
	}

	var sortStr []string
	for k, _ := range languages {
		sortStr = append(sortStr, k)
	}

	sort.Strings(sortStr)
	for _, k := range  sortStr {
		fmt.Println("Key:", k, "Value:", languages[k])
	}
}
