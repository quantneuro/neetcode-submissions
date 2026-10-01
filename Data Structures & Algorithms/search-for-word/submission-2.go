func exist(board [][]byte, word string) bool {
	dir:=[][]int{
		{-1,0},
		{1,0},
		{0,-1},
		{0,1},
	}
	mark:=make([][]bool,len(board))

	for i:=0;i<len(board);i++{
		mark[i]=make([]bool,len(board[0]))
	}

	var tra func(i int, j int,word string)bool

	tra = func(i int, j int, word string)bool{
		if len(word)==0{
			return true
		}
		if i<0 ||j<0 || i>=len(board)||j>=len(board[0]){
			return false
		}
		if mark[i][j]{
			return false
		}
		
		
		if word[0]==board[i][j]{
			for _,v:=range dir{
				mark[i][j]=true

				val:=tra(i+v[0],j+v[1],word[1:])
				mark[i][j]=false
				if val{
					return true
				}
				
			}

		}
		return false

	}

	for i:=0;i<len(board);i++{

		for j:=0;j<len(board[0]);j++{

			val:=tra(i,j,word)
			if val {
				return true
			}
		}
	}
	return false



}
