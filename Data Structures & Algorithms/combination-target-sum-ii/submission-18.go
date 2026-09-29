import("slices")
func combinationSum2(candidates []int, target int) [][]int {
	res:=[][]int{}
	var combi func(candidates []int, cur []int,sum int)

	combi = func(candidates []int, cur []int,sum int){

		

		if sum==target{
			temp:=make([]int,len(cur))
			copy(temp,cur)
			res=append(res,temp)
			return 
		}
		if len(candidates)==0{
			return 
		}
		if sum>target{
			return
		}
		// sum+=candidates[0]
		
		cur=append(cur,candidates[0])
		sum+=cur[len(cur)-1]
		combi(candidates[1:],cur,sum)
		sum-=cur[len(cur)-1]
		cur=cur[:len(cur)-1]

		for len(candidates)>1 && candidates[0]==candidates[1]{
			candidates=candidates[1:]
		}
		combi(candidates[1:],cur,sum)
	}

	slices.Sort(candidates)
	combi(candidates,[]int{},0)

	return res
}
