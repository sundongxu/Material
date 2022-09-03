package LeetCode

import (
	"path/filepath"
	"strings"
)

/**
 * 71. Simplify Path
 * 描述：
 * 难度：Medium
 * 类型：Stack
 */
func simplifyPath(path string) string {
	return filepath.Clean(path)
}

func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	stack := make([]string, 0)
	var res string
	for i := 0; i < len(arr); i++ {
		//cur := arr[i]
		cur := strings.TrimSpace(arr[i]) // 更加严谨的做法应该还要去掉末尾的空格
		if cur == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // 往上一级目录，去掉栈顶元素即切片末尾元素
			}
		} else if cur != "." && len(cur) > 0 {
			// "."往当前目录不用管，若cur为有效目录名，则将其加入栈
			stack = append(stack, arr[i])
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	res = strings.Join(stack, "/")
	return "/" + res
}
