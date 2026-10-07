func canFinish(numCourses int, prerequisites [][]int) bool {
    need:=make(map[int][]int)
	check:=make(map[int]bool)
	// done:=make(map[int]bool)


	for _,v:=range prerequisites{
			need[v[1]]=append(need[v[1]],v[0])
	}

	var dfs func(n int)bool

	dfs=func(n int)bool{
		if check[n]{
			return false
		}
		check[n]=true
		v,e:=need[n]

		if e{
			// if len(v)==0{ return true }
			for i:=0;i<len(v);i++{
				result:=dfs(v[i])
				if !result { return false}
			}
		}
		// else{
		// 	return true
		// }
		check[n]=false

		return true


	}

	for i:=0;i<numCourses;i++{
		result:=dfs(i)
		if !result { return false }
	}
	return true

}
