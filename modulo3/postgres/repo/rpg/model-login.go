package rpg

import "time"

type Login struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}
