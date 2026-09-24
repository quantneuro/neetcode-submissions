func maxAreaOfIsland(grid [][]int) int {
	rs:=len(grid)
	cs:=len(grid[0])

    var area func(r,c int)
	
	currentarea:=0
	maxarea:=0
	
	area=func(r,c int){
		if r<0 || c<0 || r>=rs || c>=cs {
			return
		}
		if grid[r][c]!=1{
			return
		}
		grid[r][c]=-1
		currentarea++
		area(r+1,c)
		area(r-1,c)
		area(r,c+1)
		area(r,c-1)
	}

	for r:=0;r<rs;r++{
		for c:=0;c<cs;c++{
			if grid[r][c]==1{
				currentarea=0
				area(r,c)
				maxarea = max(currentarea,maxarea)
			}
		}
	}
	return maxarea
}
