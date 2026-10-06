func solve(board [][]byte) {
    
	queue:=[][2]int{}

	rs:=len(board)
	cs:=len(board[0])

	for i:=0;i<rs;i++{
		for j:=0;j<cs;j++{
			if i==0 || j==0 || i==rs-1 || j==cs-1{
				if board[i][j]=='O'{
					queue=append(queue,[2]int{i,j})
				}
			}
		}
	}
// print(len(queue))
	var bfs func()
	dir:=[][]int{{-1,0},{1,0},{0,-1},{0,1}}

	bfs = func(){
		// print(len(queue))
		for len(queue)>0{
			r:=queue[0][0]
			c:=queue[0][1]
			queue=queue[1:]

			if board[r][c]=='O' {
				board[r][c]='#'
		
			for _,v:=range dir{
				nr:=r+v[0]
				nc:=c+v[1]

				if nc>=0 && nr>=0 && nr<rs && nc<cs  {
					
					queue=append(queue,[2]int{nr,nc})
					
				}
			}

		}
		}

	}
	bfs()
	for i:=0;i<rs;i++{
		for j:=0;j<cs;j++{
			
				if board[i][j]=='O'{
					board[i][j]='X'
				}else if board[i][j]=='#'{
					board[i][j]='O'
				}
			
		}
	}

	
}
