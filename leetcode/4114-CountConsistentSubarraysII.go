package main

// 4114. Count Consistent Subarrays II
// You are given an integer array nums.

// A subarray is called consistent if the bitwise AND of all its elements is equal to the minimum value of the subarray.

// Return the number of consistent subarrays in nums.

// A subarray is a contiguous, non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [5,3,1,2]
// Output: 6
// Explanation:
// The following subarrays are consistent:
//  Subarray Values Bitwise AND Minimum Value
//  nums[0..0] [5] 5 5
//  nums[1..1] [3] 3 3
//  nums[2..2] [1] 1 1
//  nums[3..3] [2] 2 2
//  nums[1..2] [3, 1] 3 AND 1 = 1 1
//  nums[0..2] [5, 3, 1] 5 AND 3 AND 1 = 1 1
// Therefore, the number of consistent subarrays is 6.

// Example 2:
// Input: nums = [1,1,1]
// Output: 6
// Explanation:
// For every subarray, both the bitwise AND and the minimum value are equal to 1. 
// Therefore, the answer is 6.

// Example 3:
// Input: nums = [0]
// Output: 1
// Explanation:
// The only subarray is [0], whose bitwise AND and minimum value are both 0. 
// Therefore, the answer is 1.

// Constraints:
//     1 <= nums.length <= 2 * 10^{​​​​​​​5}
//     0 <= nums[i] <= 10^{6}

import "fmt"

func countConsistent(nums []int) int64 {
    res, n := int64(0), len(nums)
    left, stack := make([]int, n), []int{ -1 } // 哨兵
    for i, v := range nums {
        for len(stack) > 1 && nums[stack[len(stack)-1]]&v == v {
            stack = stack[:len(stack)-1]
        }
        left[i] = stack[len(stack)-1] // nums[left[i]] 不是 v 的超集
        stack = append(stack, i)
    }
    stack = []int{n}
    for i := n - 1; i >= 0; i-- {
        x := nums[i]
        for len(stack) > 1 && nums[stack[len(stack)-1]] != x && nums[stack[len(stack)-1]]&x == x {
            stack = stack[:len(stack)-1]
        }
        right := stack[len(stack)-1] // nums[right] 不是 x 的超集
        stack = append(stack, i)
        res += int64(i-left[i]) * int64(right-i)
    }
    return res
}

var last [1_000_001]int32
var vals, starts [30]int
var gen int32

func countConsistent1(nums []int) int64 {
    gen += 1 << 18
    res, n := int64(0), len(nums)
    for i, v := range nums {
        last[v] = int32(i) + gen
        for j := range n {
            vals[j] &= v
        }
        vals[n] = v
        starts[n] = i
        n++
        k := 0
        for j := range n {
            if k > 0 && vals[k-1] == vals[j] {
                continue
            }
            vals[k] = vals[j]
            starts[k] = starts[j]
            k++
        }
        n = k
        starts[k] = i + 1
        for j := range k {
            st, end := starts[j], starts[j+1]-1
            res += max(0, int64(min(end, int(last[vals[j]]-gen))-st+1))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,3,1,2]
    // Output: 6
    // Explanation:
    // The following subarrays are consistent:
    //  Subarray Values Bitwise AND Minimum Value
    //  nums[0..0] [5] 5 5
    //  nums[1..1] [3] 3 3
    //  nums[2..2] [1] 1 1
    //  nums[3..3] [2] 2 2
    //  nums[1..2] [3, 1] 3 AND 1 = 1 1
    //  nums[0..2] [5, 3, 1] 5 AND 3 AND 1 = 1 1
    // Therefore, the number of consistent subarrays is 6.
    fmt.Println(countConsistent([]int{5,3,1,2})) // 6 
    // Example 2:
    // Input: nums = [1,1,1]
    // Output: 6
    // Explanation:
    // For every subarray, both the bitwise AND and the minimum value are equal to 1. 
    // Therefore, the answer is 6.
    fmt.Println(countConsistent([]int{1,1,1})) // 6 
    // Example 3:
    // Input: nums = [0]
    // Output: 1
    // Explanation:
    // The only subarray is [0], whose bitwise AND and minimum value are both 0. 
    // Therefore, the answer is 1.
    fmt.Println(countConsistent([]int{0})) // 1 

    fmt.Println(countConsistent([]int{1,2,3,4,5,6,7,8,9})) // 15
    fmt.Println(countConsistent([]int{9,8,7,6,5,4,3,2,1})) // 15
    fmt.Println(countConsistent([]int{4,3})) // 2

    fmt.Println(countConsistent1([]int{5,3,1,2})) // 6 
    fmt.Println(countConsistent1([]int{1,1,1})) // 6 
    fmt.Println(countConsistent1([]int{0})) // 1 
    fmt.Println(countConsistent1([]int{1,2,3,4,5,6,7,8,9})) // 15
    fmt.Println(countConsistent1([]int{9,8,7,6,5,4,3,2,1})) // 15
    fmt.Println(countConsistent1([]int{4,3})) // 2
}