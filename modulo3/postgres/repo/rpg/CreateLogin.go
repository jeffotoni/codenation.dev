package rpg

import (
	"database/sql"
)

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
