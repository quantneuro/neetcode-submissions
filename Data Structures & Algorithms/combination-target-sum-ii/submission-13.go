import("slices")
func combinationSum2(candidates []int, target int) [][]int {
	res:=[][]int{}
	var combi func(candidates []int, cur []int)

	combi = func(candidates []int, cur []int){

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
		if len(candidates)==0{
			return 
		}

		cur=append(cur,candidates[0])

		combi(candidates[1:],cur)
		cur=cur[:len(cur)-1]

		if len(candidates)>1 && candidates[0]==candidates[1]{
			candidates=candidates[1:]
		}
		combi(candidates[1:],cur)
	}

	slices.Sort(candidates)
	combi(candidates,[]int{})

	return res
}
