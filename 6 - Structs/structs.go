package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Luciano"
	u.idade = 25
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	usuario2 := usuario{"Luciano", 25, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Luciano"}
	fmt.Println(usuario3)
}
