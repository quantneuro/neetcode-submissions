// func topKFrequent(nums []int, k int) []int 
 
//  bucketarray:=[n][]{}//number of rows here are same as size of nums; each row number is a frequency
 //can't but n in runtime here
 //Just use make 
 
 //n+1 beacuae say a nums{2,2,2} size is 3 and k=1 and o we have 2 with freq 3
 //now since we store index as frequency we need the 3rd index as well, if we had done only n then index would end at 2{0,1,2} , so need that n+1{0,1,2,3}


    //since index is frequency, FYI : frequncy can't be bigger than the size of the array itself, buckeyarray index is the number's frequency because if we have multiple numbers with same frequency we can just push them  
 

// answer:=[k]int{}

func topKFrequent(nums []int, k int) []int {
	freq:=make(map[int]int)
	n:=len(nums)

	for _,val:=range nums{
		freq[val]++
	}


	//make the bucket
	bucket:=make([][]int,n+1)

	for number,numberfreq:=range freq{
		bucket[numberfreq]=append(bucket[numberfreq],number)
	}

	answer:=[]int{}
	collected:=0
	for i:=len(bucket)-1;i>0;i--{
		if collected==k{
			break
		}
		for j:=0;j<len(bucket[i]);j++{
			answer=append(answer,bucket[i][j])
			collected++
		}

	}
	return answer


}