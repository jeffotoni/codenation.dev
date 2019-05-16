// Back-End in Go server
// @jeffotoni
// 2019-01-04

package rpg

import (
	"database/sql"
	"github.com/github.com/jeffotoni/codenation.dev/modulo3/postgres/pkg/pg"
	//"log"
)

func Insert(cnpj, razaosocial string, ativo bool) (ok bool, msg string) {

	ok = false
	msg = "sucesso"

	if len(cnpj) <= 0 {
		return ok, "campo cnpj obrigatorio"
	}

	if len(razaosocial) <= 0 {
		return ok, "campo razaosocial obrigatorio"
	}

	// connect
	var Db = pg.PostDb.Pgdb

	// Db...
	if interf := pg.Connect(); interf != nil {
		Db = interf.(*sql.DB)
	} else {
		return ok, "Error ao conectar na base de dados!"
	}

	insert := `INSERT INTO empresa(cnpj,razaosocial,ativo)values($1,$2,$3)`
	tx, err := Db.Begin()
	if err != nil {
		//log.Println(err)
		return ok, err.Error()
	}

	defer tx.Rollback()
	stmt, err := tx.Prepare(insert)
	if err != nil {
		//log.Println(err)
		return ok, err.Error()
	}

	defer stmt.Close() // danger!

	_, err = stmt.Exec(cnpj, razaosocial, ativo)
	if err != nil {
		//log.Println(err)
		return ok, err.Error()
	}
	err = tx.Commit()
	if err != nil {
		//log.Println(err)
		return ok, err.Error()
	}

	return true, msg
}
