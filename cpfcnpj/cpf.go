package cpfcnpj

//validar CPF
func ValidateCPF(cpf string) bool {
	fmt_cpf := removeNonDigits(cpf)

	const (
		size     = 9
		position = 10
	)

	d := fmt_cpf[:size]
	digit := calculateDigit(d, position)
	d = d + digit
	digit = calculateDigit(d, position+1)

	return fmt_cpf == d+digit
}
