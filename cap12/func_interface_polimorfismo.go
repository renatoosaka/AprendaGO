package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesArrancados int
	salario          float64
}
type arquiteto struct {
	pessoa
	tipoConstrucao string
	tamanhoLoucura string
}

func (r dentista) oiBomDia() {
	fmt.Println("Oi, Bom Dia (dentista)")
}

func (r arquiteto) oiBomDia() {
	fmt.Println("Oi, Bom Dia (arquiteto)")
}

type humano interface {
	oiBomDia()
}

func serHumano(h humano) {
	h.oiBomDia()

	switch h.(type) {
	case dentista:
		fmt.Println("arranquei ", h.(dentista).dentesArrancados, "dentes")
	case arquiteto:
		fmt.Println("meu nivel de loucura é ", h.(arquiteto).tamanhoLoucura)
	}
}

func main() {
	dentinho := dentista{pessoa{"Dentista", "Sobrenome", 32}, 4, 10000}
	predinho := arquiteto{pessoa{"Arquiteto", "Sobrenome", 32}, "é dificil", "max"}

	fmt.Println(dentinho)
	fmt.Println(predinho)

	serHumano(dentinho)
	serHumano(predinho)
}
