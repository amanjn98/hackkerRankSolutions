package Array

func FindMaxSum(arr []int) int {
	sum1:=0;
	sum2:=0
	max:=0;
	for i:=0;i<len(arr);i=i+2{
		sum1:=sum1+arr[i]
		if i+1<len(arr){
		sum2=sum2+arr[i+1]
		}
		if sum1>sum2{
			max=sum1
		}else{
			max=sum2
		}

	}
	return max;
}
