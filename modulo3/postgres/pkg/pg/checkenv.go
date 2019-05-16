package pg

func PgCheckEnv() bool {
	msg := "Config Database é obrigatório!"
	if len(DB_HOST) == 0 {
		println(msg)
		println("DB_HOST é obrigatorio")
		return false
	}

	if len(DB_USER) == 0 {
		println(msg)
		println("DB_USER é obrigatorio")
		return false
	}

	if len(DB_PASSWORD) == 0 {
		println(msg)
		println("DB_PASSWORD é obrigatorio")
		return false
	}

	if len(DB_BANCO) == 0 {
		println(msg)
		println("DB_BANCO é obrigatorio")
		return false
	}

	if len(DB_PORT) == 0 {
		println(msg)
		println("DB_PORT é obrigatorio")
		return false
	}

	return true
}
