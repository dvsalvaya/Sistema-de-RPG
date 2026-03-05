package personagem

import (
	"RPG/modelos"
	"database/sql"
	"fmt"

	_ "github.com/glebarez/sqlite"
)

func CriarBanco() {
	db, err := sql.Open("sqlite3", "sqlite/fichas.db")
	if err != nil {
		fmt.Println("Erro ao criar o banco de dados:", err)
		return
	}
	defer db.Close()

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS personagens (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT,
		forca INTEGER,
		vida INTEGER,
		classe INTEGER
		determinacao INTEGER
		inteligencia INTEGER
		exp INTEGER
		level INTEGER
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		fmt.Println("Erro ao criar a tabela:", err)
		return
	}
	fmt.Println("Banco de dados e tabela criados com sucesso!")
}

func InserirPersonagem(nome string, forca int, vida int, classe int, determinacao int, inteligencia int, exp int, level int) {
	db, err := sql.Open("sqlite3", "sqlite/fichas.db")
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	defer db.Close()

	insertQuery := `
	INSERT INTO personagens (nome, forca, vida, classe, determinacao, inteligencia, exp, level)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`
	_, err = db.Exec(insertQuery, nome, forca, vida, classe, determinacao, inteligencia, exp, level)
	if err != nil {
		fmt.Println("Erro ao inserir personagem:", err)
		return
	}
	fmt.Println("Personagem inserido com sucesso!")
}

func ListarPersonagens() {
	db, err := sql.Open("sqlite3", "sqlite/fichas.db")
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, nome, forca, vida, classe, determinacao, inteligencia, exp, level FROM personagens")
	if err != nil {
		fmt.Println("Erro ao listar personagens:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nome string
		var forca int
		var vida int
		var classe int
		var determinacao int
		var inteligencia int
		var exp int
		var level int

		err = rows.Scan(&id, &nome, &forca, &vida, &classe, &determinacao, &inteligencia, &exp, &level)
		if err != nil {
			fmt.Println("Erro ao ler personagem:", err)
			return
		}
		fmt.Printf("ID: %d, Nome: %s, Força: %d, Vida: %d, Classe: %d, Determinação: %d, Inteligência: %d, Exp: %.2f, Level: %d\n",
			id, nome, forca, vida, classe, determinacao, inteligencia, exp, level)
	}
}

func DeletarPersonagem(id int) {
	db, err := sql.Open("sqlite3", "sqlite/fichas.db")
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	defer db.Close()

	deleteQuery := "DELETE FROM personagens WHERE id = ?"
	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		fmt.Println("Erro ao deletar personagem:", err)
		return
	}
	fmt.Println("Personagem deletado com sucesso!")
}

// A função não atualiza nada
func AtualizarPersonagem(id int) {
	perso := BuscarPersonagemPorID(id)
	db, err := sql.Open("sqlite3", "sqlite/fichas.db")
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	defer db.Close()

	updateQuery := `
	UPDATE personagens
	SET nome = ?, forca = ?, vida = ?, classe = ?, determinacao = ?, inteligencia = ?, exp = ?, level = ?
	WHERE id = ?;
	`
	_, err = db.Exec(updateQuery, perso.Nome, perso.Forca, perso.Vida, perso.Classe, perso.Determinacao, perso.Inteligencia, perso.Exp, perso.Level, perso.ID)
	if err != nil {
		fmt.Println("Erro ao atualizar personagem:", err)
		return
	}
	fmt.Println("Personagem atualizado com sucesso!")
}

func BuscarPersonagemPorID(id int) modelos.Personagem {
	db, err := sql.Open("sqlite3", "sqlite/fichas.db")
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		//return
	}
	defer db.Close()

	//var nome string
	//var forca int
	//var vida int
	//var classe int
	//var determinacao int
	//var inteligencia int
	//var exp int
	//var level int

	var perso modelos.Personagem

	query := "SELECT nome, forca, vida, classe, determinacao, inteligencia, exp, level FROM personagens WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&perso.Nome, &perso.Forca, &perso.Vida, &perso.Classe, &perso.Determinacao, &perso.Inteligencia, &perso.Exp, &perso.Level)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Personagem não encontrado.")
		} else {
			fmt.Println("Erro ao buscar personagem:", err)
		}
		//return
	}
	return perso
	//fmt.Printf("ID: %d, Nome: %s, Força: %d, Vida: %d, Classe: %d, Determinação: %d, Inteligência: %d, Exp: %.2f, Level: %d\n",
	//	id, nome, forca, vida, classe, determinacao, inteligencia, exp, level)
}

func RetirarPersonagem(id int) modelos.Personagem {
	db, err := sql.Open("sqlite3", "sqlite/fichas.db")
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	Query := "SELECT id, nome, forca, vida, classe, determinacao, inteligencia, exp, level FROM personagens WHERE id = ?"
	var perso modelos.Personagem
	err = db.QueryRow(Query, id).Scan(&perso.ID, &perso.Nome, &perso.Forca, &perso.Vida, &perso.Classe, &perso.Determinacao, &perso.Inteligencia, &perso.Exp, &perso.Level)
	if err != nil {
		fmt.Println("Erro ao retirar personagem:", err)
	}
	return perso
}

func Up_level(id int) {
	perso := BuscarPersonagemPorID(id)
	switch perso.Classe {
	case 1:
		perso.Forca = perso.Forca + 20
		perso.Vida = perso.Vida + 50
	case 2:
		perso.Inteligencia = perso.Inteligencia + 20
		perso.Vida = perso.Vida + 30
	}
	perso.Determinacao++
	perso.Exp = 0
	perso.Level++
	AtualizarPersonagem(id)
}
