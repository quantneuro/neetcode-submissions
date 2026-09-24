
type dllNode struct{
	Key int
	Value int
	Next *dllNode
	Prev *dllNode

}

type LRUCache struct {
	cache map[int]*dllNode 
    capacity int 
	head *dllNode
	tail*dllNode
	// size int// No need since len(cache) will work
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		cache:make(map[int]*dllNode),
		capacity:capacity,
		head:nil,
		tail:nil,
	}
}

func (this *LRUCache) Get(key int) int {

	// Find the node directly using the map
	address, exists := this.cache[key]

	// Key does not exist
	if !exists {
		return -1
	}

	// Node is already the most recently used
	// No DLL movement required
	if address == this.head {
		return address.Value
	}

	// Node is currently the least recently used (tail)
	if address == this.tail {
		temp := address

		// Move tail one node backward
		this.tail = this.tail.Prev
		this.tail.Next = nil

		// Disconnect old tail from its previous position
		temp.Prev = nil

		// Put this node before the current head
		temp.Next = this.head
		this.head.Prev = temp

		// This node is now the most recently used
		this.head = temp

		return temp.Value
	}

	// If we reach here, node is somewhere in the middle
	temp := address

	// Bypass temp from its current position
	temp.Prev.Next = temp.Next
	temp.Next.Prev = temp.Prev

	// Disconnect temp from its old position
	temp.Prev = nil
	temp.Next = nil

	// Attach temp before the current head
	temp.Next = this.head
	this.head.Prev = temp

	// temp is now the most recently used
	this.head = temp

	return temp.Value
}

func (this *LRUCache) Put(key int, value int) {
    
	if this.head ==nil {
		//head is a pointer storing a node
		this.head=&dllNode{
			Key:key,
			Value:value,
		}
		this.tail = this.head

		//update map
		this.cache[key]=this.head
		
		return 
	}
	address,e:=this.cache[key]

	if e && address==this.head && this.tail==address{
		this.head.Value=value
		return
	}

	if e && address==this.head{
		this.head.Value=value
		return 
	}

	if e && address==this.tail{
		temp:=this.tail
		temp.Value=value
		// temp.Key=key
		this.tail=this.tail.Prev
		this.tail.Next=nil
		temp.Prev=nil
		
		temp.Next=this.head
		this.head.Prev=temp
		this.head=temp
		this.cache[key]=temp
		return 

	}

	
	
	if e {
		temp:=&dllNode{}
		temp=address
		temp.Value=value
		temp.Prev.Next=temp.Next
		temp.Next.Prev=temp.Prev
		temp.Prev=nil
		temp.Next=nil

		this.head.Prev = temp
		temp.Next = this.head
		this.head = this.head.Prev 
		this.cache[key]=temp
		return 

		
	}else{
		temp:=dllNode{}
		temp.Value=value
		temp.Key=key

		if len(this.cache)>=this.capacity{
			delete(this.cache,this.tail.Key)
			if this.capacity == 1{
				this.head=nil
				this.tail=nil
				
				this.head=&temp
				this.tail=&temp
				this.cache[key]=&temp
				return 

			}else{
			this.tail=this.tail.Prev
			this.tail.Next=nil
			}
		}
		
		this.head.Prev = &temp
		temp.Next = this.head
		this.head = this.head.Prev 
		this.cache[key]=&temp
		return 
		
	}

	
}
