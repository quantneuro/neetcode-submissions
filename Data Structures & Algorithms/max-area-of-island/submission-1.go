func maxAreaOfIsland(grid [][]int) int {
	rs:=len(grid)
	cs:=len(grid[0])
	var all func(r,c int)
	currentarea:=0
	maxarea:=0
    
	all=func(r,c int){
		if r<0|| r>=rs|| c<0|| c>=cs{
			return
		}
		if grid[r][c]!=1{
			return
		}
		grid[r][c]=-1

		all(r+1,c)
		all(r-1,c)
		all(r,c+1)
		all(r,c-1)
		currentarea++
		return 

	}

	for i:=0;i<rs;i++{
		for j:=0;j<cs;j++{
			if grid[i][j]==1{
				currentarea=0
				all(i,j)
				maxarea=max(maxarea,currentarea)
			}
		}
	}

	return maxarea
	
}
