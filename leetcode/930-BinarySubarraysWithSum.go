package main

// 930. Binary Subarrays With Sum
// Given a binary array nums and an integer goal, return the number of non-empty subarrays with a sum goal.
// A subarray is a contiguous part of the array.

// Example 1:
// Input: nums = [1,0,1,0,1], goal = 2
// Output: 4
// Explanation: The 4 subarrays are bolded and underlined below:
// [(1),(0),(1), 0,  1]
// [(1),(0),(1),(0), 1]
// [ 1, (0),(1),(0),(1)]
// [ 1,  0, (1),(0),(1)]

// Example 2:
// Input: nums = [0,0,0,0,0], goal = 0
// Output: 15
 
// Constraints:
//     1 <= nums.length <= 3 * 10^4
//     nums[i] is either 0 or 1.
//     0 <= goal <= nums.length

import "fmt"

// 滑动窗口
func numSubarraysWithSum(nums []int, goal int) int {
	freq, sum, res := make([]int, len(nums) + 1), 0, 0
	freq[0] = 1
	for _, v := range nums {
        // 不断的加入右边的值，直到总和等于 goal
		t := sum + v - goal
		if t >= 0 {
            // 如果遇到比 goal 多的情况，多出来的值就在 freq 中查表，看多出来的值可能是由几种情况构成的。
            // 一旦和与 goal 相等以后，之后比 goal 多出来的情况会越来越多(因为在不断累积，总和只会越来越大)，不断的查 freq 表就可以了。
			// 总和有多余的，需要减去 t，除去的方法有 freq[t] 种
			res += freq[t]
		}
		sum += v
        // 在 freq 中不断的记下能使得和为 sum 的组合方法数，例如 freq[1] = 2 ，代表和为 1 有两种组合方法，(可能是 1 和 1，0 或者 0，1
		freq[sum]++
		//fmt.Printf("v = %v freq = %v sum = %v res = %v t = %v\n", v, freq, sum, res, t)
	}
	return res
}

// best solution
func numSubarraysWithSum1(nums []int, goal int) int {
    res := 0
	n := len(nums)
	for l1, l2, r, s1, s2 := 0, 0, 0, 0, 0; r < n; r++ {
		s1 += nums[r]
		s2 += nums[r]
		for l1 <= r && s1 > goal {
			s1 -= nums[l1]
			l1++
            //fmt.Printf("s1 = %v s2 = %v l1 = %v l2 = %v r = %v\n", s1, s2, l1, l2,r)
		}
		for l2 <= r && s2 >= goal {
			s2 -= nums[l2]
			l2++
            //fmt.Printf("s1 = %v s2 = %v l1 = %v l2 = %v r = %v\n", s1, s2, l1, l2,r)
		}
		res += l2 - l1
	}
	return res
}

func main() {
    fmt.Println(numSubarraysWithSum([]int{1,0,1,0,1},2)) // 4
    fmt.Println(numSubarraysWithSum([]int{0,0,0,0,0},0)) // 15

    fmt.Println(numSubarraysWithSum1([]int{1,0,1,0,1},2)) // 4
    fmt.Println(numSubarraysWithSum1([]int{0,0,0,0,0},0)) // 15
}