func numIslands(grid [][]byte) int {

	dir:=[][]int{
		{1,0},
		{-1,0},
		{0,1},
		{0,-1},
	}
	row:=len(grid)
	column:=len(grid[0])

	islands:=0
	var count func(i,j int)

	count=func(i,j int){
		if i<0 || j<0 || i>=row || j>=column{
			return
		}
		if grid[i][j]=='0' || grid[i][j]=='#'{
			return
		}
		grid[i][j]='#'

		for _,v:=range dir{
			i+=v[0]
			j+=v[1]
			count(i,j)
			i-=v[0]
			j-=v[1]
		}
		return
	}
    print("tere")
	for i:=0;i<row;i++{
		for j:=0;j<column;j++{
			if grid[i][j]=='1'{
				print("here")
				count(i,j)
				islands++
			}
		}
	}
	return islands

}
