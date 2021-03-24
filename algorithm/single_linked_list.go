package algorithm

import (
	"strconv"
	"strings"
)

type Node struct {
	Next  *Node
	Value int
}

func ConstructList(es []int) *Node {
	if len(es) <= 0 {
		return nil
	}
	head := &Node{
		Value: es[0],
	}
	cur := head
	for _, e := range es[1:] {
		cur.Next = &Node{
			Value: e,
		}
		cur = cur.Next
	}
	return head
}

func FormatList(head *Node) string {
	if head == nil {
		return ""
	}
	sb := strings.Builder{}
	for head != nil {
		sb.WriteString(" -> ")
		sb.WriteString(strconv.FormatInt(int64(head.Value), 10))
		head = head.Next
	}
	return sb.String()
}

//翻转单链表
func Reverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	cur := head.Next
	next := cur
	prev.Next = nil
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

/*
L0 -> L1 -> L2 -> L3 -> L4 -> ,,, -> Ln-1 -> Ln
L0 -> Ln -> L1 -> Ln-1 -> L2 -> Ln-2
*/

func ReorderList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	p1, p2 := head, head

	for p1 != nil {
		p2 = p2.Next

		p1 = p1.Next
		if p1 != nil {
			p1 = p1.Next
		}
	}

	p3 := head
	for p3.Next != p2 {
		p3 = p3.Next
	}
	p3.Next = nil

	p2 = Reverse(p2)
	p1 = head

	for p2 != nil && p1 != nil {
		p3 = p2.Next
		p2.Next = p1.Next
		p1.Next = p2

		p1 = p1.Next.Next
		p2 = p3
	}

	return head
}
