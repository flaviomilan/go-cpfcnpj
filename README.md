# go-cpfcnpj
Validador de CPF/CNPJ em Golang

Este pacote tem o objetivo de fornecer funções para validar CPF e CNPJ.

## Instalação

Para fazer a instalação, basta instalar o pacote com o go get.

```shell
go get github.com/flaviomilan/go-cpfcnpj
```

## Utilização

Existem duas funções para validação: `ValidateCPF` e `ValidateCNPJ`. Ambas recebem o documento como string e retornam um booleano.

### Exemplo

```golang
package main

import (
	"fmt"

	"github.com/flaviomilan/go-cpfcnpj/cpfcnpj"
)

func main() {
	cpf := cpfcnpj.ValidateCPF("086.293.010-31")
	cnpj := cpfcnpj.ValidateCNPJ("91.839.416/0001-53")

	fmt.Println(cpf, cnpj)
}
```