package rpg

import "database/sql"

// CreateUser inserts a user record into the database and returns a User object
func (m *LoginMapper) FindLoginByEmailMapper(email string) (*Login, error) {

	u := &Login{}

	err := m.DB.QueryRow(`
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
