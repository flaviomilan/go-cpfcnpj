package cpfcnpj

import "testing"

func TestValidarCNPJ(t *testing.T) {
	for _, cnpj := range []string{"34.948.284/0001-64", "55.085.855/0001-54"} {
		if ValidateCNPJ(cnpj) != true {
			t.Errorf("Validação errada de cnpj válido %s", cnpj)
		}
	}
}

func BenchmarkValidarCNPJ(*testing.B) {
	for _, cnpj := range []string{"34.948.284/0001-64", "55.085.855/0001-54"} {
		ValidateCNPJ(cnpj)
	}
}
