func topKFrequent(nums []int, k int) []int {
 freq:=make(map[int]int)
 n:=len(nums)


 for _,val:=range nums{
	freq[val]++
 }
//  bucketarray:=[n][]{}//number of rows here are same as size of nums each row number is a frequency
 //can't but n in runtime here
 //Just use make 
 bucketarray:=make([][]int,n+1)//n+1 beacuae say a nums{2,2,2} size is 3 and k=1 and o we have 2 with freq 3
 //now since we store index as frequency we need the 3rd index as well, if we had done only n then index would end at 2{0,1,2} , so need that n+1{0,1,2,3}

 for number,numfreq := range freq{
	bucketarray[numfreq]=append(bucketarray[numfreq],number)
 }

// answer:=[k]int{}
answer :=[]int{}
collected:=0

 for i:=n;i>=0;i--{
	if collected == k{
		break
	}
	for j:=0;j<	len(bucketarray[i]);j++{
		if collected == k {
			break
		}
		answer=append(answer,bucketarray[i][j])
		collected++
	}
}
return answer
}
