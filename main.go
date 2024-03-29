package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

func initFunc() {
	fmt.Println(version)
	kingpin.Version(version)
	//for _, kv := range info.Settings {
	//	if kv.Value == "" {
	//		continue
	//	}
	//	switch kv.Key {
	//	case "vcs.revision":
	//		Revision = kv.Value
	//	case "vcs.time":
	//		LastCommit, _ = time.Parse(time.RFC3339, kv.Value)
	//	case "vcs.modified":
	//		DirtyBuild = kv.Value == "true"
	//	}
	//}
}
func main() {
	//SlidingWindow.MaxSatisfied([]int{7, 8, 8, 6}, []int{0, 1, 0, 1}, 3)
	//trust := [][]int{{1, 3}, {2, 3}, {3, 1}}
	//fmt.Println(Array.FindJudge(3, trust))
	//fmt.Println(Graph.FindAllPeople(4, [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}, 3))
	initFunc()
}

//func main() {
//	//matrix := [][]rune{
//	//	{'O', 'M', 'O', 'O', 'X'},
//	//	{'O', 'X', 'X', 'O', 'M'},
//	//	{'O', 'O', 'O', 'O', 'O'},
//	//	{'O', 'X', 'X', 'X', 'O'},
//	//	{'O', 'O', 'M', 'O', 'O'},
//	//	{'O', 'X', 'X', 'M', 'O'},
//	//}
//	//weight := []int{31, 26, 33, 21, 40}
//	//matrix := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'0', '0', '1', '1', '1'}}
//	//strs := []string{"10", "0001", '1', '0', "11001"}
//	//arr := []int{1, 2, 3}
//	//k := 4
//	//
//	//subsets := Recursion.GenerateSubsetsWithProductLessThanOrEqual(arr, k)
//	//fmt.Println(subsets)
//	//ipArray := []string{}
//	//Recursion.PrintUniqueSubset("abcde", "", &ipArray)
//	//ipArray = Recursion.RemoveDuplicates(ipArray)
//	//fmt.Print(ipArray)
//	//k := 12
//	//Trial.TestQuantity()
//	//nums := []int{9, 10, 9, -7, -4, -8, 2, -6}
//	//SlidingWindow.MaxSlidingWindow(nums, 5)
//	//t := Tree.Initialize()
//	//fmt.Println(Tree.LevelOrder(&t))
//	//fmt.Println(Dp.CountSubsetSum(12, arr))
//	//.PermutationString("XY", 0, len("XY")-1)
//	//root := Tree.Initialize()
//	//Tree.PrintPaths(&root, -5)
//	// Implement your solution here
//	result := make([]int, X)
//	result[0] = 0
//	if X == 1 {
//		fmt.
//			result[1] = 1
//		for i := 2; i < X; i++ {
//			result[i] = result[i-1] + result[i-2]
//		}
//		return result
//
//	}

// }
func CheckPowerOfTwo(n int) int {
	//added one corner case if n is zero it will also consider as power 2
	if n == 0 {
		return 1
	}
	return n & (n - 1)
}

//func main() {
//	//var n = 257
//	//flag := CheckPowerOfTwo(n)
//	//if flag == 0 {
//	//	fmt.Printf("Given %d number is the power of 2.\n", n)
//	//} else {
//	//	fmt.Printf("Given %d number is not the power of 2.\n", n)
//	//}
//	s := strings.Split("12.16.20", ".")
//	fmt.Println(s)
//	majorVersion := strings.Join(s[:2], ".")
//	fmt.Println(majorVersion)
//}

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

// Send and receive integers to a channel using go routines.
// In the main function call a function "generator" which will add 10 integers to a channel; return the channel from the functions.
// In the main function call a function "receiver"; pass the channel to this function and print all the 10 integers.
//func main() {
//	ch := make(chan int, 1)
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//	go generator(ch, wg)
//	wg.Add(1)
//	go reciever(ch, wg)
//	wg.Wait()
//	close(ch)
//}
//
//func generator(ch chan int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	add := 0
//	for i := 0; i < 10; i++ {
//		fmt.Println("Writing data to a channel")
//		add = add + i
//	}
//	ch <- add
//
//}
//
//func reciever(ch chan int, wg *sync.WaitGroup) {
//	for i := <-ch; i >= 0; i-- {
//		fmt.Println(i)
//	}
//	wg.Done()
//
//}
