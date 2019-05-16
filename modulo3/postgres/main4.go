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
  orm = &rpg.LoginMapper{}
  err error
)

func init() {

  DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
    DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_BANCO, DB_SSL)

  // println(DBINFO)
  orm.DB, err = sql.Open(DB_SORCE, DBINFO)
  if err != nil {
    // log.Println(err)
    return
  }
  //defer db.Close()

  err = orm.DB.Ping()
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

  l, err := orm.FindLoginByEmailMapper(email)

  fmt.Println(err)
  fmt.Println(l)
  fmt.Println("id: ", l.ID)
  fmt.Println("name: ", l.Name)
  fmt.Println("email", l.Email)

}
