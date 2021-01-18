package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ordenaPorIdade []user

func (l ordenaPorIdade) Len() int {
	return len(l)
}

func (l ordenaPorIdade) Less(i, j int) bool {
	return l[i].Age < l[j].Age
}

func (l ordenaPorIdade) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type ordenaPorSobrenome []user

func (l ordenaPorSobrenome) Len() int {
	return len(l)
}

func (l ordenaPorSobrenome) Less(i, j int) bool {
	return l[i].Last < l[j].Last
}

func (l ordenaPorSobrenome) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(ordenaPorIdade(users))

	fmt.Println(users)

	sort.Sort(ordenaPorSobrenome(users))

	fmt.Println(users)
}
