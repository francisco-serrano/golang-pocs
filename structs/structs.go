package structs

import (
	"fmt"
	"reflect"
)

type Author struct {
	name      string
	branch    string
	language  string
	Particles int
}

func Run() {
	a1 := Author{
		name:      "Moana",
		branch:    "CSE",
		language:  "Python",
		Particles: 38,
	}

	a2 := Author{
		name:      "Moana",
		branch:    "CSE",
		language:  "Python",
		Particles: 38,
	}

	a3 := Author{
		name:      "Dona",
		branch:    "CSE",
		language:  "Python",
		Particles: 38,
	}

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println(a1 == a2)
	fmt.Println(a2 == a3)

	fmt.Println(reflect.DeepEqual(a1, a2))
}
