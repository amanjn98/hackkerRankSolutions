package main

func utopianTree(n int32) int32 {
	// Write your code here
	if n==0{
		return 1
	}
	growthValue:=2
	var i int32
	for  i=2;i<=n;i++{
		if i%2==0{
			growthValue=growthValue+1
		}else{
			growthValue=growthValue*2
		}

	}
	return int32(growthValue)
}
