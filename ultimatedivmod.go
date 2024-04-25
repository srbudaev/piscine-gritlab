package piscine

func UltimateDivMod(a *int, b *int) {
	aValue := *a
	bValue := *b

	*a = aValue / bValue
	*b = aValue % bValue
}
