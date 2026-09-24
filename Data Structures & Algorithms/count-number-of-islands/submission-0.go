func numIslands(grid [][]byte) int {
    var neighbors func(r int, c int)
	//The problem is specifically self-reference during the declaration.
	neighbors= func(r int, c int){
		if r>=len(grid) || c >=len(grid[0]) || r<0 || c<0{
			return
		}
		if grid[r][c]!='1'{
			return
		}
		
		grid[r][c]='#'
		neighbors(r+1,c)
		neighbors(r-1,c)
		neighbors(r,c+1)
		neighbors(r,c-1)

		}
    
	totalislands:=0

	for r:=0;r<len(grid);r++{
		for c:=0;c<len(grid[0]);c++{
			if grid[r][c]=='1'{
				neighbors(r,c)
				totalislands++
			}	

		}
	}
	return totalislands
}
