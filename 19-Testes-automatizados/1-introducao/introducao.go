package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("Av. Paulista")

	fmt.Println(tipoEndereco)
}
