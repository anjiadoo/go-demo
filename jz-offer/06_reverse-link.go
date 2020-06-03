package jz_offer

/**
 * 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	pre := new(ListNode)
	//反转链表
	for head != nil {
		pre, head, head.Next = head, head.Next, pre
	}

	//遍历链表
	nums := make([]int, 0) //申请一个数组
	for pre.Next != nil {
		nums = append(nums[:], pre.Val)
		pre = pre.Next
	}
	return nums
}
