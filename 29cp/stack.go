package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	data := s.top.data
	s.top = s.top.next
	return data, true
}

func (s *Stack) Push(data int) {
	n := &Node{data, s.top}
	s.top = n
}
func (s *Stack) Peek() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	return s.top.data, true
}

func main() {
	s := Stack{}
	val, pres := s.Peek()
	if pres != false {
		fmt.Println("value is ", val)
	}
	s.Push(10)
	val, pres = s.Peek()
	if pres != false {
		fmt.Println("value is ", val)
	}
	s.Push(20)
	s.Push(30)
	s.Push(40)
	val, pres = s.Peek()
	if pres != false {
		fmt.Println("value is ", val)
	}
	val, pres = s.Pop()
	fmt.Println("Popped value is ", val)
	val, pres = s.Pop()
	fmt.Println("Popped value is ", val)
	val, pres = s.Pop()
	fmt.Println("Popped value is ", val)

}
