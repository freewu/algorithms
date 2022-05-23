package main

/**
82 · Single Number
Description
Given 2 * n + 1 numbers, every numbers occurs twice except one, find it.

n ≤ 100

Example 1:
	Input: [1,1,2,2,3,4,4]
	Output: 3
	Explanation:
	Only 3 appears once

Example 2:
Input: [0,0,1]
Output: 1
Explanation:
	Only 1 appears oncears once

Challenge
	One-pass, constant extra space.
*/

import "fmt"

/**
 * @param a: An integer array
 * @return: An integer
 */
func SingleNumber(a []int) int {
    // write your code here
	m := make(map[int]bool,len(a) / 2 + 1)
	for i := 0; i < len(a); i++  {
		if _, ok := m[a[i]]; ok { // 如果存在过一次设置为 true
			m[a[i]] = true
		} else {
			m[a[i]] = false
		}
	}
	// 循环出 值为 false 的
	for k, v := range m {
		if v == false {
			return k
		}
	}
	return 0
}

// 使用异或运算的特性来处理
func SingleNumber1(a []int) int {
    // write your code here
	m := a[0]
	for i := 1; i < len(a); i++  {
		m = m ^ a[i]
	}
	return m
}

func main() {
	fmt.Printf("SingleNumber([]int{ 1,1,2,2,3,4,4 }) = %v\n",SingleNumber([]int{ 1,1,2,2,3,4,4 })) // 3
	fmt.Printf("SingleNumber([]int{ 0,0,1 }) = %v\n",SingleNumber([]int{ 0,0,1 })) // 1

	fmt.Printf("SingleNumber1([]int{ 1,1,2,2,3,4,4 }) = %v\n",SingleNumber1([]int{ 1,1,2,2,3,4,4 })) // 3
	fmt.Printf("SingleNumber1([]int{ 0,0,1 }) = %v\n",SingleNumber1([]int{ 0,0,1 })) // 1
}

