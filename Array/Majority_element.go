func majorityElement(nums []int) int {
    var max int
    var count int 
    for i:=0;i<len(nums);i++{
       if count==0{
        max=nums[i]
       }
       if max==nums[i]{
        count++
       }else{
        count--
       }
    }
    return max
}