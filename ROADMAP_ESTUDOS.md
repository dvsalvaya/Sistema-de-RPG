# 📚 Roadmap de Estudos: Construindo o Sistema VTT com Django

Este roadmap foi desenhado especialmente para **estudantes de Ciência da Computação** com conhecimentos básicos em Python. Como criar uma aplicação web interativa em tempo real exige dominar alguns novos paradigmas (como o fluxo Request/Response e WebSockets), este guia dividirá o aprendizado desde o alicerce até as funcionalidades avançadas.

---

## Nível 1: Fundações Fortes (Semana 1)
Antes de pularmos para o Django, precisamos garantir que o uso de Python no contexto de backends seja fluído.

### 1.1 Python e Orientação a Objetos (Revisão)
Como o Django é fortemente baseado em Classes (Models, Class Based Views), é crucial dominar:
- Herança clássica e múltipla (`super()`).
- Magic methods comuns (`__str__`, `__init__`).
- Decoradores (`@property`, `@classmethod`).

### 1.2 Ambientes Virtuais e Pip
Em projetos web, isolar dependências é vital.
- Entenda o que é e como usar o comando `python -m venv venv`.
- Como gerar e utilizar o arquivo `requirements.txt` (`pip freeze > requirements.txt`).

---

## Nível 2: O Coração do Django (Semana 2 e 3)
Agora entraremos no mundo Web padrão. O Django segue o padrão arquitetural **MVT (Model-View-Template)**.

### 2.1 Modelos e ORM (O Banco de Dados Sem SQL)
Como o banco original `fichas.db` usava SQL puro, no Django iremos transpor isso para classes.
- **O que estudar:** Como gerar `Models`, criar campos (CharFields, IntegerFields, ForeignKeys).
- **Relacionamentos:** Entender `ForeignKey` (Um para muitos, ex: Muitas Fichas pertencem a um Usuário) e `ManyToManyField` (ex: Jogadores em uma Campanha).
- **Comandos Chave:** `makemigrations`, `migrate` e a interação no `python manage.py shell`.

### 2.2 Views (Controladores) e URLs
- **O que estudar:** Como vincular uma rota de URL (`urls.py`) a uma função (`views.py`) que responde.
- Estude tanto *Function-Based Views (FBV)* quanto a abstração mais poderosa das *Class-Based Views (CBV)* (ex: ListView, DetailView).

### 2.3 Templates e Frontend Básico
- **O que estudar:** A linguagem de templates do Django (Diferenciar entre estático e dinâmico: `{{ variavel }}`, `{% for item in lista %}`).
- Integrar um framework CSS de sua escolha (Bootstrap ou Tailwind) usando `{% load static %}`.

### 2.4 Autenticação e Django Admin
O Django já nos entrega isso quase pronto. Economizaremos semanas de esforço aproveitando essas ferramentas nativas.
- Usar o portal de administração gratuito do Django para editar Fichas em ambiente de homologação.
- Lidar com Login, Registro, Logout e Restrições de Acesso (`@login_required`).

---

## Nível 3: O Desafio do Tempo Real (Semana 4 e 5)
Um VTT só é divertido quando o mestre rola o dado e o resultado aparece quase que na mesma hora no computador do jogador. Para isso, o fluxo HTTP normal de "pergunta-e-resposta" não é eficiente; precisamos usar **WebSockets**.

### 3.1 Django Channels: A Base do Async
O Django sozinho atende uma requisição e morre. O *Channels* expande o Django para suportar conexões WebSocket que ficam ativas constantemente.
- **O que estudar:** Os conceitos de `Consumers` (o equivalente às Views para os WebSockets) assíncronos em Python (`async def`).

### 3.2 Redis e Camada de Canal (Channel Layer)
Para que uma rolagem de dados que o "Usuário A" fez apareça na tela do "Usuário B", os processos precisam se conversar e compartilhar memória.
- **O que estudar:** O papel do Servidor Redis como um mecanismo de *Pub/Sub* (Publicação e Assinatura) no Channels.
- Como instalar o Redis usando *Docker* (Uma excelente oportunidade para aprender o básico de containers).

### 3.3 A Interface JS no Frontend
WebSockets precisam de uma linguagem no navegador para manter a conexão aberta e injetar o resultado no HTML.
- **O que estudar:** Padrões básicos do JavaScript com a API `new WebSocket()`. Como enviar objetos JSON pelo socket e recebê-los para atualizar o DOM dinamicamente (sem recarregar a tela).

---

## Nível 4: Polimento, Segurança e Deploy (Semana 6)
O projeto finalizado está local. O que falta para a comunidade usá-lo globalmente?

### 4.1 Permissões Avançadas e Validações
- Como impedir que um usuário edite as fichas que não pertencem a ele.
- Proteção contra edições maliciosas e verificação de integridade no Backend.

### 4.2 Arquitetura de Produção (Deploy)
Levar tudo para nuvem envolve a troca das ferramentas de dev.
- Mudar do `SQLite3` para o poderoso `PostgreSQL`.
- Servir arquivos estáticos (CSS/JS/Imagens) com `WhiteNoise`.
- Usar `Gunicorn` acompanhado de um proxy como `Nginx`.
- Plataformas para hospedagem recomendadas para testes: Render, Railway ou Heroku.

---

### Dica de Ouro para Estudantes 🚀
Construam o projeto em **pequenos ciclos viáveis (Metodologia Ágil)**.
Não tente configurar o Banco, o WebSockets e o Frontend de uma única vez. 
1. Faça o seu personagem salvar e exibir numa tela feia e estática (MVC Simples).
2. Adicione Autenticação para prender a ficha a você.
3. Arrude o HTML com Bootstrap para ficar bonito.
4. Adicione Javascript Vanilla + Websocket e, finalmente, rode um dado e veja atualizar.

Erros irão acontecer; usem esses bloqueios para entender profundamente a ferramenta sob a camada de abstração. Bom código!
