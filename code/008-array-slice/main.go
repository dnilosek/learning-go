package main

import "fmt"

func main() {
	// Arrays, dont use these unless you really need to
	var x [5]int
	fmt.Println(x)
	x[1] = 1
	fmt.Println(x)
	fmt.Printf("Length of x is %d\n", len(x))

	//Slice

	// x := type{values} -- compsite literal
	y := []int{4, 1, 2, 42, 0}
	fmt.Println(y)

	for idx, val := range y {
		fmt.Println(idx, val)
	}

	// Slice a slice
	fmt.Println(y[1:])
	fmt.Println(y[1:5]) // not inclusive end

	// Append to a slice
	y = append(y, 10, 20)
	fmt.Println(y)

	// Append a slice to a slice
	y = append(y, y...) //... unfurl the slice
	fmt.Println(y)

	// Delete from slice (2->3)
	y = append(y[:2], y[4:]...)
	fmt.Println(y)

	// Make with capacity
	z := make([]int, 10, 100)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	var jamesdat [][]string
	jamesdat = append(jamesdat, []string{"James", "Bond", "Shaken, not stirred"})
	jamesdat = append(jamesdat, []string{"Miss", "Moneypenny", "Helloooooo James"})

	for i, xs := range jamesdat {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index postion: %v \t value: \t %v\n", j, val)
		}
		fmt.Println("\n")
	}
}
