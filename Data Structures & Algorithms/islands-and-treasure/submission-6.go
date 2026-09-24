func islandsAndTreasure(grid [][]int) {
	rs:=len(grid)
	cs:=len(grid[0])
	queue:=[][2]int{}
	INF:=math.MaxInt32
	for r:=0;r<rs;r++{
		for c:=0;c<cs;c++{
			if grid[r][c]==0{
				queue=append(queue,[2]int{r,c})
			}
		}
	}
	dir:=[][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for len(queue)!=0{

		r:=queue[0][0]
		c:=queue[0][1]
		queue=queue[1:]

		
		
		val:=grid[r][c]
		// if r<0||c<0||r>=rs ||c>=cs || grid[r][c]!=INF{
		// 	continue 
		// }else{	
			for _,v:=range dir{
				row:=r+v[0]
				column:=c+v[1]
				
				if row>=0 && row<rs && column>=0 && column<cs && grid[row][column]==INF {
					grid[row][column]=val+1
					queue=append(queue,[2]int{row,column})

				}
			}

		}
		
		
	}
// }
// func islandsAndTreasure(grid [][]int) {
// 	rs:=len(grid)
// 	cs:=len(grid[0])
// 	visited:=make(map[[2]int]bool)
// 	// cur:=0
//     var dfs func(r,c,current int)

// 	dfs=func(r,c,current int){
// 		if r<0||c<0||r>=rs ||c>=cs{
// 			return 
// 		}
// 		if grid[r][c]==-1{
// 			return 
// 		}
// 		if visited[[2]int{r,c}]{
// 			return 
// 		}
// 		current++
// 		visited[[2]int{r,c}]=true
// 		grid[r][c]=min(grid[r][c],current)

// 		dfs(r+1,c,current)
// 		dfs(r-1,c,current)
// 		dfs(r,c+1,current)
// 		dfs(r,c-1,current)
// 		// visited[[2]int{r,c}]=false



// 	}

// 	for r:=0;r<rs;r++{
// 		for c:=0;c<cs;c++{
// 			if grid[r][c]==0{

// 				// current=0
// 				dfs(r,c,-1)
// 				clear(visited)
				
// 			}
// 		}
// 	}
// }
