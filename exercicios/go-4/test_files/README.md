# Superstars do Github

Usando APIs externas, testes, entre outros, você deverá descobrir quais são os projetos usando Go que possuem mais estrelas no Github.

## Tópicos

Neste desafio você aprenderá:

- Go
- Acessar APIs externas
- Testar APIs
- Gerar JSON
- Escrever em arquivos

## Requisitos
​
Para este desafio você precisará de:

- Go versão 1.9 (ou superior)
- Git


## Detalhes

Crie um aplicativo de linha de comando que use a API do Github para listar os 10 repositórios Go no Github que possuem mais estrelas. Gere um arquivo chamado *stars.json* com o conteúdo, conforme o exemplo abaixo:


``` json
[
    {
        "name": "moby/moby",
        "description": "Moby Project",
        "url": "https://github.com/moby/moby",
        "stars":49409
    },
    {
        "name": "golang/go",
        "description": "The Go programming language",
        "url": "https://github.com/golang/go",
        "stars":43563
    }
]
```