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

func main()  {
	obj5 := &ListNode{5,nil}
	obj4 := &ListNode{4,obj5}
	obj3 := &ListNode{3,obj4}
	obj2 := &ListNode{2,obj3}
	obj1 := &ListNode{1,obj2}
	obj := &ListNode{0,obj1}
	obj.testResult()
	result := removeNthFromEnd(obj, 2)
	result.testResult()
}
