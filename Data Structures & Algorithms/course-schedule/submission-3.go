func canFinish(numCourses int, prerequisites [][]int) bool {
    pre:=make(map[int][]int)
	rs:=len(prerequisites)
	// cs:=len(prerequisites[0])
	for i:=0;i<rs;i++{
		pre[prerequisites[i][0]]=append(pre[prerequisites[i][0]],prerequisites[i][1])
	}

	dfschain:=make(map[int]bool)

	safe:=make(map[int]bool)

	var check func(n int)bool

	check=func(n int)bool{

		if dfschain[n]{
			return false
		}
		if safe[n]{
			return safe[n]
		}

		dfschain[n]=true
		v,e := pre[n]
		
		if e{
			// if len(v)==0{
			// 	//no prereq
			// 	dfschain[n]=false
			// 	safe[n]=true
			// 	return true
			// }
			for i:=0;i<len(v);i++{
				val:=check(v[i])
				if !val{
					return false
				} 
			}
			
		}
		// else{
		// 	safe[n]=true
		// 	dfschain[n]=false
		// 	return true
		// }
		dfschain[n]=false
			safe[n]=true
			return true

	}
	var res bool
	for i:=0;i<numCourses;i++{
		res=check(i)
		if !res{
			return res
		}
		
	}
	return res
}
