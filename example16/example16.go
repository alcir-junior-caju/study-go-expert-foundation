package example16

import "fmt"

func ShowTypeAsserts() {
	var assertVar interface{} = "Wesley Willians"
	println(assertVar.(string))
	assertRes, assertOk := assertVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", assertRes, assertOk)
	assertRes2, assertOk2 := assertVar.(int)
	fmt.Printf("O valor de res2 é %v e o resultado de ok2 é %v", assertRes2, assertOk2)
}
