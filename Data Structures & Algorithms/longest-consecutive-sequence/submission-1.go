func longestConsecutive(nums []int) int {
	check:=make(map[int]bool)

	

	for _,val:=range nums{
		check[val]=true
	}
	maxx:=0
	
	for _,value:= range nums{
		count:=0
		if !check[value-1]{
			count++
			for check[value +1]{
				value++
				count++
                maxx=max(count,maxx)
			}
		}
		
		}
		return maxx

}


