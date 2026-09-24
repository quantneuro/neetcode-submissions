func numIslands(grid [][]byte) int {
    rs:=len(grid)
	cs:=len(grid[0])
	totalislands:=0

	var all func(r,c int)bool

	all=func(r,c int)bool{
		if r<0 ||r>=rs|| c<0||c>=cs{
			return false
		}
		if grid[r][c]=='0'{
			return false
		}
		if grid[r][c]=='#'{
			return false
		}
		println("#")
		grid[r][c]='#'
		all(r+1,c)
		all(r-1,c)
		all(r,c+1)
		all(r,c-1)
		

		return true

	}

	for r:=0;r<rs;r++{
		for c:=0;c<cs;c++{
			println("c")
			if grid[r][c]=='1'{
				val:=all(r,c)
				println("call")
				if val{
					totalislands++
				}
			}
		}
	}

	return totalislands
}
