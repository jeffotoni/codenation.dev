//
// conexao com postgresql
//
package pg

///
import (
	"database/sql"
	"fmt"
	//"log"
	//"strings"
	//"sync"

	_ "github.com/lib/pq"
)

func Connect1() (db *sql.DB, err error) {

	DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_BANCO, DB_SSL)

	println(DBINFO)

	db, err = sql.Open(DB_SORCE, DBINFO)
	if err != nil {
		//log.Println(err)]
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
