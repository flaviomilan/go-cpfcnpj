package cpfcnpj

//faz o cálculo do digito verificador
func digitoVerificador(acc int) int {
	v := 11 - (acc % 11)

	if v > 9 {
		return 0
	}
	return v
}

//retorna os digitos verificadores do CPF
func validarDigitosCPF(cpf string) bool {
	acc_1, acc_2, mult, d1, d2 := 0, 0, 10, 0, 0
	for i, char := range cpf {
		if i < 9 {
			acc_1 += mult * runeNumberToInt(char)
			acc_2 += (mult + 1) * runeNumberToInt(char)
			mult -= 1
		} else if i == 9 {
			acc_2 += 2 * runeNumberToInt(char)
			d1 = runeNumberToInt(char)
		} else {
			d2 = runeNumberToInt(char)
		}
	}
	return digitoVerificador(acc_1) == d1 && digitoVerificador(acc_2) == d2
}

//validar CPF
func ValidarCPF(cpf string) bool {
	fmt_cpf := formatarInput(cpf)

	if len(fmt_cpf) != 11 {
		return false
	}

	return validarDigitosCPF(fmt_cpf)
}
