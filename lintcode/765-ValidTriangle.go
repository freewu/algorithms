package main

import "fmt"

/**
765 · Valid Triangle
Description
Given three integers a, b, c, return true if they can form a triangle.

Example 1:

	Input : a = 2, b = 3, c = 4
	Output : true

Example 2:

	Input : a = 1, b = 2, c = 3
	Output : false

解题思路

	任意两条边之和大于第三条边，构成三角形
	三角形两条边相等，等腰三角形
	三角形三边相等，等边三角形
 */

/**
 * @param a: a integer represent the length of one edge
 * @param b: a integer represent the length of one edge
 * @param c: a integer represent the length of one edge
 * @return: whether three edges can form a triangle
 */
func IsValidTriangle(a int, b int, c int) bool {
	// write your code here
	// 任意两条边之和大于第三条边，构成三角形
	// 三角形两条边相等，等腰三角形
	// 三角形三边相等，等边三角形
	return a + b > c && a + c > b && b + c > a
}

func main() {
	fmt.Printf("%v\n",IsValidTriangle(2,3,4)) // true
	fmt.Printf("%v\n",IsValidTriangle(1,2,3)) // false
}