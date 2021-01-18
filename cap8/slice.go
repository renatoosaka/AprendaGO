package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 6)
	fmt.Println(slice2)

	fmt.Println(slice[3])
	slice[3] = 52412
	fmt.Println(slice[3])

	fmt.Println(slice)

	slice3 := []string{"banana", "maçã", "jaca"}

	// for idx, value := range slice3 {
	// 	fmt.Printf("No índice %v temos o valor: %v\n", idx, value)
	// }

	slice3 = append(slice3, "melancia")

	for idx, value := range slice3 {
		fmt.Printf("No índice %v temos o valor: %v\n", idx, value)
	}

	slice4 := []int{20, 21, 23}

	total := 0

	for _, value := range slice4 {
		total += value
	}

	fmt.Printf("Total: %d\n", total)

	sabores := []string{"pepperoni", "mussarela", "abacaxi", "quatro queijos", "marguerita"}

	fatia0 := sabores[0:2]
	fatia1 := sabores[1:3]
	fatia2 := sabores[2:len(sabores)]

	fmt.Println(sabores)
	fmt.Println(fatia0)
	fmt.Println(fatia1)
	fmt.Println(fatia2)

	for x := 0; x < len(sabores); x++ {
		fmt.Println(sabores[x])
	}

	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)

}
