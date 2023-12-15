package sale

import (
	"fmt"
)

type ItemNode struct {
	Name	string
	Price	int
	next	*ItemNode
}

type Stack struct {
	head	*ItemNode
}

func Init() *Stack {
	return &Stack{nil}
}

func (s *Stack) Push(item string, price int) {
	if s.head == nil {
		s.head = &ItemNode{item, price, nil}
	} else {
		curr := s.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = &ItemNode{item, price, nil}
	}
}

func (s *Stack) Pop() *ItemNode {
	if s.head == nil {
		return nil
	} else {
		curr := s.head
		prev := curr

		if curr.next == nil {
			s.head = nil
			return curr 
		}

		for curr.next != nil {
			prev = curr
			curr = curr.next
			
		}

		fmt.Println(curr, prev)

		prev.next = nil
		return curr 
	}
}

func (s *Stack) Peek() *ItemNode {
	if s.head == nil {
		return nil
	} else {
		curr := s.head
		for curr.next != nil {
			curr = curr.next
		}
		return curr
	}
}

func (s *Stack) Display() {
	curr := s.head

	for curr != nil {
		fmt.Printf("Item: %s, Sale: %d\n", curr.Name, curr.Price)
		curr = curr.next
	}
}
