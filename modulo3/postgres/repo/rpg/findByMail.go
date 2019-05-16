package rpg

import (
	"database/sql"
)

func FindLoginByEmail(db *sql.DB, email string) (*Login, error) {
	u := &Login{}

	err := db.QueryRow(`
		SELECT id, name, email, created_at
		FROM login
		WHERE email = $1
	`, email).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}
