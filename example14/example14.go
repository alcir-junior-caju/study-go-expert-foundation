package example14

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func ShowPointersAndStructs() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
}
