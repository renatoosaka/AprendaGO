package main

import (
	"encoding/json"
	"os"
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

	encoder := json.NewEncoder(os.Stdout)

	encoder.Encode(jamesBond)
}
