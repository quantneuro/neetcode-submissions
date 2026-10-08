type DLLNode struct{
	key int
	value int 
	prev *DLLNode
	next *DLLNode
}

type LRUCache struct {
	cache map[int]*DLLNode
	c int
	head *DLLNode
	tail *DLLNode
}


func Constructor(capacity int) LRUCache {
	lrucache:=LRUCache{
		cache: make(map[int]*DLLNode),
		c:capacity,
		head:&DLLNode{},
		tail:&DLLNode{},
	}
	lrucache.head.next=lrucache.tail
	lrucache.tail.prev=lrucache.head
    return lrucache
}
func (this *LRUCache) Remove(current *DLLNode){
	nodenext:=current.next
	nodeprev:=current.prev

	current.prev.next=nodenext
	current.next.prev=nodeprev

	current.prev=nil
	current.next=nil
	
}
func (this *LRUCache) Insert(current *DLLNode){
	prevofheadsnext := this.head.next.prev
	headnext := this.head.next
	//newnode udpate
	current.prev=prevofheadsnext
	current.next=headnext
	//head update
	this.head.next=current
	// prevofheadsnext=current // changes only local value
	headnext.prev=current
}

func (this *LRUCache) Get(key int) int {
	keyaddress,ok:=this.cache[key]
    if !ok {
		return -1
	}
	this.Remove(keyaddress)
	this.Insert(keyaddress)
	return keyaddress.value


}

func (this *LRUCache) Put(key int, value int) {
	keyaddress,ok:=this.cache[key]
    if ok {
		this.Remove(keyaddress)
		delete(this.cache,key)
	}
	newnode:=&DLLNode{key:key,value:value,}
	this.cache[key]=newnode
	this.Insert(newnode)

	if len(this.cache)>this.c{
		lru:=this.tail.prev//since tail is dummy just before it is lru
		this.Remove(lru)
		delete(this.cache,lru.key)
	}
    
}









