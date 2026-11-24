package main

// 3411. Maximum Subarray With Equal Products
// You are given an array of positive integers nums.

// An array arr is called product equivalent if prod(arr) == lcm(arr) * gcd(arr), where:
//     prod(arr) is the product of all elements of arr.
//     gcd(arr) is the GCD of all elements of arr.
//     lcm(arr) is the LCM of all elements of arr.

// Return the length of the longest product equivalent subarray of nums.

// Example 1:
// Input: nums = [1,2,1,2,1,1,1]
// Output: 5
// Explanation: 
// The longest product equivalent subarray is [1, 2, 1, 1, 1], where prod([1, 2, 1, 1, 1]) = 2, gcd([1, 2, 1, 1, 1]) = 1, and lcm([1, 2, 1, 1, 1]) = 2.

// Example 2:
// Input: nums = [2,3,4,5,6]
// Output: 3
// Explanation: 
// The longest product equivalent subarray is [3, 4, 5].

// Example 3:
// Input: nums = [1,2,3,1,4,5,1]
// Output: 5

// Constraints:
//     2 <= nums.length <= 100
//     1 <= nums[i] <= 10

import "fmt"

func maxLength(nums []int) int {
    n := len(nums)
    if n <= 2 { 
        return n 
    }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, l, r := 2, 0, 1
    gcf, lcm := nums[0], nums[0]
    for r < n {
        gcf = gcd(lcm, nums[r])
        for gcf != 1 && l < r {
            for gcd(nums[l], nums[r]) == 1 {
                lcm /= nums[l]
                l++
            }
            lcm /= nums[l]
            gcf /= gcd(nums[l], nums[r])
            l++
        }
        lcm *= nums[r]
        r++
        res = max(res, r - l)
    }
    return res
}

func maxLength1(nums []int) int {
    res, n, multi, l, r := 2, len(nums), 1, 0, 0
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for r < n {
        for gcd(multi, nums[r]) > 1 {
            multi /= nums[l]
            l++
        }
        multi *= nums[r]
        res = max(res, r - l + 1)
        r++
    }
    return res
}

func maxLength2(nums []int) int {
    const MASK = 0x5218341210
    prime := func(x int) int { return MASK >> (uint(x-1) * 4 & 63) & 0xF }
    i, mask, res := 0, 0, 2
    for j, v := range nums {
        m := prime(v)
        for mask&m != 0 {
            mask &^= prime(nums[i])
            i++
        }
        mask |= m
        res = max(res, j-i+1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1,2,1,1,1]
    // Output: 5
    // Explanation: 
    // The longest product equivalent subarray is [1, 2, 1, 1, 1], where prod([1, 2, 1, 1, 1]) = 2, gcd([1, 2, 1, 1, 1]) = 1, and lcm([1, 2, 1, 1, 1]) = 2.
    fmt.Println(maxLength([]int{1,2,1,2,1,1,1})) // 5
    // Example 2:
    // Input: nums = [2,3,4,5,6]
    // Output: 3
    // Explanation: 
    // The longest product equivalent subarray is [3, 4, 5].
    fmt.Println(maxLength([]int{2,3,4,5,6})) // 3
    // Example 3:
    // Input: nums = [1,2,3,1,4,5,1]
    // Output: 5
    fmt.Println(maxLength([]int{1,2,3,1,4,5,1})) // 5

    fmt.Println(maxLength([]int{1,2,3,4,5,6,7,8,9})) // 3
    fmt.Println(maxLength([]int{9,8,7,6,5,4,3,2,1})) // 3

    fmt.Println(maxLength1([]int{1,2,1,2,1,1,1})) // 5
    fmt.Println(maxLength1([]int{2,3,4,5,6})) // 3
    fmt.Println(maxLength1([]int{1,2,3,1,4,5,1})) // 5
    fmt.Println(maxLength1([]int{1,2,3,4,5,6,7,8,9})) // 3
    fmt.Println(maxLength1([]int{9,8,7,6,5,4,3,2,1})) // 3

    fmt.Println(maxLength2([]int{1,2,1,2,1,1,1})) // 5
    fmt.Println(maxLength2([]int{2,3,4,5,6})) // 3
    fmt.Println(maxLength2([]int{1,2,3,1,4,5,1})) // 5
    fmt.Println(maxLength2([]int{1,2,3,4,5,6,7,8,9})) // 3
    fmt.Println(maxLength2([]int{9,8,7,6,5,4,3,2,1})) // 3
}