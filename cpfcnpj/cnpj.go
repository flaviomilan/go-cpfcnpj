package cpfcnpj

//validar CNPJ
func ValidateCNPJ(cnpj string) bool {
	fmt_cnpj := removeNonDigits(cnpj)

	const (
		size     = 12
		position = 5
	)

	d := fmt_cnpj[:size]
	digit := calculateDigit(d, position)
	d = d + digit
	digit = calculateDigit(d, position+1)

	return fmt_cnpj == d+digit
}
