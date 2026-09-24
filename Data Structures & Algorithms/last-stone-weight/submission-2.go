type someheap []int

func(h someheap) Len()int{return len(h)}
func(h someheap) Less(i,j int)bool {return h[i]>h[j]}
func(h someheap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h*someheap)Push(x any){
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

func lastStoneWeight(stones []int) int {
	h:=someheap(stones)//var h someheap=someheap(stones)
	heap.Init(&h)

	for len(h)>1{
		v1:=heap.Pop(&h).(int)
		v2:=heap.Pop(&h).(int)
		
		if v2<v1{
			v1=v1-v2
			heap.Push(&h,v1)
		}else if v2==v1{}
	}

	if len(h)==0{return 0}
	return h[0]


}
