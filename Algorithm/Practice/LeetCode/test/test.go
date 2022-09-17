package main

import "fmt"

func main() {
	result := partition("aab")
	fmt.Println("result=", result)
}

func partition(s string) [][]string {
	result, path := make([][]string, 0), make([]string, 0)
	dfs(s, path, &result, 0)
	return result
}

// prev表示前一个隔板，start表示当前隔板
func dfs(s string, path []string, result *[][]string, start int) {
	if start == len(s) {
		// 终止条件
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
		return
	}
	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			path = append(path, s[start:i+1])
			dfs(s, path, result, i+1)
			path = path[:len(path)-1]
		}
	}
}

func isPalindrome(s string, start, end int) bool {
	for start < end && s[start] == s[end] {
		start++
		end--
	}
	return start >= end
}
