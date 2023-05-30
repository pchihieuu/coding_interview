package codinginterview

func longestCommon(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	index := 0
	for {
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i-1]) || index >= len(strs[i]) || strs[i-1][index] != strs[i][index] {
				return strs[0][:index]
			}
		}
		index++
	}
	return strs[0][:index]
}
