# codenation.dev

AceleraDev Full Stack - Go, 4 SEMANAS DE GO E DEPOIS REACT.

### Público-alvo
Curso de programação usando a linguagem Go abordando conceitos teóricos e práticos é voltado para estudantes, profissionais ou entusiastas que desejam adquirir conhecimentos em uma das linguagens que mais cresce atualmente afim de melhores colocações no mercado de trabalho ou auto conhecimento.

### Objetivo
O objetivo deste curso é capacitar os profissionais para o desenvolvimento de aplicações em Go. Para isso, o curso abordará os principais conceitos e tecnologias utilizadas no jeito Go de fazer as coisas. Iremos reforçar as boas práticas de programação, falar um pouco de métodos ágeis e abordar em um contexto geral problemas e como podemos resolvê-los usando Go.
Após a conclusão do curso “Programação Go”, o aluno estará apto a desenvolver aplicações na linguagem Go, com condição de continuar e explorar ainda mais os aspectos da linguagem.

## MÓDULO 01 Introdução a Linguagem

- [Overview](#Overview)
    - [Golang](#Golang)
      - [Linguagem Go](#Linguagem-Go)
      - [Onde posso utilizar Go](#Onde-posso-utilizar-Go)      
      - [O inicio de tudo](#O-inicio-de-tudo)
      - [Por que meu binário hello world é tão grande](#Por-que-meu-binário-hello-world-é-tão-grande)
- [Introdução a instalação](#Introdução-a-instalação)
  - [Instalação](#Instalação)
    - [Linux](#Linux)
    - [$GOPATH](#gopath)
    - [Teste nossa instalação](#Teste-nossa-instalação)
    - [Área de trabalho](#área-de-trabalho)
    - [Fora de $GOPATH](#Fora-de-GOPATH)
- [Instalação com Docker](#Instalação-com-Docker)
	- [Instalar Docker para Go](#Instalar-Docker-para-Go)
	- [Compile seu aplicativo dentro do contêiner Docker](#Compile-seu-aplicativo-dentro-do-contêiner-Docker)
	- [Cross-compile seu aplicativo dentro do contêiner Docker](#Cross-compile-seu-aplicativo-dentro-do-contêiner-Docker)
- [Introdução Golang](#Introdução-Golang)
  - [Linguagem Golang](#Linguagem-Golang)
    - [Keywords](#Keywords)
    - [Operadores e Pontuação](#Operadores-e-Pontuação)
    - [Println Print](#Println-Print)
    - [Bufio NewWriter](#Bufio-NewWriter)
    - [Func Main](#Func-Main)
- [Comandos Go](#Comandos-Go)
   - [Introdução aos comandos Go](#Introdução-aos-comandos-Go)
   - [Go Run](#Go-Run) 
   - [Go Build](#Go-Build)
   - [Instalar Go](#Instalar-Go)
   - [Go Get](#Go-Get)
   - [Go Mod](#Go-Mod)
   - [Go Mod Init](#Go-Mod-Init)
   - [Go Mod Vendor](#Go-Mod-Vendor)
   - [GO-1.11-Módulo](#Go-111-módulo)
   - [Go Test](#Go-Test)

## MÓDULO 02 Aprofundando na Linguagem

- [Tipos](#Tipos)
   - [Numeric Types](#Numeric-Types)
   - [String Types](#String-Types)
   - [Pointer Types](#Pointer-Types)
   - [Array Types](#Array-Types)
   - [Slice Types](#Slice-Types)
   - [Struct Types](#Struct-Types)
   - [Struct In C](#Struct-In-C)
   - [Struct Type Tags Json](#Struct-Type-Tags-Json)
   - [Fatih Structs to Map](#Fatih-Structs-to-Map)
   - [Map Types](#Map-Types)
   - [Map Literals Continued](#Map-Literals-Continued)
   - [Channel Types](#Channel-Types)
   - [Blank Identifier](#Blank-Identifier)
   - [Tipos de interface](#Tipos-de-interface)
	 - [Aqui está uma interface como um método](#Aqui-está-uma-interface-como-um-método)
	 - [Interface como tipo](#Interface-como-tipo)
- [Estruturas de Controle](#Estruturas-de-Controle)
  - [Ao-Controle](#Ao-Controle)
    - [Retorno de Controle](#Retorno-de-Controle)
    - [Controle Goto](#Controle-Goto)
    - [Control if Else](#Control-if-Else)
    - [Control for break continue](#Control-for-break-continue)
    - [Control switch case break](#Control-switch-case-break)
    - [Etiqueta de Controle](#Etiqueta-de-Controle)
    - [Faixa de Controle](#Faixa-de-Controle)
- [Erros](#Erros)
  - [Introdução aos Erros](#Introdução-aos-Erros)
    - [Como funciona o controle de erros](#Como-funciona-o-controle-de-erros)
    - [Novos Erros](#Novos-Erros)
    - [Erros Personalizados](#Erros-Personalizados)    
    - [fmt Errorf](#fmt-Errorf)
- [Funções](#Funções)
  - [Introução as Funções](#Introdução-as-Funções)
    - [Retornar vários valores](#Retornar-vários-valores) 
    - [Funções Variadic](#Funções-Variadic) 
    - [Funções como um parâmetro](#Funções-como-um-parâmetro) 
    - [Fechamentos](#Fechamentos)
    - [Recursão](#Recursão)
    - [Funções Asynchronous](#Funções-Asynchronous)
- [Defer](#Defer)
- [Exercício 1](#Exercício-1)


## Overview

### Golang

#### Linguagem Go

Go é uma linguagem poderosa quando se trata de **concorrência e alto desempenho**, com uma arquitetura limpa e eficiente. Ela cresce ano após ano e todos os dias as comunidades crescem ainda mais.

Alguns paradigmas foram quebrados para torná-lo uma linguagem de **alto desempenho**, onde a concorrência é um dos seus pontos fortes. O Go facilita a criação de programas que aproveitam ao máximo as máquinas multicore e em rede, enquanto o novo sistema de tipos permite que você crie programas flexíveis e modulares.

É uma linguagem rápida e **estaticamente compilada** que se parece com uma linguagem **interpretada dinamicamente**. Este recurso **Go** se torna uma linguagem única como o assunto é web.

Go é uma linguagem de programação compilada, concorrrente, com tipagem forte e estaticamente tipada. É uma linguagem de **"Uso Geral"** que pode ser usada para resolver vários problemas e em diferentes áreas. Problemas envolvendo concorrência, aplicações web, aplicações de alto desempenho, desenvolvimento de APIs, soquetes de comunicação etc ... É onde a linguagem está se tornando cada vez mais proeminente no mercado e nas comunidades.

```bash
 - Compilado
 - Linguagem estática
 - Linguagem estática com tipos dinâmicos (Em Go seria o uso de tipo como interface)
 - Estaticamente tipada (declarado o tipo ele não muda)
 - Tipagem forte (um mesmo dado não é tratado como se fosse de outro tipo)
 - Compilada estaticamente
 - Concorrente
 - Simples
 - Produtivo
 - GC (Garbage Coletor)
 - Runtime (Implementa coleta de lixo, concorrencia, gerenciamento de pilha e outros 
 recursos críticos da linguagem em tempo de execução, algoritimo implementado é Dijkstra)

Obs: É importante entender, no entanto, que o tempo de execução do 
Go não inclui uma máquina virtual. Os programas Go são compilados 
para o código de máquina nativo.
 ```

##### Onde posso utilizar Go:

```bash
. Web backend (com diversos frameworks disponíveis)
. Web Assembly (um dos frameworks vugu)
. Microservices (alguns frameworks: Go Micro, Go Kit, Gizmo, Kite)
. Fragments services (Termo citado pelo @jeffotoni em um grupo de discussão de microservices)
. Lambdas (FaaS example)
. Client Server
. IoT (alguns frameworks)
. Boots (alguns aqui)
. Aplicações Client que usam tecnologia Web 
. Desktop Usando Qt+QML, Lib Nativa Win (example Qt, widgets Qt, Qml)
```

### O inicio de tudo

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


#### Por que meu binário hello world é tão grande?

O vinculador no gc toolchain cria binários vinculados estaticamente por padrão. Portanto, todos os binários Go incluem o tempo de execução Go, juntamente com as informações do tipo em tempo de execução necessárias para oferecer suporte a verificações de tipos dinâmicos, reflexos e até mesmo rastreamentos de pilha em tempo de pânico.

Um simples programa C "hello, world" compilado e linkado estaticamente usando o gcc no Linux é de cerca de 750 kB, incluindo uma implementação do printf. Um programa Go equivalente usando fmt.Printf pesa alguns megabytes, mas isso inclui suporte a tempo de execução mais poderoso e informações de tipo e depuração.

Um programa Go compilado com gc pode ser vinculado ao sinalizador -ldflags=-w para desabilitar a geração de DWARF, removendo as informações de depuração do binário, mas sem nenhuma outra perda de funcionalidade. Isso pode reduzir substancialmente o tamanho binário.

Exemplo:

```bash
$ go build -ldflags=-w -o helo hello.go
```

## Introdução a instalação

Em golang a instalação é muito simples e prática, para Linux, Mac e Windows.

Basta copiar os arquivos para o diretório correto para cada sistema operacional e exportar os caminhos para o ambiente e solicitar, golang está instalado.

Vamos dar uma olhada em como fazemos isso.

### Instalação
---

Vamos baixar o arquivo, descompactá-lo e instalá-lo em/usr/local/go, se tivermos golang já instalado na máquina teremos que remover o existente para deixar nossa instalação como única.
Vamos criar nosso diretório em nosso espaço de trabalho e testar para ver se tudo correu bem

#### Linux

```bash
$ sudo rm -rf/usr/local/go
$ wget https://dl.google.com/go/go1.12.4.linux-amd64.tar.gz
$ sudo tar C/usr/local -xzf vai $ VERSION. $ OS- $ ARCH.tar.gz
```

#### $GOPATH

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

#### Teste nossa instalação

Vamos executar a versão para ver se tudo está correto.

```bash
$ go version
go version go1.12.4 linux/amd64
```

Verifique se o Go está instalado corretamente configurando um espaço de trabalho e construindo um programa simples, da seguinte maneira.

Crie o seu **espaço de trabalho** diretório, $HOME/go. (Se você quiser usar um diretório diferente, precisará definir a variável de ambiente $GOPATH.)

Em seguida, faça o diretório src/hello dentro de sua área de trabalho e, nesse diretório, crie um arquivo chamado hello.go que se pareça com:

#### Área de trabalho

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

#### Fora de $GOPATH

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

## Instalação com Docker

Se não quisermos instalar diretamente em nosso sistema operacional golang, podemos instalá-lo em um contêiner docker.

Podemos carregar um contêiner docker com o idioma instalado e compilar e executar nossos programas a partir desse contêiner.

Vamos verificar como podemos fazer isso abaixo.

Mais informações e detalhes você pode visitar este link: [hub.docker] (https://hub.docker.com/_/golang)

### Instalar Docker para Go

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

### Cross-compile seu aplicativo dentro do contêiner Docker
Se você precisar compilar seu aplicativo para uma plataforma diferente de linux/amd64 (como windows/386):

```bash
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=windows \
-e GOARCH=386 golang:1.12.4 go build -v
```

#### Exemplo main.go

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

Saída:

```bash
My first program being compiled by a docker container!
```

Confira uma versão:

```bash
$ docker run --rm -v "$PWD":/usr/src/main -w /usr/src/main golang:1.12.4 go versio
```

Saída:

```bash
go version go1.12.4 linux/amd64
```


## Introdução Golang
---

Go é uma linguagem de propósito geral, projetada com a programação de sistemas em mente. É fortemente tipado e colecionador de lixo, e tem suporte explícito para programação concorrente. 
Os programas são construídos a partir de pacotes, cujas propriedades permitem o gerenciamento eficiente de dependências.

A gramática é compacta e regular, permitindo fácil análise por ferramentas automáticas, como ambientes de desenvolvimento integrados.

### Linguagem Golang
---

#### Keywords

As seguintes palavras-chave são reservadas e não podem ser usadas como identificadores:

```bash
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

#### Operadores e Pontuação

As seqüências de caracteres a seguir representam operadores (incluindo operadores de atribuição) e pontuação: 

```bash
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
```

#### Println Print

Vamos aprender como enviar dados para a tela que é, na verdade, **stdout** saída padrão, veremos mais detalhes sobre  **stdout** e **stdin**.

Vamos conhecer **print, println and fmt.Println**

As implementações atuais fornecem várias funções internas úteis durante o bootstrapping. Essas funções são documentadas para integridade, mas não garantem a permanência na linguagem. Eles não retornam um resultado. 

Restrição de implementação: **print** e **println** não precisa aceitar tipos arbitrários de argumentos, mas a impressão de tipos boolean, numeric, e string types deve ser suportada. 

**println is an built-in function** (into the runtime) que pode eventualmente ser removido, enquanto o **fmt package** está na biblioteca padrão, que persistirá.

```bash
Function   Behavior

print      prints all arguments; formatting of arguments is implementation-specific
println    like print but prints spaces between arguments and a newline at the end
```

Usando Print:

```go
// test print
package main

func main() {
   print("debugging my system with print")
}
```

Saída:

```bash
debugging my system with print
```

Usando println:

```go
// test println
package main

func main() {
   println("debugging my system with println")
}
```

Saída:

```bash
debugging my system with println
```

Usando fmt.Println:

```go
package main

import "fmt"

func main() {
   fmt.Println("debugging my system with fmt.Println")
}
```

Saída:

```bash
debugging my system with fmt.Println
```

O objetivo de iniciar e executar o comando print, println ou fmt.Println é nos ajudar com os testes que faremos a partir de agora em todas as etapas do nosso aprendizado Go.


#### Bufio NewWriter

```bash
bufio.Writer
```

Fazer muitas gravações pequenas pode prejudicar o desempenho. Cada gravação é, em última instância, um syscall e se fazer com freqüência pode sobrecarregar a CPU. Dispositivos como discos funcionam melhor lidando com dados alinhados ao bloco. Para evitar a sobrecarga de muitas pequenas operações de gravação, o Golang é fornecido com o bufio.Writer. Os dados, em vez de ir diretamente para o destino (implementando a interface io.Writer) são acumulados primeiro dentro do buffer e enviados quando o buffer está cheio:

Vamos visualizar como o buffering funciona com nove gravações (um caractere cada) quando o buffer tem espaço para quatro caracteres:

```bash
producer         buffer           destination (io.Writer)
 
   a    ----->   a
   b    ----->   ab
   c    ----->   abc
   d    ----->   abcd
   e    ----->   e      ------>   abcd
   f    ----->   ef               abcd
   g    ----->   efg              abcd
   h    ----->   efgh             abcd
   i    ----->   i      ------>   abcdefgh
```

Confira o exemplo abaixo:

```go
package main

import (
	"bufio"
	"os"
)

// creating the write object pointer
// so that we can receive value in every
// scope of our program
var writer *bufio.Writer

func main() {
	// All screen output will be redirected
	// to bufio.NewWriter
	writer = bufio.NewWriter(os.Stdout)
	s := "How many stars does Orion have?\n"
	var b byte = 'H'

	writer.WriteString(s)
	writer.WriteByte(b)
	writer.WriteString("\n")

	// when all the functions finishes it closes
	// the buffer and sends to the.Stdout
	defer writer.Flush()
}
```

Saída:

```bash
How many stars does Orion have?
H
```

#### Func Main

```go
package main

import "fmt"

func main() {
  fmt.Printf("hello, Gophers\n")
}
```

Então **build** com o **go tool**: 

```go
$ cd $HOME/go/src/hello
$ go build
```

Ou podemos compilar assim:

```go
$ cd $HOME/go/src/hello
$ go build -o hello hello.go
```

O comando acima irá construir um executável chamado hello no diretório ao lado do seu código-fonte. Execute para ver a saudação:

```go
$ ./hello
hello, Gophers
```

Verifique também o comando **run** com o go: 

```go
$ go run hello.go
hello, Gophers
```

Se você ver o **"hello, Gophers"** mensagem, em seguida, sua instalação Go **is working**.

Você pode executar o **go install** para instalar o binário no diretório **bin** da sua área de trabalho ou **ir limpar -i** para removê-lo.

Exemplo: go install

```go
$ pwd
$ $HOME/go/src/hello
$ cd $HOME/go/src/hello
$ go install
$ ls -lhs $HOME/go/bin
-rwxrwxr-x 1 user user 2,9M nov  8 03:11 hello
```

Example: go clean -i

```go
$ go clean -i 
$ ls -lhs $HOME/go/bin
```

## Comandos Go
---

### Introdução aos Comandos Go

Em golang, temos um arsenal para nos ajudar quando se trata de compilar, testar, documentar, gerenciar perfis, etc.

```bash
bug         start a bug report
build       compile packages and dependencies
clean       remove object files and cached files
doc         show documentation for package or symbol
env         print Go environment information
fix         update packages to use new APIs
fmt         gofmt (reformat) package sources
generate    generate Go files by processing source
get         download and install packages and dependencies
install     compile and install packages and dependencies
list        list packages or modules
mod         module maintenance
run         compile and run Go program
test        test packages
tool        run specified go tool
version     print Go version
vet         report likely mistakes in packages
```

Use "go help" para mais informações sobre um comando.


### Go Run
---

Uso:

```bash
go run [build flags] [-exec xprog] package [arguments...]
```

Executar compila e executa o pacote principal chamado Go. Normalmente, o pacote é especificado como uma lista de arquivos de origem .go, mas também pode ser um caminho de importação, um caminho do sistema de arquivos ou um padrão que corresponda a um único pacote conhecido, como em 'go run .' ou 'go run my/cmd'.

Por padrão, 'go run' executa o binário compilado diretamente: 'a.out arguments...'. se o sinalizador -exec for fornecido, 'go run' invoca o binário usando xprog:

Se o sinalizador -exec não for fornecido, GOOS ou GOARCH será diferente do padrão do sistema, e um programa chamado go_ $ GOOS_ $ GOARCH_exec pode ser encontrado no caminho de busca atual, 'go run' invoca o binário usando esse programa, por exemplo, argumentos 'go_nacl_386_exec a.out...'. Isso permite a execução de programas compilados quando um simulador ou outro método de execução estiver disponível.

O status de saída de Executar não é o status de saída do binário compilado.

Para mais informações sobre sinalizadores de compilação, consulte 'go help build'. Para mais informações sobre como especificar pacotes, consulte 'go help packages'.

Veja abaixo um exemplo:

```go
// test println
package main

func main() {
   println("Debugging my system with println")
}
```
Go run:

```bash
go run println.go
```

Saída:

```bash
Debugging my system with println
```

### Go Build
---

Build compila os pacotes nomeados pelos caminhos de importação, junto com suas dependências, mas não instala os resultados.

Ao compilar pacotes, o build ignora os arquivos que terminam em '_test.go'.

O -o flag, permitido somente ao compilar um único pacote, obriga a compilação a gravar o executável ou objeto resultante no arquivo de saída nomeado, em vez do comportamento padrão descrito nos dois últimos parágrafos.

O -i flag instala os pacotes que são dependências do destino.

```go
$ go build [-o output] [-i] [build flags] [packages]
```
Veja um exemplo:

```go
package main

import "fmt"

func main() {
  fmt.Println("Workshop Golang 2019.")
}
```

Saída:

```bash
Workshop Golang 2019.
```

Compilação normal

```go
go build -o hello hello.go
```

Saída:

```bash
$ ls -lh 
-rwxrwxr-x 1 root root **1,9M** jan 18 12:42 hello
-rw-rw-r-- 1 root root   75 jan 17 12:04 hello.go
```

Deixando o arquivo menor após a compilação

```go
go build -ldflags="-s -w" hello.go
```

Saída:

```bash
$ ls -lh 
-rwxrwxr-x 1 root root **1,3M** jan 18 12:42 hello
-rw-rw-r-- 1 root root   75 jan 17 12:04 hello.go
```

### Instalar Go 
---

Instalar pacotes e dependências

Uso:

```bash
$ go install [-i] [build flags] [packages]
```

Instale compila e instala os pacotes nomeados pelos caminhos de importação.

O **-i flag** instala as dependências dos pacotes nomeados também.

Para mais sobre o build flags, confira 'go help build'. Para mais informações sobre como especificar pacotes, confira 'go help packages'.

### Go Get
---

O **'go get'** comando muda o comportamento dependendo se o comando go está sendo executado no modo com reconhecimento de módulo ou no modo GOPATH herdado. Este texto de ajuda, acessível como 'go help module-get' mesmo em legado GOPATH mode, descreve 'go get' como opera em module-aware mode.

Uso:

```bash
$ go get [-d] [-m] [-u] [-v] [-insecure] [build flags] [packages]
```

Obtenha downloads dos pacotes nomeados pelos caminhos de importação, junto com suas dependências. Em seguida, instale os pacotes nomeados, como 'go install'.

Veja as bandeiras aceitas abaixo:

```bash
The -d flag instructs get to stop after downloading the packages; that is, it instructs get not to install the packages.

The -f flag, valid only when -u is set, forces get -u not to verify that each package has been checked out from the source control repository implied by its import path. This can be useful if the source is a local fork of the original.

The -fix flag instructs get to run the fix tool on the downloaded packages before resolving dependencies or building the code.

The -insecure flag permits fetching from repositories and resolving custom domains using insecure schemes such as HTTP. Use with caution.

The -t flag instructs get to also download the packages required to build the tests for the specified packages.

The -u flag instructs get to use the network to update the named packages and their dependencies. By default, get uses the network to check out missing packages but does not use it to look for updates to existing packages.

The -v flag enables verbose progress and debug output.
```

Exemplos:

```bash
$ go get -v github.com/guptarohit/asciigraph
$ go get -u github.com/mxk/go-sqlite
$ go get -v github.com/google/uuid
$ go get -v github.com/sirupsen/logru
```


### Go Mod
---

Um módulo é uma coleção de pacotes Go relacionados. Módulos são a unidade de intercâmbio de código-fonte e controle de versão. O comando go tem suporte direto para trabalhar com módulos, incluindo gravação e resolução de dependências em outros módulos. Os módulos substituem a antiga abordagem baseada em GOPATH para especificar quais arquivos de origem são usados em uma determinada compilação.

Uso:

```bash
$ go mod <command> [arguments]
```
Um módulo é definido por uma árvore de arquivos de origem Go com um **go.mod** arquivo no diretório raiz da árvore. O diretório que contém o arquivo go.mod é chamado de raiz do módulo. Normalmente, a raiz do módulo também corresponderá a uma raiz de repositório de código-fonte (mas, em geral, não precisa). O módulo é o conjunto de todos os pacotes Go na raiz do módulo e em seus subdiretórios, mas excluindo subárvores com seus próprios arquivos go.mod.

O "module path" é o prefixo do caminho de importação correspondente à raiz do módulo. O arquivo go.mod define o caminho do módulo e lista as versões específicas de outros módulos que devem ser usados ao resolver importações durante uma construção, fornecendo seus caminhos e versões de módulo.

Por exemplo, este go.mod declara que o diretório que o contém é a raiz do módulo com o caminho example.com/m, e também declara que o módulo depende de versões específicas de golang.org/x/text and gopkg.in/yaml.v2: 

```bash
$ go mod init github.com/user/gomyproject

require (
  golang.org/x/text v0.3.0
  gopkg.in/yaml.v2 v2.1.0
)
```
O arquivo go.mod também pode especificar substituições e versões excluídas que só se aplicam ao construir o módulo diretamente; eles são ignorados quando o módulo é incorporado em uma construção maior. Para mais informações sobre o arquivo go.mod, consulte 'go help go.mod'.
Para iniciar um novo módulo, basta criar um arquivo go.mod na raiz da árvore de diretórios do módulo, contendo apenas uma instrução do módulo. O comando 'go mod init' pode ser usado para fazer isso:

```bash
$ go mod init github.com/user/gomyproject
```
Em um projeto que já utiliza uma ferramenta de gerenciamento de dependências existente como **godep, glide ou dep, 'go mod init'** também adicionará instruções requeridas que correspondam à configuração existente.

Uma vez que o arquivo go.mod existe, nenhuma etapa adicional é necessária: vá com comandos como **'go build'**, **'go test' **, ou até mesmo **'go list'** adicionará automaticamente novas dependências conforme necessário para satisfazer as importações.

Os comandos são: 

```bash
download    download modules to local cache
edit        edit go.mod from tools or scripts
graph       print module requirement graph
init        initialize new module in current directory
tidy        add missing and remove unused modules
vendor      make vendored copy of dependencies
verify      verify dependencies have expected content
why         explain why packages or modules are needed
```
Use "go help mod <command>" para mais informações sobre um comando.


### Go Mod Init

Inicializar novo módulo no diretório atual

Uso:

```bash
$ go mod init [module]
```

Init inicializa e grava um novo **go.mod** no diretório atual, na verdade, criar um novo módulo com raiz no diretório atual. O arquivo go.mod não deve existir. Se possível, o init irá adivinhar o caminho do módulo a partir dos comentários de importação (consulte 'go help importpath') ou da configuração do controle de versão. Para substituir essa suposição, forneça o caminho do módulo como um argumento.


```bash
$ go mod init github.com/user/gomyproject2

require (
  github.com/dgrijalva/jwt-go v3.2.0+incompatible
  github.com/didip/tollbooth v4.0.0+incompatible
  github.com/go-sql-driver/mysql v1.4.1
  github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
  golang.org/x/crypto v0.0.0-20190103213133-ff983b9c42bc
  golang.org/x/time v0.0.0-20181108054448-85acf8d2951c // indirect
)
```

### Go Mod Vendor

O comando go mod mod vendor baixará todas as dependências para o diretório "vendor".
Ao usar o go mod mod, os pacotes não estão no seu diretório.

```bash
$ cd gomyproject2
$ go mod vendor
```

Saída:

```bash
$ ls -lh vendor
total 8,0K
drwxrwxr-x 3 root root 4,0K jan 27 01:47 github.com
-rw-rw-r-- 1 root root  137 jan 27 01:47 modules.txt
```

### GO 1.11 Módulo

Go 1.11 inclui suporte preliminar para os módulos Go, incluindo um novo comando 'go get' que reconhece os módulos. Nós pretendemos continuar revisando este suporte, preservando a compatibilidade, até que ele possa ser declarado oficial (não mais preliminar), e então, em um ponto posterior, podemos remover o suporte para o trabalho no GOPATH e o antigo comando 'go get'.

A maneira mais rápida de aproveitar o novo suporte ao módulo Go 1.11 é verificar seu repositório em um diretório fora de GOPATH / src, criar um arquivo go.mod (descrito na próxima seção) e executar comandos go a partir desse arquivo árvore.

Para um controle mais refinado, o suporte a módulos no Go 1.11 respeita uma variável de ambiente temporária, GO111MODULE, que pode ser definida como um dos três valores de string: off, on ou auto (o padrão). Se GO111MODULE = off, então o comando go nunca usa o novo suporte ao módulo. Em vez disso, ele procura nos diretórios de fornecedores e no GOPATH para localizar dependências; agora nos referimos a isso como "modo GOPATH". Se GO111MODULE = on, então o comando go requer o uso de módulos, nunca consultando o GOPATH. Nós nos referimos a isso como o comando sendo ciente do módulo ou em execução no "modo de reconhecimento de módulo". Se GO111MODULE = auto ou não estiver definido, o comando go ativa ou desativa o suporte a módulos com base no diretório atual. O suporte a módulos é ativado somente quando o diretório atual está fora de GOPATH / src e ele contém um arquivo go.mod ou está abaixo de um diretório contendo um arquivo go.mod

No modo ciente de módulo, GOPATH não define mais o significado de importações durante uma compilação, mas ainda armazena dependências baixadas (em GOPATH / pkg / mod) e comandos instalados (em GOPATH / bin, a menos que GOBIN esteja definido).


Confira abaixo como usamos o comando:

```bash
$ GO111MODULE=on go run myprogram.go
$ GO111MODULE=on go build myprogram.go
```
When our project is not in our **$GOPATH** it is not necessary to use **GO111MODULE**, but when our project is in **$GOPATH** and we want to use **"go mod"** we need to inform this to the compiler using **GO111MODULE**...


### Go Test
---

Pacotes de teste

Uso:

```go
$ go test [build/test flags] [packages] [build/test flags & test binary flags]
```

Go **test** automatiza o teste dos pacotes nomeados pelos caminhos de importação. Imprime um resumo dos resultados do teste no formato:

```bash
=== RUN   TestWhatever
--- PASS: TestWhatever (0.00s)
PASS
ok    command-line-arguments  0.001s
```

O pacote de teste é executado lado a lado com o comando go test. O teste de pacote deve ter o sufixo "\ _test.go".
Podemos dividir os testes em vários arquivos seguindo esta convenção. Por exemplo: "myprog1_test.go" and "myprog2_test.go".

Devemos colocar nossas funções de teste nesses arquivos de teste.

Cada função de teste é uma função pública exportada cujo nome começa com **"Test"**, aceita um ponteiro para um objeto **testing.T** e não retorna nada. Como isso:

Exemplo 1 / myprog1_test:

```go
package main

import "testing"

func TestWhatever(t *testing.T) {
    // Your test code goes here
}
```

```bash
$ go test -v
```

Saída:

```bash
=== RUN   TestWhatever
--- PASS: TestWhatever (0.00s)
PASS
ok    command-line-arguments  0.001s
```

O objeto T fornece vários métodos que podemos usar para indicar falhas ou erros de log.

Exexmplo 2 / myprog2_test:

```go
package main

import "testing"

func TestSum(t *testing.T) {
  x := 1 + 1
  if x != 11 { // forcing the error
    t.Error("Error! 1 + 1 is not equal to 2, I got", x)
  }
}
```

```bash
$ go test -v
```

Saída:

```bash
=== RUN   TestSum
-- FAIL: TestSum (0.00s)
    myprog1_test.go:12: Error! 1 + 1 is not equal to 2, I got 2
FAIL
FAIL  command-line-arguments  0.001s
```

Neste exemplo, faremos uma examinação de como seria em nossos projetos.

Neste programa vou passar parâmetro em tempo de compilação ou em nossa execução para facilitar e também servir como exemplo o uso de **"ldflags"** que podemos usar em **go run -ldflags** e **go build -ldflags**

De um check-in em nosso código abaixo / main.go:

```go
import "strconv"

import (
  "github.com/jeffotoni/goworkshopdevops/examples/tests/pkg/math"
)

var (
  x, y   string
  xi, yi int
)

func init() {
  xi, _ = strconv.Atoi(x)
  yi, _ = strconv.Atoi(y)
}

func main() {
  println(math.Sum(xi, yi))
}
```

Agora temos uma função Soma em um pacote que criamos em **pkg/math/sum.go**

```go
package math

func Sum(x, y int) int {
  return x + y
}
```

Criamos nosso arquivo de teste em **pkg/math/sum_test.go**

```go
package math

import "testing"

func TestSum(t *testing.T) {
  type args struct {
    x int
    y int
  }
  tests := []struct {
    name string
    args args
    want int
  }{
    // TODO: Add test cases.
    {"test_1: ", args{2, 2}, 4},
    {"test_2: ", args{-2, 6}, 4},
    {"test_3: ", args{-4, 8}, 4},
    {"test_4: ", args{5, 7}, 12},
    {"test_5: ", args{8, 8}, 15}, // forcing the error
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := Sum(tt.args.x, tt.args.y); got != tt.want {
        t.Errorf("Sum() = %v, want %v", got, tt.want)
      }
    })
  }
}
```

```bash
$ cd pkg/math/
$ go test -v
```

Saída:

```bash
=== RUN   TestSum
=== RUN   TestSum/test_1:_
=== RUN   TestSum/test_2:_
=== RUN   TestSum/test_3:_
=== RUN   TestSum/test_4:_
=== RUN   TestSum/test_5:_
--- FAIL: TestSum (0.00s)
    --- PASS: TestSum/test_1:_ (0.00s)
    --- PASS: TestSum/test_2:_ (0.00s)
    --- PASS: TestSum/test_3:_ (0.00s)
    --- PASS: TestSum/test_4:_ (0.00s)
    --- FAIL: TestSum/test_5:_ (0.00s)
        sum_test.go:29: Sum() = 16, want 15
FAIL
exit status 1
FAIL  github.com/jeffotoni/goworkshopdevops/examples/tests/pkg/pkg/math  0.001s
```

Converte para json a saída dos testes

```bash
$ go test -v -json
```

Verifique sua saída na tela do seu terminal para ver a saída do json.

---

Agora que salvamos nosso pkg / math / sum.go, vamos fazer um main.go fazendo a chamada neste pacote.
Mas primeiro vamos executar o go mod para gerenciar nossos pacotes e versões corretamente.

Verifique o comando abaixo:

```bash
$ go mod init github.com/jeffotoni/goworkshopdevops/examples/tests/pkg
```

Saída:

```bash
go: finding github.com/jeffotoni/goworkshopdevops/examples/tests/pkg/math latest
go: finding github.com/jeffotoni/goworkshopdevops/examples/tests latest
go: finding github.com/jeffotoni/goworkshopdevops/examples latest
go: finding github.com/jeffotoni/goworkshopdevops latest
go: downloading github.com/jeffotoni/goworkshopdevops v0.0.0-20190127023912-a2fa53299867
0
```

Agora podemos fazer **build** ou **run** em nosso **main.go**.

Vamos rodar para rodar usando o flag **"- ldflags"** para passar parâmetros para o nosso código em tempo de compilação.

```bash
$ go run -ldflags "-X main.x=2 -X main.y=3" main.go
```

Saída:

```bash
5
```

```bash
$ go run -ldflags "-X main.x=7 -X main.y=3" main.go
```

Saída:

```bash
10
```

## MÓDULO 02 Aprofundando na Linguagem
---

## Tipos
---

A type determines a set of values together with operations and methods specific to those values. A type may be denoted by a type name, if it has one, or specified using a type literal, which composes a type from existing types. 

The language predeclares certain type names. Others are introduced with type declarations. Composite types—array, struct, pointer, function, interface, slice, map, and channel types—may be constructed using type literals.

Each type T has an underlying type: If T is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is T itself. Otherwise, T's underlying type is the underlying type of the type to which T refers in its type declaration. 

Example:
```bash
type (
    A1 = string
    A2 = A1
)

type (
    B1 string
    B2 B1
    B3 []B1
    B4 B3
)
```

The underlying type of string, A1, A2, B1, and B2 is string. The underlying type of []B1, B3, and B4 is []B1. 


### Numeric Types

A numeric type represents sets of integer or floating-point values. The predeclared architecture-independent numeric types are: 

```bash
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```
The value of an n-bit integer is n bits wide and represented using two's complement arithmetic.

There is also a set of predeclared numeric types with implementation-specific sizes:

```bash
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

To avoid portability issues all numeric types are defined types and thus distinct except byte, which is an alias for uint8, and rune, which is an alias for int32. Conversions are required when different numeric types are mixed in an expression or assignment. For instance, int32 and int are not the same type even though they may have the same size on a particular architecture. 

### String Types

Um tipo de string representa o conjunto de valores de string. Um valor de string é uma seqüência (possivelmente vazia) de bytes. As strings são imutáveis: uma vez criadas, é impossível alterar o conteúdo de uma string. O tipo de string pré-declarado é string; é um tipo definido.

O comprimento de uma string s (seu tamanho em bytes) pode ser descoberto usando a função interna len. O comprimento é uma constante de tempo de compilação se a string for uma constante. Os bytes de uma string podem ser acessados por índices inteiros de 0 a len (s) -1. É ilegal tomar o endereço de tal elemento; Se s [i] é o i'th byte de uma string, & s [i] é inválido.

```go
package main

import "fmt"

type S string

var (
  String = "@jeffotoni"
)

func main() {
  var text string
  var str S

  mypicture := "@Photograph-jeffotoni"
  str = "@workshop-devOpsBh"
  text = "@jeffotoni-golang"

  fmt.Println(str)
  fmt.Println(String)
  fmt.Println(text)
  fmt.Println(mypicture)

  // example string
  s := "日本語"
  fmt.Printf("Glyph:             %q\n", s)
  fmt.Printf("UTF-8:             [% x]\n", []byte(s))
  fmt.Printf("Unicode codepoint: %U\n", []rune(s))
}
```
Saída:

```go
@workshop-devOpsBh
@jeffotoni
@jeffotoni-golang
@Photograph-jeffotoni
Glyph:             "日本語"
UTF-8:             [e6 97 a5 e6 9c ac e8 aa 9e]
Unicode codepoint: [U+65E5 U+672C U+8A9E]
```
### Pointer Types

Campos de estrutura podem ser acessados através de um ponteiro de estrutura.

Para acessar o campo X de uma struct quando tivermos o struct pointer p poderíamos escrever (* p) .X. No entanto, essa notação é incômoda, de modo que a linguagem nos permite escrever apenas p.X, sem a referência explícita.

Um tipo de ponteiro indica o conjunto de todos os ponteiros para variáveis de um determinado tipo, chamado tipo de base do ponteiro. O valor de um ponteiro não inicializado é nulo.

```bash
PointerType = "*" BaseType .
BaseType    = Type .
```

```bash
*Point
*[4]int
```

Exemplo:

```go
package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
  fmt.Println(p.Y)
}
```
Saída:

```bash
{1000000000 2}
2
```
Para cada tipo declarado, seja por você ou pela própria linguagem, você obtém gratuitamente um tipo de ponteiro de complemento que você pode usar para compartilhar. Já existe um tipo embutido chamado int, então existe um tipo de ponteiro de complemento chamado * int.

Todos os tipos de ponteiro possuem as mesmas duas características. Primeiro, eles começam com o personagem *. Segundo, todos eles têm o mesmo tamanho de memória e representação, que são 4 ou 8 bytes que representam um endereço. Em arquiteturas de 32 bits (como o playground), ponteiros exigem 4 bytes de memória e em arquiteturas de 64 bits (como sua máquina), eles exigem 8 bytes de memória.

Exemplo:

```go
package main

func main() {

  var a int
  inc := &a
  *inc = 2
  *inc++
  println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
```

Saída:

```bash
inc:  Value Of[ 0xc000036778 ]  Addr Of[ 0xc000036780 ]  Value Points To[ 3 ]
```

Para um operando x do tipo T, a operação de endereço & x gera um ponteiro do tipo * T para x. O operando deve ser endereçável, isto é, uma variável, indicação indireta de ponteiro ou operação de indexação de fatia; ou um seletor de campo de um operável struct endereçável; ou uma operação de indexação de matriz de uma matriz endereçável. Como uma exceção ao requisito de endereçamento, x também pode ser um literal composto (possivelmente entre parênteses). Se a avaliação de x causar um pânico em tempo de execução, a avaliação de & x também ocorrerá.

Para um operando x do tipo ponteiro * T, o indicador indireto * x denota a variável do tipo T apontada por x. Se x for nulo, uma tentativa de avaliar * x causará um pânico em tempo de execução.

```bash
&x
&a[f(2)]
&Point{2, 3}
*p
*pf(x)

var x *int = nil
*x   // causes a run-time panic
&*x  // causes a run-time panic
```

Veja o exemplo abaixo:

```go
package main

import "fmt"

func main() {
  var a int = 100  /* actual variable declaration */
  var pa *int      /* pointer variable declaration */
  var pointer *int /* pointer variable declaration */

  pa = &a /* store address of a in pointer variable*/

  fmt.Printf("Address of a variable: %x\n", &a)

  /* address stored in pointer variable */
  fmt.Printf("Address stored in ip variable: %x\n", pa)

  /* access the value using the pointer */
  fmt.Printf("Value of *ip variable: %d\n", *pa)

  /* succeeds if p is not nil */
  if pa != nil {
    fmt.Println("success is not nil")
  }

  /* succeeds if p is null */
  if (pointer) == nil {
    fmt.Println("success pointer is nil")
  }
}
```

Saída:

```bash
Address of a variable: c0000160c8
Address stored in ip variable: c0000160c8
Value of *ip variable: 100
success is not nil
success pointer is nil
```

### Array Types

Uma matriz é uma sequência numerada de elementos de um único tipo, denominada tipo de elemento. O número de elementos é chamado de comprimento e nunca é negativo.

```bash
ArrayType   = "[" ArrayLength "]" ElementType .
ArrayLength = Expression .
ElementType = Type .
```

O comprimento é parte do tipo da matriz; deve avaliar a uma constante não negativa representável por um valor do tipo int. O comprimento do array a pode ser descoberto usando a função interna len. Os elementos podem ser endereçados pelos índices inteiros 0 a len (a) -1. Os tipos de matriz são sempre unidimensionais, mas podem ser compostos para formar tipos multidimensionais.

```
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

O comprimento de uma matriz é parte do seu tipo, portanto, as matrizes não podem ser redimensionadas. Isso parece limitante, mas não se preocupe; Go fornece uma maneira conveniente de trabalhar com matrizes.

Exemplo:

```go
package main

import "fmt"

func main() {

  // var a []string // wrong

  // An array of 10 integers
  var a1 [5]int
  a1[0] = 5
  a1[1] = 4
  a1[2] = 3
  a1[3] = 2
  a1[4] = 1
  fmt.Println(a1)
}
```

Saída:

```bash
[5 4 3 2 1]
```

### Slice Types

Uma fatia é um descritor de um segmento contíguo de uma matriz subjacente e fornece acesso a uma seqüência numerada de elementos dessa matriz. Um tipo de fatia indica o conjunto de todas as fatias de matrizes de seu tipo de elemento. O valor de uma fatia não inicializada é nulo.

Uma matriz tem um tamanho fixo. Uma fatia, por outro lado, é uma visão flexível, de tamanho dinâmico, dos elementos de uma matriz. Na prática, as fatias são muito mais comuns que as matrizes.

O tipo **[] T** é uma fatia com elementos de **type T**.

Uma fatia é formada especificando dois índices, um limite baixo e alto, separados por dois pontos:

```bash
a[low : high]
```

Isso seleciona um intervalo half-open que inclui o primeiro elemento, mas exclui o último.

A expressão a seguir cria uma fatia que inclui os elementos de 1 a 3 de um:

```bash
a[1:4]
```

Exemplo:

```go
package main

import "fmt"

func main() {
  primes := [7]int{2, 3, 5, 7, 11, 13, 14}
  
  var p []int = primes[2:5]
  fmt.Println(p)
}
```
Saída:

```bash
[5 7 11]
```

Um novo valor de fatia inicializado para um determinado tipo de elemento T é feito usando a função interna make, que usa um tipo de fatia e parâmetros que especificam o comprimento e, opcionalmente, a capacidade. Uma fatia criada com make sempre aloca uma nova matriz oculta à qual o valor da fatia retornada se refere. Isto é, executando.

```bash
make([]T, length, capacity)
```

Produz a mesma fatia como alocar uma matriz e fatiá-la, então essas duas expressões são equivalentes:

```bash
make([]int, 50, 100)
new([100]int)[0:50]
```

Como matrizes, as fatias são sempre unidimensionais, mas podem ser compostas para construir objetos de dimensões mais altas. Com matrizes de matrizes, as matrizes internas são, por construção, sempre o mesmo comprimento; no entanto, com fatias de fatias (ou conjuntos de fatias), os comprimentos internos podem variar dinamicamente. Além disso, as fatias internas devem ser inicializadas individualmente.

Fatias podem ser criadas com a função built-in make; É assim que você cria matrizes de tamanho dinâmico.

A função make aloca uma matriz zerada e retorna uma fatia que se refere a essa matriz:

```go
a := make([]int, 4)  // len(a)=4
```
Exemplo:

```go
package main

import "fmt"

func main() {
  a := make([]int,4)
  a[0]=12
  fmt.Println("a", a)

  b := make([]int, 0, 5)
  fmt.Println("b", b)
  
  c := b[:2]
  fmt.Println("c", c)
}
```
Saída:

```bash
a [12 0 0 0]
b []
c [0 0]
```

Uma fatia do tipo T é declarada usando [] T. Por exemplo, aqui está como você pode declarar uma fatia do tipo int -

```go
// Slice of type `int`
var slice []int

// Slice of type `string`
var slice []string

// Slice of type `string` with parameter variadic
var lang = [...]string{"Erlang", "Elixir", "Haskell", "Clojure", "Scala"}
```

Você pode criar uma fatia usando um literal de fatia como este -

```go
// Creating a slice using a slice literal
var s = []int{3, 5, 7, 9, 11, 13, 17}
```

A expressão do lado direito da declaração acima é um literal de fatia. O literal de fatia é declarado como um literal de matriz, exceto que você não especifica nenhum tamanho entre colchetes [].

Quando você cria uma fatia usando um literal de fatia, ela primeiro cria uma matriz e, em seguida, retorna uma referência de fatia a ela.

Vamos ver:

```go
package main

import "fmt"

func main() {
	// Creating a slice using a slice literal
	var s = []string{"@jeffotoni", "@awsbrasil", "@devopsbh", "@go_br"}

	// Short hand declaration
	t := []int{2, 4, 8, 16, 32, 64}

	fmt.Println("s = ", s, len(s))
	fmt.Println("t = ", t, len(t))
}
```
Saída:

```bash
s =  [@jeffotoni @awsbrasil @devopsbh @go_br] 4
t =  [2 4 8 16 32 64] 6
```

Como uma fatia é um segmento de uma matriz, podemos criar uma fatia a partir de uma matriz.

Para criar uma fatia de uma matriz a, especificamos dois índices baixo (limite inferior) e alto (limite superior) separados por dois pontos

```bash
// Obtaining a slice from an array `a`
a[low:high]
```

A expressão acima seleciona uma fatia da matriz a. A fatia resultante inclui todos os elementos, começando do índice baixo para o alto, mas excluindo o elemento no índice alto.

```go
package main

import "fmt"

func main() {
	// Creating a slice using a slice literal
	var groups = [5]string{"@awsbrasil", "@devopsbh", "@go_br", "@devopsbr", "@docker"}

	// Creating a slice from the array
	var s []string = groups[2:5]

	s2 := s[1:3]
	s3 := s[:3]
	s4 := s[2:]
	s5 := s[:]

	fmt.Println("Array groups = ", groups, "len:", len(groups), "cap:", cap(groups))
	fmt.Println("Slice s = ", s, "len:", len(s), "cap:", cap(s))
	fmt.Println("Slice s = ", s2, "len:", len(s2), "cap:", cap(s2))
	fmt.Println("Slice s = ", s3, "len:", len(s3), "cap:", cap(s3))
	fmt.Println("Slice s = ", s4, "len:", len(s4), "cap:", cap(s4))
	fmt.Println("Slice s = ", s5, "len:", len(s5), "cap:", cap(s5))
}
```

Saída:

```bash
Array groups =  [@awsbrasil @devopsbh @go_br @devopsbr @docker] len: 5 cap: 5
Slice s =  [@go_br @devopsbr @docker] len: 3 cap: 3
Slice s =  [@devopsbr @docker] len: 2 cap: 2
Slice s =  [@go_br @devopsbr @docker] len: 3 cap: 3
Slice s =  [@docker] len: 1 cap: 1
Slice s =  [@go_br @devopsbr @docker] len: 3 cap: 3
```

A função copy () copia elementos de uma fatia para outra. Sua assinatura se parece com isso:

```go
func copy(dst, src []T) int
```

São necessárias duas fatias - uma fatia de destino e uma fatia de origem. Em seguida, copia elementos da origem para o destino e retorna o número de elementos copiados.

Observe que os elementos são copiados somente se a fatia de destino tiver capacidade suficiente.

```go
package main

import "fmt"

func main() {

	src := []string{"Erlang", "Elixir", "Haskell", "Clojure", "Scala"}
	dest := make([]string, 2)

	numElementsCopied := copy(dest, src)

	fmt.Println("src = ", src)
	fmt.Println("dest = ", dest)
	fmt.Println("Number of elements copied from src to dest = ", numElementsCopied)
}
```

Saída:

```bash
src =  [Erlang Elixir Haskell Clojure Scala]
dest =  [Erlang Elixir]
Number of elements copied from src to dest =  2
```
A função append () acrescenta novos elementos no final de uma determinada fatia. A seguir, a assinatura da função append.

```go
func append(s []T, x ...T) []T
```

É preciso uma fatia e um número variável de argumentos x… T. Em seguida, ele retorna uma nova fatia contendo todos os elementos da fatia especificada, bem como os novos elementos.

Se o segmento especificado não tiver capacidade suficiente para acomodar novos elementos, um novo array subjacente será alocado com maior capacidade. Todos os elementos da matriz subjacente da fatia existente são copiados para essa nova matriz e, em seguida, os novos elementos são anexados.

No entanto, se a fatia tiver capacidade suficiente para acomodar novos elementos, a função append () reutilizará sua matriz subjacente e anexará novos elementos à mesma matriz.

Vamos ver um exemplo:

```go
package main

import "fmt"

func main() {
	slice1 := []string{"Clojure", "Scala", "Elixir"}
	slice2 := append(slice1, "Assembly", "Rust", "Go")
	fmt.Printf("slice1 = %v, len = %d, cap = %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 = %v, len = %d, cap = %d\n", slice2, len(slice2), cap(slice2))
	slice1[0] = "C++"
	fmt.Println("\nslice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	
	// slice nil
	var s []string

	// Appending to a nil slice
	s = append(s, "Java", "C", "Lisp", "Haskell")
	fmt.Printf("\ns = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
```

Saída:

```bash
slice1 = [Clojure Scala Elixir], len = 3, cap = 3
slice2 = [Clojure Scala Elixir Assembly Rust Go], len = 6, cap = 6

slice1 =  [C++ Scala Elixir]
slice2 =  [Clojure Scala Elixir Assembly Rust Go]

s = [Java C Lisp Haskell], len = 4, cap = 4
```

### Struct Types

Uma struct é uma sequência de elementos nomeados, chamados de campos, cada um com um nome e um tipo. Nomes de campos podem ser especificados explicitamente (IdentifierList) ou implicitamente (EmbeddedField). Dentro de uma estrutura, os nomes de campos não vazios devem ser exclusivos.

```go
StructType    = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
EmbeddedField = [ "*" ] TypeName .
Tag           = string_lit .

// An empty struct.
struct{}

// A struct with 6 fields.
struct {
  x, y int
  u float32
  _ float32  // padding
  A *[]int
  F func()
}
```

Um campo declarado com um tipo, mas sem nome de campo explícito, é chamado de campo incorporado. Um campo incorporado deve ser especificado como um nome de tipo T ou como um ponteiro para um nome de tipo não-interface * T, e T em si não pode ser um tipo de ponteiro. O nome do tipo não qualificado atua como o nome do campo.

```go
// A struct with four embedded fields of types T1, *T2, P.T3 and *P.T4
struct {
  T1        // field name is T1
  *T2       // field name is T2
  P.T3      // field name is T3
  *P.T4     // field name is T4
  x, y int  // field names are x and y
}
```

A declaração a seguir é ilegal porque os nomes de campo devem ser exclusivos em um tipo de estrutura:

```go
struct {
  T     // conflicts with embedded field *T and *P.T
  *T    // conflicts with embedded field T and *P.T
  *P.T  // conflicts with embedded field T and *T
}
```
Uma estrutura é uma coleção de campos. 

```go
package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{10, 201}
  v.X = 4
  fmt.Println(v)
}
```

Saída:

```bash
{4 201}
```

### Struct in C 

Antes de vermos a origem de tudo em C, vamos ver o código escrito em C e os comentários de como seria em Golang.
E nós temos isso escrito em Golang, para que você possa ver como você estava dinâmico em Go, a herança C era simplesmente fantástica e leve.

```c
#include <stdio.h>
#include <stdlib.h>

struct Login {
    char *User;
    char *Email;
};

/**
  // In Go would look like this
  type Login struct {
    User string
    Email string
  }
 */

int main() {
/** 
 // In Go 
func main() {
*/
    
    struct Login login, *pointerToLogin;

    /**
    // In Go
    login := Login{User:"jeffotoni", Email:"jef@m.com"}
    */

    login.User = "jeffotoni";
    login.Email = "jef@m.com";

    pointerToLogin = malloc(sizeof(struct Login));
    pointerToLogin->User = "pike";
    pointerToLogin->Email = "pike@g.com";

    /**
    var pointerToLogin  = new(Login)
    pointerToLogin.User  = "pike"
    pointerToLogin.Email = "pike@g.com"
    */
    printf("login vals: %s %s\n", login.User, login.Email);
    printf("pointerToLogin: %p %s %s\n", pointerToLogin, pointerToLogin->User, pointerToLogin->Email);

    /**
     fmt.Printf("login vals: %s %s\n", login.User, login.Email)
     fmt.Printf("pointerToLogin: %v %s %s\n", pointerToLogin, pointerToLogin.User, pointerToLogin.Email);
    */

    free(pointerToLogin);
    return 0;
}
```

Compile o programa em C:

```
$ gcc -o struct-1-c struct-1-c.c
```

Saída:

```bash
login vals: jeffotoni jef@m.com
pointerToLogin: 0x5627788a0260 pike pike@g.com
```

Confira agora o Código em Go abaixo:

```go
package main

import "fmt"
// #include <stdio.h>
// #include <stdlib.h>

// struct Login {
//     char *User;
//     char *Email;
// };

//In Go would look like this
type Login struct {
    User  string
    Email string
}

// int main() {

// In Go
func main() {

    //struct Login login, *pointerToLogin;

    // In Go
    login := Login{User: "jeffotoni", Email: "jef@m.com"}

    login.User = "jeffotoni"
    login.Email = "jef@m.com"

    // pointerToLogin = malloc(sizeof(struct Login));
    // pointerToLogin->User = "pike";
    // pointerToLogin->Email = "pike@g.com";

    var pointerToLogin = new(Login)
    pointerToLogin.User = "pike"
    pointerToLogin.Email = "pike@g.com"

    //printf("login vals: %s %s\n", login.User, login.Email);
    //printf("pointerToLogin: %p %s %s\n", pointerToLogin, pointerToLogin->User, pointerToLogin->Email);

    fmt.Printf("login vals: %s %s\n", login.User, login.Email)
    fmt.Printf("pointerToLogin: %v %s %s\n", pointerToLogin, pointerToLogin.User, pointerToLogin.Email)

    //free(pointerToLogin);
    return
}
```

```bash
login vals: jeffotoni jef@m.com
pointerToLogin: &{pike pike@g.com} pike pike@g.com
```

### Struct Type Tags Json

Usamos tags json com omitempty ou não para representar nossa string json, quando geradas pelo **json.Marshal**, que veremos logo abaixo.

Exemplo:

```go
import (
	//"encoding/json"
	"fmt"
)

// We use the omitempty tag in 'json: field, omitempty'
// when we want the field not to appear after
// making the marshal if it is empty.
type D struct {
	Height int
	Width  int `json:"width,omitempty"`
}

// Fields that do not have the "omitempty" tag will display 
// the same field being empty when generating 
// json through marshal.
type Cat struct {
	Breed    string `json:"breed,omitempty"`
	WeightKg int    `json:WeightKg`
	Size     D      `json:"size,omitempty"`
}

func main() {
	d := Cat{
		//Breed:    "Persa",
		WeightKg: 23,
		//Size:     D{20, 60},
	}
	//b, _ := json.Marshal(d)
	//fmt.Println(string(b))
	fmt.Println(d)
}
```

Saída:

```bash
{ 23 {20 60}}
```

Neste exemplo temos uma estrutura "ApiLogin" e dentro dela temos And1 \ * struct ", ou seja, referenciando um ponteiro de outra struct, beautiful não é.
Para inicializar e acessar o campo "And1", temos que usar o operador "&" e criar a estrutura em sua inicialização, porque ela ainda não foi definida, então ficaria assim: **"And1: & struct {City string} {City: "BH"}"**
```go
type ApiLogin struct {
	And1  *struct {
		City string
	}
}
```

Neste exemplo, só damos tempo ao ponteiro de uma estrutura já definida acima.
Para acessá-lo não precisamos criá-lo novamente, basta se referir a ele com "&" e assim "And2: & Address {}", muito legal, certo?

```go
type ApiLogin struct {
	And2 *Address
}
```

O exemplo abaixo é uma maneira de exibir várias formas de inicialização usando struct.

```go
package main

import "fmt"

type Address struct {
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Zipcode      string `json:"zipcode"`
}

type ApiLogin struct {
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Login string `json:"login"`
	Email string `json:"email"`

	// anonymous
	And1 *struct {
		City string
	}

	// pointer
	And2 *Address

	// list Address
	And3 []Address
}

func main() {

	// different ways to inizialize a struct
	//
	//

	apilogin1 := &ApiLogin{Name: "@jeffotoni", Cpf: "093.393.334-34", And1: &struct{ City string }{City: "BH"}}
	fmt.Println(apilogin1)
	fmt.Println(apilogin1.Name)
	fmt.Println(apilogin1.And1)
	fmt.Println(apilogin1.And1.City)

	apilogin2 := &ApiLogin{Name: "@jeffotoni", Cpf: "093.393.334-34", And2: &Address{City: "BH"}}
	fmt.Println(apilogin2)
	fmt.Println(apilogin2.Name)
	fmt.Println(apilogin2.And2)
	fmt.Println(apilogin2.And2.City)

	apilogin3 := &ApiLogin{Name: "@jeffotoni", Cpf: "093.393.334-34", And1: &struct{ City string }{City: "BH"}, And2: &Address{City: "BH"}}
	fmt.Println(apilogin3)
	fmt.Println(apilogin3.Name)
	fmt.Println(apilogin3.And1)
	fmt.Println(apilogin3.And1.City)
	fmt.Println(apilogin3.And2)
	fmt.Println(apilogin3.And2.City)

	var apilogin4 ApiLogin
	fmt.Println(apilogin4)

	apilogin5 := ApiLogin{}
	fmt.Println(apilogin5)

	apilogin6 := &ApiLogin{}
	fmt.Println(apilogin6)

	apilogin7 := new(ApiLogin)
	fmt.Println(apilogin7)

	// another way to feed the struct
	g1add := Address{City: "Belo Horizonte"}
	g2add := Address{City: "Curitiba"}

	// declaring as list
	gall := []Address{}

	// add items
	gall = append(gall, g1add)
	gall = append(gall, g2add)

	fmt.Println(gall)

	// initializes Struct
	apil3 := ApiLogin{}

	// recive same type
	apil3.And3 = gall

	// show struct
	fmt.Println(apil3)

	// another way to initialize and feed the struct list
	apil3.And3 = []Address{{City: "Sao Paulo"}, {City: "Brasilia"}}

	// show struct
	fmt.Println(apil3)
}
```

Saída:

```bash
&{@jeffotoni 093.393.334-34   0xc00000e1e0 <nil> []}
@jeffotoni
&{BH}
BH
&{@jeffotoni 093.393.334-34   <nil> 0xc000060150 []}
@jeffotoni
&{BH  }
BH
&{@jeffotoni 093.393.334-34   0xc00000e300 0xc0000601b0 []}
@jeffotoni
&{BH}
BH
&{BH  }
BH
{    <nil> <nil> []}
{    <nil> <nil> []}
&{    <nil> <nil> []}
&{    <nil> <nil> []}
[{Belo Horizonte  } {Curitiba  }]
{    <nil> <nil> [{Belo Horizonte  } {Curitiba  }]}
{    <nil> <nil> [{Sao Paulo  } {Brasilia  }]}
```

O exemplo abaixo é para demonstrar a facilidade que temos quando manipulamos estruturas em Golang.
Estamos inicializando a estrutura com valores e exibindo na tela.

Dê uma olhada:

```go
package main

import "fmt"

type jsoninput []struct {
	Data string `json:"data"`
}

func main() {

	data := &jsoninput{{Data: "some data"}, {Data: "some more data"},
		{Data: "some more data"}}
	fmt.Println(data)

	// output:
	//  &[{some data} {some more data} {some more data}]
}
```

```bash
&[{some data} {some more data} {some more data}]
```

Em nosso exemplo abaixo é outra maneira de fazer array de struct, fizemos a declaração no momento de inicializar nossa struct, isso faz com que a struct não seja pego apenas em aceitar array de struct.
Nós acrescentamos nos campos e a mágica acontece.

Vamos dar uma olhada no código completo:

```go
package main

import (
	"fmt"
)

type User struct {
	FirstName string `tag_name:"firstname"`
	LastName  string `tag_name:"lastname"`
	Age       int    `tag_name:"age"`
}

func main() {
	// create instance
	// slice struct
	users := []*User{}

	user := new(User)
	user.FirstName = "Jefferson"
	user.LastName = "otoni"
	user.Age = 350
	users = append(users, user)

	user = new(User)
	user.FirstName = "Pike"
	user.LastName = "Hob"
	user.Age = 55
	users = append(users, user)

	fmt.Println(users)

	for k, v := range users {
		fmt.Println(k, v)
		fmt.Println(v.FirstName)
		fmt.Println(v.LastName)
		fmt.Println(v.Age)
	}
}
```

Saída:

```bash
[0xc000060150 0xc000060180]
0 &{Jefferson otoni 350}
Jefferson
otoni
350
1 &{Pike Hob 55}
Pike
Hob
55
```

Exemplo de estrutura AWS Sqs Json

```go
type SqsJson struct {
	Type      string `json:"type"`
	MessageId string `json:"messageid"`
	TopicArn  string `json:"topicarn"`
	Message   string `json:"message"`
	//Message          JsonMessage
	Timestamp        string `json:"timestamp"`
	SignatureVersion string `json:"signatureversion"`
	Signature        string `json:"signature"`
	SigningCertURL   string `json:"signingcerturl"`
	UnsubscribeURL   string `json:"unsubscribeurl"`
}

type JsonMessage struct {
	NotificationType string `json:"notificationType"`
	Bounce           struct {
		BounceType        string `json:"bounceType"`
		BounceSubType     string `json:"bounceSubType"`
		BouncedRecipients []struct {
			EmailAddress   string `json:"emailAddress"`
			Action         string `json:"action"`
			Status         string `json:"status"`
			DiagnosticCode string `json:"diagnosticCode"`
		} `json:"bouncedRecipients"`
		Timestamp    time.Time `json:"timestamp"`
		FeedbackID   string    `json:"feedbackId"`
		RemoteMtaIP  string    `json:"remoteMtaIp"`
		ReportingMTA string    `json:"reportingMTA"`
	} `json:"bounce"`

	Mail struct {
		Timestamp        time.Time `json:"timestamp"`
		Source           string    `json:"source"`
		SourceArn        string    `json:"sourceArn"`
		SourceIP         string    `json:"sourceIp"`
		SendingAccountID string    `json:"sendingAccountId"`
		MessageID        string    `json:"messageId"`
		Destination      []string  `json:"destination"`
		HeadersTruncated bool      `json:"headersTruncated"`
		Headers          []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"headers"`

		CommonHeaders struct {
			From    []string `json:"from"`
			ReplyTo []string `json:"replyTo"`
			To      []string `json:"to"`
			Subject string   `json:"subject"`
		} `json:"commonHeaders"`
	} `json:"mail"`
}
```

Quando o assunto é struct temos várias possibilidades de lidar e trabalhar com esse recurso em Golang, ele está praticamente embutido em tudo que vamos construir em Golang, Structs é algo poderoso em Go manipular dados e enviar dados o tempo todo entre canais usando goroutine, seja em filas, gravações de banco de dados, json e leituras de banco de dados, GraphQL, REST API, SOAP etc ...


### Fatih Structs to Map

Abaixo está uma biblioteca que trabalha diretamente com struct, convertendo para o Maps.
Eu não uso na produção, mas para o nosso curso é interessante analisar.
Sabemos que quanto mais nativos somos em Golang, será uma boa opção, mas às vezes precisaremos de algumas libs para nos ajudar.
Este projeto não é mais mantido e é arquivado. Sinta-se à vontade para bifurcar e fazer suas próprias alterações, se necessário.

Structs contém vários utilitários para trabalhar com estruturas Go (Golang). Ele foi inicialmente usado por mim para converter uma estrutura em uma interface {string] do mapa {}. Com o tempo, adicionei outros utilitários para estruturas. É basicamente um pacote de alto nível baseado em primitivas do pacote reflect. Sinta-se à vontade para adicionar novas funções ou melhorar o código existente.

Instalar

```go
go get github.com/fatih/structs
```

Para testar baixo, gire o código abaixo:

```go

type Server struct {
	Name        string `json:"name,omitempty"`
	ID          int
	Enabled     bool
	Users       []string // not exported
	http.Server          // embedded
}

func main() {

	// Create a new struct type:

	server := &Server{
		Name:    "gopher",
		ID:      12345678,
		Users:   []string{"jeffotoni", "pike", "dennis", "ken"},
		Enabled: true,
	}

	// struct
	fmt.Println(server)

	// create struct fatih
	s := structs.New(server)
	m := s.Map() // Get a map[string]interface{}
	fmt.Println(m)

	v := s.Values() // Get a []interface{}
	fmt.Println(v)

	f := s.Fields() // Get a []*Field
	fmt.Println(f)

	n := s.Names() // Get a []string
	fmt.Println(n)

	name := s.Field("Name") // Get a *Field based on the given field name

	// Get the underlying value,  value => "gopher"
	value := name.Value().(string)
	fmt.Println(value)

	tagValue := name.Tag("json")
	fmt.Println(tagValue)

	f1, ok := s.FieldOk("Name") // Get a *Field based on the given field name
	fmt.Println(f1, ok)

	n2 := s.Name() // Get the struct name
	fmt.Println(n2)

	h := s.HasZero() // Check if any field is uninitialized
	fmt.Println(h)
}
```

Saída:

```bash
&{gopher 12345678 true [jeffotoni pike dennis ken] 
{ <nil> <nil> 0s 0s 0s 0s 0 map[] <nil> <nil> 0 0 {{0 0} 0} <nil> {0 0} map[] map[] <nil> []}}
map[Server:map[TLSConfig:<nil> ReadHeaderTimeout:0s IdleTimeout:0s Addr: Handler:<nil> 
]MaxHeaderBytes:0 TLSNextProto:map[] ConnState:<nil> ErrorLog:<nil> ReadTimeout:0s WriteTimeout:0s] 
Name:gopher ID:12345678 Enabled:true Users:[jeffotoni pike dennis ken]]
[gopher 12345678 true [jeffotoni pike dennis ken]  <nil> <nil> 0s 0s 0s 0s 0 map[] <nil> <nil>]
[0xc000106000 0xc000106090 0xc000106120 0xc0001061b0 0xc000106240]
[Name ID Enabled Users Server]
gopher
name,omitempty
&{{0x626180 0xc0000f6000 408} {Name  0x626180 json:"name,omitempty" 0 [0] false} structs} true
Server
true
```

### Map Types

Um mapa é um grupo não ordenado de elementos de um tipo, chamado de tipo de elemento, indexado por um conjunto de chaves exclusivas de outro tipo, chamado de tipo de chave. O valor de um mapa não inicializado é nulo.

```bash
MapType     = "map" "[" KeyType "]" ElementType .
KeyType     = Type .
```

Os operadores de comparação == e! = Devem ser totalmente definidos para operandos do tipo de chave; Assim, o tipo de chave não deve ser uma função, mapa ou fatia. Se o tipo de chave for um tipo de interface, esses operadores de comparação devem ser definidos para os valores de chave dinâmica; falha causará um pânico em tempo de execução.

```bash
map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}
```

O número de elementos do mapa é chamado seu comprimento. Para um mapa m, ele pode ser descoberto usando a função interna len e pode mudar durante a execução. Elementos podem ser adicionados durante a execução usando atribuições e recuperados com expressões de índice; eles podem ser removidos com a função integrada de exclusão.

Um novo valor de mapa vazio é feito usando a função interna make, que usa o tipo de mapa e uma sugestão de capacidade opcional como argumentos:

```bash
make(map[string]int)
make(map[string]int, 100)
```

A capacidade inicial não vincula seu tamanho: os mapas crescem para acomodar o número de itens armazenados neles, com exceção dos mapas nulos. Um mapa nulo é equivalente a um mapa vazio, exceto que nenhum elemento pode ser adicionado.

Alguns exemplos de inicialização de mapa:

```go
package main

import "fmt"

type linkResult struct {
	body string
	urls []string
}

type linkFetcher map[string]*linkResult

func main() {
	// Required to initialize
	// the map with values
	var m1 map[string]int
	var m2 = make(map[string]int)
	var m3 = map[string]int{"population": 500000}
	var m4 = m3
	var m5 map[string]string
	/* create a map*/
	m5 = make(map[string]string)
	fmt.Println(m1, m2, m3, m4, m5)

	var l = linkFetcher{
		"https://golang.org/": &linkResult{
			"The Go Programming Language",
			[]string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": &linkResult{
			"Packages",
			[]string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		}}

	fmt.Println(l)
}
```
Saída:

```bash
map[] map[] map[population:500000] map[population:500000] map[]
```

Um mapa é declarado usando a seguinte sintaxe:

```go
var m map[KeyType]ValueType

// For example, Here is how you can declare a map of string keys to int values
var m map[string]int
```

O valor zero de um mapa é nulo. Um mapa nulo não possui chaves. Além disso, qualquer tentativa de adicionar chaves a um mapa nulo resultará em um erro de execução.

Vamos ver um exemplo:

```go
package main

import "fmt"

func main() {

	// Required to initialize
	// the map with values
	var m map[string]int
	fmt.Println(m)
	if m == nil {
		fmt.Println("is nil")
	}
	// Attempting to add keys
	// to a nil map will result in a runtime error
	//m["population"] = 500000
	//fmt.Println(m)
}
```

Saída:

```bash
map[]
is nil
```

Você pode inicializar um mapa usando a função interna make (). Você só precisa passar o tipo do mapa para a função make () como no exemplo abaixo. A função retornará um mapa inicializado e pronto para uso.

```go
// Initializing a map using the built-in make() function
var m = make(map[string]int)
```

Exemplo:

```go
package main

import "fmt"

func main() {

	// Required to initialize
	// the map with values
	var m = make(map[string]int)
	fmt.Println(m)
	if m == nil {
		fmt.Println("is nil")
	}
	m["population"] = 500000
	fmt.Println(m)
}
```
Saída:

```bash
map[]
map[population:500000]
```

O exemplo abaixo introduz a criação de um mapa obtendo um struct feito.
Quando usamos uma estrutura vazia?
Existem alguns cenários em que temos uma grande quantidade de acesso em nossa API, mas quando digo grande é> 10k = 10k de solicitações por segundo, nesse cenário quando fazemos nosso manipulador, podemos implementar um canal recebendo uma estrutura vazia { } para que possamos colocar em um canal e processar tudo com mais segurança.
Mostraremos mais adiante esta abordagem muito legal.

Exemplo:

```go
package main

import "fmt"

func main() {
	s := "key"
	seen := make(map[string]struct{}) // set of strings
	// ...
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
		// ...first time seeing s...
	}
	fmt.Println(seen)
}
```

Saída:

```go
map[key:{}]
```
Também podemos fazer com que nossa chave de mapa receba uma estrutura vazia.
Bem, nós sabemos que ele recebe qualquer tipo, ou seja, a estrutura pode ser completa sem que seja feita também.

Exemplo:

```go
type T struct{}

func main() {
	s := T{}
	seen := make(map[struct{}]struct{}) // set of strings
	// ...
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
		// ...first time seeing s...
	}
	fmt.Println(seen)
}
```

Saída:

```go
map[{}:{}]
```

### Map Literals Continued

Um map literal é uma maneira muito conveniente de inicializar um mapa com alguns dados. Você só precisa passar os pares de valores-chave separados por dois pontos dentro de chaves.

```go
package main

import "fmt"

func main() {
	var m = map[string]string{
		"Brasil": "Brasilia",
		"EUA":    "Washington, D.c",
		"Italy":  "Roma",
		"France": "Paris",
		"Japan":  "Toquio",
	}

	fmt.Println(m)
}
```

Saída:

```bash
map[Italy:Roma France:Paris Japan:Toquio Brasil:Brasilia EUA:Washington, D.c]
```

Assim, você pode verificar a existência de uma chave em um mapa usando a seguinte atribuição de dois valores.
A variável booleana ok será verdadeira se a chave existir e, caso contrário, será falsa.

```go
value, ok := m[key]
```

Considere o seguinte mapa, por exemplo:

```go
var C = map[string]string{
		"Brasil": "Brasilia",
		"EUA":    "Washington, D.c",
		"Italy":  "Roma",
		"France": "Paris",
		"Japan":  "Toquio",
	}
```

```go
capital, ok := C["EUA"]  // "Washington, D.c", true
```

No entanto, se você tentar acessar uma chave que não existe, o mapa retornará uma string vazia "" (valor zero de strings) e false

```go
capital, ok := C["África do Sul"]  // "", false
```
Você pode excluir uma chave de um mapa usando a função integrada delete (). A sintaxe é assim.

```go
// Delete the `key` from the `map`
delete(map, key)
```
Exemplo:

```go
package main

import "fmt"

func main() {
	var country = map[string]string{
		"Brasil": "Brasilia",
		"EUA":    "Washington, D.c",
		"Italy":  "Roma",
		"France": "Paris",
		"Japan":  "Toquio",
	}
	delete(country, "Japan")
	delete(country, "Italy")
	fmt.Println(country)
}
```

Saída:

```bash
map[Brasil:Brasilia EUA:Washington, D.c France:Paris]
```
 
Se o tipo de nível superior for apenas um nome de tipo, você poderá omiti-lo dos elementos do literal. 

```go
package main

import "fmt"

type Login struct {
	User, Login, Email string
}

// passing a struct as parameter
// for our struct map
var m = map[string]Login{
	"jeffotoni": {"jeffotoni", "jeff", "jeff@gm.com"},
	"Google":    {"root", "super", "google@gm.com"},
}

func main() {
	fmt.Println(m)
}
```go

```bash
map[jeffotoni:{jeffotoni jeff jeff@gm.com} Google:{root super google@gm.com}]
```

### Channel Types

Um canal fornece um mecanismo para a execução simultânea de funções para comunicação, enviando e recebendo valores de um tipo de elemento especificado. O valor de um canal não inicializado é nulo.

```bash
ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .
```

O operador <- opcional especifica a direção do canal, enviar ou receber. Se nenhuma direção for dada, o canal é bidirecional. Um canal pode ser restrito apenas para enviar ou receber apenas por conversão ou atribuição.

```bash
chan T          // can be used to send and receive values of type T
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints
```
O operador <- se associa ao canal mais à esquerda possível:

```bash
chan<- chan int    // same as chan<- (chan int)
chan<- <-chan int  // same as chan<- (<-chan int)
<-chan <-chan int  // same as <-chan (<-chan int)
chan (<-chan int)
```

Um novo valor de canal inicializado pode ser feito usando a função interna make, que aceita o tipo de canal e uma capacidade opcional como argumentos:

```bash
make(chan int, 100)
```

A capacidade, em número de elementos, define o tamanho do buffer no canal. Se a capacidade for zero ou ausente, o canal não será buffer e a comunicação será bem-sucedida somente quando o remetente e o receptor estiverem prontos. Caso contrário, o canal é armazenado em buffer e a comunicação é bem-sucedida sem bloqueio se o buffer não estiver cheio (envia) ou não estiver vazio (recebe). Um canal nulo nunca está pronto para comunicação.

Um canal pode ser fechado com a função interna fechada. O formulário de atribuição de valor múltiplo do operador de recebimento informa se um valor recebido foi enviado antes do canal ser fechado.

A single channel may be used in send statements, receive operations, and calls to the built-in functions cap and len by any number of goroutines without further synchronization. Channels act as first-in-first-out queues. Por exemplo, se uma goroutine enviar valores em um canal e uma segunda goroutine os receber, os valores serão recebidos na ordem enviada.

Deixe-me mostrar-lhe um exemplo:

```go

package main

import (
  "fmt"
  "os"
  "time"
)

type Promise struct {
  Result chan string
  Error  chan error
}

var (
  ch1  = make(chan *Promise)  // received a pointer from the structure
  ch2  = make(chan string, 1) // allows only 1 channels
  ch3  = make(chan int, 2)    // allows only 2 channels
  ch4  = make(chan float64)   // has not been set can freely receive
  ch5  = make(chan []byte)    // by default the capacity is 0
  ch6  = make(chan bool, 1)   // non-zero capacity
  ch7  = make(chan time.Time, 2)
  ch8  = make(chan struct{}, 2)
  ch9  = make(chan struct{})
  ch10 = make(map[string](chan int)) // map channel
  ch11 = make(chan error)
  ch12 = make(chan error, 2)
  // receives a zero struct
  ch14 <-chan struct{}
  ch15 = make(<-chan bool)          // can only read from
  ch16 = make(chan<- []os.FileInfo) // // can only write to
  // holds another channel as its value
  ch17 = make(chan<- chan bool) // // can read and write to
)

// Parameters of Func
// (jobs <-chan int, results chan<- int)

// Receives Value, only read
// jobs <-chan int //receives the value

// Receives Channel, only write
// results chan<- int // receive channel
// or
// results chan int // receive channel

// Receives Channel variadic
// results ...<-chan int

func main() {

  ch2 <- "okay"
  defer close(ch2)
  fmt.Println(ch2, &ch2, <-ch2)

  ch7 <- time.Now()
  ch7 <- time.Now()
  fmt.Println(ch7, &ch7, <-ch7)
  fmt.Println(ch7, &ch7, <-ch7)
  defer close(ch7)

  ch3 <- 1 // okay
  ch3 <- 2 // okay
  // deadlock
  // ch3 <- 3 // does not accept any more values, if you do it will error : deadlock
  defer close(ch3)
  fmt.Println(ch3, &ch3, <-ch3)
  fmt.Println(ch3, &ch3, <-ch3)

  ch10["lambda"] = make(chan int, 2)
  ch10["lambda"] <- 100
  defer close(ch10["lambda"])
  fmt.Println(<-ch10["lambda"])
}
```

Saída:

```bash
0xc000056180 0x55bb00 okay
0xc0000561e0 0x55bb28 2019-01-25 15:11:41.982906669 -0200 -02 m=+0.000147197
0xc0000561e0 0x55bb28 2019-01-25 15:11:41.982906922 -0200 -02 m=+0.000147409
0xc00001e0e0 0x55bb08 1
0xc00001e0e0 0x55bb08 2
100
```

### Blank Identifier

O identificador em branco é representado pelo caractere de sublinhado **_**. Ele serve como um espaço reservado anônimo em vez de um identificador regular (non-blank) e tem um significado especial em declarações, como um operando e em atribuições.

Exemplo:

```bash
// function statement
func f() (int, string, error)

// function return
_, _, _ := f()
```

### Tipos de interface

**Uma interface são duas coisas:**
 - é um conjunto de métodos
 - mas também é um tipo

O __interface {} type__, a interface vazia é a interface que tem __no métodos__

Como não há nenhuma palavra-chave implements, todos os tipos implementam pelo menos zero métodos e a satisfação de uma interface é feita automaticamente, todos os tipos satisfazem a interface vazia.
Isso significa que, se você escrever uma função que usa um valor {} de interface como um parâmetro, você poderá fornecer essa função com qualquer valor.

Exemplo:

```go
func DoSomething(v interface{}) {
   // ...
}

var Msg interface{}

type Stringer interface {
    String() string
}
```

#### Aqui está uma interface como um método

Um tipo de interface especifica um conjunto de métodos chamado sua interface. Uma variável do tipo de interface pode armazenar um valor de qualquer tipo com um conjunto de métodos que seja qualquer superconjunto da interface. Tal tipo é dito para implementar a interface. O valor de uma variável não inicializada do tipo de interface é nulo.


```bash
InterfaceType      = "interface" "{" { MethodSpec ";" } "}" .
MethodSpec         = MethodName Signature | InterfaceTypeName .
MethodName         = identifier .
InterfaceTypeName  = TypeName .
```

Como com todos os conjuntos de métodos, em um tipo de interface, cada método deve ter um nome exclusivo não vazio.

```go
// A simple File interface
interface {
  Read(b Buffer) bool
  Write(b Buffer) bool
  Close()
}
```

Mais de um tipo pode implementar uma interface. Por exemplo, se dois tipos S1 e S2 tiverem o conjunto de métodos

```bash
func (p T) Read(b Buffer) bool { return … }
func (p T) Write(b Buffer) bool { return … }
func (p T) Close() { … }
```

(onde T significa S1 ou S2) então a interface File é implementada por S1 e S2, independentemente de quais outros métodos S1 e S2 possam ter ou compartilhar.

Um tipo implementa qualquer interface que inclua qualquer subconjunto de seus métodos e, portanto, pode implementar várias interfaces distintas. Por exemplo, todos os tipos implementam a interface vazia:

```bash
interface{}
```

Da mesma forma, considere esta especificação de interface, que aparece dentro de uma declaração de tipo para definir uma interface chamada Locker:

```go
type Locker interface {
  Lock()
  Unlock()
}
```

Se S1 e S2 também implementarem

```bash
func (p T) Lock() { … }
func (p T) Unlock() { … }
```

Eles implementam a interface do Locker, bem como a interface do arquivo.

Uma interface T pode usar um nome de tipo de interface (possivelmente qualificado) E no lugar de uma especificação de método. Isso é chamado de interface de incorporação E em T; adiciona todos os métodos (exportados e não exportados) de E para a interface T.

```go
type ReadWriter interface {
  Read(b Buffer) bool
  Write(b Buffer) bool
}

type File interface {
  ReadWriter  // same as adding the methods of ReadWriter
  Locker      // same as adding the methods of Locker
  Close()
}

type LockedFile interface {
  Locker
  File        // illegal: Lock, Unlock not unique
  Lock()      // illegal: Lock not unique
}
```

Um tipo de interface T não pode incorporar a si mesmo ou a qualquer tipo de interface que incorpore T, recursivamente.

```go
// illegal: Bad cannot embed itself
type Bad interface {
  Bad
}

// illegal: Bad1 cannot embed itself using Bad2
type Bad1 interface {
  Bad2
}
type Bad2 interface {
  Bad1
}
```

Exemplo:

```go
package main

import (
  "fmt"
)

type R struct {
  R string
}

type Iread interface {
  Read() string
}

func (r *R) Read() string {
  return fmt.Sprintf("Only: call Read")
}

func Call(ir Iread) string {
  return fmt.Sprintf("Read: %s", ir.Read())
}

func main() {
  var iread Iread
  r := R{"hello interface"}
  // A way to use Interface
  iread = &r
  fmt.Println(iread, r)
  fmt.Println(iread.Read())

  // Second way to access interface
  r2 := R{"hello interface call"}
  fmt.Println(Call(&r2))
}
```

Saída:

```bash
&{hello interface} {hello interface}
Only: call Read
Read: Only: call Read
```

####  Interface como tipo

Interfaces como tipo __ interface {} __ significa que você pode colocar valor de qualquer tipo, incluindo seu próprio tipo personalizado. Todos os tipos em Go satisfazem uma interface vazia (interface {} é uma interface vazia).
No seu exemplo, o campo Msg pode ter valor de qualquer tipo.


```go
var val interface{} // element type of m is assignable to val
``` 

```go
type Empty interface {
    /* it has no methods */
}

// Because, Empty interface has no methods, 
// following types satisfy the Empty interface
var a Empty

a = 60
a = 10.5
a = "Lambda Man"
```

Interfaces como tipos, veja outro exemplo abaixo:

```go
package main

import (
  "fmt"
)

type MyStruct struct {
  Msg interface{}
}

func main() {
  b := MyStruct{}
  // string
  b.Msg = "5"
  fmt.Printf("%#v %T \n", b.Msg, b.Msg) // Output: "5" string

  // int
  b.Msg = 5
  fmt.Printf("%#v %T", b.Msg, b.Msg) //Output:  5 int

  // map
  b.Msg = map[string]string{"population": "500000", "language": "sueco"}
  fmt.Printf("%#v %T", b.Msg, b.Msg) //Output:  5 int
}
```

### Exercício 1

Exercício:
Preencha o struct JsonMessage AWS acima, inicialize a estrutura e preencha os campos, e faça um fmt.Println para exibir os campos preenchidos.
Para ser mais legível, você pode separar em cada estrutura do tipo struct.

## Estruturas de Controle
---

### Ao Controle

As estruturas de controle são:

__For, If, else, else if__

E algumas declarações entre elas: __break, continue, switch, case and goto__.

#### Controle Return

Declarações controlam a execução.

```bash
Statement =
  Declaration | LabeledStmt | SimpleStmt |
  GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
  FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
  DeferStmt .

SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
```

Uma instrução final impede a execução de todas as instruções que aparecem lexicalmente após o mesmo bloco. As seguintes declarações estão terminando:

1. Uma instrução "return" ou "goto".

Return:

```go
package main

func main() {
  println(Lambda())
  return
}

func Lambda() string {

  return "Lambda"
}
```

Saída:

```bash
Lambda
```

#### Controle Goto

Goto:

```go
package main

import "fmt"

func main() {
  n := 0

LOOP1:
  n++
  if n == 10 {
    println("fim")
    return
  }

  if n%2 == 0 {
    goto LOOP2
  } else {

    fmt.Println("n", n, "LOOP1 here...")
    goto LOOP1
  }

LOOP2:
  fmt.Println("n", n, "LOOP2 here...")
  goto LOOP1

}
```

Saída:

```bash
n 1 LOOP1 here...
n 2 LOOP2 here...
n 3 LOOP1 here...
n 4 LOOP2 here...
n 5 LOOP1 here...
n 6 LOOP2 here...
n 7 LOOP1 here...
n 8 LOOP2 here...
n 9 LOOP1 here...
fim
```

#### Control if Else

2. Uma declaração "if" em que:
      - a ramificação "else" está presente, e
      - ambas as ramificações são declarações finais.
      
```go
package main

func main() {
  n := 100
  if n > 0 && n <= 55 {
    println("n > 0 or n <= 55")
  } else if n > 56 && n < 70 {
    println("n > 56 and n < 70")
  } else {

    if n >= 100 {
      println(" else here.. n > 100")
    } else {
      println(" else here.. n > 70")
    }
  }
}
```

Saída:

```bash
else here.. n > 100
```

#### Control For Break Continue

3. Uma declaração "for" em que:
      - não há declarações de "break" referentes à declaração "for" e
      - a condição de loop está ausente.
      - existem "continue"
      - Uma instrução "break"  termina a execução da instrução "for", "switch" ou "select" mais interna dentro da mesma

```go
package main

func main() {
  // will be looping infinitely
  // for {
  // }

  // will run only once and exit
  for {
    break
  }

  n := 5
  for n > 0 {
    n--
    println(n)
  }
  // Output:
  // 4
  // 3
  // 2
  // 1
  // 0

  // declaring i no and increasing i
  for i := 0; i < 5; i++ {
    println(i)
  }
  // Output:
  // 0
  // 1
  // 2
  // 3
  // 4

  n = 5
  for i := 0; i < n; i++ {
    if i <= 2 {
      continue
    } else {
      println("i > 2 = ", i)
    }
  }

  // Output:
  // i > 2 =  3
  // i > 2 =  4

  n = 5
  for i := 0; i < n; i++ {
    if i == 2 {
      break
    } else {
      println("i: ", i)
    }
  }
  // Output:
  // i:  0
  // i:  1
  
  // infinitely
  for ; ; i++ {
    println("i: ", i)
  }
  // Output:
  // i:  1
  // i:  2
  // ..
  // ..
}
```

#### Control Switch Case Break

4. Uma declaração "switch" em que:
      - não há declarações de "break" referentes à declaração "switch",
      -existe um "case" padrão e
      - as listas de instruções em cada caso, incluindo o padrão, terminam em uma instrução final, ou uma declaração "fallthrough" possivelmente rotulada.

```go
package main

func main() {
  j := 10
  i := 0
  switch j {
  case 11:
    println("here: 11")
    break
  default:
    println("here default")
    break
  }

  // infinitely
  for ; ; i++ {

    switch i {
    case 5:
      goto LABELS
    case i:
      println("i: ", i)
      break
    default:
      println("default: ", i)
    }
  }

LABELS:
  f()

}

func f() {
  println("goto fim")
}
```

Saída:

```bash
here default
i:  0
i:  1
i:  2
i:  3
i:  4
goto fim
```

#### Etiqueta de Controle

5. Uma instrução rotulada rotulando uma instrução final.

```go
package main

func main() {
  i := 0
  // infinitely
  for ; ; i++ {
    for {
      if i == 10 {
        goto LABEL
      }
      i++
    }
  }

LABEL:
  f(i)

}

func f(i int) {
  println("label fim i:", i)
}
```

Saída:

```bash
label fim i: 10
```

Todas as outras declarações não terminam.

Uma lista de instruções termina em uma instrução final se a lista não estiver vazia e sua instrução final não vazia estiver sendo finalizada. 

Uma instrução "for" com uma cláusula "range" itera todas as entradas de uma matriz, fatia, string ou mapa ou valores recebidos em um canal. Para cada entrada, ele atribui valores de iteração a variáveis de iteração correspondentes, se presentes, e, em seguida, executa o bloco.

```bash
RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .
```

#### Controle Range

A expressão à direita na cláusula "range" é chamada de expressão de intervalo, que pode ser uma matriz, ponteiro para uma matriz, fatia, string, mapa ou canal que permite operações de recebimento. Como em uma atribuição, se presentes, os operandos à esquerda devem ser endereçáveis ou expressões de índice de mapa; eles denotam as variáveis de interação. Se a expressão de intervalo for um canal, no máximo uma variável de iteração será permitida, caso contrário, pode haver até duas. Se a última variável de iteração for o identificador em branco, a cláusula range é equivalente à mesma cláusula sem esse identificador.

```bash
Range expression                          1st value          2nd value

array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
string          s  string type            index    i  int    see below  rune
map             m  map[K]V                key      k  K      m[k]       V
channel         c  chan E, <-chan E       element  e  E
```

Veja um exemplo abaixo, com vários usos usando o Range:

```go

package main

import "fmt"

func main() {

  // string arrays / slice
  var lang = [...]string{"Erlang", "Elixir", "Haskell", "Clojure", "Scala"}

  // screen list
  fmt.Println(lang)

  // list the positions srtring arrays
  for k, v := range lang {
    fmt.Println(k, v)
  }

  /* create a map*/
  countryCapitalMap := map[string]string{"Brasil": "Brasilia", "EUA AMERICA": "Washington, D.C.", "France": "Paris", "Italy": "Roma", "Japan": "Tokyo"}

  /* print map using key-value*/
  for country, capital := range countryCapitalMap {
    fmt.Println("Capital of", country, "is", capital)
  }

  // channel
  jobs := make(chan int, 3)

  // for channel
  for j := 1; j <= 3; j++ {
    jobs <- j
  }
  // println(<-jobs)
  // println(<-jobs)
  // println(<-jobs)

  // close
  /* date is required for range to work*/
  close(jobs)

  /* This syntax is valid too. */
  for range jobs {
  }

  /* it is mandatory to close the channels to be able to scroll */
  for ch := range jobs {
    println(ch)
  }

  // it is not an array struct, it will range from error.
  sa := struct{ nick string }{"@jeffotoni"}
  fmt.Println(sa.nick)

  // here the range will be able to list all struct
  a := []struct{ nick string }{{"@devopsbr"}, {"@go_br"}, {"@awsbrasil"}, {"@go_br"}, {"@devopsbh"}}
  for i, v := range a {
    fmt.Println(i, v.nick)
  }

  // struct pointer
  var testdata *struct {
    a *[3]int
  }
  for i := range testdata.a {
    // testdata.a is never evaluated; len(testdata.a) is constant
    // i ranges from 0 to 2
    fmt.Println(i)
  }

  // new example interface and range
  var key string
  var val interface{} // element type of m is assignable to val
  m := map[string]int{"mon": 0, "tue": 1, "wed": 2, "thu": 3, "fri": 4, "sat": 5, "sun": 6}
  for key, val = range m {
    fmt.Println(key, val)
  }
}
```

Saída:

```bash
[Erlang Elixir Haskell Clojure Scala]
0 Erlang
1 Elixir
2 Haskell
3 Clojure
4 Scala
Capital of Brasil is Brasilia
Capital of EUA AMERICA is Washington, D.C.
Capital of France is Paris
Capital of Italy is Roma
Capital of Japan is Tokyo
@jeffotoni
0 @devopsbr
1 @go_br
2 @awsbrasil
3 @go_br
4 @devopsbh
0
1
2
sat 5
sun 6
mon 0
tue 1
wed 2
thu 3
fri 4
```
## Erros
---
 
 O erro de tipo pré-declarado é definido como

 ```bash
 type error interface {
	Error() string
}
 ```
É a interface convencional para representar uma condição de erro, com o valor inexistente representando nenhum erro. Por exemplo, uma função para ler dados de um arquivo pode ser definida:

```bash
func Read(f *File, b []byte) (n int, err error)
```

### Introdução aos Erros

O tratamento de erros em Golang é muito simples de lidar. Não há tentativa, captura ou exceções.
O erro é tratado com todas as chamadas de alguma função.
Como mostramos "error" é uma interface e muito utilizada.

Quando falamos em lidar com erros em Golang tudo é muito simples, uma boa prática é retornar um erro nas funções que criamos e tratá-las.

Vamos ver na prática como isso funciona.

```go
package main

import "fmt"

func main() {
	var error error
	fmt.Println(error)
}
```

Saída:

```bash
<nil>
```
#### Como funciona o controle de erros

Exemplo:

```go
package main

import (
	"encoding/json"
	"fmt"
)

var Error error
var b []byte

func main() {
	fmt.Println(Error)
	cpu := make(chan int)
	// Create JSON from the instance data.
	b, Error = json.Marshal(cpu)
	if Error != nil {
		fmt.Println(Error)
	} else {
		fmt.Println(string(b))
	}
}
```

Saída:

```bash
<nil>
json: unsupported type: chan int
```
#### Novos Erros

```go
package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(fvalue float64) (float64, error) {
	if fvalue < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt [" + fmt.Sprintf("%.2f", fvalue) + "]")
	}
	return math.Sqrt(fvalue), nil
}

func main() {
	result, err := Sqrt(-33)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sqrt (-33) = [", result, "]")
	}

	result, err = Sqrt(81)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sqrt(81) = [", result, "]")
	}
}
```
Saída:

```bash
Math: negative number passed to Sqrt [-33.00]
Sqrt(81) = [ 9 ]
```
#### Erros Personalizados

```go
package main

import "fmt"

// It's possible to use custom types as `error`s by
// implementing the `Error()` method on them.
type MyError struct {
    line int
    msg  string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("%d - %s", e.line, e.msg)
}

func MyFunc(line int) (int, error) {
    if line < 100 {

        // In this case we use `&MyError` syntax to build
        // a new struct, supplying values for the two
        // fields `arg` and `prob`.
        return -1, &MyError{line, "can't work with it"}
    }
    return line, nil
}

func main() {

    // The one loops below test out each of our
    // error-returning functions.
    for _, i := range []int{200, 99} {
        if r, e := MyFunc(i); e != nil {
            fmt.Println("MyFunc failed:", e)
        } else {
            fmt.Println("MyFunc worked in line: ", r)
        }
    }
}
```

Saída:

```bash
MyFunc worked in line:  200
MyFunc failed: 99 - can't work with it
```

#### fmt Errorf

```go
package main

import (
    "fmt"
    "math"
)

func circleArea(radius float64) (float64, error) {
    if radius < 0 {
        return 0, fmt.Errorf("Failed, radius %0.2f is less than zero", radius)
    }
    return math.Pi * radius * radius, nil
}

func main() {
    radius := -80.0
    area, err := circleArea(radius)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Area of circle: %0.2f", area)
}
```
Saída:

```bash
Area calculation failed, radius -80.00 is less than zero
```

## Funções
---

Declarando e Chamando Funções em Golang. 
Em Golang, declaramos uma função usando a palavra-chave func. 
Uma função tem um nome, uma lista de parâmetros de entrada separados por vírgula junto com seus tipos, o (s) tipo (s) de resultado e um corpo. 
Os parâmetros de entrada e os tipos de retorno são opcionais para uma função.

Exemplo de declarar e chamar funções em Golang:

```go
func Sum(x float64, y float64) float64 {
	return (x + y) / 2
}
```

### Introdução as Funções

Go exige retornos explícitos, ou seja, não retorna automaticamente o valor da última expressão.
Quando você tem vários parâmetros consecutivos do mesmo tipo, você pode omitir o nome do tipo para os parâmetros do tipo semelhante até o parâmetro final que declara o tipo.
Um tipo de função denota o conjunto de todas as funções com os mesmos tipos de parâmetros e resultados. O valor de uma variável não inicializada do tipo de função é nulo.

Algumas possibilidades:

```bash
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
```

```go
package main

func F1(name string) string {
    return "Hello, " + name
}

func main() {

    println(F1("@go_br"))
}
```

Saída:

```bash
Hello, @go_br
```

#### Retornar vários valores

Go tem suporte embutido para vários valores de retorno. Esse recurso é usado frequentemente no Go idiomático, por exemplo, para retornar valores de resultado e erro de uma função.

```go
package main

import "fmt"

// The `(int, int)` in this function signature shows that
// the function returns 2 `int`s.
func F1() (string, int, error) {
    return "@go_br", 100, nil
}

func main() {

    // Here we use the 2 different return values from the
    // call with _multiple assignment_.
    a, b, err := F1()
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(err)

    // If you only want a subset of the returned values,
    // use the blank identifier `_`.
    a, _, err = F1()
    fmt.Println(a)
    fmt.Println(err)
}
```

Saída:

```bash
@go_br
100
<nil>
@go_br
<nil>
```

#### Funções Variadic

Funções Variadic podem ser chamadas com qualquer número de argumentos à direita. Por exemplo, fmt.Println é uma função variadic comum.
Aqui está uma função que levará um número arbitrário de ints como argumentos.
Funções variadic podem ser chamadas da maneira usual com argumentos individuais.

```go
package main

import (
	"fmt"
	"strings"
)

// This variadic function takes an arbitrary number of ints as arguments.
func Show(names ...string) {
	fmt.Print("The Len of ", len(names)) // Also a variadic function.
	appends := ""
	for _, name := range names {
		appends += name + ","
	}
	appends = strings.Trim(appends, ",")
	fmt.Println(" is: [", appends, "]") // Also a variadic function.
}

func main() {

	// Variadic functions can be called in the usual way with individual
	// arguments.
	Show("C", "C++")
	Show("Clojure", "Elixir", "Scala")

	// If you already have multiple args in a slice, apply them to a variadic
	// function using func(slice...) like this.
	nums := []string{"Algol", "C", "C++", "Golang"}
	Show(nums...)
}
```

Saída:

```bash
The Len of 2 is: [ C,C++ ]
The Len of 3 is: [ Clojure,Elixir,Scala ]
The Len of 4 is: [ Algol,C,C++,Golang ]
```

#### Funções como um parâmetro

Você pode passar a função como parâmetro para uma função Go. Aqui está um exemplo de função de passagem como parâmetro para outra função Go.

```go
package main

import "fmt"

// convert types take an int
// and return a string value.
type fn func(int) string

func f1(param int) string {
	return fmt.Sprintf("param is %v", param)
}

func f2(param int) string {
	return fmt.Sprintf("param is %v", param)
}

func test(f fn, val int) {
	fmt.Println(f(val))
}

func main() {
	test(f1, 432)
	test(f2, 874)
}
```

Saída:

```bash
param is 432
param is 874
```

```go
package main

import "fmt"

// --------------------------------
func Square(num int) int {
	return num * num
}

func Mapp(f func(int) int, List []int) []int {
	var a = make([]int, len(List), len(List))
	for index, val := range List {

		a[index] = f(val)
	}
	return a
}

func main() {
	list := []int{454, 455, 86, 988}
	result := Mapp(Square, list)
	fmt.Println(result)
}
```

saída:

```bash
[206116 207025 7396 976144]
```

#### Fechamentos

Go suporta funções anônimas, que podem formar fechamentos. Funções anônimas são úteis quando você deseja definir uma função inline sem precisar nomeá-la.
Esta função intSeq retorna outra função, que definimos anonimamente no corpo do intSeq. A função retornada se fecha sobre a variável i para formar um fechamento.

Exemplo:

```go
package main

import "fmt"

func PlusX() func(v int) int {
    return func(v int) int {
        return v + 5
    }
}

func plusXandY(x int) func(v int) int {
    return func(v int) int {
        return v + x
    }
}

func main() {
    p := PlusX()
    fmt.Printf("5+15: %d\n", p(15))

    px := plusXandY(6)
    fmt.Printf("6+10: %d\n", px(10))
}
```

Saída:

```bash
5+15: 20
6+10: 16
```

#### Recursão

Go suporta funções recursivas. Aqui está um exemplo fatorial clássico.

Um exemplo simples:

```go
package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
```

Saída:

```bash
5040
```

Listando todos os diretórios de subdiretórios:

```go
package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// Listing our example directory
// Listing our example directory recursively
func main() {

	// Capturing our path that is in the environment
	gopath := os.Getenv("PWD")

	// directory we want to list
	gopath += "/examples"

	// Making call in function
	list := ListDir(gopath)

	// listing the function return
	for i, p := range list {
		fmt.Printf("[%d:%s===%s]\n", i, path.Dir(p), path.Base(p))
	}
}

// This function uses pkg filepath.Walk, it is
// a recursive function, where it will go through
// our directory and its subfolders.
func ListDir(rootpath string) []string {

	list := make([]string, 0)

	// recursive call
	// This function receives a function as parameter and after going through all levels it ends.
	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".git" && filepath.Ext(path) != ".svn" {
			list = append(list, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return list
}
```

Saída:

```bash
[0:$(pwd)/examples/bufio.writer===main.go]
[1:$(pwd)/examples/error===error1.go]
[2:$(pwd)/examples/error===error2.go]
[3:$(pwd)/examples/error===error3.go]
[4:$(pwd)/examples/error===error4.go]
...
...
```

#### Funções Assíncronas

No golang para executar funções assíncronas, usamos a palavra-chave **"go"**, que é responsável por colocar as funções a serem executadas simultaneamente.
Uma instrução **"go"** inicia a execução de uma chamada de função como um encadeamento de controle concorrente independente, ou goroutine, dentro do mesmo espaço de endereço. 

```go
GoStmt = "go" Expression .
```

```go
go Server()
go func(ch chan<- bool) { for { sleep(10); ch <- true }} (c)
```

Exemplo:

```go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.Intn(500)

func pinger() {
	time.Sleep(time.Duration(r) * time.Microsecond)
	fmt.Println("pinger")
}

func ponger() {
	time.Sleep(time.Duration(r) * time.Microsecond)
	fmt.Println("ponger")
}

func printer() {
	time.Sleep(time.Duration(r) * time.Microsecond)
	fmt.Println("printer")
}

func main() {
	// making functions
	// calls asynchronously
	go pinger()
	go ponger()
	go printer()

	// Waiting to press any key to end
	var input string
	fmt.Scanln(&input)
}
```
Saída 1:

```bash
ponger
pinger
printer
```

saída 2:

```bash
pinger
ponger
printer
```

## Defer

Uma "defer" invoca uma função cuja execução é adiada para o momento em que a função circundante retorna, seja porque a função circundante executou uma instrução de retorno, atingiu o fim de seu corpo de função ou porque a gorout correspondente está em pânico.

```go
DeferStmt = "defer" Expression .
```

A expressão deve ser uma chamada de função ou método; não pode ser entre parênteses. Chamadas de funções internas são restritas como para instruções de expressão.

Cada vez que uma instrução de "defer" é executada, o valor da função e os parâmetros da chamada são avaliados como de costume e salvos novamente, mas a função real não é invocada. Em vez disso, as funções diferidas são chamadas imediatamente antes da função circundante retornar, na ordem inversa em que foram adiadas. Se um valor de função diferido for avaliado como nulo, a execução entra em pane quando a função é invocada, e não quando a instrução "defer" é executada. 

Exemplos:

```go
defer unlock(l)
defer myFunc()
defer close(channel)
defer fmt.Print(x)
defer db.Close()
defer f.Close()
defer res.Body.Close()
```
## Exercício 01 

### Módulo 01

Utilizando Go e testes unitários você deverá determinar os dez maiores estados brasileiros em extensão territorial.

### Tópicos

Neste desafio você aprenderá:

- Go
- Testes unitários

### Requisitos
Para este desafio você precisará de:

- Go versão 1.9 (ou superior)
- Git
