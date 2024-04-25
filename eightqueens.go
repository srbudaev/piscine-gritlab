package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var table [8]int
	for i := 0; i < 8; i++ {
		table[i] = 1
	}
	solveQ(0, table)
}

func solveQ(index int, table [8]int) {
	if index == 8 {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(table[i] + 48))
		}
		z01.PrintRune('\n')
		return

	}
	for i := 1; i <= 8; i++ {
		table[index] = i
		if checkQ(index, table) {
			solveQ(index+1, table)
		}
	}
}

func checkQ(index int, table [8]int) bool {
	for i := 0; i < index; i++ {
		diff := i - index
		if table[i] == table[index] || table[i]+diff == table[index] || table[i]-diff == table[index] {
			return false
		}

	}
	return true
}
