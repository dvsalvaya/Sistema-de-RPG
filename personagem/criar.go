package personagem

import (
	"RPG/modelos"
	"fmt"
)

func Criar_personagem() modelos.Personagem {
	var novo_personagem modelos.Personagem

	fmt.Print("Digite o nome do personagem\n-> ")
	fmt.Scan(&novo_personagem.Nome)

	fmt.Print("Digite a idade do personagem\n-> ")
	fmt.Scan(&novo_personagem.Idade)

	fmt.Print("Escolha a classe do personagem\n1 - Guerreiro\n2 - Mago\n-> ")
	fmt.Scan(&novo_personagem.Classe)

	fmt.Print("Digite a determinação do personagem\n-> ")
	fmt.Scan(&novo_personagem.Determinacao)

	fmt.Print("Digite a força do personagem\n-> ")
	fmt.Scan(&novo_personagem.Forca)

	fmt.Print("Digite a vida do personagem\n-> ")
	fmt.Scan(&novo_personagem.Vida)

	fmt.Print("Digite a inteligencia do personagem\n-> ")
	fmt.Scan(&novo_personagem.Inteligencia)

	novo_personagem.Exp = 0.0
	novo_personagem.Level = 1
	return novo_personagem
}

func Up_level(perso modelos.Personagem) {
	switch perso.Classe {
	case 1:
		perso.Forca = perso.Forca + 20
		perso.Vida = perso.Vida + 50
	case 2:
		perso.Inteligencia = perso.Inteligencia + 20
		perso.Vida = perso.Vida + 30
	}
	perso.Determinacao++
	perso.Exp = 0.00
	perso.Level++
}
