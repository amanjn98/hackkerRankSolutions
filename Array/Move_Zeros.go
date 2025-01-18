func moveZeroes(nums []int)  {
    if len(nums)==1{
        return 
    }
    zp:=0
    for i:=0;i<len(nums);i++{
        if nums[i]!=0{
         nums[zp],nums[i]=nums[i],nums[zp]
         zp++
    }
    }
}