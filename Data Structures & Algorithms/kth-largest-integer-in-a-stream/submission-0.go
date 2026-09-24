// import "container/heap" 
type heapp []int

func(h heapp) Len() int{ return len(h)}
func(h heapp) Less(i,j int) bool { return h[i]<h[j]}
func(h heapp) Swap(i,j int){ h[i],h[j]=h[j],h[i]}
func(h *heapp) Push(x any){
	value :=x.(int)
	*h = append(*h, value)
} 
func(h *heapp) Pop()any{
	old:=*h
	n:=len(old)
	value:=old[n-1]
	*h=old[:n-1]
	return value

}

type KthLargest struct {
	k int 
	nums []int
	hp heapp
    
}


func Constructor(k int, nums []int) KthLargest {

	kthlargest:=KthLargest{
		k:k,
		nums:nums,
		hp:heapp(nums),
	}
	heap.Init(&kthlargest.hp)
	return kthlargest
    
}


func (this *KthLargest) Add(val int) int {
	heap.Push(&this.hp,val)//the & gives the address of the hp
	for len(this.hp)>this.k{
		heap.Pop(&this.hp)
	}
	return this.hp[0]
	
}
