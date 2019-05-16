package main

import (
  "fmt"
  "github.com/github.com/jeffotoni/codenation.dev/modulo3/postgres/repo/rpg"
)

func main() {

  var codemp int64 = 2
  // var cnpj, razaosocial string
  // var ativo bool

  // cnpj = "08776968000472"
  // razaosocial = "empresa codenatio three"
  // ativo = true

  // // // vamos inserir
  // ok, msg := rpg.Insert(cnpj, razaosocial, ativo)

  // fmt.Println("ok:: ", ok)
  // fmt.Println("msg:: ", msg)

  ok, msg := rpg.Select(codemp)
  fmt.Println("ok:: ", ok)
  fmt.Println("msg:: ", msg)
}
