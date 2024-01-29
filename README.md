# Desafio PingPang

Este repositório contém uma implementação simples do desafio "PingPang" em Go. O desafio consiste em percorrer os números de 1 a N e imprimir "Ping" se o número for divisível por 3, "Pang" se divisível por 5, "PingPang" se divisível por ambos, e o próprio número caso contrário.

## Estrutura do Código

### main.go
O arquivo `main.go` contém a implementação principal do desafio. A função `PingPang` aceita um argumento inteiro `n` e retorna uma string contendo a saída do desafio para os números de 1 a N. O código utiliza um loop `for` e uma estrutura `switch` para avaliar as condições e construir a string de saída.

A função `main` chama `PingPang` com um valor arbitrário (100 neste caso) e imprime o resultado.

### main_test.go
O arquivo `main_test.go` contém testes unitários para a função `PingPang` usando o pacote de teste padrão do Go. Dois conjuntos de testes são definidos:

1. **TestPingPang**: Verifica se a saída da função `PingPang` é igual à saída esperada para vários valores de entrada.
2. **TestPingPangStringContains**: Verifica se a saída da função `PingPang` contém substrings específicas para determinados valores de entrada.

## Executando os Testes
Para executar os testes, basta utilizar o comando `go test` no terminal dentro do diretório do projeto.

```bash
go test
```

## Contexto e Aprendizado
Este projeto foi criado como parte de um exercício de aprendizado para praticar a implementação de algoritmos simples em Go e familiarizar-se com os conceitos básicos da linguagem.

Os testes unitários são uma prática comum no desenvolvimento de software, e eles garantem que as alterações no código não quebrem o comportamento esperado. Além disso, o uso do pacote `testing` do Go é uma boa introdução aos testes na linguagem.

O código também demonstra o uso de estruturas de controle, como loops e switches, em Go, assim como a manipulação de strings.

Sinta-se à vontade para explorar, contribuir ou utilizar este projeto como um recurso educacional para aprender Go.
