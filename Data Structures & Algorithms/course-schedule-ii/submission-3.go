func findOrder(numCourses int, prerequisites [][]int) []int {
    
	need:=make(map[int][]int)

	dfschain:=make(map[int]bool)

	safe:=make(map[int]bool)
	res:=[]int{}

	for _,v:=range prerequisites{
		need[v[0]]=append(need[v[0]],v[1])
	}

	var check func(n int)bool

	check=func(n int)bool {
		if dfschain[n]{
			return false
		}
		if safe[n]{
			return safe[n]
		}
		dfschain[n]=true
		v,e:=need[n]

		if e{
			for i:=0;i<len(v);i++{
				val:=check(v[i])
				if !val{
					return false
					}
				}
		}
		dfschain[n]=false
		safe[n]=true
		res=append(res,n)
		return true
	}

	for i:=0;i<numCourses;i++ {
		val:=check(i)
		if !val{
			return []int{}
		}
	}
	return res
}
