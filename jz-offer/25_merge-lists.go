package jz_offer

// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
//示例1：
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	dummnyhead := &ListNode{Val: -1, Next: nil}
	cur := dummnyhead
	if l1.Val <= l2.Val {
		cur.Next = l1
		cur = cur.Next
		cur.Next = mergeTwoLists(l1.Next, l2)
	} else {
		cur.Next = l2
		cur = cur.Next
		cur.Next = mergeTwoLists(l1, l2.Next)
	}
	return dummnyhead.Next
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1,}
	curr := dummyHead
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				curr.Next = l1
				curr = curr.Next
				l1 = l1.Next
			} else {
				curr.Next = l2
				curr = curr.Next
				l2 = l2.Next
			}
		} else if l1 != nil {
			curr.Next = l1
			break
		} else {
			curr.Next = l2
			break
		}

	}
	return dummyHead.Next
}
