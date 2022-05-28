package eclesio

type Letra struct {
	letra string
	prox  *Letra
}

func (l *Letra) linkar(proximos ...*Letra) {
	l.prox = proximos[0]

	for index, proximo := range proximos {
		proximo.prox = proximos[index+1]
	}
}

func nova_Letra(letra string) *Letra {
	return &Letra{letra: letra}
}

func nova_lista_linkada()

func Imprimir_Palavra(palavra []*Letra) string {

}
