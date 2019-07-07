package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "test.txt"
	//filePath := "/usr/test.txt"

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(err) // same as err.Error()
		return
	}
	defer f.Close()

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
}
