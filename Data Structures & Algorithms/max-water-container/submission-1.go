func maxArea(heights []int) int {

	l:=0
	r:=len(heights)-1
	area:=0
	// current:=0
	for l<r{
		// current = (r-l)*min(heights[l],heights[r])
		// area=max(area,current)
        area=max(area,(r-l)*min(heights[l],heights[r]))
		if heights[l]<heights[r]{
			l++
		} else{
			r--
		}
	}
	return area
}
