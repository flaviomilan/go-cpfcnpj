package cpfcnpj

import "testing"

func TestToInt(t *testing.T) {
	r := toInt('3')
	if r != 3 {
		t.Errorf("Conversão de rune inválida %d", r)
	}
}

func TestRemoveNonDigits(t *testing.T) {
	r := removeNonDigits("asdfa123asdf58i1023")
	if r != "123581023" {
		t.Errorf("Formatação de input inválida %s", r)
	}
}
