// ORM EM GO: http://xorm.io
// ORM EM GO: http://gorm.io
// POSTGRES

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	//"github.com/github.com/jeffotoni/codenation.dev/modulo3/postgres/repo/rpg"
	"os"
	"time"
)

// CREATE TABLE
//  id uuid DEFAULT public.gen_random_uuid() NOT NULL,
//  password text NOT NULL,
//  email text NOT NULL,
//  name text NOT NULL,
//  created_at timestamp

/// FICAR NO  MODEL
type Login struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdat"`
}

// FICAR NO CONFIG
// PKG/PG
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
	//db2 *sql.DB
	err error
)

// PKG/PG
func connect() (db *sql.DB) {

	// fmt.Println("init start")
	DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_BANCO, DB_SSL)

	println(DBINFO)

	db, err = sql.Open(DB_SORCE, DBINFO)
	// error
	if err != nil {
		fmt.Println("open: ", err)
		return
	}

	// defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("ping: ", err)
		return
	}

	println("connected!")
	return
}

// REPO
func CreateLogin(db *sql.DB, name, email, password string) (*Login, error) {

	u := &Login{}

	err := db.QueryRow(`
		INSERT INTO login(name, email, password)
		VALUES($1, $2, crypt($3, gen_salt('bf', 8)))
		RETURNING id, name, email, created_at
	`, name, email, password).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func main() {

	db2 := connect()

	name := "GiovaniAbel"
	email := "giovani@gm.com"
	password := "1234"

	l, err := CreateLogin(db2, name, email, password)

	if err != nil {
		fmt.Println("err create:: ", err)
		return
	}

	//fmt.Println(l)
	fmt.Println(l.ID)
	fmt.Println(l.Name)
	fmt.Println(l.Email)
	// insert
	// delete
	// select
	// begin
	fmt.Println("main start", db2)
}
