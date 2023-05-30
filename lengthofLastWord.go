package codinginterview

func lengthOfLastWord(s string) int {
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && result == 0 {
			continue
		}
		if s[i] == ' ' && result > 0 {
			break
		}
		if s[i] != ' ' {
			result++
		}
	}
	return result
}
