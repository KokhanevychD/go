package problem

// leetcode Add Two Numbers.
// Runtime: 12 ms, faster than 70.07% of Go online submissions for Add Two Numbers.
// Memory Usage: 6 MB, less than 5.01% of Go online submissions for Add Two Numbers.

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers return some shit
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result []int
	temp := 0
	for true {
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		result = append([]int{temp % 10}, result[0:]...)
		if temp/10 > 0 {
			temp /= 10
		} else {
			temp = 0
		}
		if l1 == nil && l2 == nil && temp == 0 {
			break
		}
	}
	var resLinkedlist *ListNode

	for idx, num := range result {
		if idx > 0 {
			resLinkedlist = &ListNode{num, resLinkedlist}
		} else {
			resLinkedlist = &ListNode{num, nil}
		}

	}
	return resLinkedlist
}
