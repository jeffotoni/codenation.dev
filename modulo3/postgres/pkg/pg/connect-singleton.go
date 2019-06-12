//
// conexao com postgresql
//
package pg

///
import (
	"database/sql"
	"fmt"
	"log"
	//"strings"
	"sync"

	_ "github.com/lib/pq"
)

var once sync.Once
var dbLocal *sql.DB

var pchan = make(chan string)

type PgStruct struct {
	Pgdb *sql.DB
}

type StatusMsg struct {
	Msg string `json:msg`
	Db  string `json:db`
}

// cache sync.Map
type cache struct {
	mm sync.Map
	sync.Mutex
}

var (
	err    error
	PostDb PgStruct
)

var (
	pool = &cache{}
)

// put sync.Map
func (c *cache) put(key, value interface{}) {

	c.Lock()
	defer c.Unlock()
	c.mm.Store(key, value)
}

// get sync.Map
func (c *cache) get(key interface{}) interface{} {

	c.Lock()
	defer c.Unlock()

	v, _ := c.mm.Load(key)
	return v

}

// setLoad... fn func() interface{}
func (c *cache) loadStore(key interface{}, fc func() (interface{}, error)) (interface{}, error) {

	c.Lock()
	defer c.Unlock()

	if v, ok := c.mm.Load(key); ok {
		return v, nil
	}

	// treat or error
	val, err := fc()

	// error return
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	c.mm.Store(key, val)
	return val, nil
}

// conectando de forma segura usando goroutine
func Connect() interface{} {

	if dbPg := pool.get(DB_BANCO); dbPg != nil {

		// return objeto conexao
		return dbPg.(*sql.DB)

	} else {

		if !PgCheckEnv() {
			return nil
		}

		// removendo aspas..
		// DB_BANCO = strings.Replace(DB_BANCO, `"`, "", -1)

		DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_BANCO, DB_SSL)

		// println(DBINFO)
		// func for execute
		// in loadStore
		// when two or more
		// goroutine at moment
		fn := func() (interface{}, error) {

			once.Do(func() {
				PostDb.Pgdb, err = sql.Open(DB_SORCE, DBINFO)
			})

			if err != nil {
				log.Println(err.Error())
				return nil, err
			}

			if ok2 := PostDb.Pgdb.Ping(); ok2 != nil {
				log.Println("connect error...: ", ok2)
				return nil, err
			}

			//log.Println("connect return sucess:: client [" + DB_BANCO + "]")
			return PostDb.Pgdb, nil
		}

		// get connect
		// load cache loadStore
		sqlDb, err := pool.loadStore(DB_BANCO, fn)

		if err != nil {
			// error
			return nil
		}

		if sqlDb != nil {
			return sqlDb.(*sql.DB)

		} else {
			return nil
		}
	}
}

// singleton versao
// resumida e eficiente
func ConnectNew() (db *sql.DB) {

	if dbLocal != nil {
		return dbLocal

	} else {

		DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_BANCO, DB_SSL)

		once.Do(func() {
			dbLocal, err = sql.Open(DB_SORCE, DBINFO)
		})

		if err != nil {
			log.Println("Erro ao tentar conectar Postgres: ", err)
			return dbLocal
		}

		return dbLocal
	}
}
