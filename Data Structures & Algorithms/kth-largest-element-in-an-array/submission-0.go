type someheap []int 

func (h someheap)Len()int{return len(h)}
func (h someheap)Less(i,j int)bool{return h[i]<h[j]}
func (h someheap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *someheap) Push(x any){
	value:=x.(int)
	*h=append(*h,value)
}
func (h *someheap) Pop()any{
	old:=*h
	n:=len(old)
	value:=old[n-1]
	*h=old[:n-1]
	return value
}


func findKthLargest(nums []int, k int) int {
	h:=someheap(nums)
	heap.Init(&h)

	for len(h)>k{heap.Pop(&h)}
	return heap.Pop(&h).(int)

}
