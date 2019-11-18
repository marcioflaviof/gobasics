package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
		imprimirLogs()

	case 0:
		fmt.Println("Té mais")
		os.Exit(0)

	default:
		fmt.Println("Digitou errado, tente novamente")
		os.Exit(-1)

	}

}

func iniciarMonitoramento() {

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
		registraLog(site, true)

	} else {
		fmt.Printf("Deu ruim, site %s fora do ar %d\n", site, resp.StatusCode)
		registraLog(site, false)
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

func registraLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	file.WriteString(time.Now().Format("02/01/2006 -- 15:04:05") + " / " + site + " - online: " + strconv.FormatBool(status) + "\n")

}

func imprimirLogs() {

	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(file))
}
