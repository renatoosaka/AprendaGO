package main

import (
	"encoding/json"
	"fmt"
)

type info struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	ContaBancaria float64 `json:"ContaBancaria"`
}

func main() {
	str := []byte(`{"Nome":"Anakin","Sobrenome":"Skywalker","Idade":50,"Profissao":"Vilão","ContaBancaria":1203.22}`)

	var pessoa info

	err := json.Unmarshal(str, &pessoa)

	if err != nil {
		fmt.Println("Erro", err)
	}

	fmt.Println(pessoa)
}
