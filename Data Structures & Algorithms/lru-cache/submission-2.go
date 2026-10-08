type LRUCache struct {
    capacity int
	cache map[int]*list.Element
	list *list.List
}
type pair struct{
  key,val int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		capacity:capacity,
		cache: make(map[int]*list.Element),
		list:list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	//check map
	if v,ok := this.cache[key];ok{
		this.list.MoveToFront(v)
		return v.Value.(pair).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if v,ok := this.cache[key];ok{
		this.list.MoveToFront(v)
		v.Value = pair{key,value}//positional filling 
		return
	}
	newelement:=this.list.PushFront(pair{key,value})//pushfront returns a pointer to the new element/thing made
	this.cache[key]=newelement

	if this.list.Len()>this.capacity{
		lru:=this.list.Back()
		this.list.Remove(lru)
		delete(this.cache,lru.Value.(pair).key)//every value in map is element and elemnt is any type stored
		//so back.Value is any type we do back.Value.(entry) to flip it to entry type , then we take key from it 
	}



}

// type DLLNode struct{
// 	key int
// 	value int 
// 	prev *DLLNode
// 	next *DLLNode
// }

// type LRUCache struct {
// 	cache map[int]*DLLNode
// 	c int
// 	head *DLLNode
// 	tail *DLLNode
// }


// func Constructor(capacity int) LRUCache {
// 	lrucache:=LRUCache{
// 		cache: make(map[int]*DLLNode),
// 		c:capacity,
// 		head:&DLLNode{},
// 		tail:&DLLNode{},
// 	}
// 	lrucache.head.next=lrucache.tail
// 	lrucache.tail.prev=lrucache.head
//     return lrucache
// }
// func (this *LRUCache) Remove(current *DLLNode){
// 	nodenext:=current.next
// 	nodeprev:=current.prev

// 	current.prev.next=nodenext
// 	current.next.prev=nodeprev

// 	current.prev=nil
// 	current.next=nil
	
// }
// func (this *LRUCache) Insert(current *DLLNode){
// 	prevofheadsnext := this.head.next.prev
// 	headnext := this.head.next
// 	//newnode udpate
// 	current.prev=prevofheadsnext
// 	current.next=headnext
// 	//head update
// 	this.head.next=current
// 	// prevofheadsnext=current // changes only local value
// 	headnext.prev=current
// }

// func (this *LRUCache) Get(key int) int {
// 	keyaddress,ok:=this.cache[key]
//     if !ok {
// 		return -1
// 	}
// 	this.Remove(keyaddress)
// 	this.Insert(keyaddress)
// 	return keyaddress.value


// }

// func (this *LRUCache) Put(key int, value int) {
// 	keyaddress,ok:=this.cache[key]
//     if ok {
// 		this.Remove(keyaddress)
// 		delete(this.cache,key)
// 	}
// 	newnode:=&DLLNode{key:key,value:value,}
// 	this.cache[key]=newnode
// 	this.Insert(newnode)

// 	if len(this.cache)>this.c{
// 		lru:=this.tail.prev//since tail is dummy just before it is lru
// 		this.Remove(lru)
// 		delete(this.cache,lru.key)
// 	}
    
// }


