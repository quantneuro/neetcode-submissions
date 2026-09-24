// import("container/heap")
type someheap []int

func(h someheap)Len()int{return len(h)}
func(h someheap)Less(i,j int) bool{return h[i]<h[j]}
func(h someheap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func(h *someheap)Push(x any){
	val:=x.(int)
	*h=append(*h,val)
}
func(h *someheap)Pop ()any{
	old:=*h
	n:=len(old)
	val:=old[n-1]
	*h=old[:n-1]
	return val
}


type KthLargest struct {
    n []int
	kth int
}


func Constructor(k int, nums []int) KthLargest {
    return KthLargest{
		n:nums,
		kth:k,
	}
}


func (this *KthLargest) Add(val int) int {
	this.n=append(this.n,val)
    h:=someheap(this.n)
	heap.Init(&h)
	

	for len(h)>this.kth{
		heap.Pop(&h)
	}
	value:=heap.Pop(&h).(int)
	return value
}
