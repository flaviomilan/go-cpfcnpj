package cpfcnpj

import "testing"

func TestDigitoVerificador(t *testing.T) {
	r := digitoVerificador(248)
	if r != 5 {
		t.Errorf("Digito verificador %d calculado errado", r)
	}
	r = digitoVerificador(258)
	if r != 6 {
		t.Errorf("Digito verificador %d calculado errado", r)
	}
}

func TestDigitoVerificadorIgualAZero(t *testing.T) {
	r := digitoVerificador(297)
	if r != 0 {
		t.Errorf("Digito verificador %d calculado errado", r)
	}
}

func TestValidarCPF(t *testing.T) {
	for _, cpf := range []string{"999.300.510-07", "982.714.570-39", "666.204.260-93", "306.052.520-09"} {
		if ValidarCPF(cpf) != true {
			t.Errorf("Validação errada de cpf válido %s", cpf)
		}
	}
}

func BenchmarkValidarCPF(*testing.B) {
	for _, cpf := range []string{"999.300.510-07", "982.714.570-39", "666.204.260-93", "306.052.520-09"} {
		ValidarCPF(cpf)
	}
}
