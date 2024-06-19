package main

// 907. Sum of Subarray Minimums
// Given an array of integers arr, find the sum of min(b), where b ranges over every (contiguous) subarray of arr. 
// Since the answer may be large, return the answer modulo 10^9 + 7.

// Example 1:
// Input: arr = [3,1,2,4]
// Output: 17
// Explanation: 
// Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4]. 
// Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.
// Sum is 17.

// Example 2:
// Input: arr = [11,81,94,43,3]
// Output: 444

// Constraints:
//     1 <= arr.length <= 3 * 10^4
//     1 <= arr[i] <= 3 * 10^4

import "fmt"

// stack
func sumSubarrayMins(arr []int) int {
    stack := [][2]int{[2]int{arr[0], 1}}
    res, sum, mod := arr[0], arr[0], 1_000_000_007
    for i := 1; i < len(arr); i++ {
        v0 := arr[i]
        v1 := 1
        for len(stack) > 0 && v0 <= stack[len(stack) - 1][0] {
            v1 += stack[len(stack) - 1][1]
            sum -= stack[len(stack) - 1][0] * stack[len(stack) - 1][1]
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, [2]int{v0, v1})
        sum = (sum + v0 * v1) % mod
        res = (res + sum) % mod
    }
    return res
}

func sumSubarrayMins1(arr []int) int {
    arr = append(arr, -1)
    res, stack := 0, []int{ -1 }
    for r, x := range arr {
        for len(stack) > 1 && arr[stack[len(stack)-1]] >= x {
            i := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res += arr[i] * (r-i) * (i - stack[len(stack)-1])
        }
        stack = append(stack, r)
    }
    return res % 1_000_000_007
}

func main() {
    // Example 1:
    // Input: arr = [3,1,2,4]
    // Output: 17
    // Explanation: 
    // Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4]. 
    // Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.
    // Sum is 17.
    fmt.Println(sumSubarrayMins([]int{3,1,2,4})) // 17
    // Example 2:
    // Input: arr = [11,81,94,43,3]
    // Output: 444
    fmt.Println(sumSubarrayMins([]int{11,81,94,43,3})) // 444

    fmt.Println(sumSubarrayMins1([]int{3,1,2,4})) // 17
    fmt.Println(sumSubarrayMins1([]int{11,81,94,43,3})) // 444
}