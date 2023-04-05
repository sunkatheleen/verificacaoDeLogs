package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	//"io/ioutil"
	"bufio"
	"net/http"
	"os"
	"strconv"
	"time"
)

const monitoramentos = 3
const delay = 10

// "reflect"

func main() {

	exibeIntrodução()
	//leSitesDoArquivo()

	for {

		exibeMenu()

		comando := leComando()

		//nome := devolveNome() criado apenas para fazer testes
		//fmt.Println(nome)

		// fmt.Println("O tipo de variável versão é", reflect.TypeOf(versão))

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
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não reconheço este comando")
			os.Exit(-1)

		}

	}

}

//func devolveNome() string {
//nome := "Katheleen"
//return nome
//}

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
	fmt.Println("")

	return comandoLido

}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://www.alura.com.br", "https://www.facebook.com/",
	//	"https://www.linkedin.com/feed/", "https://twitter.com/"}

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
		//fmt.Println(sites)

		//for i := 0; i < len(sites); i++ { escrita da forma tradicional

	}

	fmt.Println("")

}

func testaSite(site string) {

	resp, err := http.Get(site)
	//fmt.Println(resp)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		registraLog(site, true)
		fmt.Println("Site:", site, "Foi carregado com sucesso!")
	} else {
		fmt.Println("O site:", site, "não foi carregado com sucesso. Status Code",
			resp.StatusCode)
		registraLog(site, false)
	}

}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	//fmt.Println(string(arquivo))

	leitor := bufio.NewReader(arquivo)

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

		//if err != nil {
		//fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(sites)

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "-online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()

}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
