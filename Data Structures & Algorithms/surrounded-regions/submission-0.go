func solve(board [][]byte) {
    rs:=len(board)
	cs:=len(board[0])
	visited:=make(map[[2]int]bool)
	var regions func(r,c int)

	regions=func(r,c int){
		if r<0||c<0||r>=rs||c>=cs{
			return 
		}
		if board[r][c]!='O'{
			return 
		}
		// if r==0 || c==0 || r==rs-1 || c==cs-1{
		// 	return true
		// }
		

		if visited[[2]int{r,c}]{
			return 
		}

		visited[[2]int{r,c}]=true

		regions(r+1,c)
		regions(r-1,c)
		regions(r,c+1)
		regions(r,c-1)

		// return u||d||l||ri

	}

	for r:=0;r<rs;r++{
		for c:=0;c<cs;c++ {
			if board[r][c]=='O'{
				if r==0 || c==0 || r==rs-1 || c==cs-1{
				regions(r,c)
				}
			}
		}
	}
	for r:=0;r<rs;r++{
		for c:=0;c<cs;c++ {
			if !visited[[2]int{r,c}]{
				board[r][c]='X'
			}
		}
	}


}
