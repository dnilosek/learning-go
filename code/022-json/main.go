package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string //To marshal it must start with captial
	Last  string `json:"lastName"` //Tag to change JSON field
	Age   int
}

func main() {

	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	s := `[{"First":"James","lastName":"Bond","Age":32},{"First":"Miss","lastName":"Moneypenny","Age":27}]`
	bs_in := []byte(s)

	var person_in []person
	err = json.Unmarshal(bs_in, &person_in)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(person_in)

}
