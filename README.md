# 🎲 Plataforma VTT RPG Web

> Uma plataforma colaborativa e em tempo real para gerenciamento de mesas de RPG, conectando Mestres e Jogadores de forma rápida e intuitiva.

![Python Badge](https://img.shields.io/badge/Backend-Python-blue)
![Django Badge](https://img.shields.io/badge/Framework-Django-darkgreen)
![Status Badge](https://img.shields.io/badge/Status-Em_Desenvolvimento-yellow)
![License Badge](https://img.shields.io/badge/License-Open_Source-green)

## 📖 Sobre o Projeto

Este projeto consiste na evolução de um conceito inicial de Virtual Tabletop (VTT) focado para desktop, migrando agora para uma robusta **Aplicação Web**. Nosso objetivo é criar uma experiência conectada e em tempo real para grupos de RPG, permitindo que a gestão de personagens, campanhas e rolagens de dados seja acessível por qualquer navegador, sem perder a velocidade e praticidade.

A nova arquitetura é totalmente baseada em **Python e Django**, desenhada para ser modular, escalável e de fácil contribuição para a comunidade open source.

## 🎯 Motivação e Objetivos

O principal objetivo é oferecer uma alternativa leve aos grandes VTTs do mercado, mantendo a simplicidade sem sacrificar ferramentas essenciais. A escolha do **Django** permite que desenvolvedores possam rapidamente expandir a plataforma através de padrões estabelecidos (MTV/MVC), facilitando a colaboração de estudantes e profissionais.

Queremos construir um sistema genérico que sirva de base sólida para qualquer sistema de RPG futuro, começando com estruturas flexíveis de atributos e inventário.

## ✨ Funcionalidades Principais (Em Desenvolvimento)

### 🛡️ Autenticação e Gestão de Mesas
- **Perfis Multi-usuário:** Diferenciação entre "Mestres de Mesa" e "Jogadores".
- **Sistema de Campanhas:** Mestres podem criar salas online e convidar jogadores, gerenciando múltiplos grupos em um só lugar.

### 📝 Fichas de Personagem Dinâmicas e Genéricas
- **Moldes Modulares:** Atributos flexíveis para acomodar sistemas genéricos.
- **Acesso Contínuo:** Todas as fichas são armazenadas de forma segura no banco de dados, acessíveis de qualquer dispositivo.

### 🎲 Interação em Tempo Real (WebSockets)
- **Rolador de Dados Online:** Suporte completo aos dados clássicos (d4, d6, d8, d10, d12, d20, d100).
- **Log de Sessão e Chat:** Todas as rolagens e ações da equipe são transmitidas simultaneamente para a mesa usando *Django Channels*.

### 🧭 Painel do Mestre (Ferramentas Avançadas)
- **Gestão de NPC e Combate:** Acesso a fichas simplificadas para controle de inimigos.
- **Anotações Privadas:** Controle total sobre o que os jogadores podem enxergar.

---

## 🛠️ Tecnologias Utilizadas

A stack foi escolhida pela sua maturidade, capacidade de expansão e para fins de aprendizado aprofundado:

### Backend
- **Linguagem:** [Python](https://www.python.org/) 3.10+
- **Framework Opcional/Principal:** [Django](https://www.djangoproject.com/)
- **Assincronicidade e Tempo Real:** Django Channels + Redis
- **Banco de Dados:** SQLite (Desenvolvimento) / PostgreSQL (Produção)

### Frontend
- **Interface:** HTML5, CSS3, JavaScript (Opcionalmente HTMX ou Bootstrap/Tailwind para estilização rápida - *a definir*).

---

## 🚀 Como Rodar o Projeto (Desenvolvimento)

### Pré-requisitos
- [Python 3.10+](https://www.python.org/downloads/) instalado.
- [Git](https://git-scm.com/) instalado.
- (Recomendado) Conhecimento básico de Ambientes Virtuais (`venv` ou `pipenv`).

### Passo a Passo

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/SEU_USUARIO/Sistema-de-RPG.git
   cd Sistema-de-RPG
   ```

2. **Crie e ative um ambiente virtual:**
   ```bash
   # Windows
   python -m venv venv
   venv\Scripts\activate

   # Linux/macOS
   python3 -m venv venv
   source venv/bin/activate
   ```

3. **Instale as dependências:**
   *(O arquivo `requirements.txt` será fornecido em breve)*
   ```bash
   pip install -r requirements.txt
   ```

4. **Realize as as migrações do banco de dados:**
   ```bash
   python manage.py migrate
   ```

5. **Execute o servidor de desenvolvimento:**
   ```bash
   python manage.py runserver
   ```
   *O projeto estará disponível em `http://127.0.0.1:8000/`*

---

## 🗺️ Roadmap de Desenvolvimento

- [ ] **Fase 1: Configuração Inicial** - Setup do Django, modelos base de Usuários e Perfis.
- [ ] **Fase 2: Gestão Básica** - CRUD de Fichas Genéricas e criação de Campanhas/Salas.
- [ ] **Fase 3: O Frontend** - Protótipo das telas principais e estilização básica.
- [ ] **Fase 4: Tempo Real** - Implementação do Django Channels e Redis para rolagens de dados visíveis na rede.
- [ ] **Fase 5: Mestre da Mesa** - Painel exclusivo e controle de Sessões.
- [ ] **Fase 6: Deploy** - Configuração final e hospedagem na nuvem.

Consulte também o arquivo [ROADMAP_ESTUDOS.md](./ROADMAP_ESTUDOS.md) se você está aprendendo a stack do projeto e deseja entender como as tecnologias implementadas aqui funcionam.

---

## 🤝 Como Contribuir

Contribuições de todos os tamanhos são muito bem-vindas! Como o projeto está no seu início, buscamos ajuda com:

- Definição da arquitetura e design das models estruturais.
- Planejamento do design e componentes visuais de Frontend.
- Revisão de código e testes lógicos focados em balanceamento de sistemas RPG.

*Por favor, abra uma _Issue_ discutindo o que você gostaria de modificar antes de enviar um _Pull Request_.*
