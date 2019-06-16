package main

import "fmt"

func main() {

	m := map[string]int{
		"a": 0,
		"b": 1,
	}
	fmt.Println(m)

	v, ok := m["a"]
	fmt.Println(v, ok)

	v, ok = m["c"]
	fmt.Println(v, ok)

	if v, ok := m["a"]; ok {
		fmt.Println("Comma okay idiom", v, ok)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "a")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
