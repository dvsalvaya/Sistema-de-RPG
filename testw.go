package main

import (
	"RPG/personagem"
	"fmt"
)

func main() {
	fmt.Println("Bem-vindo ao RPG!")
	var escolha int
	fmt.Print("1 - Salvar personagem\n2 - Editar personagem\n3 - Subir de nível\n4 - Salvar_json\n-> ")
	fmt.Scan(&escolha)
	switch escolha {
	case 1:
		perso := personagem.Criar_personagem()
		personagem.InserirPersonagem(perso.Nome, perso.Forca, perso.Vida, perso.Classe, perso.Determinacao, perso.Inteligencia, perso.Exp, perso.Level)
		fmt.Printf("Personagem criado: %+v\nID: %d\nNome: %s\nVida: %d\nClasse: %d\nForça: %d\nDeterminação: %d\nInteligência: %d\nExp: %d\nLevel: %d\n", perso, perso.ID, perso.Nome, perso.Vida, perso.Classe, perso.Forca, perso.Determinacao, perso.Inteligencia, perso.Exp, perso.Level)
		fmt.Println("Personagem salvo no banco de dados!")
	case 2:
		var id int
		personagem.ListarPersonagens()
		fmt.Println("Insira o id")
		fmt.Scan(&id)
		personagem.AtualizarPersonagem(id)
		fmt.Println("Personagem atualizado no banco de dados!")
	case 3:
		var id int
		fmt.Println("Insira o id")
		fmt.Scan(&id)
		personagem.Up_level(id)
	case 4:
		var id int
		fmt.Println("Insira o id")
		fmt.Scan(&id)
		ficha_json, err := personagem.Salvar_json(personagem.BuscarPersonagemPorID(id))
		if err != nil {
			fmt.Println("Erro ao salvar personagem em JSON:", err)
			return
		}
		fmt.Printf("Personagem salvo em JSON: %s\n", string(ficha_json))
	case 8:
		personagem.CriarBanco()
	default:
		fmt.Println("Opção inválida.")
		personagem.ListarPersonagens()
	}
}
