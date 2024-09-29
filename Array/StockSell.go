package Array

func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			res += profit
		}
	}
	return res
}
