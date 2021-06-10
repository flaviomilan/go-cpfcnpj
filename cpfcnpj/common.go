package cpfcnpj

import (
	"regexp"
)

func runeNumberToInt(r rune) int {
	return int(r) - 48
}

//formatar CPF removendo
//tudo o que não for número
func formatarInput(cpf string) string {
	re, _ := regexp.Compile("[^0-9]+")
	cpf = re.ReplaceAllString(cpf, "")
	return cpf
}
