package main

import "fmt"

const QUEUE_SIZE = 5

type Node struct {
	val   string
	left  *Node
	right *Node
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

type Hash map[string]*Node

type Cache struct {
	queue Queue
	hash  Hash
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.right = tail
	tail.left = head

	return Queue{
		head:   head,
		tail:   tail,
		length: 0,
	}
}

func NewCache() Cache {
	return Cache{
		queue: NewQueue(),
		hash:  make(Hash),
	}
}

func (c *Cache) Check(ele string) {
	node, ok := c.hash[ele]

	if ok {
		c.Remove(node)
	} else {
		node = &Node{val: ele}
	}
	c.Add(node)
	c.hash[ele] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove %s\n", n.val)

	left := n.left
	right := n.right

	right.left = left
	left.right = right
	c.queue.length--

	delete(c.hash, n.val)
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add %s\n", n.val)

	temp := c.queue.head.right
	n.left = c.queue.head
	n.right = temp
	c.queue.head.right = n
	temp.left = n

	c.queue.length++
	if c.queue.length > QUEUE_SIZE {
		c.Remove(c.queue.tail.left)
	}
}

func (c *Cache) Display() {
	c.queue.Display()
}

func (q *Queue) Display() {
	node := q.head.right
	fmt.Printf("%d ", q.length)

	for i := 0; i < q.length; i++ {
		fmt.Printf("%s", node.val)
		if i < q.length-1 {
			fmt.Printf(" <-> ")
		}
		node = node.right
	}
	fmt.Println()
}

func main() {
	fmt.Println("Cache Starting...")

	cache := NewCache()
	pushValue := []string{
		"Alice",
		"Bob",
		"Alice",
		"Charlie",
		"David",
		"Emma",
		"Frank",
		"Grace",
		"Henry",
		"Isabella",
		"Jack",
	}

	for _, ele := range pushValue {
		cache.Check(ele)
		cache.Display()
	}
}
