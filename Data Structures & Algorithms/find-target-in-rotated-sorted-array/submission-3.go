func search(nums []int, target int) int {
	l:=0
	r:=len(nums)-1
	mid:=0

	for l<=r{
		mid=l+(r-l)/2
		if nums[mid]<nums[r]{
			r=mid
		}else if nums[mid]>nums[r]{
			l=mid+1
		}else{
			break
		}
	}
	//mid is the index of the smallest value aka the pivot 
	println(mid)
	l1:=mid
	r1:=len(nums)-1
	l2:=0
	r2:=mid-1

	for l1<=r1{
		mid1:=l1+(r1-l1)/2

		if target>nums[mid1]{
			l1=mid1+1
		}else if target<nums[mid1]{
			r1=mid1-1
		}else{
			return mid1
		}
	}

	for l2<=r2{
		mid2:=l2+(r2-l2)/2

		if target>nums[mid2]{
			l2=mid2+1
		}else if target<nums[mid2]{
			r2=mid2-1
		}else{
			return mid2
		}
	}
	return -1
	
}
