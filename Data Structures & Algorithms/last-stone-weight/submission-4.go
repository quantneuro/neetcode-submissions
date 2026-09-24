type someheap []int

func(h someheap)Len()int{return len(h)}
func(h someheap)Less(i,j int)bool{return h[i]>h[j]}
func(h someheap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func(h *someheap)Push(x any){
	value:=x.(int)
	*h=append(*h,value)
}

func(h *someheap)Pop()any{
	old:=*h
	n:=len(old)
	value:=old[n-1]
	*h=old[:n-1]
	return value
}

func lastStoneWeight(stones []int) int {
	h:=someheap(stones)
	heap.Init(&h)
	for len(h)>1{
		s1:=heap.Pop(&h).(int)
		s2:=heap.Pop(&h).(int)

	if s1>s2{
		s1=s1-s2
		heap.Push(&h,s1)
	}
	}

	if len(h)==0{return 0}
	return h[0]


}
