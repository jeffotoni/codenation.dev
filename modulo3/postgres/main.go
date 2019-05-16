// ORM EM GO: http://xorm.io
// POSTGRES
//

package main

import (
	"database/sql"
	"fmt"
	"github.com/github.com/jeffotoni/codenation.dev/modulo3/postgres/repo/rpg"
	"os"
)

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

	//ok, msg := rpg.Select(codemp)
	ok, msg := rpg.Select2(db, codemp)
	fmt.Println("ok:: ", ok)
	fmt.Println("msg:: ", msg)
}
