func pacificAtlantic(heights [][]int) [][]int {
    qp:=[][2]int{}
	qa:=[][2]int{}
	rs:=len(heights)
	cs:=len(heights[0])

	for c:=0;c<cs;c++{
		qp=append(qp,[2]int{0,c})
		qa=append(qa,[2]int{rs-1,c})
	}

	for r:=1;r<rs;r++{
		qp=append(qp,[2]int{r,0})
	}
	for r:=0;r<rs-1;r++{
		qa=append(qa,[2]int{r,cs-1})
	}

	directions:=[][]int{{1,0},{-1,0},{0,1},{0,-1}}
	res:=[][]int{}
	p:=[][2]int{}
	for _,pair:=range qp{
		p=append(p,pair)
	}
	
	seenp:=make(map[[2]int]bool)
	
	      
	for _, pair := range qp {
		seenp[pair] = true
	}

	for len(qp)!=0{
		cv:=qp[0]
		qp=qp[1:]
		r:=cv[0]
		c:=cv[1]
		
		for _,d:=range directions{
			nr:=r+d[0]
			nc:=c+d[1]
			if nr<0 || nc<0 || nr>=rs || nc>=cs || heights[nr][nc] < heights[r][c]{
				continue
			}

			
			if !seenp[[2]int{nr,nc}]{
			p=append(p,[2]int{nr,nc})
			qp=append(qp,[2]int{nr,nc})
			}
			seenp[[2]int{nr,nc}]=true


		}

	}
	a:=[][2]int{}
	seena:=make(map[[2]int]bool)
	
	      
	for _, pair := range qa {
		seena[pair] = true
	}
	for _,pair:=range qa{
		a=append(a,pair)
	}
	for len(qa)!=0{
		cv:=qa[0]
		qa=qa[1:]
		r:=cv[0]
		c:=cv[1]
		
		for _,d:=range directions{
			nr:=r+d[0]
			nc:=c+d[1]
			if nr<0 || nc<0 || nr>=rs || nc>=cs || heights[nr][nc] < heights[r][c]{
				continue
			}

			
			if !seena[[2]int{nr,nc}]{
			a=append(a,[2]int{nr,nc})
			qa=append(qa,[2]int{nr,nc})
			}
			seena[[2]int{nr,nc}]=true


		}

	}
	//now I have p and a 
	for _,pairp:=range p{
		for _,paira:=range a{
			if pairp==paira{
				res = append(res, []int{pairp[0], pairp[1]})
			}
		}
	}
	return res
	
	
}
