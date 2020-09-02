package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string	`json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
}

func jsonToStruct() {
	jsonStr := `{"name":"xiaoming", "age":19, "gender":"M"}`

	var people People
	if err := json.Unmarshal([]byte(jsonStr), &people); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(people)

	}
}

func structToJson() {
	people := People{
		"xiaoming",
		10,
		"M",
	}

	jsonBytes, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}

func jsonToMap() {
	jsonStr := `{"name":"xiaoming", "age":19, "gender":"M"}`

	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mapResult)

	}
}

func mapToJson() {
	mapResult := make(map[string]interface{})
	mapResult["name"] = "xiaoming"
	mapResult["age"] = 18
	mapResult["gender"] = "F"

	jsonBytes, err := json.Marshal(mapResult)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}

func main() {

	//jsonToStruct()

	//structToJson()

	//jsonToMap()

	mapToJson()
}
