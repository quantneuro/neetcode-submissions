func findOrder(numCourses int, prerequisites [][]int) []int {
    need:=make(map[int][]int)
	
	for _,v:=range prerequisites{
		need[v[0]]=append(need[v[0]],v[1])
	} 
	cycle:=make(map[int]bool)
	done:=make(map[int]bool)

	res:=[]int{}
	var dfs func(n int)bool

	dfs=func(n int)bool{
		if cycle[n]{return false}
		if done[n]{return true}
		cycle[n]=true
		
		if v,ok:=need[n];ok{
			for i:=0;i<len(v);i++{
				val:=dfs(v[i])
				if !val{return false}
			}
		}
		
		cycle[n]=false
		done[n]=true
		res=append(res,n)
		return true

	}

	for i:=0;i<numCourses;i++{
		val :=dfs(i)
		if !val{
			return []int{}
		}
	}
	return res

}
