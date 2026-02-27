package personagem

import (
	"RPG/modelos"
	"encoding/json"
)

func Salvar_json(ficha modelos.Personagem) ([]byte, error) {
	ficha_json, err := json.Marshal(ficha)
	return ficha_json, err
}

func Carregar_json(input []byte) (modelos.Personagem, error) {
	var ficha modelos.Personagem
	err := json.Unmarshal(input, &ficha)
	return ficha, err
}
