package linkedlist

import "testing"

/*
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
*/
func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	current := head

	for current != nil {
		next = current.Next // Lưu trữ node tiếp theo
		current.Next = prev // Đảo ngược con trỏ
		prev = current      // Di chuyển prev lên node hiện tại
		current = next      // Di chuyển current lên node tiếp theo
	}

	return prev // prev sẽ là node mới đầu tiên của linked list
}

func Test_reserveList(t *testing.T) {
	listNode := []int{1, 2, 3, 4, 5}
	head := createList(listNode)
	printList(reverseList(head))
}
