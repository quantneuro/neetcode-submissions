func permute(nums []int) [][]int {
	n:=len(nums)
	res:=[][]int{}
	var calc func(nums []int,cur []int)
	
	calc=func(nums []int,cur []int){
		if len(cur)==n{
			temp:=make([]int,len(cur))
			copy(temp,cur)
			res=append(res,temp)
			return
		}

		for i:=0;i<len(nums);i++{
			tempnums:=[]int{}
			cur=append(cur,nums[i])
			tempnums=append(tempnums,nums[:i]...)
			tempnums=append(tempnums,nums[i+1:]...)
			calc(tempnums,cur)
			cur=cur[:len(cur)-1]
		}


	}
	calc(nums,[]int{})
	return res
}
