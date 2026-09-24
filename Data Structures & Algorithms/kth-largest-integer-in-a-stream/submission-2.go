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
    n someheap
	kth int
}


func Constructor(k int, nums []int) KthLargest {
	h:=someheap(nums)
	heap.Init(&h)
	obj:=KthLargest{
		n:h,
		kth:k,
	}
    return obj
}


func (this *KthLargest) Add(val int) int {
	heap.Push(&this.n,val)

	for len(this.n)>this.kth{
		heap.Pop(&this.n)
	}
	value:=heap.Pop(&this.n).(int)
	return value
}
