package main

import "fmt"

func main() {
	var fatura int
	var nubank int
	var c6 int
	var caixa int
	var recebiveis int

	fmt.Println("Digite o valor da sua fatura: ")
	fmt.Scan(&fatura)
	fmt.Println("Digite o seu saldo no nubank: ")
	fmt.Scan(&nubank)
	fmt.Println("Digita o seu saldo no c6 bank: ")
	fmt.Scan(&c6)
	fmt.Println("Digita o seu saldo no CAIXA: ")
	fmt.Scan(&caixa)
	fmt.Println("Digite o valor que estão te devendo: ")
	fmt.Scan(&recebiveis)

	saldo := nubank + c6 + caixa + recebiveis
	debito := fatura

	dinheiro_restante := saldo - debito

	fmt.Printf("Seu dinheiro restante é: %v \n", dinheiro_restante)

}
