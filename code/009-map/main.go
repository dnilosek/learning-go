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

	jamesmap := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	jamesmap["q_q"] = []string{"cool stuff", "more cool stuff"}

	for key, val := range jamesmap {
		fmt.Printf("Key: %v\n", key)
		for i, v := range val {
			fmt.Printf("Index: %v \t Value: %v\n", i, v)
		}
	}

	delete(jamesmap, "q_q")
	for key, val := range jamesmap {
		fmt.Printf("Key: %v\n", key)
		for i, v := range val {
			fmt.Printf("Index: %v \t Value: %v\n", i, v)
		}
	}

}
