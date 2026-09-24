func findDuplicate(nums []int) int {
	slow:=0
	fast:=0
    for fast<len(nums){
		slow=nums[slow]
		fast=nums[nums[fast]]

		if fast == slow {
			slow=0
			for {
				slow=nums[slow]
				fast=nums[fast]
				if slow ==fast{
					break
				}
			}
			return slow
		}
	}
	return -1
}
