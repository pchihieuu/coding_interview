package codinginterview

func removeElement(nums []int, val int) int {
	i := 0
	j := 0
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}
