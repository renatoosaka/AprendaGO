package main

import (
	"fmt"
	"log"
)

type erroEspecial struct {
	especial string
}

func (e *erroEspecial) Error() string {
	return fmt.Sprintf("Erro Especial %v", (*e).especial)
}

func main() {
	e := erroEspecial{"erro customizado"}

	geraErro(&e)
}

func geraErro(e *erroEspecial) {
	log.Println((*e))
}
