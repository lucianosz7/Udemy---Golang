package enderecos

import "strings"

func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoLetraMinuscula, " ")[0]

	enderecoTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoTipoValido = true
		}
	}

	if enderecoTipoValido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo inválido"
}
