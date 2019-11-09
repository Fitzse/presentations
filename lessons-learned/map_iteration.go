package main

import "fmt"

func main() {
	myMap := map[string]int{
		"a": 1,
		"b": 7,
		"c": 3,
	}

	for k, v := range myMap {
		fmt.Printf("Key: %v Value: %v\n", k, v)
	}
}
