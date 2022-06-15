package utils

import "testing"

func TestRemoveMask(t *testing.T) {
	var tests = []struct {
		CPFCNPJ string
		Retorno string
	}{
		{
			CPFCNPJ: "03.093.215/0001-92",
			Retorno: "03093215000192",
		},
		{

			CPFCNPJ: "03.09V3.2E15/0001A-92",
			Retorno: "0309V32E150001A92",
		},
	}

	for _, test := range tests {
		retorno := RemoveMask(test.CPFCNPJ)

		if retorno != test.Retorno {

			t.Fatalf("Recebido: %s, Esperado: %s", retorno, test.Retorno)

		}

	}

}
