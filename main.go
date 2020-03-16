package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

//PagaBoleto conta
func PagaBoleto(conta verificarConta, valor float64) {
	conta.Saca(valor)
}

type verificarConta interface {
	Saca(valor float64) string
}

func main() {
	cliente1 := clientes.Titular{
		Nome:      "Carlos",
		CPF:       "123.456.789.10",
		Profissao: "Dev",
	}

	cc := contas.ContaCorrente{
		Titular:       cliente1,
		NumeroAgencia: 123,
		NumeroConta:   10,
	}

	fmt.Println(cc.Depositar(100))
	fmt.Println(cc.GetSaldo())

	cp := contas.ContaCorrente{
		Titular:       cliente1,
		NumeroAgencia: 123,
		NumeroConta:   10,
	}

	fmt.Println(cp.Depositar(3300))
	fmt.Println(cp.GetSaldo())

	PagaBoleto(&cc, 60)
	fmt.Println(cc.GetSaldo())
}
