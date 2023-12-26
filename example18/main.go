package example18

import (
	"fmt"
	"study-go-expert-foundation/example18/matematica"

	"github.com/google/uuid"
)

func ShowPackages() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Println(carro.Andar())
	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)
	fmt.Println(uuid.New())
}
