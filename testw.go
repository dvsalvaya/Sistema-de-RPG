package main

import (
	"RPG/personagem"
	"fmt"
)

func main() {
	pero := personagem.Criar_personagem()
	fmt.Printf("Nome: %s\nForça: %d", pero.Nome, pero.Forca)
}
