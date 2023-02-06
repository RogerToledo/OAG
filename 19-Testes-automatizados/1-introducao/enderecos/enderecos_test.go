package enderecos

import "testing"

func TestTipoEndereco(t *testing.T) {

	var tests = []struct {
		descricao string
		esperado string
		recebido string
	}{
		{
			descricao: "Deve retornar Rua",
			esperado: "Rua",
			recebido: TipoEndereco("Rua arujá"),
		},
		{
			descricao: "Deve retornar Praça",
			esperado: "Praça",
			recebido: TipoEndereco("PRAÇA dos Bandeirantes"),
		},
		{
			descricao: "Deve retornar Avenida",
			esperado: "Avenida",
			recebido: TipoEndereco("Avenida Paulista"),
		},
		{
			descricao: "Deve retornar Tipo inválido quando parâmetro for Av.",
			esperado: "Tipo inválido",
			recebido: TipoEndereco("Av. Paulista"),
		},
		{
			descricao: "Deve retornar Tipo inválido quando parâmetro endereço for vazio",
			esperado: "Tipo inválido",
			recebido: TipoEndereco(""),
		},
	}
		
	for _, test := range tests {
		t.Log(test.descricao)

		if test.esperado != test.recebido {
			t.Errorf("Tipo de endereço Esperado: %v - Recebido: %v", test.esperado, test.recebido)
		}
	}
}