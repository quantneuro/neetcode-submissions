

func help(board[][]byte,word string,row int ,column int)bool{
	if len(word)==0{
		return true
	}
	if row > len(board)-1 || column > len(board[0])-1 || row<0 ||column<0{
		return false
	}
	left,right,up,down:=false,false,false,false

	if board[row][column]==word[0]{
			val:=board[row][column]
			board[row][column]='#'
			tempword:=word[1:]
			down=help(board, tempword,row+1,column)
			up=help(board,tempword,row-1,column)
			right=help(board,tempword,row,column+1)
			left=help(board,tempword,row,column-1)
			board[row][column]=val
	}
	return left || right || up || down
	
	
}



func exist(board [][]byte, word string) bool {

	left,right,up,down:=false,false,false,false
	for row:=0 ;row<len(board);row++{
		for column:=0;column<len(board[0]);column++{

			if board[row][column]==word[0]{
				val:=board[row][column]
				board[row][column]='#'
				tempword:=word[1:]
				down=help(board, tempword,row+1,column)
				up=help(board,tempword,row-1,column)
				right=help(board,tempword,row,column+1)
				left=help(board,tempword,row,column-1)
				board[row][column]=val
				}
				if up==true || down==true || left==true || right==true{
					break 
				}

			}
		if up==true || down==true || left==true || right==true{
				break
		}
	

	}
	return left || right || up || down
}
