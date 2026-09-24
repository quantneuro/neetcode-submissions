func isValidSudoku(board [][]byte) bool {
row:=make([]map[byte]bool,9)

column:=make([]map[byte]bool,9)

boxes:=make([]map[byte]bool,9)

for i:=range row{
	row[i]=make(map[byte]bool)
	column[i]=make(map[byte]bool)
	boxes[i]=make(map[byte]bool)
}
var value byte
for r:=0;r<9;r++{

	for c:=0;c<9;c++{

		if board[r][c]=='.'{
			continue
		}
		value=board[r][c]
		box:=(r/3)*3+(c/3)

		if row[r][value]==true || column[c][value]==true || boxes[box][value]==true{
			return false
		}

		row[r][value]=true
		column[c][value]=true
		boxes[box][value]=true
	}
}
return true

}
