func search(nums []int, target int) int {
 mid:=0
 l:=0
 r:=len(nums)-1

 for l<=r{
	mid=l+(r-l)/2

	if(target>nums[mid]){
		l=mid+1
	}else if(target<nums[mid]){
		r=mid-1
	}else{
		return mid
	}
 }
 return -1
}
