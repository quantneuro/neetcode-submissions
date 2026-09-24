func findMin(nums []int) int {
	
	r:=len(nums)-1
	var l int =0
	var mid int =0
	for l<=r {
		mid=l+(r-l)/2
		if(nums[mid]>nums[r]){
//Explaining see if mid was minimum then nums[right] needs to bigger than nums[mid], aka continuos? so we know now that mid is not the minimum so safe to skip 
			//hence we skip mid and go to mid+1
			l=mid+1
		} else if(nums[mid]<nums[r]){
			//why not skip here? becasue the condition say that nums[mid]<nums[r] hence it is continuos but what if mid was minimum? 
			//It does not give us the safey that mid can't be smallest, since we already know that is smaller for sure from r but sitll need to check if it
			//is the smallest so r =mid
			r=mid
		}else{
			
			return nums[mid]
		}
	}
	return -1
}
