package linkedin

import "fmt"

type No struct {
	Valor   int
	Proximo *No
}

func (n *No) Add(prox *No) {
	n.Proximo = prox
}

func imprimir(n *No) {
	for n != nil {
		fmt.Printf("%d -> ", n.Valor)

		n = n.Proximo
	}

	fmt.Println()
}
