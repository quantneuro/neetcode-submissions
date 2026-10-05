func maxAreaOfIsland(grid [][]int) int {
	r:=len(grid)
	c:=len(grid[0])
    dir:=[][]int{{1,0},{-1,0},{0,1},{0,-1},}
	var area func(i,j int)
	curarea:=0
	maxarea:=0

	area=func(i,j int){
		if i<0 ||j<0 || i>=r || j>=c {
			return
		}
		if grid[i][j]==0 || grid[i][j]==-1{
			return
		}
		grid[i][j]=-1
		curarea++
		for _,v:=range dir{
			i+=v[0]
			j+=v[1]
			area(i,j)
			i-=v[0]
			j-=v[1]
		} 
	}

	for i:=0;i<r;i++{
		for j:=0;j<c;j++{
			
			if grid[i][j]==1{
				curarea=0
				area(i,j)
				println(curarea,"current area")
				maxarea=max(curarea,maxarea)

			}
		}
	}
	return maxarea
}
