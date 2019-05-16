package pg

import (
	"os"
)

const LayoutDateLog = "2006-01-02 15:04:05"
const LayoutDate = "2006-01-02"
const LayoutHour = "15:04:05"

var (
	DB_HOST     = os.Getenv("DB_HOST")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_BANCO    = os.Getenv("DB_BANCO")
	DB_PORT     = os.Getenv("DB_PORT")

	DB_SSL   = "disable"
	DB_SORCE = "postgres"
)
