package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	ContaBancaria float64
}

func main() {
	jamesBond := pessoa{"James", "Bond", 43, "Agente Secreto", 1000.01}

	darthVader := pessoa{"Anakin", "Skywalker", 50, "Vilão", 1203.22}

	j, err := json.Marshal(jamesBond)

	if err != nil {
		fmt.Println("Erro", err)
	}

	d, err := json.Marshal(darthVader)

	if err != nil {
		fmt.Println("Erro", err)
	}

	fmt.Println(string(j))
	fmt.Println(string(d))
}
