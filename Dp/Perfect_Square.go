package Dp

func NumSquares(n int) int {
	if n == 1 {
		return 1
	}
	arr := []int{}
	i := 1
	for {
		if i*i < n {
			arr = append(arr, i*i)
		}
		if i*i > n {
			break
		}
		if i*i == n {
			return 1
		}
		i++
	}
	return Minimum_Num_Coins(n, arr)
}
