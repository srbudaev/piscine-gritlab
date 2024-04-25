package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	numDigits := countDigits(n)
	digits := make([]int, numDigits)
	for i := numDigits - 1; i >= 0; i-- {
		digits[i] = n % 10
		n /= 10
	}
	sortDigits(digits)
	for _, digit := range digits {
		z01.PrintRune(rune(digit + '0'))
	}
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func sortDigits(digits []int) {
	for i := 0; i < len(digits)-1; i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}
}
