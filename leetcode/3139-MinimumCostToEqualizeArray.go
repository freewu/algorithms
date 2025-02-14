package main

// 3139. Minimum Cost to Equalize Array
// You are given an integer array nums and two integers cost1 and cost2. 
// You are allowed to perform either of the following operations any number of times:
//     1. Choose an index i from nums and increase nums[i] by 1 for a cost of cost1.
//     2. Choose two different indices i, j, from nums and increase nums[i] and nums[j] by 1 for a cost of cost2.

// Return the minimum cost required to make all elements in the array equal.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [4,1], cost1 = 5, cost2 = 2
// Output: 15
// Explanation:
// The following operations can be performed to make the values equal:
// Increase nums[1] by 1 for a cost of 5. nums becomes [4,2].
// Increase nums[1] by 1 for a cost of 5. nums becomes [4,3].
// Increase nums[1] by 1 for a cost of 5. nums becomes [4,4].
// The total cost is 15.

// Example 2:
// Input: nums = [2,3,3,3,5], cost1 = 2, cost2 = 1
// Output: 6
// Explanation:
// The following operations can be performed to make the values equal:
// Increase nums[0] and nums[1] by 1 for a cost of 1. nums becomes [3,4,3,3,5].
// Increase nums[0] and nums[2] by 1 for a cost of 1. nums becomes [4,4,4,3,5].
// Increase nums[0] and nums[3] by 1 for a cost of 1. nums becomes [5,4,4,4,5].
// Increase nums[1] and nums[2] by 1 for a cost of 1. nums becomes [5,5,5,4,5].
// Increase nums[3] by 1 for a cost of 2. nums becomes [5,5,5,5,5].
// The total cost is 6.

// Example 3:
// Input: nums = [3,5,3], cost1 = 1, cost2 = 3
// Output: 4
// Explanation:
// The following operations can be performed to make the values equal:
// Increase nums[0] by 1 for a cost of 1. nums becomes [4,5,3].
// Increase nums[0] by 1 for a cost of 1. nums becomes [5,5,3].
// Increase nums[2] by 1 for a cost of 1. nums becomes [5,5,4].
// Increase nums[2] by 1 for a cost of 1. nums becomes [5,5,5].
// The total cost is 4.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^6
//     1 <= cost1 <= 10^6
//     1 <= cost2 <= 10^6

import "fmt"

func minCostToEqualizeArray(nums []int, cost1 int, cost2 int) int {
    res, mx, sum, mod := 1 << 61, 0, 0, 1_000_000_007
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    if len(nums) == 2 { return abs(nums[0] - nums[1]) * cost1 % mod }
    for _, v := range nums {
        mx = max(mx, v)
    }
    arr, mx1 := []int{}, 0
    for _, v := range nums {
        t := mx - v
        mx1 = max(mx1, t)
        arr = append(arr, t)
    }
    for _, v := range arr {
        sum += v
    }
    if cost1 * 2 <= cost2 || len(arr) == 1 { return sum * cost1 % mod }
    rest, val := sum - mx1, 0
    if mx1 >= rest {
        val = (mx1 - rest) * cost1 + rest * cost2
    } else {
        sum = rest + mx1
        val = sum / 2 * cost2 + (sum % 2) * cost1
    }
    for res > val {
        res = val
        mx1++
        rest += len(arr) - 1
        if mx1 >= rest {
            val = (mx1 - rest) * cost1 + rest * cost2
        } else {
            sum = rest + mx1
            val = sum / 2 * cost2 + (sum % 2) * cost1
        }
    }
    return res % mod
}

func minCostToEqualizeArray1(nums []int, cost1 int, cost2 int) int {
    mx, mv, count, n, mod := 0, 0, 0, len(nums), 1_000_000_007
    for _, v := range nums {
        if v > mx {
            mx = v
        }
    }
    for _, v := range nums {
        d := mx - v
        if count += d; d > mv {
            mv = d 
        }
    }
    min := func(x, y int) int { if x < y { return x; }; return y; }
    find := func(mv, count, c1, c2 int) int {
        switch {
        case c1 << 1 <= c2:
            return count * c1 
        case mv << 1 <= count: 
            return (count >> 1) * c2 + (count & 1) * c1 
        }
        return (count - mv) * c2 + (mv << 1 - count) * c1 
    }
    res := find(mv, count, cost1, cost2)
    if n <= 2 { return res % mod }
    for (mv << 1) > count { 
        mv, count = mv + 1, count + n
        res = min(res, find(mv, count, cost1, cost2))
    }
    mv, count = mv + 1, count + n
    res = min(res, find(mv, count, cost1, cost2)) 
    return res % mod
}

func main() {
    // Example 1:
    // Input: nums = [4,1], cost1 = 5, cost2 = 2
    // Output: 15
    // Explanation:
    // The following operations can be performed to make the values equal:
    // Increase nums[1] by 1 for a cost of 5. nums becomes [4,2].
    // Increase nums[1] by 1 for a cost of 5. nums becomes [4,3].
    // Increase nums[1] by 1 for a cost of 5. nums becomes [4,4].
    // The total cost is 15.
    fmt.Println(minCostToEqualizeArray([]int{4,1}, 5, 2)) // 15
    // Example 2:
    // Input: nums = [2,3,3,3,5], cost1 = 2, cost2 = 1
    // Output: 6
    // Explanation:
    // The following operations can be performed to make the values equal:
    // Increase nums[0] and nums[1] by 1 for a cost of 1. nums becomes [3,4,3,3,5].
    // Increase nums[0] and nums[2] by 1 for a cost of 1. nums becomes [4,4,4,3,5].
    // Increase nums[0] and nums[3] by 1 for a cost of 1. nums becomes [5,4,4,4,5].
    // Increase nums[1] and nums[2] by 1 for a cost of 1. nums becomes [5,5,5,4,5].
    // Increase nums[3] by 1 for a cost of 2. nums becomes [5,5,5,5,5].
    // The total cost is 6.
    fmt.Println(minCostToEqualizeArray([]int{2,3,3,3,5}, 2, 1)) // 6
    // Example 3:
    // Input: nums = [3,5,3], cost1 = 1, cost2 = 3
    // Output: 4
    // Explanation:
    // The following operations can be performed to make the values equal:
    // Increase nums[0] by 1 for a cost of 1. nums becomes [4,5,3].
    // Increase nums[0] by 1 for a cost of 1. nums becomes [5,5,3].
    // Increase nums[2] by 1 for a cost of 1. nums becomes [5,5,4].
    // Increase nums[2] by 1 for a cost of 1. nums becomes [5,5,5].
    // The total cost is 4.
    fmt.Println(minCostToEqualizeArray([]int{3,5,3}, 1, 3)) // 4

    fmt.Println(minCostToEqualizeArray([]int{1,1000000,999999}, 1000000, 1)) // 1999997
    fmt.Println(minCostToEqualizeArray([]int{1,2,3,4,5,6,7,8,9}, 1, 3)) // 36
    fmt.Println(minCostToEqualizeArray([]int{9,8,7,6,5,4,3,2,1}, 1, 3)) // 36

    fmt.Println(minCostToEqualizeArray1([]int{4,1}, 5, 2)) // 15
    fmt.Println(minCostToEqualizeArray1([]int{2,3,3,3,5}, 2, 1)) // 6
    fmt.Println(minCostToEqualizeArray1([]int{3,5,3}, 1, 3)) // 4
    fmt.Println(minCostToEqualizeArray1([]int{1,1000000,999999}, 1000000, 1)) // 1999997
    fmt.Println(minCostToEqualizeArray1([]int{1,2,3,4,5,6,7,8,9}, 1, 3)) // 36
    fmt.Println(minCostToEqualizeArray1([]int{9,8,7,6,5,4,3,2,1}, 1, 3)) // 36
}