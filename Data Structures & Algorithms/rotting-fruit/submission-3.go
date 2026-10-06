func orangesRotting(grid [][]int) int {
queue:=[][2]int{}
rl:=len(grid)
cl:=len(grid[0])
for i:=0;i<rl;i++{
	for j:=0;j<cl;j++{
		if grid[i][j]==2{
			queue=append(queue,[2]int{i,j})
		}
	}
}
time:=0

dir:=[][]int{{1,0},{-1,0},{0,1},{0,-1},}
for len(queue)!=0 {
	curl:=len(queue)
	change:=false
	for i:=0 ;i<curl;i++{
	r:=queue[0][0]
	c:=queue[0][1]
	queue=queue[1:]
	
	for _,v:=range dir{
		r+=v[0]
		c+=v[1]
		if r>=0 && c>=0 && r<=rl-1 && c<=cl-1 && grid[r][c]==1{
			grid[r][c]=2
			queue=append(queue,[2]int{r,c})
			change=true
		}
		r-=v[0]
		c-=v[1]
	}
	

}
if change {time++}
}

for i:=0;i<rl;i++{
	for j:=0;j<cl;j++{
		if grid[i][j]==1{
			return -1
		}
	}
}
return time


}
