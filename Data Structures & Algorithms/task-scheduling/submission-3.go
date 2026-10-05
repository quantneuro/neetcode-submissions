type task struct{
	taskid byte
	frequency int
	cooldown int

}

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




func leastInterval(tasks []byte, n int) int {

	var least someheap
	
	freq:=make(map[byte]int)

	for _,val:=range tasks{
		freq[val]++
	}

	for k,v := range freq{
		least=append(least,task{
			taskid:k,
			frequency: v,
			cooldown:0,
		})
	}
	
	heap.Init(&least)

	count:=0
	queue:=[]task{}
	for len(least)!=0 || len(queue)!=0{
		var current task
		if len(least)!=0{
		current=heap.Pop(&least).(task)
		}
		count++
		if current.frequency > 0{current.frequency--}
		
		for i:=len(queue)-1;i>=0;i--{
			queue[i].cooldown--
			if queue[i].cooldown==0{
				heap.Push(&least,queue[i])
				queue=append(queue[:i],queue[i+1:]...)
			}
		}

		if current.frequency>0{
			if n==0{
			heap.Push(&least,current)
			}else{
			current.cooldown=n
			queue=append(queue,current)
			}
		}
		
		

		
	}
	return count


}



