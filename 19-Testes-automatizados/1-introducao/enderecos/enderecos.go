package enderecos

import "strings"

func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "praça"}

	enderecoMinusculo := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoMinusculo, " ")[0]

	tipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			tipoValido = true
		}
	} 

	if tipoValido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo inválido"
}