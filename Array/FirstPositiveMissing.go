package Array

func FirstMissingPositive(nums []int) int {
	hashmap := make(map[int]int, len(nums))
	count := 1
	for i := 0; i < len(nums); i++ {
		_, ok := hashmap[nums[i]]
		if !ok {
			hashmap[nums[i]] = 1
		} else {
			hashmap[nums[i]]++
		}
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := hashmap[i]; !ok {
			return i
		}
	}
	return count
}

// AnotherSolution The solution below uses a simple strategy of swapping the values to make sure we get the element at the right index
// If the element value is not in range of 1 to n or the element is already at it's right index e.g '1' being at index 0,
// then we move ahead Otherwise, we keep swapping the element value with it's supposed index until we get the correct value
// at the current index At then end, we loop over the array again to see which number isn't at it's right index and
// return that value, else we return n+1 for case where all elements are present in array.
func AnotherSolution(nums []int) int {
	length := len(nums)
	i := 0
	for i < length {
		if nums[i] > length || nums[i] <= 0 || nums[nums[i]-1] == nums[i] {
			i++
			continue
		}
		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
	}
	for i := 0; i < length; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return length + 1
}
