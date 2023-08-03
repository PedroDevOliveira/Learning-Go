package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		showIntroduction()
		var command int = readCommand()
		//if comando == 1 {
		//	fmt.Println("Monitorando")
		//} else if comando == 2 {
		//	fmt.Println("Exibindo Logs")
		//} else if comando == 0 {
		//	fmt.Println("Saindo do programa")
		//} else {
		//	fmt.Println("Comando não reconhecido")
		//}

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo Logs")
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}

}

func showIntroduction() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	// fmt.Println("O endereço de memória da minha variável comando é: ", &command)
	fmt.Println("O comando escolhido foi:", command)
	return command
}

func startMonitoring() {
	fmt.Println("Monitorando")
	var url string = "https://alura.com.br"
	var res, _ = http.Get(url)

	if res.StatusCode == 200 {
		fmt.Println("Site:", url, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", url, "está com problemas")
	}

}
