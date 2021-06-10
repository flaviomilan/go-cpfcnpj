package main

import (
	"fmt"

	"github.com/flaviomilan/go-cpfcnpj/cpfcnpj"
)

func main() {
	fmt.Println("Validar CPF/CNPJ")
	cpfcnpj.ValidarCPF()
}
