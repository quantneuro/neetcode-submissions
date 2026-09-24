func ssum(nums []int,current[]int,res *[][]int,target int,sum int ){
	// sum:=0
	// for _,v:=range current{
	// 	sum+=v
	// }

	if sum>target{
		return
	}
	if sum==target{
		temp:=make([]int,len(current))
		copy(temp,current)
		*res=append(*res,temp)
		return
	}
	if len(nums)==0 {
		
		return
		
	}
	
	current=append(current,nums[0])
	sum+=current[len(current)-1]
	ssum(nums,current,res,target,sum)
	sum-=current[len(current)-1]
	current=current[:len(current)-1]
	ssum(nums[1:],current,res,target,sum)
	
}
func combinationSum(nums []int, target int) [][]int {
	res:=[][]int{}

	ssum(nums,[]int{},&res,target,0)
	return res
    
}

// func ssum(nums []int,current[]int,res *[][]int,target int){
// 	sum:=0
// 	for _,v:=range current{
// 		sum+=v
// 	}

// 	if sum>target{
// 		return
// 	}
// 	if sum==target{
// 		temp:=make([]int,len(current))
// 		copy(temp,current)
// 		*res=append(*res,temp)
// 		return
// 	}
// 	if len(nums)==0 {
		
// 		return
		
// 	}
	
// 	current=append(current,nums[0])
// 	ssum(nums,current,res,target)
// 	current=current[:len(current)-1]
// 	ssum(nums[1:],current,res,target)
	
// }
// func combinationSum(nums []int, target int) [][]int {
// 	res:=[][]int{}

// 	ssum(nums,[]int{},&res,target)
// 	return res
    
// }
