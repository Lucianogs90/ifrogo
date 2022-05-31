package desafio

type Elo struct {
	letra string
	prox  *Elo
}

func (e *Elo) linkar(elos []*Elo) {
	if len(elos) == 0 {
		return
	}

	primeiro := elos[0]

	e.prox = primeiro

	primeiro.linkar(elos[1:])
}

func criar_corrente(letras []string) []*Elo {
	corrente := []*Elo{}

	for _, letra := range letras {
		elo := &Elo{letra: letra}
		corrente = append(corrente, elo)
	}

	return corrente
}

func imprimir_corrente(corrente []*Elo) string {
	palavra := ""

	for _, elo := range corrente {
		palavra += elo.letra
	}

	return palavra
}
