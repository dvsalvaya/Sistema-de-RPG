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

	novo_personagem.Exp = 0
	novo_personagem.Level = 1

	//No futuro tem que mudar essa logica, pq assim pode vim dois persogens com o mesmo ID
	return novo_personagem
}
