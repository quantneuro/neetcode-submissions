func trap(height []int) int {
	l:=0
	r:=len(height)-1
	
	totalwater:=0
	lm,rm:=height[l],height[r]
	
	for l<r { 
		if lm<rm{
			l++
			lm=max(lm,height[l])
			totalwater += lm - height[l]
			
		}else if rm<=lm{
			r--
			rm=max(rm,height[r])
			totalwater += rm - height[r]
			
		}
	}
	return totalwater 
	
}
