func islandsAndTreasure(grid [][]int) {
	rs:=len(grid)
	cs:=len(grid[0])
	visited:=make(map[[2]int]bool)
	// cur:=0
    var dfs func(r,c,current int)

	dfs=func(r,c,current int){
		if r<0||c<0||r>=rs ||c>=cs{
			return 
		}
		if grid[r][c]==-1{
			return 
		}
		if visited[[2]int{r,c}]{
			return 
		}
		current++
		visited[[2]int{r,c}]=true
		grid[r][c]=min(grid[r][c],current)

		dfs(r+1,c,current)
		dfs(r-1,c,current)
		dfs(r,c+1,current)
		dfs(r,c-1,current)
		visited[[2]int{r,c}]=false



	}

	for r:=0;r<rs;r++{
		for c:=0;c<cs;c++{
			if grid[r][c]==0{

				// current=0
				dfs(r,c,-1)
				
			}
		}
	}
}
