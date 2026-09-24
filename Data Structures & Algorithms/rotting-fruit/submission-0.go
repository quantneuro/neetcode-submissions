func orangesRotting(grid [][]int) int {
    queue:=[][2]int{}
	minutes:=0
	for r:=0;r<len(grid);r++{
		for c:=0;c<len(grid[0]);c++{
			if grid[r][c]==2{
				queue=append(queue,[2]int{r,c})
			}
		}
	}

	for len(queue)!=0{
		currentfruit:=queue[0]

		queue=queue[1:]

		r:=currentfruit[0]
		c:=currentfruit[1]

		directions:=[][2]int{{1,0},{-1,0},{0,1},{0,-1}}
		change:=false
		for _,v:=range directions{
			nr:=r+v[0]
			nc:=c+v[1]

			if nr<0 || nc<0 || nc>=len(grid[0]) || nr>=len(grid){
				continue
			}
			if grid[nr][nc]==1{
				grid[nr][nc]=2
				queue=append(queue,[2]int{nr,nc})
				change = true
			}
		}
		if change{
		minutes++
		}

	}
	if minutes==0{return -1}
	return minutes
	
}
