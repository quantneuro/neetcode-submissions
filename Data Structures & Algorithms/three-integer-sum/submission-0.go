import "slices"
func threeSum(nums []int) [][]int {
	ans:=[][]int{}
	slices.Sort(nums)

	i:=0
	
	for ;i<=len(nums)-3;i++{
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		j:=i+1
		k:=len(nums)-1
		
		for j<k {
			if 0>nums[i]+nums[j]+nums[k]{
				j++
			}else if 0< nums[i]+nums[j]+nums[k]{
				k--
			} else {
				ans=append(ans,[]int{nums[i],nums[j],nums[k]})
				j++
				k--
			}
		}
		
		
	}
	return ans
}
