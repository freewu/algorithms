package main

// 2941. Maximum GCD-Sum of a Subarray
// You are given an array of integers nums and an integer k.

// The gcd-sum of an array a is calculated as follows:
//     Let s be the sum of all the elements of a.
//     Let g be the greatest common divisor of all the elements of a.
//     The gcd-sum of a is equal to s * g.

// Return the maximum gcd-sum of a subarray of nums with at least k elements.

// Example 1:
// Input: nums = [2,1,4,4,4,2], k = 2
// Output: 48
// Explanation: We take the subarray [4,4,4], the gcd-sum of this array is 4 * (4 + 4 + 4) = 48.
// It can be shown that we can not select any other subarray with a gcd-sum greater than 48.

// Example 2:
// Input: nums = [7,3,9,4], k = 1
// Output: 81
// Explanation: We take the subarray [9], the gcd-sum of this array is 9 * 9 = 81.
// It can be shown that we can not select any other subarray with a gcd-sum greater than 81.

// Constraints:
//     n == nums.length
//     1 <= n <= 10^5
//     1 <= nums[i] <= 10^6
//     1 <= k <= n

import "fmt"

func maxGcdSum(nums []int, k int) int64 {
    type Pair [2]int
    res, n := 0, len(nums)
    prefix := make([]int, n + 1)
    prefix[0] = 0
    for i, v := range nums {
        prefix[i+1] = prefix[i] + v
    }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    f := []Pair{}
    for i := 0; i < n; i++ {
        g := []Pair{}
        for _, p := range f {
            j, x := p[0], p[1]
            y := int(gcd(x, nums[i]))
            if len(g) == 0 || g[len(g)-1][1] != y {
                g = append(g, Pair{j, y})
            }
        }
        f = g
        f = append(f, Pair{i, nums[i]})
        for _, p := range f {
            j, v := p[0], p[1]
            if i - j + 1 >= k {
                res = max(res, (prefix[i + 1] - prefix[j]) * v)
            }
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [2,1,4,4,4,2], k = 2
    // Output: 48
    // Explanation: We take the subarray [4,4,4], the gcd-sum of this array is 4 * (4 + 4 + 4) = 48.
    // It can be shown that we can not select any other subarray with a gcd-sum greater than 48.
    fmt.Println(maxGcdSum([]int{2,1,4,4,4,2}, 2)) // 48
    // Example 2:
    // Input: nums = [7,3,9,4], k = 1
    // Output: 81
    // Explanation: We take the subarray [9], the gcd-sum of this array is 9 * 9 = 81.
    // It can be shown that we can not select any other subarray with a gcd-sum greater than 81.
    fmt.Println(maxGcdSum([]int{7,3,9,4}, 1)) // 81
}