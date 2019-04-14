# HARD - LRU Cache [![Build Status](https://api.travis-ci.org/arian-amador/GoLeetCode.svg)](https://travis-ci.org/arian-amador/GoLeetCode)

### Problem

Design and implement a data structure for `Least Recently Used (LRU) cache`. It should support the following operations: get and put.

`get(key)` - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.

`put(key, value)` - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

##### Example

```Go
LRUCache cache = new LRUCache( 2 /* capacity */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
```

### Solutions

The key concept learned from this exercise is the ability to use a linked list to maintain an order along with a lookup table for fast lookups/access. The lookup table allows us to map a `key` to a memory address for a `node` on the linked list.

```Go
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
```

The linked list needs to be able to look at both the elements before and after. This is the reason we need it to be a doubly linked list.

An LRU cache allocates the max amount of memory up front. Once we've reached our capacity the least recently used node on the list. This node is the one closest to the tail of the list.

```Go
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
```

Elements that are accessed using Get() or updated/created using Put() are updated to the front of the list.

##### Additional Resources

- [Cache replacement policies](https://www.wikiwand.com/en/Cache_replacement_policies)

---

#### Hope you find this useful!
