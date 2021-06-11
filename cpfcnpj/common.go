package cpfcnpj

import (
	"regexp"
	"strconv"
)

//formatar CPF removendo
//tudo o que não for número
func removeNonDigits(cpf string) string {
	re, _ := regexp.Compile("[^0-9]+")
	cpf = re.ReplaceAllString(cpf, "")
	return cpf
}

// toInt converte uma rune para inteiro
func toInt(r rune) int {
	return int(r - '0')
}

//calcular digito do documento
func calculateDigit(doc string, position int) string {

	var sum int
	for _, r := range doc {

		sum += toInt(r) * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}
