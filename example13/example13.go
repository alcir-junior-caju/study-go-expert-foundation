package example13

func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func ShowWhenUsedPointers() {
	minhaVar1 := 10
	minhaVar2 := 20
	soma(&minhaVar1, &minhaVar2)

	println(minhaVar1)
	println(minhaVar2)
}
