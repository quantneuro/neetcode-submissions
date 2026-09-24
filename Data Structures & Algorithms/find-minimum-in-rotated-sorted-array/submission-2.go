func findMin(nums []int) int {
	
	r:=len(nums)-1
	var l int =0
	var mid int =0
	for l<=r {
		mid=l+(r-l)/2
		if(nums[mid]>nums[r]){
			l=mid+1
		} else if(nums[mid]<nums[r]){
			r=mid
		}else{
			
			return nums[mid]
		}
	}
	return -1
}
