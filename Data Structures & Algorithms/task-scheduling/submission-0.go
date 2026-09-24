type item struct{
	task byte
	frequency int
	cooldown int
}
type someheap []item

func (h someheap)Len()int{return len(h)}
func (h someheap)Less(i,j int) bool{ return h[i].frequency > h[j].frequency}
func (h someheap)Swap(i,j int) {h[i],h[j]=h[j],h[i]}
func (h *someheap)Push(x any){
	value:=x.(item)
	*h=append(*h,value)
}
func(h *someheap) Pop()any{
	old:=*h
	n:=len(old)
	value:=old[n-1]
	*h=old[:n-1]
	return value
}

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	freq :=make(map[byte]int)
	
	for _,v:=range tasks{
		freq[v]++
	}
	var h someheap
	for task,count:=range freq{
		
		var temp item
		temp=item{
			task:task,
			frequency:count,
			cooldown:0,
		}
		
		h=append(h,temp)//remember append should be before initialization 
	}
	heap.Init(&h)

	cycles:=0
	queue:=[]item{}
	for len(h) != 0 || len(queue) != 0 {
		var current item
		if len(h)!=0{
			current=heap.Pop(&h).(item)
		}
		cycles++
		current.frequency--

		for i:=len(queue)-1;i>=0;i--{//if I do i,it :=range queue , then it is just a copy ! so any changes on it won't work 
			queue[i].cooldown--
			if queue[i].cooldown==0{
				heap.Push(&h,queue[i])
				queue=append(queue[:i],queue[i+1:]...)
			}
		}
		if current.frequency>0{
			current.cooldown=n
			queue=append(queue,current)
		}
	}
	return cycles
}
