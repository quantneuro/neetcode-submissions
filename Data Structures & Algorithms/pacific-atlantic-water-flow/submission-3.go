func pacificAtlantic(heights [][]int) [][]int {
    res:=[][]int{}

	rs:=len(heights)
	cs:=len(heights[0])

	a:=[][2]int{}
	p:=[][2]int{}
	//row have same column
	for c:=0;c<cs;c++{
		p=append(p,[2]int{0,c})
		a=append(a,[2]int{rs-1,c})
	}
	//first column for pacific
	//starting at 1 becasue I don't want that to repeat in pacific
	for r:=1;r<rs;r++{
		p=append(p,[2]int{r,0})
	}
	//last column for atlantic
	for r:=0;r<rs-1;r++{
		a = append(a, [2]int{r, cs-1})
	}

	dir:=[][]int{{1,0},{-1,0},{0,1},{0,-1},}
	seena:=make(map[[2]int]bool)

	for _, pair := range a {
       seena[pair] = true
   }
   seenp := make(map[[2]int]bool)
   for _, pair := range p {
       seenp[pair] = true
   }


	var bfs func(o [][2]int,seen map[[2]int]bool)

	bfs=func(o [][2]int,seen map[[2]int]bool){

		for len(o)!=0{
			r:=o[0][0]
			c:=o[0][1]
			o=o[1:]
			
			for _,d:=range dir{
				nr:=r+d[0]
				nc:=c+d[1]
				// Can this nr nc value reach the Pacific through the r and c I already know reaches Pacific?”
				// So  BFS is not simulating water movement. It is tracing reachability in reverse from the ocean.
				//
				if nr>=0 && nc>=0 && nr<=rs-1 && nc<=cs-1 && heights[nr][nc]>=heights[r][c]{
					if !seen[[2]int{nr,nc}]{
           			o=append(o,[2]int{nr,nc})
					}
					seen[[2]int{nr, nc}] = true

				}
			}

		}
	}
	bfs(p, seenp); bfs(a, seena)

	for r:=0;r<rs;r++{
		for c:=0;c<cs;c++{
	
			if seenp[[2]int{r,c}] && seena[[2]int{r,c}] { 
				res = append(res, []int{r,c}) 
				}
		}
	}

	return res



	
}
