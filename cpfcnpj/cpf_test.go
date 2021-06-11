package cpfcnpj

import "testing"

func TestValidarCPF(t *testing.T) {
	for _, cpf := range []string{"999.300.510-07", "982.714.570-39", "666.204.260-93", "306.052.520-09"} {
		if ValidateCPF(cpf) != true {
			t.Errorf("Validação errada de cpf válido %s", cpf)
		}
	}
}

func BenchmarkValidarCPF(*testing.B) {
	for _, cpf := range []string{"999.300.510-07", "982.714.570-39", "666.204.260-93", "306.052.520-09"} {
		ValidateCPF(cpf)
	}
}
