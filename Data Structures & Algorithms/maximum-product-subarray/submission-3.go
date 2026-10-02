func maxProduct(nums []int) int {

	maxsofar:=nums[0]
	minsofar:=nums[0]
	res:=nums[0]

	for i:=1;i<len(nums);i++{
		
		currentmax:=maxsofar*nums[i]//current value with last max value
		currentmin:=minsofar*nums[i]
		currrentvalue:=nums[i]
		maxsofar=max(currrentvalue,max(currentmax,currentmin))
		minsofar=min(currrentvalue,min(currentmax,currentmin))

		res=max(maxsofar,res)
		
	}
	return res
    
}


// func maxProduct(nums []int) int {

// 	maxsofar:=nums[0]

// 	for i:=0;i<len(nums);i++{
// 		cur:=1
// 		for j:=0;j<len(nums);j++{
// 			cur*=nums[j]
			
// 			maxsofar=max(cur,maxsofar)
// 		}
// 	}
// 	return maxsofar
    
// }
