package main

// 2584. Split the Array to Make Coprime Products
// You are given a 0-indexed integer array nums of length n.

// A split at an index i where 0 <= i <= n - 2 is called valid if the product of the first i + 1 elements 
// and the product of the remaining elements are coprime.
//     For example, if nums = [2, 3, 3], then a split at the index i = 0 is valid because 2 and 9 are coprime, while a split at the index i = 1 is not valid because 6 and 3 are not coprime. 
//     A split at the index i = 2 is not valid because i == n - 1.

// Return the smallest index i at which the array can be split validly or -1 if there is no such split.

// Two values val1 and val2 are coprime if gcd(val1, val2) == 1 where gcd(val1, val2) is the greatest common divisor of val1 and val2.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/12/14/second.PNG" />
// Input: nums = [4,7,8,15,3,5]
// Output: 2
// Explanation: The table above shows the values of the product of the first i + 1 elements, the remaining elements, and their gcd at each index i.
// The only valid split is at index 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/12/14/capture.PNG" />
// Input: nums = [4,7,15,8,3,5]
// Output: -1
// Explanation: The table above shows the values of the product of the first i + 1 elements, the remaining elements, and their gcd at each index i.
// There is no valid split.

// Constraints:
//     n == nums.length
//     1 <= n <= 10^4
//     1 <= nums[i] <= 10^6

import "fmt"
import "math"
// Time Limit Exceeded 74 / 75
func findValidSplit(nums []int) int {
    helper := func(x int) map[int]int {
        res := make(map[int]int)
        for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
            for x % i == 0 {
                res[i]++
                x /= i
            }
        }
        if x != 1 {
            res[x]++
        }
        return res
    }
    n := len(nums)
    count := make([]map[int]int, n, n)
    for i, v := range nums {
        count[i] = helper(v)
    }
    mp := make(map[int]bool)
    for i := 0; i < n - 1; i++ {
        for prime, _ := range count[i] {
            mp[prime] = true
        }
        flag := false
        for j := i + 1; j < n && !flag; j++ {
            for prime, _ := range count[j] {
                if mp[prime] {
                    flag = true
                }
            }
        }
        if !flag {
            return i
        }
    }
    return -1
}

func findValidSplit1(nums []int) int {
    first := map[int]int{}
    n := len(nums)
    last := make([]int, n)
    for i := range last {
        last[i] = i
    }
    for i, x := range nums {
        for j := 2; j <= x / j; j++ {
            if x % j == 0 {
                if k, ok := first[j]; ok {
                    last[k] = i
                } else {
                    first[j] = i
                }
                for x % j == 0 {
                    x /= j
                }
            }
        }
        if x > 1 {
            if k, ok := first[x]; ok {
                last[k] = i
            } else {
                first[x] = i
            }
        }
    }
    mx := last[0]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, x := range last {
        if mx < i {
            return mx
        }
        mx = max(mx, x)
    }
    return -1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/12/14/second.PNG" />
    // Input: nums = [4,7,8,15,3,5]
    // Output: 2
    // Explanation: The table above shows the values of the product of the first i + 1 elements, the remaining elements, and their gcd at each index i.
    // The only valid split is at index 2.
    fmt.Println(findValidSplit([]int{4,7,8,15,3,5})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/12/14/capture.PNG" />
    // Input: nums = [4,7,15,8,3,5]
    // Output: -1
    // Explanation: The table above shows the values of the product of the first i + 1 elements, the remaining elements, and their gcd at each index i.
    // There is no valid split.
    fmt.Println(findValidSplit([]int{4,7,15,8,3,5})) // -1

    fmt.Println(findValidSplit([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(findValidSplit([]int{9,8,7,6,5,4,3,2,1})) // 7

    fmt.Println(findValidSplit1([]int{4,7,8,15,3,5})) // 2
    fmt.Println(findValidSplit1([]int{4,7,15,8,3,5})) // -1
    fmt.Println(findValidSplit1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(findValidSplit1([]int{9,8,7,6,5,4,3,2,1})) // 7
}