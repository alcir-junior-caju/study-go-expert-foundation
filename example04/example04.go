package example04

import (
	"fmt"
)

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Wesley"
	e float64 = 1.2
	f ID      = 1
)

func ShowTypesAndFmt() {
	fmt.Printf("O tipo de E é %T", f)
}
