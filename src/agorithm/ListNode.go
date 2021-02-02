package main

import "fmt"

// 删除链表倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode{
	// 定义哨兵节点result
	result := &ListNode{}
	result.next = head
	// 目标指针cur
	cur := result
	// 目标指针cur的前一个指针pre
	var pre *ListNode
	i := 1
	for head != nil{
		// 删除倒数n的节点 可以按照删除正数的节点来思考，设置一个N-1间隔，当 >= n时，头节点也随之移动  理解成一个窗口
		if i >= n{
			pre = cur
			cur = cur.next
		}
		head = head.next
		i++
	}
	pre.next = pre.next.next
	return result.next
}

type ListNode struct {
	value int
	next *ListNode
}

// Create n就是0位,指向第一位
func (n *ListNode) Create(list []int) {
	cur := n
	for i := 1; i < len(list); i++ {
		//1.实例化node对象
		next := &ListNode{value: list[i]}
		//2.当前指向性生成的对象
		cur.next = next
		//3.当前对象更为最新
		cur = next
	}
}

// 打印
func (n *ListNode) testResult() {
	item := n.next
	fmt.Print(n.value)
	for item != nil {
		fmt.Print("->", item.value)
		item = item.next
	}
	fmt.Println()
}

// 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	result := preHead
	for l1 != nil && l2 != nil {
		if l1.value < l2.value {
			preHead.next = l1
			l1 = l1.next
		}else {
			preHead.next = l2
			l2 = l2.next
		}
		// 使当前元素下移，才能连接成一个链表，如果使用result 直接保存，只能保存最后一个元素
		preHead = preHead.next
	}
	// 上面两个链表任意一个走到终点都会退出 这里再去补齐另一个未遍历完节点的元素
	if l1 != nil {
		preHead.next = l1
	}
	if l2 != nil {
		preHead.next = l2
	}
	return result.next
}

// 环形链表
func hasCycle(head *ListNode) bool {
	if head == nil{
		return false
	}
	// 快指针 每次走两步
	fast := head.next
	for head != nil && fast != nil && fast.next != nil {
		if head.value == fast.value {
			return true
		}
		// 慢指针，每次走一步
		head = head.next
		fast = fast.next.next
	}
	return false
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{0, nil}
	// 用这个哨兵节点指向我们的新链表 如果只用一个result 不用list 最终result一直next时 它代表的只是最后一个节点
	result := list
	tmp := 0
	// tmp != 0 在两个链表都遍历完之后 还需要进位，比如 555 + 555 = 1110 最终千位进一
	for l1 != nil || l2 != nil || tmp != 0{
		if l1 != nil {
			tmp += l1.value
			l1 = l1.next
		}
		if l2 != nil {
			tmp += l2.value
			l2 = l2.next
		}
		list.next = &ListNode{tmp % 10, nil}
		tmp = tmp / 10
		list = list.next
	}
	return result.next
}



func main11()  {
	//obj5 := &ListNode{5,nil}
	//obj4 := &ListNode{4,obj5}
	//obj3 := &ListNode{3,obj4}
	//obj2 := &ListNode{2,obj3}
	//obj1 := &ListNode{1,obj2}
	//obj := &ListNode{0,obj1}
	//obj.testResult()
	//result := removeNthFromEnd(obj, 2)
	//result.testResult()

	//obj5 := &ListNode{4,nil}
	//obj4 := &ListNode{3,obj5}
	//obj3 := &ListNode{1,obj4}
	//
	//obj2 := &ListNode{4,nil}
	//obj1 := &ListNode{2,obj2}
	//obj := &ListNode{1,obj1}
	//
	//obj.testResult()
	//obj3.testResult()
	//result := mergeTwoLists(obj, obj3)
	//result.testResult()

	obj5 := &ListNode{3,nil}
	obj4 := &ListNode{4,obj5}
	obj3 := &ListNode{2,obj4}

	obj2 := &ListNode{4,nil}
	obj1 := &ListNode{6,obj2}
	obj := &ListNode{5,obj1}

	obj.testResult()
	obj3.testResult()
	result := addTwoNumbers(obj, obj3)
	result.testResult()


}
