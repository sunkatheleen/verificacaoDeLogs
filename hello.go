package main

import (
	"fmt"
	"os"
)

// "reflect"

func main() {

	exibeIntrodução()
	exibeMenu()

	nome := "Katheleen"
	versão := 1.1
	idade := 19

	fmt.Println("Olá Sra.", nome, "sua idade é", idade, "anos")
	fmt.Println("Este programa está na versão", versão)

	// fmt.Println("O tipo de variável versão é", reflect.TypeOf(versão))

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")

	comando := leComando()

	//var comando int
	// fmt.Scanf("%d", &comando) onde %d representa um número inteiro
	//fmt.Scan(&comando)
	// fmt.Println("O endereço da minha variável comando é", &comando)
	//fmt.Println("O comando escolhido foi", comando)

	// if comando == 1 {
	// fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// fmt.Println("Exibindo logs")
	// } else if comando == 0 {
	// fmt.Println("Saindo do programa")
	// } else {
	// fmt.Println("Não reconheço este comando")
	// }

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não reconheço este comando")
		os.Exit(-1)

	}

}

func exibeIntrodução() {

	nome := "Katheleen"
	versão := 1.1
	idade := 19

	fmt.Println("Olá Sra.", nome, "sua idade é", idade, "anos")
	fmt.Println("Este programa está na versão", versão)
}

func exibeMenu() {

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")

}

func leComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido

}
