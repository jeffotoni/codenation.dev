# codenation.dev

AceleraDev Full Stack - Go, 4 SEMANAS DE GO E DEPOIS REACT.

### Público-alvo
Curso de programação usando a linguagem Go abordando conceitos teóricos e práticos é voltado para estudantes, profissionais ou entusiastas que desejam adquirir conhecimentos em uma das linguagens que mais cresce atualmente afim de melhores colocações no mercado de trabalho ou auto conhecimento.

### Objetivo
O objetivo deste curso é capacitar os profissionais para o desenvolvimento de aplicações em Go. Para isso, o curso abordará os principais conceitos e tecnologias utilizadas no jeito Go de fazer as coisas. Iremos reforçar as boas práticas de programação, falar um pouco de métodos ágeis e abordar em um contexto geral problemas e como podemos resolvê-los usando Go.
Após a conclusão do curso “Programação Go”, o aluno estará apto a desenvolver aplicações na linguagem Go, com condição de continuar e explorar ainda mais os aspectos da linguagem.

### Linguagem Go
A linguagem Go foi lançada em 2009 com propósito de facilitar a resolução de problemas quando o assunto é desenvolvimento em camadas de rede, escalabilidade, desempenho, produtividade e o mais importante concorrência. Go teve influências de diversas linguagens de programação e paradigmas diferentes, e se sobressaiu ao juntar o que tinham de melhor e criou algo novo e enxuto, com o mínimo necessário para resolver os problemas propostos, acredito que isto podemos chamar de inovação. 


MÓDULO 1 - Introdução da Linguagem
## MÓDULO 01 Install and Commands Golang

- [Overview](#overview)
- [Introdução a Instalação](#Introdução-Instalação)
  - [Instalação](#Instalação)
    - [Linux](#linux)
    - [$GOPATH](#gopath)
    - [Test your installation](#test-your-installation)
    - [Workspace](#workspace)
    - [Outside GOPATH](#outside-gopath)
- [Instalação Docker](#Instalação-docker)
	- [Instalar Docker to Golang](#instalar-docker-to-golang)
	- [Compile your app Inside the Docker Container](#compile-your-app-inside-the-docker-container)
	- [Cross-compile Your app Inside the Docker Container](#cross-compile-your-app-inside-the-docker-container)
- [Introduction Golang](#introduction-golang)
  - [Golang Language](#golang-language)
    - [Keywords](#keywords)
    - [Operators and Punctuation](#operators-and-punctuation)
    - [Println Print](#println-print)
    - [Bufio NewWriter](#bufio-newWriter)
    - [Func Main](#func-main)
- [Go Commands](#go-commands)
   - [Go Commands Introduction](#go-commands-introduction)
   - [Go Run](#go-run) 
   - [Go Build](#go-build)
   - [Go Install](#go-install)
   - [Go Get](#go-get)
   - [Go Mod](#go-mod)
   - [Go Mod Init](#go-mod-init)
   - [Go Mod Vendor](#go-mod-vendor)
   - [GO111MODULE](#go111module)
   - [Go Test](#go-test)
  


### Overview

Existem inúmeras linguagens de programação e cada uma nasceu com um propósito: “resolver problemas”. As linguagens são ferramentas onde teremos que saber utilizá-las no momento certo. Falando “como desenvolvedor” quanto mais poliglota conseguir ser melhor será para sua carreira profissional e para um melhor entendimento e compreensão da diversidade deste ecossistema.

O objetivo deste curso é apresentar o que é a linguagem Go e, porque ela é tão poderosa. Apresentando alguns conceitos e pontos importantes sobre a linguagem Go.

A equipe de desenvolvedores do Go dizem que sua criação é uma tentativa de tornar os programadores mais produtivos. Melhorando o processo de desenvolvimento de software no Google. O primeiro objetivo foi criar uma linguagem melhor para enfrentar os desafios da simultaneidade escalável, ou seja software que lida com muitas preocupações simultaneamente, um exemplo seria a coordenação de mil servidores back-end enviando tráfego de rede todo o tempo.

Manter a linguagem Go pequena permite objetivos mais importantes. Ser pequeno torna o Go mais fácil de aprender, mais fácil de entender, mais fácil de implementar, mais fácil de reimplementar, mais fácil de depurar, mais fácil de ajustar e mais fácil de evoluir. Fazer menos permite mais. É uma expressão utilizada pela equipe de desenvolvimento do Go: “Do Less. Enable More” seria o equilíbrio entre os universo de problemas existentes e o que Go poderá ajudar a resolver bem tais problemas. Go explicitamente não foi projetado para resolver todos os problemas em vez disso fizeram o suficiente para possamos criar as próprias soluções personalizadas com facilidade, mas deixando bem claro que Go não pode fazer tudo.

Não temos dúvidas que desempenho ou simultaneidade são características importantes e muito relevante na linguagem Go mas Simplicidade, Legibilidade e Produtividade tem um peso muito grande para que tudo ocorra de forma equilibrada e saudável.

 - Existem idiomas que são um pouco mais rápidos que o Go, mas certamente não são tão simples quanto o Go. 
 - Existem linguagens que tornam a concorrência seu objetivo mais elevado, mas não são tão legíveis nem produtivas.
 - Desempenho e simultaneidade são atributos importantes, mas não tão importantes quanto simplicidade, legibilidade e produtividade.
 
 A **Simplicidade** é pré-requisito para confiabilidade, escrevemos códigos para máquinas certo? Não necessariamente, escrevemos código para outros humanos conseguirem manter e da continuidade o que iniciamos a máquina somente executa podemos dizer que  algo secundário uma conseguencia.
 
 A **Legibilidade** é essencial para a manutenção, e o software não puder ser mantido, ele será reescrito; e essa pode ser a última vez que sua empresa investirá em Go. Se você está escrevendo um programa para si mesmo, talvez ele tenha que ser executado apenas uma vez, ou você seja a única pessoa que já o verá, então faça o que já funciona para você. Mas se esse for um software com o qual mais de uma pessoa contribuirá ou que será usado por pessoas durante um período de tempo suficiente que os requisitos, recursos ou o ambiente é executado em alterações, seu objetivo deve ser para o seu programa para ser sustentável .

A **Produtividade** depende diretamente do design, que é a arte de organizar o código para funcionar hoje e ser mutável para sempre. A piada diz que o Go foi projetado enquanto aguardava a compilação de um programa em C++. Compilação rápida é uma característica fundamental do Go e uma ferramenta chave de recrutamento para atrair novos desenvolvedores. 

Go foi projetado pelo Google em 2007 para melhorar a produtividade de programação em uma era de multicore, rede de máquinas e grandes bases de código. Os designers queriam abordar críticas de outras línguas em uso no Google, mas manter suas características úteis.

Os criadores Rob Pike, Ken Thompson e Robert Griesemer mantiveram a sintaxe de Go semelhante ao C. No final de 2008 Russ Cox juntou-se a equipe e ajudou a mudar a linguagem e as bibliotecas de protótipo para realidade.

**A linguagem Go** foi lançada em 2009 com propósito de facilitar a resolução de problemas quando o assunto é desenvolvimento em camadas de rede, escalabilidade, desempenho, produtividade e o mais importante concorrência. O próprio Rob Pike declarou que “Go foi projetado para tratar de um conjunto de problemas de engenharia de software a que estávamos expostos na construção de grandes softwares de servidor”.

Go teve influências de diversas linguagens de programação e paradigmas diferentes dentre elas: Alef, APL, BCPL, C, CSP, Limbo, Modula, Newsqueak, Oberon, occam, Pascal, Smalltalk e Cristal, percebe-se que utilizaram do que tinham de melhor e criou algo novo e enxuto, com o mínimo necessário para resolver os problemas propostos, sem perder sua simplicidade. Acredito que isto podemos chamar de inovação. Go inovou ao quebrar os paradigmas de linguagens e implementar algo novo de forma simples e muito poderosa.
 
### Introdução Instalação

Em golang a instalação é muito simples e prática, para Linux, Mac e Windows.

Basta copiar os arquivos para o diretório correto para cada sistema operacional e exportar os caminhos para o ambiente e solicitar, golang está instalado.

Vamos dar uma olhada em como fazemos isso.

### Instalação
---

Vamos baixar o arquivo, descompactá-lo e instalá-lo em/usr/local/go, se tivermos golang já instalado na máquina teremos que remover o existente para deixar nossa instalação como única.
Vamos criar nosso diretório em nosso espaço de trabalho e testar para ver se tudo correu bem

### Linux

```bash
$ sudo rm -rf/usr/local/go
$ wget https://dl.google.com/go/go1.12.4.linux-amd64.tar.gz
$ sudo tar C/usr/local -xzf vai $ VERSION. $ OS- $ ARCH.tar.gz
```

### $GOPATH

$GOPATH é o golang em seu $HOME, isso é necessário para que seus projetos usem o pkg e construam corretamente. Isso era obrigatório para todas as versões anteriores à versão 1.11. O legal é que a partir de agora não teremos que criar projetos no $GOPATH, podemos criar em qualquer outro diretório que não esteja no $GOPATH.

Este é o link para a proposta de versão [Proposta: Módulos Go Versionados] (https://go.googlesource.com/proposal/+/master/design/24301-versioned-go.md/) ou [Go 1.11 Modules] ( https://github.com/golang/go/wiki/Modules/)

Vamos detalhar como trabalhar com **go mod**, foi uma das melhores experiências que tive para projetos de versionamento usando Golang.

Vamos configurar nosso ambiente para rodar o Go. Adicione **/usr/local/go/bin** à variável de ambiente PATH. Você pode fazer isso adicionando esta linha ao seu **/etc/profile** (para uma instalação em todo o sistema) ou **$HOME/.profile**.

```bash
$ export PATH = $PATH:/usr/local/go/bin
```

** Nota**: as alterações feitas em um arquivo de perfil podem não se aplicar até a próxima vez que você fizer login no seu computador. Para aplicar as alterações imediatamente, basta executar os comandos do shell diretamente ou executá-los a partir do perfil usando um comando como o source $HOME/.profile.

```bash
$ echo "exportar GOPATH = $HOME/go" >> $HOME/.profile
$ echo "export PATH = $ CAMINHO:/usr/local/go/bin" >> $HOME/.profile
$ echo "export PATH = $PATH: $GOPATH/bin" >> $HOME/.profile

or

$ echo "exportar GOPATH = $HOME/go" >> $HOME/.zshrc
$ echo "export PATH = $ CAMINHO:/usr/local/go/bin" >> $HOME/.zshrc
$ echo "export PATH = $PATH: $GOPATH/bin" >> $HOME/.zshrc
```

### Teste nossa instalação

Vamos executar a versão para ver se tudo está correto.

```bash
$ go version
go version go1.12.4 linux/amd64
```

Verifique se o Go está instalado corretamente configurando um espaço de trabalho e construindo um programa simples, da seguinte maneira.

Crie o seu **espaço de trabalho** diretório, $HOME/go. (Se você quiser usar um diretório diferente, precisará definir a variável de ambiente $GOPATH.)

Em seguida, faça o diretório src/hello dentro de sua área de trabalho e, nesse diretório, crie um arquivo chamado hello.go que se pareça com:

### Área de trabalho

O espaço de trabalho é o nosso local de trabalho, onde organizaremos nossos diretórios com nossos projetos. Como mostrado acima, até **Go versão 1.12** fomos forçados a fazer tudo no espaço de trabalho. $GOPATH abaixo do Projeto.

**Exemplo Hello**
```bash
$ export GOPATH=$HOME/go
$ mkdir $HOME/go
$ mkdir $HOME/go/src
$ mkdir $HOME/go/src/hello
$ vim $HOME/go/src/hello/hello.go
```

```bash
$GOPATH/
  |-src
    |-hello
      |-hello.go
```

**Example Project**
```bash
$ export GOPATH=$HOME/go
$ mkdir $HOME/go/src/project1
$ mkdir $HOME/go/src/project1/my-pkg
$ mkdir $HOME/go/src/project1/my-cmd
$ mkdir $HOME/go/src/project1/my-vendor
$ mkdir $HOME/go/src/project1/my-logs
$ mkdir $HOME/go/src/project1/my-models
$ mkdir $HOME/go/src/project1/my-repo
$ mkdir $HOME/go/src/project1/my-handler
```

```bash
$GOPATH/
  |-src
    |-github.com/user/project1/
        |-cmd (of project1)
          |-main.go
        |-vendor
        |-logs
        |-models
        |-repo
        |-handler
    |-github.com/user/project2/
      ....
      ....
```

```bash
$GOPATH/
  | -src
    | -github.com/user/project1/
        | -cmd (do projeto1)
          | -main.go
        | -vendor
        | -logs
        | -models
        | -repo
        | manipulador
    | -github.com/user/project2/
      ....
      ....
```

A variável de ambiente $GOPATH informa a ferramenta Go onde sua área de trabalho está localizada.

```ir
$ go get github.com/user/project1
```

O comando **go get** busca repositórios de origem da Internet e os coloca em sua área de trabalho.
Os caminhos do pacote são importantes para a ferramenta Ir. Usar "github.com/..." significa que a ferramenta sabe como buscar seu repositório.

No cenário acima, tudo teria que ficar em nosso **$GOPATH** para que nossos projetos funcionassem corretamente.

### Fora de $GOPATH

Agora podemos fazer nossos projetos sem estar em $GOPATH, podemos, por exemplo, fazê-lo em qualquer diretório.

**Projeto fora do GOPATH**

```bash
$ export GOPATH=$HOME/go
$ mkdir $HOME/2019/project1
$ mkdir $HOME/2019/project1/my-pkg
$ mkdir $HOME/2019/project1/my-cmd
$ mkdir $HOME/2019/project1/my-logs
$ mkdir $HOME/2019/project1/my-models
$ mkdir $HOME/2019/project1/my-repo
$ mkdir $HOME/2019/project1/my-handler
```
```bash
$HOME/
  |-2019
    |-github.com/user/project1/
      |-cmd
        |-main.go
      |-vendor
      |-logs
      |-models
      |-repo
      |-handler
```

Nós podemos colocar o nosso projeto em qualquer diretório agora.

```bash
$HOME/
  |-any-directory
    |-github.com/user/project1/
      |-cmd
        |-main.go
      |-vendor
      |-logs
      |-models
      |-repo
      |-handler
```

Para o cenário acima, teremos que usar **go mod** em nosso projeto para que todos os pacotes externos possam funcionar corretamente, assim poderemos gerenciá-los corretamente e versão.
Mais informações podem ser encontradas aqui: [Wiki Go Modules] (https://github.com/golang/go/wiki/Modules)

Exemplo prático de como você irá proceder:
```go
$ go mod init github.com/user/project1
```

**Note**: 
Quando usarmos o go mod em $GOPATH, teremos que habilitar o uso de GO111MODULE=on, para que ele possa trabalhar dentro da estrutura $GOPATH.

Então, nosso programa pode compilar com sucesso.
```go
$ GO111MODULE=on go run cmd/main.go
$ GO111MODULE=on go build -o project1 cmd/main.go
```

**Projeto fora do GOPATH e Local sem github**

Vamos criar nosso projeto sem github, somente local com nossos diretórios

```bash
$HOME/
  |-codenation
    |-project1
      |-main.go
      |-go.mod
      |-pkg
         |-math
	    - go.mod
	 |-util
	    - go.mod
      |-vendor
      |-logs
      |-models
      |-repo
      |-handler
```
 
Percebe-se que temos go.mod agora no raiz e em cada pkg que desejamos configurar para que possamos importa-los.

O arquivo go.mod na raiz ficaria assim:
```go
module project1

require (
	pkg/math v0.0.0
	pkg/util v0.0.0
)

replace (
	pkg/math => ./pkg/math
	pkg/util => ./pkg/util
)

go 1.12
```

Os arquivos que encontra-se dentro de cada pacote ficaria assim:
pkg/math
```go
module math
```

pkg/util
```go
module util
```
Prontinho agora é testar
```go
GO111MODULE=on go run main.go
```

```go
GO111MODULE=on go build main.go
```

### Instalação com Docker

Se não quisermos instalar diretamente em nosso sistema operacional golang, podemos instalá-lo em um contêiner docker.

Podemos carregar um contêiner docker com o idioma instalado e compilar e executar nossos programas a partir desse contêiner.

Vamos verificar como podemos fazer isso abaixo.

Mais informações e detalhes você pode visitar este link: [hub.docker] (https://hub.docker.com/_/golang)

### Instalar docker para Go

```bash
$ docker pull golang
```

### Compile seu aplicativo dentro do contêiner Docker

Pode haver ocasiões em que não é apropriado executar seu aplicativo em um contêiner. Para compilar, mas não executar seu aplicativo dentro da Instância do Docker, você pode escrever algo como:
```bash
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.12.4 go build -v
```

Isso adicionará seu diretório atual como um volume ao contêiner, configurará o diretório de trabalho para o volume e executará o comando go build, que informará para compilar o projeto no diretório de trabalho e exibir o executável em myapp. Como alternativa, se você tiver um Makefile, poderá executar o comando make dentro do contêiner.
```bash
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.12.4 make
```

### Cross-compile Seu aplicativo dentro do contêiner Docker
Se você precisar compilar seu aplicativo para uma plataforma diferente de linux/amd64 (como windows/386):

```bash
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=windows \
-e GOARCH=386 golang:1.12.4 go build -v
```

### Exemplo main.go

Vamos fazer nosso programa de testes, vamos chamar isso de main.go

```go
package main

import "fmt"

func main(){
	fmt.Println("My first program being compiled by a docker container!")
}
```

Agora vamos rodar um programa para ver se funciona corretamente

```bash
$ docker run --rm -v "$PWD":/usr/src/main -w /usr/src/main golang:1.12.4 go run main.go
```

Output:
```bash
My first program being compiled by a docker container!
```

Check a version:
```bash
$ docker run --rm -v "$PWD":/usr/src/main -w /usr/src/main golang:1.12.4 go versio
```

Output:
```bash
go version go1.12.4 linux/amd64
```


