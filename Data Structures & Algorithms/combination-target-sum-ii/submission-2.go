import (
    "slices"
)
func help(nums []int ,target int ,current [] int,result*[][]int ){
	if len(nums)==0{
		
	
	sum:=0
	for _,v:=range current{
		sum+=v
	}
	if sum==target{
		temp:=make([]int, len(current))
		copy(temp,current)
		*result=append(*result,temp)
		// return 
	}
	// else if sum >target{//don't need it here since it's not infinite loop since the nums is decreasing !
	// //use this because you don't wanna keep moving after you have passed the target
	// 	return 
	// }
	return 
	}
	if sum >target{//don't need it here since it's not infinite loop since the nums is decreasing !
	// //use this because you don't wanna keep moving after you have passed the target
		return 
	// }
	

	
    //choose
	current=append(current,nums[0])
	

	help(nums[1:],target,current,result)

	current=current[:len(current)-1]
	for len(nums)>1 && nums[0]==nums[1]{
		nums=nums[1:]
	}

	help(nums[1:],target,current,result)


}
func combinationSum2(candidates []int, target int) [][]int {
	result:=[][]int{}
	slices.Sort(candidates)

	help(candidates,target,[]int{},&result)
	return result
}
