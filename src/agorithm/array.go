package main

import (
	"fmt"
	"sort"
)
/**
给定两个数组，编写一个函数来计算它们的交集。
 */
func intersect(nums1 []int, nums2 []int) []int  {
	m0 := map[int]int{}

	for _,v := range nums1{
		// 遍历nums1 初始化map
		m0[v] += 1
	}

	k := 0

	for _,v := range nums2{
		if m0[v] > 0 {
			m0[v] -= 1
			nums2[k] = v
			k++
		}
	}
	return nums2[0:k]
}

/**
如果给定的数组已经排好序呢？你将如何优化你的算法？

解题步骤如下：
<1>设定两个为0的指针，比较两个指针的元素是否相等。 如果指针的元素相等，我们将两个指针一起向后移动，并且将相等的元素放入空白数组。
 */
func intersect1(nums1 []int, nums2 []int) []int  {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i,j,k := 0,0,0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		}else if nums1[i] < nums2[j] {
			i++
		}else if nums1[i] == nums2[j]{
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[0:k]
}


func main()  {
	num1 := []int{1,2,3,4,4,13}
	num2 := []int{1,3,2,9,10}

	result := intersect(num1, num2)
	fmt.Println(result)

	result1 := intersect1(num1, num2)
	fmt.Println(result1)
}