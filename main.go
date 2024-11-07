package main

import (
	"./Tree"
	"fmt"
)

func main() {
	//matrix := [][]rune{
	//	{'O', 'M', 'O', 'O', 'X'},
	//	{'O', 'X', 'X', 'O', 'M'},
	//	{'O', 'O', 'O', 'O', 'O'},
	//	{'O', 'X', 'X', 'X', 'O'},
	//	{'O', 'O', 'M', 'O', 'O'},
	//	{'O', 'X', 'X', 'M', 'O'},
	//}
	//weight := []int{31, 26, 33, 21, 40}
	//matrix := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'0', '0', '1', '1', '1'}}
	//strs := []string{"10", "0001", '1', '0', "11001"}
	//arr := []int{1, 2, 3}
	//k := 4
	//
	//subsets := Recursion.GenerateSubsetsWithProductLessThanOrEqual(arr, k)
	//fmt.Println(subsets)
	//ipArray := []string{}
	//Recursion.PrintUniqueSubset("abcde", "", &ipArray)
	//ipArray = Recursion.RemoveDuplicates(ipArray)
	//fmt.Print(ipArray)
	//k := 12
	//Trial.TestQuantity()
	//nums := []int{9, 10, 9, -7, -4, -8, 2, -6}
	//SlidingWindow.MaxSlidingWindow(nums, 5)
	//t := Tree.Initialize()
	//fmt.Println(Tree.LevelOrder(&t))
	//fmt.Println(Dp.CountSubsetSum(12, arr))
	//.PermutationString("XY", 0, len("XY")-1)
	// root := Tree.Initialize()
	// Tree.BinaryPathSum(&root, 0, []int{}, 7)

	root := &Tree.TreeNode{Val: 1}
	root.Left = &Tree.TreeNode{Val: 3}
	root.Left.Left = &Tree.TreeNode{Val: 2}
	root.Right = &Tree.TreeNode{Val: 4}
	root.Right.Right = &Tree.TreeNode{Val: 5}
	root.Right.Left = &Tree.TreeNode{Val: 6}
	root.Right.Right.Right = &Tree.TreeNode{Val: 7}
	queries := []int{2, 4, 6}
	fmt.Println(Tree.TreeQueries(root, queries))
	//root := Tree.TreeNode{Val: -2}
	//root.Right = &Tree.TreeNode{Val: -3}
	//arr := []int{2, 3, 7, -9, 4, 4, 7, 3, 2, 10, 8, 15, 2, 1, -8, 10, -5, 10, -2, 21, 9, 20, 0, 4, 24, 5, 12,
	//	-10, 8, 9, 18, 13, -8, 10, -4, -3, 0, 16, -4, 8, 14, 15, -9, 0, 0, -6, 11, -3, 10, 11, 7, -1, -5, 5,
	//	11, 2, 5, 9, -2, 8, 9, -10, 6, -2, 7, 8, 3, 0, -2, 11}
	//
	//fmt.Println(Array.CanArrange(arr, 18))

}

//func substringCalculator(s string) int64 {
//	// Write your code here
//	ipArray := []string{}
//	ss := map[string]struct{}{}
//	for i := range s {
//		for j := i + 1; j <= len(s); j++ {
//			ss[s[i:j]] = struct{}{}
//		}
//	}
//	//Recursion.PrintUniqueSubset(s, "", &ipArray)
//	//result := Recursion.RemoveDuplicates(ipArray)
//	//return int64(len(result))
//}
