package contas

import (
	"banco/clientes"
	"fmt"
)

//ContaCorrente cc
type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

//Saca cc
func (cc *ContaCorrente) Saca(valor float64) string {
	if cc.saldo > valor && valor > 0 {
		cc.saldo -= valor
		return "saque efetuado"
	}
	return "saldo insuficiente ou valor do saque invalido"
}

//Depositar cc
func (cc *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		cc.saldo += valor
		return "deposito efetuado com sucesso", cc.saldo
	}
	return "valor do deposito invalido", cc.saldo
}

//Transferencia cc
func (cc *ContaCorrente) Transferencia(valor float64, contaDestino *ContaCorrente) bool {
	if valor > 0 && cc.saldo > valor {
		fmt.Println(contaDestino.Depositar(valor))
		fmt.Println(cc.Saca(valor))
		return true
	}
	return false
}

//GetSaldo cc
func (cc ContaCorrente) GetSaldo() float64 {
	return cc.saldo
}
