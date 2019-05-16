package rpg

import (
	"database/sql"
	"fmt"
	//"github.com/github.com/jeffotoni/codenation.dev/modulo3/postgres/pkg/pg"
)

func Select2(Db *sql.DB, codemp int64) (ok bool, msg string) {

	msg = fmt.Sprintf("Validando User")
	ok = true

	// CONNECT 1
	// Db, errc := pg.Connect1()
	// if errc != nil {
	// 	fmt.Println("banco: ", errc)
	// 	return false, errc.Error()
	// }

	// CONNECT SINGLETON
	// var Db = pg.PostDb.Pgdb
	// // Db...
	// if interf := pg.Connect(); interf != nil {
	// 	Db = interf.(*sql.DB)
	// } else {
	// 	return false, "Error ao conectar na base de dados!"
	// }

	row := Db.QueryRow("select cnpj,razaosocial,ativo from public.empresa where codemp=$1", codemp)
	var cnpj int64
	var razaosocial string
	var ativo bool

	err := row.Scan(&cnpj, &razaosocial, &ativo)

	if err != nil && err != sql.ErrNoRows {
		// log the error
		msg = fmt.Sprintf("Error ao tentar selecionar empresa...%s", err)
		ok = false
		return ok, msg
	}

	msg = fmt.Sprintf("%d, %s, %t", cnpj, razaosocial, ativo)
	return
}
