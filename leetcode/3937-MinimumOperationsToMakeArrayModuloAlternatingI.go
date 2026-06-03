package main 

// 3937. Minimum Operations to Make Array Modulo Alternating I
// You are given an integer array nums and an integer k.

// In one operation, you can increase or decrease any element of nums by 1.

// An array is called modulo alternating if there exist two distinct integers x and y (0 <= x, y < k) such that:
//     For every even index i, nums[i] % k == x
//     For every odd index i, nums[i] % k == y

// Return the minimum number of operations required to make nums modulo alternating.

// Example 1:
// Input: nums = [1,4,2,8], k = 3
// Output: 2
// Explanation:
// Let's choose x = 1 for even indices and y = 2 for odd indices.
// Perform the following operations:
// Increment nums[1] = 4 by 1, giving nums = [1, 5, 2, 8].
// Decrement nums[2] = 2 by 1, giving nums = [1, 5, 1, 8].
// Now, for even indices, nums[i] % k = 1, and for odd indices, nums[i] % k = 2.
// Thus, the total number of operations required is 2.

// Example 2:
// Input: nums = [1,1,1], k = 3
// Output: 1
// Explanation:
// Incrementing nums[1] by 1 gives nums = [1, 2, 1], which satisfies the condition with x = 1 and y = 2.
// Thus, the total number of operations required is 1.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 109
//     2 <= k <= 100

import "fmt"
import "slices"
import "sort"

func minOperations(nums []int, k int) int {
    n := len(nums)
    count := func(n, m, k int) int {
        m %= k
        diff := (m - n + k) % k
        return min(diff, k-diff)
    }
    // 初始化两个长度为k的切片，初始值都是0
    se1, se2 := make([]int, k),  make([]int, k)
    for i := 0; i < n; i += 2 { // 遍历偶数下标（0,2,4...）
        for j := 0; j < k; j++ {
            se1[j] += count(j, nums[i], k)
        }
    }
    for i := 1; i < n; i += 2 { // 遍历奇数下标（1,3,5...）
        for j := 0; j < k; j++ {
            se2[j] += count(j, nums[i], k)
        }
    }
    res := 1 << 61
    for j := 0; j < k; j++ { // 双重循环找最小值
        for i := 0; i < k; i++ {
            if i == j { continue }
            res = min(res, se1[j] + se2[i])
        }
    }
    return res
}

func minOperations1(nums []int, k int) int {
    if len(nums) == 1 {
        return 0
    }
    arr := [2][]int{}
    for i, v := range nums {
        arr[i % 2] = append(arr[i % 2], v % k)
    }
    calc := func(a []int, k int) (int, int, int) {
        n := len(a)
        slices.Sort(a)
        for _, x := range a {
            a = append(a, x + k)
        }
        sum := make([]int, n*2+1)
        for i, x := range a {
            sum[i+1] = sum[i] + x
        }
        // 都变成 target 的最小操作次数
        calcOp := func(target int) int {
            i := sort.SearchInts(a[:n], target)
            j := i + sort.SearchInts(a[i:i+n], target+k/2+1)
            return (sum[j] - sum[i]) - (j-i)*target + // [i, j) 中的数都减小到 target
                (n-j+i)*(target+k) - (sum[i+n] - sum[j]) // [j, i+n) 中的数都增大到 target+k
        }
        mn, mn2, bestX := 1 << 61, 1 << 61, 0
        for i, x := range a[:n] {
            if i > 0 && a[i] == a[i-1] { // 优化：相同的值无需重复计算
                continue
            }
            op := calcOp(x)
            // 维护最小次小操作次数
            if op < mn {
                mn2 = mn
                mn, bestX = op, x
            } else if op < mn2 {
                mn2 = op
            }
        }
        // 还可以都变成 bestX - 1 或者 bestX + 1
        mn2 = min(mn2, calcOp((bestX - 1 + k) % k), calcOp((bestX + 1) % k))
        return mn, mn2, bestX
    }
    mn1x, mn2x, bestX := calc(arr[0], k)
    mn1y, mn2y, bestY := calc(arr[1], k)
    if bestX != bestY {
        return mn1x + mn1y
    }
    return min(mn1x + mn2y, mn2x + mn1y)
}

func main() {
    // Example 1:
    // Input: nums = [1,4,2,8], k = 3
    // Output: 2
    // Explanation:
    // Let's choose x = 1 for even indices and y = 2 for odd indices.
    // Perform the following operations:
    // Increment nums[1] = 4 by 1, giving nums = [1, 5, 2, 8].
    // Decrement nums[2] = 2 by 1, giving nums = [1, 5, 1, 8].
    // Now, for even indices, nums[i] % k = 1, and for odd indices, nums[i] % k = 2.
    // Thus, the total number of operations required is 2.
    fmt.Println(minOperations([]int{1,4,2,8}, 3)) // 2
    // Example 2:
    // Input: nums = [1,1,1], k = 3
    // Output: 1
    // Explanation:
    // Incrementing nums[1] by 1 gives nums = [1, 2, 1], which satisfies the condition with x = 1 and y = 2.
    // Thus, the total number of operations required is 1.
    fmt.Println(minOperations([]int{1,1,1}, 3)) // 1

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9}, 3)) // 5
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1}, 3)) // 5

    fmt.Println(minOperations1([]int{1,4,2,8}, 3)) // 2
    fmt.Println(minOperations1([]int{1,1,1}, 3)) // 1
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9}, 3)) // 5
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1}, 3)) // 5
}
