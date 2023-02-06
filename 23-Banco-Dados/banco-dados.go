package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConn := "root:1234@/devbook?charset=utf8&parseTime=true&loc=Local"
	db, erro := sql.Open("mysql", stringConn)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão aberta!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}