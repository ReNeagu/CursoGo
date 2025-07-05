package enderecos

import "strings"

//TipoDeEndereco, verifica se o tipo é valido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0] //Separa as frases com o ESPAÇO " ", para contabilizar.. o numero anterior marca a posição que estamos procurando

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}
	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}
