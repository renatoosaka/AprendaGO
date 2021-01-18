package main

import (
	"fmt"
	"sort"
)

type Carro struct {
	Nome     string
	Potencia int
	Consumo  int
}

type OrdenaPorPotencia []Carro

func (l OrdenaPorPotencia) Len() int {
	return len(l)
}

func (l OrdenaPorPotencia) Less(i, j int) bool {
	return l[i].Potencia < l[j].Potencia
}

func (l OrdenaPorPotencia) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type OrdenaPorConsumo []Carro

func (l OrdenaPorConsumo) Len() int {
	return len(l)
}

func (l OrdenaPorConsumo) Less(i, j int) bool {
	return l[i].Consumo > l[j].Consumo
}

func (l OrdenaPorConsumo) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type OrdenaPorLucroProDonoDoPosto []Carro

func (l OrdenaPorLucroProDonoDoPosto) Len() int {
	return len(l)
}

func (l OrdenaPorLucroProDonoDoPosto) Less(i, j int) bool {
	return l[i].Consumo < l[j].Consumo
}

func (l OrdenaPorLucroProDonoDoPosto) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func main() {
	carros := []Carro{Carro{"DS5", 2, 1}, Carro{"Gol", 1, 3}, Carro{"Ferrari", 10, 2}}

	fmt.Println(carros)
	sort.Sort(OrdenaPorPotencia(carros))
	fmt.Println(carros)

	sort.Sort(OrdenaPorConsumo(carros))
	fmt.Println(carros)

	sort.Sort(OrdenaPorLucroProDonoDoPosto(carros))
	fmt.Println(carros)
}
