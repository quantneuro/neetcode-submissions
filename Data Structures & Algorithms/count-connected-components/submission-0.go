func countComponents(n int, edges [][]int) int {
    dfschain:=make(map[int]bool)
	seensofar:=make(map[int]bool)

	//Graph is undirected so both directions need to be considered

	neighbors:=make(map[int][]int)

	for _,v:=range edges{
		neighbors[v[0]]=append(neighbors[v[0]],v[1])
		neighbors[v[1]]=append(neighbors[v[1]],v[0])
	}

	var dfs func(n int)bool

	dfs= func(n int)bool{
		if dfschain[n]{
			return false
		}

		dfschain[n]=true

		v,e:=neighbors[n]

		if e{
			for i:=0;i<len(v);i++{
				dfs(v[i])
				// if !val{
				// 	return false
				// } 
			}
		}
		seensofar[n]=true
		dfschain[n]=false
		return true
	}

	components:=0
	for i:=0;i<n;i++{
		if seensofar[i]{
			continue
		}

		dfs(i)
		components++
	}
	return components
}
