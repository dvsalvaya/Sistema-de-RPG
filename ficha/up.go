package ficha

import "modelos"

func Up_level(perso modelos.Personagem) {
	switch perso.classe {
	case 1:
		perso.forca = perso.Forca + 20
		perso.vida = perso.Vida + 50
	case 2:
		perso.inteligencia = perso.Inteligencia + 20
		perso.vida = perso.Vida + 30
	}
	perso.Determinacao++
	perso.Exp = 0.0
	perso.Level++
}
