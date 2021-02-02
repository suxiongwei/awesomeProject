package main

import "fmt"

// 翻转字符串
func reverseString(s []byte) []byte {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left],s[right] = s[right],s[left]
		left ++
		right --
	}
	return s
}

// 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	var arr [26]int
	// a的ascii码是97
	for i,k := range s {
		arr[k - 'a'] = i
	}
	for i,k := range s {
		if arr[k - 'a'] == i{
			return i
		} else {
			// 标示该元素非目标元素,否则遇到这种peopoe就会return 3
			arr[k - 'a'] = -1
		}
	}
	return -1
}

func main12()  {
	//s := []byte{'h','e','l','l','o'}
	//result := reverseString(s)
	//for i := range result {
	//	fmt.Print(string(s[i]))
	//}

	fmt.Println(firstUniqChar("peopoe"))
}
