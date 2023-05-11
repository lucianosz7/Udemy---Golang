package enderecos

import "testing"

// Teste de Unidade

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoEndereco(t *testing.T) {
	t.Parallel()

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida dos Imigrantes", "Avenida"},
		{"BR-050", "Rodovia"},
		/* {"Square of roses", "Tipo inválido"}, */
		{"Any road", "Estrada"},
		/* {"", "Tipo inválido"}, */
	}

	for _, cenario := range cenarioDeTeste {
		tipoEnderecoRecebido := TipoEndereco(cenario.enderecoInserido)
		if tipoEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
