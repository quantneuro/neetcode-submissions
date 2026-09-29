func combinationSum(nums []int, target int) [][]int {
	res:=[][]int{}
	var combi func(nums []int,cur []int)
    combi=func(nums []int, cur []int){
		sum:=0
		for _,v:=range cur{
			sum+=v
		}
		
		if sum==target{
			temp:=make([]int,len(cur))
			copy(temp,cur)
			res=append(res,temp)
			return
		}
		if sum>target{
			return
		}
		if len(nums)==0{
			return 
		}

		cur=append(cur,nums[0])
		combi(nums,cur)
		cur=cur[:len(cur)-1]
		combi(nums[1:],cur)
	}

	combi(nums,[]int{})

	return res
}
