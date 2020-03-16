package contas

import (
	"banco/clientes"
)

//ContaPoupanca cp
type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

//Saca cp
func (cp *ContaPoupanca) Saca(valor float64) string {
	if cp.saldo > valor && valor > 0 {
		cp.saldo -= valor
		return "saque efetuado"
	}
	return "saldo insuficiente ou valor do saque invalido"
}

//Depositar cp
func (cp *ContaPoupanca) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		cp.saldo += valor
		return "deposito efetuado com sucesso", cp.saldo
	}
	return "valor do deposito invalido", cp.saldo
}

//GetSaldo cp
func (cp ContaPoupanca) GetSaldo() float64 {
	return cp.saldo
}
