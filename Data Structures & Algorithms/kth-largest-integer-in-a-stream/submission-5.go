type someheap []int

func (s someheap) Len()int{return len(s)}
func (s someheap) Swap(i,j int){s[i],s[j]=s[j],s[i]}
func (s someheap) Less(i,j int)bool{return s[i]<s[j]}
func (s *someheap)Push(x any){
	val:=x.(int)
	*s=append(*s,val)
}
func (s *someheap)Pop()any{
	old:=*s
	n:=len(old)
	val:=old[n-1]
	*s=old[:n-1]
	return val
}


type KthLargest struct {
    h someheap
	k int
}


func Constructor(k int, nums []int) KthLargest {

	// var stream KthLargest

	// for _,v:=range nums{
	// 	stream.h=append(stream.h,v)
	// }

	stream:=someheap(nums)
	heap.Init(&stream)
	obj:=KthLargest{
		h:stream,
		k:k,
	}

	return obj
}


func (this *KthLargest) Add(val int) int {
	heap.Push(&this.h,val)
	for len(this.h)>this.k{
		heap.Pop(&this.h)
	}
	return this.h[0]
 
}
