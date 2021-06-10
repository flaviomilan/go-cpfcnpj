package cpfcnpj

import "testing"

func TestDigitoVerificadorCNPJ(t *testing.T) {
	r := digitoVerificador(224)
	if r != 7 {
		t.Errorf("Digito verificador %d calculado errado", r)
	}
	r = digitoVerificador(219)
	if r != 1 {
		t.Errorf("Digito verificador %d calculado errado", r)
	}
}

func TestDigitoVerificadorCNPJMenor2(t *testing.T) {
	r := digitoVerificadorCNPJ(219)
	if r != 0 {
		t.Errorf("Digito verificador %d calculado errado", r)
	}
}

func TestValidarCNPJ(t *testing.T) {
	for _, cnpj := range []string{"34.948.284/0001-64", "55.085.855/0001-54"} {
		if ValidarCNPJ(cnpj) != true {
			t.Errorf("Validação errada de cnpj válido %s", cnpj)
		}
	}
}

func BenchmarkValidarCNPJ(*testing.B) {
	for _, cnpj := range []string{"34.948.284/0001-64", "55.085.855/0001-54"} {
		ValidarCNPJ(cnpj)
	}
}
