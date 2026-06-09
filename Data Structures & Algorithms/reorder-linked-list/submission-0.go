/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    // find middle (slow ends at middle)
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse second half
	var prev *ListNode
	curr := slow.Next
	slow.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// Merge two halves
	first, second := head, prev
	for second != nil {
		fNext, sNext := first.Next, second.Next
		first.Next = second
		second.Next = fNext
		first = fNext
		second = sNext
	}
}