package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	s_wrong := `pasword123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(string(bs))

	err = bcrypt.CompareHashAndPassword(bs, []byte(s))
	if err != nil {
		fmt.Println("NO, WRONG")
	}
	fmt.Println("You're in")

	err = bcrypt.CompareHashAndPassword(bs, []byte(s_wrong))
	if err != nil {
		fmt.Println("NO, WRONG")
	}
}
