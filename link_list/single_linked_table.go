package link_list

import "fmt"

type Node struct {
	data int
	next *Node
}

func InsertAtHead(head **Node, data int) {
	newNode := &Node{
		data: data,
	}
	if *head == nil {
		*head = newNode
		return
	}
	current := *head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}
func DeleteNode(head **Node, data int) {
	if *head == nil {
		return
	}
	if (*head).data == data {
		*head = (*head).next
		return
	}
	current := *head
	for current.next != nil && current.next.data != data {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}
func PrintList(head *Node) {
	for head != nil {
		fmt.Printf("%d--->", head.data)
		head = head.next
	}
	fmt.Println("nil")

}
