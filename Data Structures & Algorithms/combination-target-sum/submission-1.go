func help(nums [] int,target int, result *[][]int,current []int){
	if len(nums)==0{
		return
	}
	sum:=0
	for _,v:=range current{
		sum+=v
	}
	
	if sum==target{
		temp:=make([]int,len(current))
		copy(temp,current)
		*result=append(*result,temp)
		return
	}else if sum >target{
		return
	}

	//choose
	current=append(current,nums[0])
	
	//explore with same value to allow reptition
	help(nums,target,result,current)

	//unchoose
	current=current[:len(current)-1]

	//explore
	help(nums[1:],target,result,current)
	
}

func combinationSum(nums []int, target int) [][]int {
	result:=[][]int{}
	help(nums,target,&result,[]int{})
	return result
    
}
