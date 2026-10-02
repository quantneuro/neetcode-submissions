func maxProduct(nums []int) int {

	maxsofar:=nums[0]

	for i:=0;i<len(nums);i++{
		cur:=1
		for j:=0;j<len(nums);j++{
			cur*=nums[j]
			
			maxsofar=max(cur,maxsofar)
		}
	}
	return maxsofar
    
}
