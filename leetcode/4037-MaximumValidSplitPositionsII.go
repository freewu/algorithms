package main

// 4037. Maximum Valid Split Positions II
// You are given an integer array nums.

// You may remove at most one element from nums. 
// Let arr be the array of remaining elements in their original order, and let m be its length.

// A split position i of arr is valid if:
//     1. 0 <= i < m - 1, and
//     2. gcd(arr[0..i]) == gcd(arr[i + 1..m - 1]).

// An array of length 1 has no valid split positions.

// The score of arr is the number of valid split positions in it.

// Return the maximum possible score of arr.

// Here, gcd(a) denotes the greatest common divisor of all elements in the array a.

// Example 1:
// Input: nums = [10,30,15,10]
// Output: 2
// Explanation:
// One optimal solution is to remove nums[2] = 15. Then arr = [10, 30, 10].
// The split positions are:
// Split Position i | gcd(arr[0..i])    | gcd(arr[i + 1..m - 1])
// 0                | 10                | 10
// 1                | 10                | 10
// All split positions are valid. Thus, the answer is 2.

// Example 2:
// Input: nums = [2,10,14]
// Output: 1
// Explanation:
// One optimal solution is to not remove any element. Then arr = [2, 10, 14].
// The split positions are:
// Split Position i | gcd(arr[0..i])   | gcd(arr[i + 1..m - 1])
// 0                | 2                | 2
// 1                | 2                | 14
// Only the split position at index 0 is valid. Thus, the answer is 1.

// Example 3:
// Input: nums = [2,4]
// Output: 0
// Explanation:
// The only remaining array that has a split position is arr = [2, 4].
// The split positions are:
// Split Position i | gcd(arr[0..i])   | gcd(arr[i + 1..m - 1])
// 0                | 2                | 4
// There are no valid split positions. Thus, the answer is 0.

// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9​​​​​​​

import "fmt"

func maxValidSplits(nums []int) int {
    res, n := 0 ,len(nums)
    premain := make([]int, n+1)
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for i := 1; i <= n; i++ {
        premain[i] = gcd(premain[i-1], nums[i-1])
    }
    solve := func(pre, suff []int, skip int, a []int) int {
        n := len(a)
        for i := 1; i <= n; i++ {
            if i-1 == skip {
                pre[i] = pre[i-1]
                continue
            }
            pre[i] = gcd(pre[i-1], a[i-1])
        }
        for i := n - 1; i >= 0; i-- {
            if i == skip {
                suff[i] = suff[i+1]
                continue
            }
            suff[i] = gcd(suff[i+1], a[i])
        }
        curr := 0
        for i := 0; i < n-1; i++ {
            if i == skip {
                continue
            }
            if pre[i+1] == suff[i+1] {
                curr++
            }
        }
        return curr
    }
    for i := 0; i <= n; i++ {
        if i > 0 && premain[i] == premain[i-1] {
            continue
        }
        skip := i - 1
        prefix, suffix := make([]int, n+1), make([]int, n+1)
        v := solve(prefix, suffix, skip, nums)
        if v > res {
            res = v
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [10,30,15,10]
    // Output: 2
    // Explanation:
    // One optimal solution is to remove nums[2] = 15. Then arr = [10, 30, 10].
    // The split positions are:
    // Split Position i | gcd(arr[0..i])    | gcd(arr[i + 1..m - 1])
    // 0                | 10                | 10
    // 1                | 10                | 10
    // All split positions are valid. Thus, the answer is 2.
    fmt.Println(maxValidSplits([]int{10,30,15,10})) // 2
    // Example 2:
    // Input: nums = [2,10,14]
    // Output: 1
    // Explanation:
    // One optimal solution is to not remove any element. Then arr = [2, 10, 14].
    // The split positions are:
    // Split Position i | gcd(arr[0..i])   | gcd(arr[i + 1..m - 1])
    // 0                | 2                | 2
    // 1                | 2                | 14
    // Only the split position at index 0 is valid. Thus, the answer is 1.
    fmt.Println(maxValidSplits([]int{2,10,14})) // 1
    // Example 3:
    // Input: nums = [2,4]
    // Output: 0
    // Explanation:
    // The only remaining array that has a split position is arr = [2, 4].
    // The split positions are:
    // Split Position i | gcd(arr[0..i])   | gcd(arr[i + 1..m - 1])
    // 0                | 2                | 4
    // There are no valid split positions. Thus, the answer is 0.
    fmt.Println(maxValidSplits([]int{2,4})) // 0

    fmt.Println(maxValidSplits([]int{1,2,3,4,5,6,7,8,9})) // 7
    fmt.Println(maxValidSplits([]int{9,8,7,6,5,4,3,2,1})) // 7
}