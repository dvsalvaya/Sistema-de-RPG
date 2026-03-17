import sqlite3 as sq
import Dados
def criar_personagem():
    nome = input("Insere o nome")
    idade = int(input("isere a idade"))

    # Classes
    classe = input("Classe: ")
    if classe == "Guerreiro":
        vida = 100
        forca = 20
        inteligencia = 10
        mira = 15
    elif classe == "Arqueiro":
        vida = 60
        forca = 10
        inteligencia = 15
        mira = 20
    elif classe == "Mago":
        vida = 70
        forca = 10
        inteligencia = 20
        mira = 10
    else:
        print("Erro")

    ####
        return 
    perso = {#"ID" : int,
            "Nome" : f"{nome}",
            "Idade" : idade,
            "Classe" : f"{classe}",
            "Vida" : vida,
                    
            "Forca" : forca,
            "Mira" : mira,
            "Inteligencia" : inteligencia,
            "Determinação" : 15,
            "Level" : 1,
            "Exp" : 0
            }
    return perso

def criar_banco():
    with sq.connect("fichas.db") as conn:
        cur = conn.cursor()
        cur.execute(f"""
        
        CREATE TABLE IF NOT EXISTS fichas(
        id INTEGER PRIMARY KEY,
        NAME TEXT,
        IDADE INTEGER,
        CLASSE TEXT,
        VIDA INTEGER,
        FORÇA INTEGER,
        MIRA INTEGER,
        INTELIGENCIA INTEGER,
        LEVEL INTEGER,
        EXP INTEGER
                )
        """)

def salvar_personagem(perso):
    with sq.connect("fichas.db") as conn:
        cur = conn.cursor()
        cur.execute("""
    INSERT INTO fichas
    (NAME, IDADE, CLASSE, VIDA, FORÇA, MIRA, INTELIGENCIA, LEVEL, EXP)
    VALUES 
    (:Nome, :Idade, :Classe, :Vida, :Forca, :Mira, :Inteligencia, :Level, :Exp)
""", perso)
    
def exibir_personagens():
    with sq.connect("fichas.db") as conn:
        cur = conn.cursor()
        cur.execute("""
        SELECT * FROM fichas
        """)
        rows = cur.fetchall()
        for row in rows:
            print(f"""ID {row[0]}\n
        NAME {row[1]}\n
        IDADE {row[2]}\n
        CLASSE {row[3]}\n
        VIDA {row[4]}\n
        FORÇA {row[5]}\n
        MIRA {row[6]}\n
        INTELIGENCIA {row[7]}\n
        LEVEL {row[8]}\n
        EXP {row[9]}""")

#criar_banco()
salvar_personagem(criar_personagem())
exibir_personagens()
