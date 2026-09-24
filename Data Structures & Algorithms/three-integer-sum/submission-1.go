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
			if j-1>i && nums[j]==nums[j-1]{
				j++
				continue
			}
			if k+1<len(nums) && nums[k]==nums[k+1]{
				k--
				continue
			}

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
