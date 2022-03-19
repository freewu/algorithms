package main

import "fmt"

/**
42. Trapping Rain Water

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.
The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case,
6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

Constraints:

	n == height.length
	1 <= n <= 2 * 104
	0 <= height[i] <= 105

Example 1:

	Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
	Output: 6

Example 2:

	Input: height = [4,2,0,3,2,5]
	Output: 9

解题思路:
	从 x 轴开始，给出一个数组，数组里面的数字代表从 (0,0) 点开始，宽度为 1 个单位，高度为数组元素的值。
	如果下雨了，问这样一个容器能装多少单位的水
 */

func trap(height []int) int {
	res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
	for left <= right { // 从开头 0 和结尾 len(height) - 1 向中间靠拢
		if height[left] <= height[right] { // 如果 左边(开头) <= 右边(结尾) 开头向里走一步
			if height[left] > maxLeft { // 如果 左边 最大
				maxLeft = height[left]
			} else {
				fmt.Printf("maxLeft = %v\n",maxLeft)
				fmt.Printf("height[%v] = %v\n",left,height[left])
				fmt.Printf("maxLeft - height[left] = %v\n",maxLeft - height[left])
				res += maxLeft - height[left] // 这里是关键
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				fmt.Printf("maxRight = %v\n",maxRight)
				fmt.Printf("height[%v] = %v\n",right,height[right])
				fmt.Printf("maxRight - height[right] = %v\n",maxRight - height[right])
				res += maxRight - height[right] //
			}
			right--
		}
	}
	return res
}

// best solution
func trapBest(heights []int) int {
	leftBoundaries := computeLeftBoundary(heights)
	right := heights[len(heights)-1]
	res := 0
	for i := len(heights)-2; i >= 1; i-- {
		ht := min(leftBoundaries[i], right)
		if ht > heights[i] {
			res += ht - heights[i]
		}
		right = max(right, heights[i])
	}
	return res
}

func computeLeftBoundary(heights []int) []int {
	res := make([]int, len(heights))
	for i := 1; i < len(heights); i++ {
		res[i] = max(res[i-1], heights[i-1])
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Printf("trap([]int{0,1,0,2,1,0,1,3,2,1,2,1} = %v\n",trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Printf("trap([]int{4,2,0,3,2,5} = %v\n",trap([]int{4,2,0,3,2,5}))
	fmt.Printf("trapBest([]int{0,1,0,2,1,0,1,3,2,1,2,1} = %v\n",trapBest([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Printf("trapBest([]int{4,2,0,3,2,5} = %v\n",trapBest([]int{4,2,0,3,2,5}))
}
