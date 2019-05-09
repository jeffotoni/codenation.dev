# Parser de HTML

Crie um parser para extrair dados da página da linguagem Go no Github. Para isso, fará leitura de arquivos, parse de conteúdo HTML, escrita em arquivos e geração de JSON.

## Tópicos

Neste desafio você aprenderá:

- Go
- Ler arquivos
- Fazer parse de conteúdo HTML
- Gerar JSON
- Escrever em arquivos

## Requisitos
​
Para este desafio você precisará de:

- Go versão 1.9 (ou superior)
- Git


## Detalhes

Crie um aplicativo de linha de comando que processe o arquivo *golang.html* e gere um arquivo chamado *golang.json* com o conteúdo, conforme o exemplo abaixo:

``` json
{
    "name": "Go",
    "description": "Go is a programming language built to resemble a simplified version of the C programming language",
    "created_by": "Robert Griesemer, Rob Pike, Ken Thompson",
    "released_at":"November 10, 2009",
    "repositories":21022,
    "related_topics": ["docker", "bot", "godoc", "database", "language", "python", "marbles", "heroku"]
}
```