package main

import "fmt"

func diaDaSemana1(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabádo"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	case numero == 7:
		diaDaSemana = "Sabádo"
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana1(3)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
