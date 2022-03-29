package main

import "fmt"

/**
84. Largest Rectangle in Histogram
Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.

Constraints:

	1 <= heights.length <= 10^5
	0 <= heights[i] <= 10^4

Example 1:

	Input: heights = [2,1,5,6,2,3]
	Output: 10
	Explanation: The above is a histogram where width of each bar is 1.
	The largest rectangle is shown in the red area, which has an area = 10 units.

Example 2:

	Input: heights = [2,4]
	Output: 4

 */

func largestRectangleArea(heights []int) int {
	maxArea := 0
	n := len(heights) + 2
	// Add a sentry at the beginning and the end
	getHeight := func(i int) int {
		if i == 0 || n-1 == i {
			return 0
		}
		return heights[i-1]
	}
	st := make([]int, 0, n/2)
	for i := 0; i < n; i++ {
		for len(st) > 0 && getHeight(st[len(st)-1]) > getHeight(i) {
			// pop stack
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			maxArea = max(maxArea, getHeight(idx)*(i-st[len(st)-1]-1))
		}
		// push stack
		st = append(st, i)
	}
	return maxArea
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// best solution
func largestRectangleAreaBest(heights []int) int {
	l := len(heights)
	if l == 0 {
		return 0
	}
	dpLeft := make([]int, l)
	dpRight := make([]int, l)
	for i := 1 ; i < l; i++ {
		dpLeft[i] = i
		j := i
		for dpLeft[j] - 1 >= 0 && heights[dpLeft[j] - 1] >= heights[i] {
			j = dpLeft[j] - 1
		}
		dpLeft[i] = dpLeft[j]
	}
	dpRight[l - 1] = l - 1
	for i := l - 2; i >= 0; i-- {
		dpRight[i] = i
		j := i
		for dpRight[j] + 1 <= l - 1 && heights[dpRight[j] + 1] >= heights[i] {
			j = dpRight[j] + 1
		}
		dpRight[i] = dpRight[j]
	}
	result := 0
	for i := 0; i < l; i++  {
		if heights[i] * (dpRight[i] - dpLeft[i] + 1) > result {
			result = heights[i] * (dpRight[i] - dpLeft[i] + 1)
		}
	}
	return result
}

func main() {
	fmt.Printf("largestRectangleArea([]int{2,1,5,6,2,3}) = %v\n",largestRectangleArea([]int{2,1,5,6,2,3})) // 10
	fmt.Printf("largestRectangleArea([]int{2,4}) = %v\n",largestRectangleArea([]int{2,4})) // 4
	fmt.Printf("largestRectangleArea([]int{2,3}) = %v\n",largestRectangleArea([]int{2,3})) // 4

	fmt.Printf("largestRectangleAreaBest([]int{2,1,5,6,2,3}) = %v\n",largestRectangleAreaBest([]int{2,1,5,6,2,3})) // 10
	fmt.Printf("largestRectangleAreaBest([]int{2,4}) = %v\n",largestRectangleAreaBest([]int{2,4})) // 4
	fmt.Printf("largestRectangleAreaBest([]int{2,3}) = %v\n",largestRectangleAreaBest([]int{2,3})) // 4
}
