package main

import "fmt"

type person struct {
	first           string
	last            string
	favorite_icream []string
}

type vehicle struct {
	color string
	doors int
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
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

	person_map := map[string]person{
		"me":  p1,
		"you": p2,
	}

	for _, person_val := range person_map {
		fmt.Printf("First name: %v \t Last name: %v \t\n Favorite Ice Creams:\n", person_val.first, person_val.last)
		for _, icream := range person_val.favorite_icream {
			fmt.Printf("\t\t\t%v\n", icream)
		}
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "green",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}
	fmt.Println(t1)
	fmt.Println(s1)

	anon := struct {
		mapD   map[string]int
		sliceD []string
	}{
		mapD: map[string]int{
			"d1": 0,
			"d2": 1,
		},
		sliceD: []string{"d1", "d2"},
	}

	fmt.Println(anon)
}
