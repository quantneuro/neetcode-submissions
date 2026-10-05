func islandsAndTreasure(grid [][]int) {
	queue:=[][2]int{}
	r:=len(grid)
	c:=len(grid[0])
	INF:=math.MaxInt32

	for i:=0;i<r;i++{
		for j:=0;j<c;j++{
			if grid[i][j]==0{
				queue=append(queue,[2]int{i,j})
			}
		}
	}
	dir:=[][]int{{1,0},{-1,0},{0,1},{0,-1},}

	for len(queue)!=0{
		row:=queue[0][0]
		column:=queue[0][1]
		current:=grid[row][column]
		queue = queue[1:]

		for _,v:=range dir{    
			row+=v[0]
			column+=v[1]  
			if row>=0 && column>=0 && row<=r-1 && column<=c-1 && grid[row][column]==INF {
				queue=append(queue,[2]int{row,column})
				grid[row][column] = current+1
			}
			row-=v[0]
			column-=v[1] 
		}
	}
    
}
