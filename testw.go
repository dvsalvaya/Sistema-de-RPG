package main

import (
	"/home/duck/Projetos/go/rpg/ficha"
	"fmt"
)

func main() {
	pero := ficha.Criar_personagem()
	fmt.Printf("Nome: %s\nForça: %d", pero.nome, pero.forca)
}
