package codinginterview

import "math"

func maxProfit(prices []int) int {
	min := math.MaxInt32
	result := 0
	for _, v := range prices {
		if v < min {
			min = v
		} else if v-min > result {
			result = v - min
		}
	}
	return result
}

//^
//|
//|
//|  x
//|              x
//|       x
//|               x
//|         x
//|   x
//|-------------------------->
// 1  2  3  4  5  6
