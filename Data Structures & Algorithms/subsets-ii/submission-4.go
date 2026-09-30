import("slices")
func subsetsWithDup(nums []int) [][]int {
	res:=[][]int{}
	slices.Sort(nums)

	var sub func(nums []int, cur []int)
	sub = func(nums []int, cur []int){
		if len(nums)==0{
			temp:=make([]int,len(cur))
			copy(temp,cur)
			res=append(res,temp)
			return
		}

		cur=append(cur,nums[0])

		sub(nums[1:],cur)
		cur=cur[:len(cur)-1]

		for len(nums)>1 && nums[0]==nums[1]{
			nums = nums[1:]
		}	
		sub(nums[1:],cur)


	}

	sub(nums,[]int{})
	return res

}
