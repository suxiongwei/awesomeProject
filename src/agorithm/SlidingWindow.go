package main

import "fmt"

func maxSlidingWindow0(nums []int, k int) []int {
	length := len(nums)
	if length == 0 {
		return []int{0}
	}
	win := make([]int, length -k + 1)
	for i := 0; i < length - k + 1; i++{
		maxValue := 1 >> 31
		for j := i; j < i + k; j++{
			maxValue = max(maxValue, nums[j])
		}
		win[i] = maxValue
	}
	return win
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	//用切片模拟一个双端队列
	queue := []int{}
	result := []int{}
	for i := range nums {
		for i > 0 && (len(queue) > 0) && nums[i] > queue[len(queue)-1] {
			//将比当前元素小的元素祭天
			queue = queue[:len(queue)-1]
		}
		//将当前元素放入queue中
		queue = append(queue, nums[i])
		if i >= k && nums[i-k] == queue[0] {
			//维护队列，保证其头元素为当前窗口最大值 比如1 1 3 1 2 1 在i=5时 nums[i-k]=nums[2]=queue[0]=3,此时3
			queue = queue[1:]
		}
		if i >= k-1 {
			//queue[0] 始终保存的是最大的元素 放入结果数组
			result = append(result, queue[0])
		}
	}
	return result
}

func main111()  {
	//nums := []int{1,3,-1,-3,5,3,6,7}
	nums := []int{1,1,3,1,2,1}
	window := maxSlidingWindow(nums, 3)
	fmt.Println(window)
}