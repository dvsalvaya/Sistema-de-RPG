# 🎲 VTT RPG Platform (Desktop App)

> Uma aplicação Desktop leve e de alta performance para gerenciamento de mesas de RPG, focada em simplicidade e interação local.

![Go Badge](https://img.shields.io/badge/Backend-Go-blue)
![Frontend Badge](https://img.shields.io/badge/Frontend-???-lightgrey)
![Status Badge](https://img.shields.io/badge/Status-Em_Desenvolvimento-yellow)

## 📖 Sobre o Projeto

Este projeto é um **Virtual Tabletop (VTT)** em formato de aplicativo Desktop. Diferente de plataformas web pesadas, nosso foco é oferecer uma experiência nativa, rápida e offline-first, permitindo gestão de fichas e rolagens de dados sem complicações.

O sistema é construído com **Go** no backend e utiliza **Wails** para compilar como um executável nativo. O frontend ainda será definido.

## ✨ Funcionalidades Principais

### 🛡️ Gestão de Mesas e Sessões
- **Salas Locais:** Criação e gerenciamento de mesas diretamente no computador.
- **Relatório Automático:** Logs de sessão salvos localmente.

### 📝 Fichas de Personagem Dinâmicas
- **Editor Flexível:** Criação e edição de fichas.
- **Persistência Local:** Fichas salvas automaticamente no banco de dados local (SQLite).
- **Portabilidade:** Exportação de fichas em formato JSON.

### 🎲 Rolador de Dados (Dice Roller)
- **Suporte Completo:** d4, d6, d8, d10, d12, d20 e d100.
- **Rolagem Rápida:** Execução nativa via Go.

---

## 🛠️ Tecnologias Utilizadas

A arquitetura foi pensada para criar um executável único e leve.

### Core (Desktop App)
- **Framework:** [Wails](https://wails.io/) (Go + Web Frontend)
- **Linguagem Backend:** [Go (Golang)](https://go.dev/)
- **Banco de Dados:** SQLite (Arquivo local `rpg.db`)

### Frontend
- **Framework:** ??? (A definir)
- **Estilização:** ??? (A definir)

---

## 🚀 Como Rodar o Projeto

### Pré-requisitos
- [Go](https://go.dev/) 1.22+ instalado.
- [Wails](https://wails.io/) instalado (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`).
- [Node.js](https://nodejs.org/) (Necessário para build do frontend).

### Passo a Passo

1. **Clone o repositório:**
   ```bash
   git clone [https://github.com/dvsalvaya/Sistema-de-RPG.git](https://github.com/dvsalvaya/Sistema-de-RPG.git)
   cd Sistema-de-RPG
