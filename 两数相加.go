package leetcode刷题

type ListNode struct {
	     Val int
	     Next *ListNode
	}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil{
		return nil
	}
	var carry int
	head := &ListNode{}
	current := head
	var x,y int
	for l1 !=nil || l2 != nil {
		if l1 == nil {
			x = 0
		} else {
			x  = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{(x+y+carry)%10,nil}
		current = current.Next
		carry = (x+y+carry)/10
	}
	if carry > 0 {
		current.Next = &ListNode{Val:carry%10,Next:nil}
	}
	return head.Next

}
