package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

// Map Literals
var myLitralMap = map[string]Vertex{
	"Sample": Vertex{123.533, 34234.34},
	"Data":   Vertex{223.53, 222.34},
}

var myLitralMap2 = map[string]Vertex{
	"Sample": {123.533, 34234.34},
	"Data":   {223.53, 222.34},
}

func main() {
	m := make(map[string]Vertex)
	m["Bell"] = Vertex{12.345, 987.322}

	fmt.Print(m["Bell"])
	fmt.Printf("\n")
	fmt.Println(myLitralMap)
	fmt.Println(myLitralMap2)

	simple := make(map[string]int)
	simple["sonu"] = 1
	simple["rai"] = 2

	fmt.Println(simple)
	simple["sonu"] = 5

	fmt.Println(simple["sonu"])

	delete(simple, "sonu")

	_, status := simple["sonu"]

	fmt.Println(status)

	myMap := WordCount("Hello Kaushal, how are you? Hope you are doing fine")
	fmt.Println(myMap)

	newMap := map[string]string{"1": "a", "2": "b", "3": "c"}

	for key, value := range newMap {
		fmt.Println("Key: " + key + "  , value: " + value)

		delete(newMap, key)
	}

	fmt.Println(len(newMap))

}

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	fmt.Println(len(words))

	wordMap := make(map[string]int)
	for _, value := range words {
		count := wordMap[value]

		wordMap[value] = count + 1
	}

	return wordMap
}
