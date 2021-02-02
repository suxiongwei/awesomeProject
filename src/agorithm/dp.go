package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	// n + 1 是因为没有dp[0] 有dp[n]
	dp := make([]int, n + 1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++{
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
}

/**
最大子序和
 */
func maxSubArray(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	dp := make([]int, length)
	// 设置初始值
	dp[0] = nums[0]
	for i := 1; i < length; i++ {
		if dp[i - 1] < 0 {
			dp[i] = nums[i]
		}else {
			dp[i] = dp[i - 1] + nums[i]
		}
	}

	// -2147483648 取到一个极小的值
	result := -1 << 31
	for _,k := range dp{
		result = max(result, k)
	}
	return result
}

/**
最长上升子序列
 */
func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				// 在比当前index小的值里找较大的数
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

/**
最大子序和 精简版
 */
func maxSubArray1(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	dp := make([]int, length)
	result := nums[0]
	dp[0] = nums[0]
	for i := 1; i < length; i++ {
		dp[i] = max(nums[i], dp[i -1] + nums[i])
		result = max(result, dp[i])
	}
	return result
}

func main1()  {
	//fmt.Println(climbStairs(6))
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	array := maxSubArray1(nums)
	fmt.Println(array)
}
