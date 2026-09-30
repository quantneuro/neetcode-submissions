func topKFrequent(nums []int, k int) []int {
	n:=len(nums)
	buckets:=make([][]int,n+1)
	// for i,_:=range buckets{
	// 	buckets[i]=make([]int,n+1)
	// }
	freq:=make(map[int]int)
	for _,v:=range nums{
		freq[v]++
	}
	for key,value:=range freq{
		buckets[value]=append(buckets[value],key)
	}

	
	res:=[]int{}

	for i:=len(buckets)-1;i>0;i--{
		
		for _,val:=range buckets[i] {
			res=append(res,val)
			if len(res)==k{
				return res
			}
		}
		
	}
	return res
}
