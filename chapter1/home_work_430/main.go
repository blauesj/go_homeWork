package main

import "fmt"

func main() {
	node3 := Node{Val: 11, Next: &Node{Val: 12}}
	node2 := Node{Val: 7, Next: &Node{Val: 8, Next: &Node{Val: 9, Next: &Node{Val: 10}}}, Child: &node3}
	root := Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Child: &node2}}}
	a := flatten(&root)
	fmt.Println(a)
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	dealNode(root)
	return root
}

func dealNode(node *Node) (last *Node) {
	cur := node
	for cur != nil {
		if cur.Child != nil {
			last = dealNode(cur.Child)
			next := cur.Next
			cur.Next = cur.Child
			cur.Child.Prev = cur
			cur.Child = nil

			last.Next = next
			if next != nil {
				next.Prev = last
			}
		} else {
			last = cur
		}
		cur = cur.Next
	}
	return last
}
