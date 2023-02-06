package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json: "nome"`
	Raca string `json: "raca"`
	Idade uint  `json: "idade"`
}

func main() {
	cachorroJSON := `{"Nome":"Rex","Raca":"Dálmata","Idade":3}`
	var c cachorro
	if erro := json.Unmarshal([]byte(cachorroJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)
}