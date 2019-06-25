package main

import "fmt"

type person struct {
	first           string
	last            string
	favorite_icream []string
}

func main() {
	p1 := person{
		first:           "Me",
		last:            "Melast",
		favorite_icream: []string{"Yum", "YumYum"},
	}
	p2 := person{
		first:           "You",
		last:            "YouLast",
		favorite_icream: []string{"YumYum", "YumYumYum"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
