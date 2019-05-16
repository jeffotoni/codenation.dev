# Makefile

# postgres in docker
# export DB_HOST := 192.0.0.3
# export DB_USER := aluno
# export DB_PASSWORD := 1234
# export DB_BANCO := aluno
# export DB_PORT := 5432

# postgres localhost
export DB_HOST := localhost
export DB_USER := aluno
export DB_PASSWORD := 12345
export DB_BANCO := aluno
export DB_PORT := 5432

all:
	go run main.go

all2:
	go run main2.go

all3:
	go run main3.go

all4:
	go run main4.go

aula1:
	go run main-aula1.go

build:
	go build -o pgtest 
	./pgtest
	rm -f pgtest

test:
	go test -v $(PWD)/repo/rpg