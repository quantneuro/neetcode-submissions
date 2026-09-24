func validTree(n int, edges [][]int) bool {
    
	seen:=make(map[int][]int)

	for _,v:=range edges{
		//if I keep only one that will be unidirectional 
		//I need two one for each direction
		seen[v[0]] = append(seen[v[0]],v[1])
		seen[v[1]] = append(seen[v[1]], v[0])
	}

	dfschain:=make(map[int]bool)
	seensofar:=make(map[int]bool)


	var check func(n, parent int)bool


	check=func(n ,parent int)bool{
		if dfschain[n]{
			return false
		}
		dfschain[n]=true
		v,e:=seen[n]
		if e{
		for i:=0;i<len(v);i++{
			if v[i]==parent{
				continue
			}
			val:=check(v[i],n)
			if !val{
				return false
			}
		}
		}

		dfschain[n]=false
		seensofar[n]=true
		return true
	}

	for i:=0;i<n;i++{
		val:=check(i,-1)
		if !val{
			return false
		}
	}

	for i:=0;i<n;i++{
		if !seensofar[i]{
			return false
		}
	}
	return true
}
