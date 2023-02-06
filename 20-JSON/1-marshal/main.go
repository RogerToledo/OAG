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

	c := cachorro{"Rex", "Dálmata", 3 }
	fmt.Println(c)

	cachorroJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJSON)
	fmt.Printf(string(cachorroJSON))

}