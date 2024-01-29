package main

import "fmt"

func PingPang(n int) string {
	// Inicializando uma variável para acumular os resultados
	resultado := ""

	// Loop de 1 a n
	for i := 1; i <= n; i++ {
		// Estrutura switch sem expressão, atuando como um if-else
		switch {
		// Caso o número seja divisível por 3 e 5
		case i%3 == 0 && i%5 == 0:
			resultado += "PingPang\n"
		// Caso o número seja divisível por 3
		case i%3 == 0:
			resultado += "Ping\n"
		// Caso o número seja divisível por 5
		case i%5 == 0:
			resultado += "Pang\n"
		// Caso não atenda a nenhum dos casos anteriores
		default:
			resultado += fmt.Sprintf("%d\n", i)
		}
	}

	// Retornando a string acumulada
	return resultado
}

func main() {
	// Exibindo o resultado na função main
	fmt.Print(PingPang(100))
}
