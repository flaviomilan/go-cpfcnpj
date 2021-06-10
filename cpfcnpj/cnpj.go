package cpfcnpj

//faz o cálculo do digito verificador
func digitoVerificadorCNPJ(acc int) int {
	v := 11 - (acc % 11)

	if v < 2 {
		return 0
	}
	return v
}

//retorna os digitos verificadores do CNPJ
func validarDigitosCNPJ(cnpj string) bool {
	acc_1, acc_2, mult, d1, d2 := 0, 0, []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}, 0, 0
	for i, char := range cnpj {
		if i < 12 {
			acc_1 += mult[i+1] * runeNumberToInt(char)
			acc_2 += mult[i] * runeNumberToInt(char)
		} else if i == 12 {
			acc_2 += digitoVerificadorCNPJ(acc_1) * 2
			d1 = runeNumberToInt(char)
		} else {
			d2 = runeNumberToInt(char)
		}
	}
	return digitoVerificadorCNPJ(acc_1) == d1 && digitoVerificadorCNPJ(acc_2) == d2
}

//validar CNPJ
func ValidarCNPJ(cnpj string) bool {
	fmt_cnpj := formatarInput(cnpj)

	if len(fmt_cnpj) != 14 {
		return false
	}

	return validarDigitosCNPJ(fmt_cnpj)
}
