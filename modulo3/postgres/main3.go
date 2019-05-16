// ORM EM GO: http://xorm.io
// POSTGRES
//

package main

import (
  "database/sql"
  "fmt"
  "github.com/github.com/jeffotoni/codenation.dev/modulo3/postgres/repo/rpg"
  "os"
  "time"
)

type Login struct {
  ID        string
  Name      string
  Email     string
  CreatedAt time.Time
}

var (
  DB_HOST     = os.Getenv("DB_HOST")
  DB_USER     = os.Getenv("DB_USER")
  DB_PASSWORD = os.Getenv("DB_PASSWORD")
  DB_BANCO    = os.Getenv("DB_BANCO")
  DB_PORT     = os.Getenv("DB_PORT")

  DB_SSL   = "disable"
  DB_SORCE = "postgres"
)

var (
  db  *sql.DB
  err error
)

func init() {

  DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
    DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_BANCO, DB_SSL)

  println(DBINFO)

  db, err = sql.Open(DB_SORCE, DBINFO)
  if err != nil {
    // log.Println(err)
    return
  }
  //defer db.Close()

  err = db.Ping()
  if err != nil {
    //log.Println(err)
    return
  }
  println("connected!")
  return
}

// Create a Login
// Find a Login by ID
// Find a Login by Email and Password

func main() {
  //name := "Petter"
  email := "petter@g.com"
  password := "1234"

  // l, err := rpg.CreateLogin(db, name, email, password)
  // fmt.Println(l)
  // fmt.Println(err)

  // l, err = rpg.FindLogin(db, "99d63be6-02ad-4c07-ac44-63e4660f0ae2")
  l, err := rpg.FindLoginByEmail(db, email)

  fmt.Println(l)
  fmt.Println(err)

  rpg.FindLoginByEmailPass(db, email, password)

  fmt.Println(l)
  fmt.Println(err)
}
