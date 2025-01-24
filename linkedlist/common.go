package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Hàm tiện ích để in ra linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

// Hàm tiện ích để tạo linked list từ slice
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}
