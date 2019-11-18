package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const statusOk = 200

func main() {

	exibeMenu()

}

func exibeMenu() {
	for {
		fmt.Println(`Me diga o que deseja: 
	1. Iniciar o Monitoramento
	2. Exibir logs
	0. Sair`)

		possiveisEscolhas(lendoEntradas())
	}
}

func lendoEntradas() int {

	var escolha int

	fmt.Scan(&escolha)

	return escolha

}

func possiveisEscolhas(escolha int) {

	switch escolha {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Os logs são ...")

	case 0:
		fmt.Println("Té mais")
		os.Exit(0)

	default:
		fmt.Println("Digitou errado, tente novamente")
		os.Exit(-1)

	}

}

func iniciarMonitoramento() {

	// sites := []string{"http://random-status-code.herokuapp.com/", "http://www.google.com/", "http://www.facebook.com/"}
	sites := leSitesDoArquivo()

	for _, site := range sites {

		testaSite(site)

	}
}

func testaSite(site string) {

	resp, err := http.Get(site)
	if err != nil {
		log.Panic(err)
	}

	if resp.StatusCode == statusOk {
		fmt.Printf("Tudo certo, o site %s está no ar %d\n", site, resp.StatusCode)

	} else {
		fmt.Printf("Deu ruim, site %s fora do ar %d\n", site, resp.StatusCode)
	}

}

func leSitesDoArquivo() []string {

	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {

		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}
	return sites
}
