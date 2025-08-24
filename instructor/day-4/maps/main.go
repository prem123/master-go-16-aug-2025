package main

import "fmt"

func main() {
	// map[key-type]value-type
	// var m map[string]int

	// var m1 map[int]string

	m := make(map[string]int)

	m["zero"] = 0
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	key := "one"
	v, ok := m[key]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Printf("%s not found...\n", key)
	}

	// this is another example. The scope of value will be limited to if-else block
	if value, ok := m[key]; ok {
		fmt.Println(value)
	} else {
		fmt.Printf("%s not found...\n", key)
	}

	weekdays := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
	}

	fmt.Println(weekdays[2])
	delete(weekdays, 2)
	fmt.Println(weekdays)

	// iterating over a map
	//for key, val := range < array/slice/string/map>
	for k, v := range m {
		fmt.Printf("key: %s, value: %d \n", k, v)
	}
}
