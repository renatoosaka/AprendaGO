package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"joao":  12345,
		"maria": 67890,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joao"])

	amigos["gopher"] = 9876

	fmt.Println(amigos)

	//comma ok idiom
	phone, ok := amigos["jose"]

	fmt.Println(phone, ok)

	if phone, ok := amigos["maria"]; ok {
		fmt.Println(phone)
	} else {
		fmt.Println("joana nao existe")
	}

	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		19:  "idade de ir pra festa",
	}

	fmt.Println(qualquercoisa)

	total := 0

	for key, value := range qualquercoisa {
		total += key
		fmt.Println(key, value)
	}

	fmt.Println(total)

	delete(qualquercoisa, 123)

	fmt.Println(qualquercoisa)
}
