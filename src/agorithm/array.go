package main

import (
	"sort"
	"strings"
)
/**
1、给定两个数组，编写一个函数来计算它们的交集。
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

/**
最长公共前缀
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _,k := range strs{
		// 循环进行匹配
		for strings.Index(k, prefix) != 0{
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[0:len(prefix) - 1]
		}
	}
	return prefix
}

// 买卖股票的最佳时机
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit := 0
	for i := 1; i < len(prices); i++{
		tmp := prices[i] - prices[i - 1]
		if tmp > 0{
			profit += tmp
		}
	}
	return profit
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

// 旋转数组:给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
func rotate(nums []int, k int)  {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(arr []int)  {
	for i := 0; i < len(arr)/2; i++ {
		// 第一轮为第一个与最后一个换，依此类推
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// 给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
func removeElement(nums []int, val int) ([]int,int) {
	for i := 0; i < len(nums); {
		if nums[i] == val{
			nums = append(nums[:i], nums[i+1:]...)
		}else {
			i++
		}
	}
	return nums,len(nums)
}

/**
删除排序数组中的重复项，
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次。
 */
func removeDuplicates(nums []int) ([]int, int)  {
	for i := 0; i < len(nums) - 1; {
		if nums[i] == nums[i+1] {
			// nums[:i] 前闭后开 包含左边index的值
			nums = append(nums[:i], nums[i+1:]...)
		}else {
			i++
		}
	}
	return nums, len(nums)
}

func main()  {
	// 两个数组的交集
	//num1 := []int{1,2,3,4,4,13}
	//num2 := []int{1,3,2,9,10}
	//
	//result := intersect(num1, num2)
	//fmt.Println(result)
	//
	//result1 := intersect1(num1, num2)
	//fmt.Println(result1)


	// 最长公共前缀
	//strs := []string{"flower","flow","flight"}
	//prefix := longestCommonPrefix(strs)
	//fmt.Println(prefix)

	// 买卖股票的最佳时机
	//price := []int{7, 1, 5, 3, 6, 4}
	//fmt.Println(maxProfit(price))

	// 数组中的元素向右移动 k 个位置
	//nums := []int{7, 1, 5, 3, 6, 4}
	//rotate(nums, 2)
	//fmt.Println(nums)

	// 给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
	//nums := []int{3,2,2,3}
	//fmt.Println("修改前数组长度", len(nums))
	//fmt.Println("修改前数组值", nums)
	//
	//newNums,val := removeElement(nums, 3)
	//
	//fmt.Println("修改后数组长度", val)
	//fmt.Println("修改后数组值", newNums)

	// 删除排序数组中的重复项
	//nums := []int{0,0,1,1,1,2,2,3,3,4}
	//
	//fmt.Println("修改前数组长度", len(nums))
	//fmt.Println("修改前数组值", nums)
	//
	//newNums,val := removeDuplicates(nums)
	//
	//fmt.Println("修改后数组长度", val)
	//fmt.Println("修改后数组值", newNums)


}