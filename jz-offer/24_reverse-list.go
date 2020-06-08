package jz_offer

// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	//反转链表
	for head != nil {
		pre, head, head.Next = head, head.Next, pre
	}
	return pre
}
