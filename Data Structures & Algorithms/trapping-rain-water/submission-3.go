func trap(height []int) int {
	if len(height) == 0{
		return 0
	}
l:=0
r:=len(height)-1
lmax:=height[l]
rmax:=height[r]
totalwater :=0
for l<r{
	
	if height[l]<height[r]{
		totalwater+=min(lmax,rmax)-height[l]
		l++
		lmax=max(height[l],lmax)
		
	}else {
		totalwater+=min(lmax,rmax)-height[r]
		r--
		rmax=max(height[r],rmax)
		
	}


}
return totalwater
	
}
