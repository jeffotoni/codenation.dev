# Listagem de diretórios em formato JSON

Gere um relatório (em formato JSON) com a lista de arquivos de um diretório, recursivamente.

FINAL MODULO 2 aplica este exercicio

## Tópicos

Neste desafio você aprenderá:

- Go
- Ler diretórios do sistema operacional
- Ler parâmetros a partir da linha de comando
- Gerar JSON
- Escrever em arquivos

## Requisitos
​
Para este desafio você precisará de:

- Go versão 1.9 (ou superior)
- Git


## Detalhes

Crie um aplicativo de linha de comando que receba como parâmetro o nome de um diretório e gere um arquivo chamado _files.json_ com o conteúdo, conforme o exemplo abaixo:

    jsonify ./

Conteúdo do _files.json_:

``` json
[
    {
        "name": "README.md",
        "path": "./README.md"
    },
    {
        "name": "files.json",
        "path": "./files.json"
    },
    {
        "name": "main.go",
        "path": "./main.go"
    },
    {
        "name": "main_test.go",
        "path": "./main_test.go"
    }
]
```