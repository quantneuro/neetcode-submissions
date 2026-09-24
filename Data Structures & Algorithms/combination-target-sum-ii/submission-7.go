func add(candidates []int,current []int,res *[][]int,target int,sum int){

	// sum:=0
	// for _,v:=range current{
	// 	sum+=v
	// }
	if sum==target{
		temp:=make([]int,len(current))
		copy(temp,current)
		*res=append(*res,temp)
		return
	}
	if len(candidates)==0{
		return
	}
	current=append(current,candidates[0])
	sum+=current[len(current)-1]
	add(candidates[1:],current,res,target,sum)
	sum-=current[len(current)-1]
	current=current[:len(current)-1]
	
	for len(candidates)>1 && candidates[0]==candidates[1]  {
		candidates = candidates[1:]
	}
	add(candidates[1:],current,res,target,sum)
}

func combinationSum2(candidates []int, target int) [][]int {
	res:=[][]int{}

	add(candidates,[]int{},&res,target,0)
	
	return res

}
