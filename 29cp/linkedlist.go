package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

var root = new(Node)

func add(t *Node, val int) int {
	if root == nil {
		t = &Node{val, nil}
		root = t
		return 1
	}

	if t.next == nil {
		t.next = &Node{val, nil}
		return 1
	}

	return add(t.next, val)

}
func traverse(t *Node) {
	if t.next == nil {
		return
	}
	for t.next != nil {
		fmt.Println("value is ", t.val)
		t = t.next
	}

}

func main() {
	add(root, 10)
	add(root, 20)
	add(root, 30)
	add(root, 40)

	traverse(root)

}
