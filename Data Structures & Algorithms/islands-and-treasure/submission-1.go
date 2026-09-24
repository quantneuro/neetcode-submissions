func islandsAndTreasure(grid [][]int) {
	cs:=len(grid[0])
	rs:=len(grid)
	// INF:=math.MaxInt32
    var find func(r,c,counter int)
	

	find=func(r,c,counter int) {
		
		if r<0 || c<0 || r>=rs || c>=cs{
			return 
		}
		if grid[r][c]==-1{
			return 
		}
		if grid[r][c] < counter {
			return
		}

		grid[r][c] = counter

		find(r+1,c,counter+1)
		find(r-1,c,counter+1)
		find(r,c+1,counter+1)
		find(r,c-1,counter+1)
		

	}

for r:=0;r<rs;r++{

	for c:=0;c<cs;c++{
		if grid[r][c]==0{
			
			find(r,c,0)
		}
	}
}
}

// // import("math")
// func islandsAndTreasure(grid [][]int) {
// 	cs:=len(grid[0])
// 	rs:=len(grid)
// 	INF:=math.MaxInt32
//     var find func(r,c int)int
// 	distance:=0

// 	find=func(r,c int) int{
		
// 		if r<0 || c<0 || r>=rs || c>=cs{
// 			return INF
// 		}
// 		if grid[r][c]==0{
// 			return 0
// 		}

// 		if grid[r][c]!=INF{
// 			return INF
// 		}
		
// 		val:=grid[r][c]
// 		grid[r][c]=-2
// 		u:=find(r+1,c)
// 		d:=find(r-1,c)
// 		l:=find(r,c+1)
// 		ri:=find(r,c-1)
// 		grid[r][c]=val
// 		best :=min(u,d,l,ri)
// 		if best==INF{
// 			return INF
// 		}else{
// 			return 1+best
// 		}
		
// 	}
// 	values:=make(map[[2]int]int)
// 	for r:=0;r<rs;r++{

// 		for c:=0;c<cs;c++{
// 			if grid[r][c]==INF{
// 				distance=find(r,c)
// 				values[[2]int{r,c}]=distance
				
// 			}
// 		}
// 	}
// 	for r:=0;r<rs;r++{

// 		for c:=0;c<cs;c++{
// 			if grid[r][c]==INF{
// 			v:=values[[2]int{r, c}]
// 			grid[r][c]=v
// 			}
// 		}
// 	}


// }
