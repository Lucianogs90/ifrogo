package desafio

import (
	"fmt"
	"testing"
)

func TestDesafio(t *testing.T) {
	letras := []string{"e", "c", "l", "e", "s", "i", "o"}

	esperado := "eclesio"

	corrente := criar_corrente(letras)

	resultado := imprimir_corrente(corrente)

	if resultado != esperado {
		t.Fail()
		t.Errorf("Obtido: %s, esperado: %s", resultado, esperado)
	} else {
		fmt.Println("Deu certo, maluco!")
	}
}
