package piscine

func TrimAtoi(s string) int {
	var result int
	var sign int = 1
	var foundNum bool
	for _, char := range s {
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
			foundNum = true
		} else if char == '-' && !foundNum {
			sign = -1
		} else if foundNum {
			continue
		}
	}
	return result * sign
}
