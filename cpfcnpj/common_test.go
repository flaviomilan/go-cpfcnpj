package cpfcnpj

import "testing"

func TestRuneNumberToInt(t *testing.T) {
	r := runeNumberToInt('3')
	if r != 3 {
		t.Errorf("Conversão de rune inválida %d", r)
	}
}

func TestFormatarInput(t *testing.T) {
	r := formatarInput("asdfa123asdf58i1023")
	if r != "123581023" {
		t.Errorf("Formatação de input inválida %s", r)
	}
}
