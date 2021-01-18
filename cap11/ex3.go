package main

import (
	"fmt"
)

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	quatroRodas bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	hilux := caminhonete{veiculo{4, "preta"}, true}
	fusion := sedan{veiculo{4, "branca"}, true}

	fmt.Println(hilux)
	fmt.Println(fusion)
}
