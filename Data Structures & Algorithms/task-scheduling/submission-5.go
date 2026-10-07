type someheap []task

func(h someheap)Len()int{return len(h)}
func(h someheap)Less(i,j int)bool{return h[i].frequency>h[j].frequency}
func(h someheap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func(h *someheap)Push(x any){
	val:=x.(task)
	*h=append(*h,val)
}
func(h *someheap)Pop()any{
	old:=*h
	n:=len(old)
	val:=old[n-1]
	*h=old[:n-1]
	return val
} 

type task struct{
	taskid byte
	frequency int
	cooldown int

}
func leastInterval(tasks []byte, n int) int {
	t:=[]task{}
	freq:=make(map[byte]int)
	for _,v:=range tasks{
		freq[v]++
	}
	for k,v := range freq{
		
		t=append(t,task{
			taskid:k,
			frequency:v,
			cooldown:0,
		},
		)
	}

	h:=someheap(t)
	heap.Init(&h)

	cycles:=0
	queue:=[]task{}
	for len(h)!=0 ||len(queue)!=0{
		var currenttask task
		hasTask := false
		if len(h)!=0 {
			currenttask=heap.Pop(&h).(task)
			currenttask.frequency--
			hasTask = true
		}
		
	
		for i:=len(queue)-1;i>=0;i--{

				queue[i].cooldown--
			
			if queue[i].cooldown==0{
				heap.Push(&h,queue[i])
				queue=append(queue[:i],queue[i+1:]...)
			}
		}
		

		if hasTask && currenttask.frequency>0{
			
			if n==0{
				heap.Push(&h,currenttask)

			}else{
				currenttask.cooldown=n
				queue=append(queue,currenttask)
			}
		}
		cycles++
		
	}
	return cycles

}
