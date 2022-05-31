package desafio

type Elo struct {
	letra string
	prox  *Elo
}

func (l *Elo) linkar(proximos ...*Elo) {
	l.prox = proximos[0]

	for index, proximo := range proximos {
		proximo.prox = proximos[index+1]
	}
}

func nova_corrente(palavra []string) corrente []*Elo {
	for _, letra := range palavra{
		elo := &Elo{letra: letra}
		corrente.append(elo)
	}
	
	return 
}


func Imprimir_Palavra(corrente []*Elo) string {
	palavra := ""

	for _, letra := range corrente{
		palavra += letra
	}

	return palavra
}

