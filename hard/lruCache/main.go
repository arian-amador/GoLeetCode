package main

import "fmt"

func main() {
	cap := 2
	cache := Constructor(cap)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)    // returns 1
	cache.Put(3, 3) // evicts key 2
	cache.Get(2)    // returns -1 (not found)
	cache.Put(4, 4) // evicts key 1
	cache.Get(1)    // returns -1 (not found)
	cache.Get(3)    // returns 3
	cache.Get(4)    // returns 4
	cache.debug()
}

type LRUCache struct {
	max   int
	stack map[int]*Node
	head  *Node
	tail  *Node
}

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

func Constructor(capacity int) LRUCache {
	if capacity <= 0 {
		panic("Capacity needs to be greater than 0")
	}

	// Setup initial head/tail nodes
	hNode := &Node{key: -1}
	tNode := &Node{key: -1}
	hNode.next = tNode
	tNode.prev = hNode

	return LRUCache{
		max:   capacity,
		stack: make(map[int]*Node),
		head:  hNode,
		tail:  tNode,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.stack[key]; ok {
		this.insert(v)
		return v.val
	}

	return -1
}

func (this *LRUCache) Put(key, value int) {
	// We already have an entry in the lookup table
	if node, ok := this.stack[key]; ok {
		node.val = value
		this.insert(node)
		return
	}

	// Create a new node to add to the List
	node := &Node{
		key: key,
		val: value,
	}

	// Add the new [key]nodeAddress to the lookup table
	this.stack[key] = node
	this.insert(node)
}

func (this *LRUCache) insert(node *Node) {
	// First node the head points too
	hNode := this.head.next

	// Head already points to this node [Head <-> node]
	if hNode == node {
		return
	}

	// No nodes added yet and [Head <-> Tail]
	if hNode == this.tail {
		node.prev = this.head
		node.next = this.tail

		this.head.next = node
		this.tail.prev = node
		return
	}

	// Repoint surrounding nodes to each other
	if node.next != nil && node.prev != nil {
		pNode := node.prev
		nNode := node.next

		pNode.next = nNode
		nNode.prev = pNode
	}

	// Point this node to head and the current first node
	node.prev = this.head
	node.next = hNode

	// Point previous first node to this node [node <-> hNode]
	hNode.prev = node

	// Point head to this node [Head <-> node]
	this.head.next = node // Set the head .next to this new node

	this.prune()
}

func (this *LRUCache) Len() int {
	return len(this.stack)
}

func (this *LRUCache) prune() {
	if this.Len() > this.max {
		// Last node pointed to by tail [tNode <-> tail]
		tNode := this.tail.prev
		// Node before tNode [pNode <-> tNode <-> tail]
		pNode := tNode.prev

		// Repoint pNode to tail and tail to pNode
		pNode.next = this.tail
		this.tail.prev = pNode

		// Remove pointers from tNode as it is now pruned
		// This is probably uneccessary but might help with GC?
		tNode.prev = nil
		tNode.next = nil

		// Remove tNodes key from the lookup table
		delete(this.stack, tNode.key)
	}
}

func (this *LRUCache) debug() {
	fmt.Printf("[ Stack Info ] Capacity: %d, Current Size: %d\n", this.max, len(this.stack))
	fmt.Printf("[ Stack Info ] Contents: %v\n", this.stack)
	fmt.Printf("[ Stack Info ] Head: %+v, Tail: %+v\n", this.head, this.tail)

	n := this.head.next
	for n.next != nil {
		fmt.Printf("  %+v\n", n)
		n = n.next
	}
	fmt.Println()
}
