package LeetCode

import "math"

/**
 * 155. Min Stack
 * 描述：
 * 难度：Medium
 * 类型：Stack
 */

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (ms *MinStack) Push(x int) {
	ms.stack = append(ms.stack, x)
	top := ms.minStack[len(ms.minStack)-1]
	ms.minStack = append(ms.minStack, min(x, top))
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.minStack = ms.minStack[:len(ms.minStack)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
