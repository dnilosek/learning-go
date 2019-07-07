package main

import "fmt"

type customErr struct {
	DataPayload string
}

func (c customErr) Error() string {
	return fmt.Sprintf("There hath been an error?! %v", c.DataPayload)
}

func main() {
	cusErr := customErr{
		DataPayload: "All the base blong to us",
	}
	foo(cusErr)
}

func foo(e error) {
	fmt.Println(e.Error())
}
