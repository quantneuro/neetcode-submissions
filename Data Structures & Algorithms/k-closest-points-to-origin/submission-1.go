type someheap []float64

func(h someheap) Len()int{return len(h)}
func(h someheap) Less(i,j int)bool {return h[i]<h[j]}//smallest on top
func(h someheap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func(h *someheap) Push(x any){
	value:=x.(float64)
	*h=append(*h,value)
	}
func(h *someheap) Pop()any{
	old:=*h
	n:=len(old)
	value:=old[n-1]
	*h=old[:n-1]
	return value
}

func kClosest(points [][]int, k int) [][]int {

	h:=&someheap{}//h is address here
	heap.Init(h)
	check:=make(map[float64][][]int)

	for _,v:=range points{
		for i:=0;i<1;i++{
			x:=v[0]
			y:=v[1]
			//since other point 0,0 x2=0 y2=0
			val:=math.Sqrt(math.Pow(float64(x),2.0)+math.Pow(float64(y),2.0))
			heap.Push(h,val)
			check[val]=append(check[val],v)
		}
	}
	res:=[][]int{}
	for k>0{
		shortest:=heap.Pop(h).(float64)//.()int works here because interfaces
		v:=check[shortest]
		
		for len(v)>0 && k>0{
			res=append(res,v[0])
			v=v[1:]
			k--
			}
		delete(check,shortest)

//LOGIC If k gets satisfied before v empties, the loop breaks and never touches that key again — removing it or not makes no difference there.

//If v does become empty (all its points got used), nothing later needs check[shortest] again for that value, since you already gave out every point at that distance. So deleting it at that point is safe.

// ORR

//I can do
	}
	return res

}
