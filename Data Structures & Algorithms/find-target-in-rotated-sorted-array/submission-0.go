func search(nums []int, target int) int {
	n:=len(nums)
	r:=n-1
	var l int = 0
	var mid int 
	start2:=0
	
	for l<=r{
		mid=l+(r-l)/2
		if nums[mid]<nums[r] {
			r=mid
		}else if nums[mid]>nums[r] {
			l=mid+1

		}else{
			 start2=mid
			 l++
			 r=mid
		}
	}

	right2:=n-1
	right1:=(start2-1+n)%n
	start1:=0
	mid1:=0
	mid2:=0

	for start1<=right1{
		mid1=start1+(right1-start1)/2
		if nums[mid1]<target{
			start1=mid+1
		}else if nums[mid1]>target{
			right1=mid1-1
		}else{
			return mid1
		}
	}

	for start2 <=right2{
		mid2= start2+(right2-start2)
		if nums[mid2]<target{
			start2=mid2+1
		}else if nums[mid2]>target{
			right2=mid2-1
		}else{
			return mid2
		}
	}

return -1
	
}
