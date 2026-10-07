func findOrder(numCourses int, prerequisites [][]int) []int {
    need:=make(map[int][]int)
	
	for _,v:=range prerequisites{
		need[v[1]]=append(need[v[1]],v[0])
	} 
	cycle:=make(map[int]bool)
	done:=make(map[int]bool)

	res:=[]int{}
	var dfs func(n int)bool

	dfs=func(n int)bool{
		if cycle[n]{return false}
		if done[n]{return true}
		cycle[n]=true
		res=append(res,n)
		if v,ok:=need[n];ok{
			for i:=0;i<len(v);i++{
				val:=dfs(v[i])
				if !val{return false}
				
			}
		}
		cycle[n]=false
		done[n]=true
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
