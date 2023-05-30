package codinginterview

func romanToInt(s string) int {
	var result, cursor, previous int
	for i := 0; i < len(s); i++ {
		cursor = roll(s[i])
		result += cursor
		if cursor > previous {
			result = result - previous*2

		}
		previous = cursor
	}
	return result
}
func roll(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
