package main

// 3892. Minimum Operations to Achieve At Least K Peaks
// You are given a ​​​​​​​circular integer array​​​​​​​ nums of length n.

// An index i is a peak if its value is strictly greater than its neighbors:
//     1. The previous neighbor of i is nums[i - 1] if i > 0, otherwise nums[n - 1].
//     2. The next neighbor of i is nums[i + 1] if i < n - 1, otherwise nums[0].

// You are allowed to perform the following operation any number of times:
//     1. Choose any index i and increase nums[i] by 1.

// Return an integer denoting the minimum number of operations required to make the array contain at least k peaks. 
// If it is impossible, return -1.

// Example 1:
// Input: nums = [2,1,2], k = 1
// Output: 1
// Explanation:
// To achieve at least k = 1 peak, we can increase nums[2] = 2 to 3.
// After this operation, nums[2] = 3 is strictly greater than its neighbors nums[0] = 2 and nums[1] = 1.
// Therefore, the minimum number of operations required is 1.

// Example 2:
// Input: nums = [4,5,3,6], k = 2
// Output: 0
// Explanation:
// The array already contains at least k = 2 peaks with zero operations.
// Index 1: nums[1] = 5 is strictly greater than its neighbors nums[0] = 4 and nums[2] = 3.
// Index 3: nums[3] = 6 is strictly greater than its neighbors nums[2] = 3 and nums[0] = 4.
// Therefore, the minimum number of operations required is 0.

// Example 3:
// Input: nums = [3,7,3], k = 2
// Output: -1
// Explanation:
// It is impossible to have at least k = 2 peaks in this array. Therefore, the answer is -1.

// Constraints:
//     2 <= n == nums.length <= 5000
//     -10^5 <= nums[i] <= 10^5
//     0 <= k <= n​​​​​​​

import "fmt"

func minOperations(nums []int, k int) int {
    n := len(nums)
    if k > n / 2 { return -1 }
    count := 0
    for i, v := range nums {
        if nums[(i - 1 + n) % n] < v && v > nums[(i + 1) % n] {
            count++
        }
    }
    if count >= k { return 0 } // 已经有至少 k 个峰值了，无需操作    
    helper := func (a []int, k int) int {
        n := len(a)
        f := make([][]int, k + 1)
        for i := range f {
            f[i] = make([]int, n)
            if i > 0 {
                f[i][0], f[i][1] = 1 << 61, 1 << 61
            }
        }
        for left := 1; left <= k; left++ {
            for i := 1; i < n-1; i++ {
                // 选或不选
                notChoose := f[left][i]
                choose := f[left-1][i-1] + max(max(a[i-1], a[i+1])-a[i]+1, 0)
                f[left][i+1] = min(notChoose, choose)
            }
        }
        return f[k][n-1]
    }
    res1 := helper(append([]int{nums[n-1]}, nums...), k) // 如果 nums[0] 是峰顶，那么 nums[n-1] 不是峰顶
    res2 := helper(append(nums, nums[0]), k) // 如果 nums[0] 不是峰顶
    return min(res1, res2)
}

func minOperations1(nums []int, k int) int {
    n, inf := len(nums), 1 << 61
    if k == 0    { return 0 }
    if k > n / 2 { return -1 }
    cost := make([]int, n)
    for i := 0; i < n; i++ {
        left := nums[(i-1+n)%n]
        right := nums[(i+1)%n]
        need := max(left, right) + 1
        if need > nums[i] {
            cost[i] = need - nums[i]
        }
    }
    helper := func(cost []int, l, r, k int) int {
        if k == 0 { return 0  }
        length := r - l + 1
        if length <= 0 || k > (length + 1) / 2 {
            return inf
        }
        prev, curr := make([]int, k + 1),  make([]int, k + 1)
        for j := range prev {
            prev[j] = inf
        }
        prev[0] = 0
        prevPrev := make([]int, k + 1)
        for j := range prevPrev {
            prevPrev[j] = inf
        }
        prevPrev[0] = 0
        for i := l; i <= r; i++ {
            for j := range curr {
                curr[j] = inf
            }
            for j := 0; j <= k; j++ {
                if prev[j] < curr[j] {
                    curr[j] = prev[j]
                }
                if j >= 1 && prevPrev[j-1] < inf {
                    val := prevPrev[j-1] + cost[i]
                    if val < curr[j] {
                        curr[j] = val
                    }
                }
            }
            prevPrev, prev, curr = prev, curr, prevPrev
        }
        return prev[k]
    }
    res := helper(cost, 1, n - 1, k)
    res = min(res, helper(cost, 2, n - 2, k - 1) + cost[0])
    if res == inf {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,1,2], k = 1
    // Output: 1
    // Explanation:
    // To achieve at least k = 1 peak, we can increase nums[2] = 2 to 3.
    // After this operation, nums[2] = 3 is strictly greater than its neighbors nums[0] = 2 and nums[1] = 1.
    // Therefore, the minimum number of operations required is 1.
    fmt.Println(minOperations([]int{2,1,2}, 1)) // 1
    // Example 2:
    // Input: nums = [4,5,3,6], k = 2
    // Output: 0
    // Explanation:
    // The array already contains at least k = 2 peaks with zero operations.
    // Index 1: nums[1] = 5 is strictly greater than its neighbors nums[0] = 4 and nums[2] = 3.
    // Index 3: nums[3] = 6 is strictly greater than its neighbors nums[2] = 3 and nums[0] = 4.
    // Therefore, the minimum number of operations required is 0.
    fmt.Println(minOperations([]int{4,5,3,6}, 2)) // 0
    // Example 3:
    // Input: nums = [3,7,3], k = 2
    // Output: -1
    // Explanation:
    // It is impossible to have at least k = 2 peaks in this array. Therefore, the answer is -1.
    fmt.Println(minOperations([]int{3,7,3}, 2)) // -1   

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9}, 2)) // 2 
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1}, 2)) // 2 

    fmt.Println(minOperations1([]int{2,1,2}, 1)) // 1
    fmt.Println(minOperations1([]int{4,5,3,6}, 2)) // 0
    fmt.Println(minOperations1([]int{3,7,3}, 2)) // -1   
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 2 
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 2 
}