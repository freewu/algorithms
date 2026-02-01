package main

// 3826. Minimum Partition Score
// You are given an integer array nums and an integer k.

// Your task is to partition nums into exactly k subarrays and return an integer denoting the minimum possible score among all valid partitions.

// The score of a partition is the sum of the values of all its subarrays.

// The value of a subarray is defined as sumArr * (sumArr + 1) / 2, where sumArr is the sum of its elements.

// Example 1:
// Input: nums = [5,1,2,1], k = 2
// Output: 25
// Explanation:
// We must partition the array into k = 2 subarrays. One optimal partition is [5] and [1, 2, 1].
// The first subarray has sumArr = 5 and value = 5 × 6 / 2 = 15.
// The second subarray has sumArr = 1 + 2 + 1 = 4 and value = 4 × 5 / 2 = 10.
// The score of this partition is 15 + 10 = 25, which is the minimum possible score.

// Example 2:
// Input: nums = [1,2,3,4], k = 1
// Output: 55
// Explanation:
// Since we must partition the array into k = 1 subarray, all elements belong to the same subarray: [1, 2, 3, 4].
// This subarray has sumArr = 1 + 2 + 3 + 4 = 10 and value = 10 × 11 / 2 = 55.​​​​​​​
// The score of this partition is 55, which is the minimum possible score.

// Example 3:
// Input: nums = [1,1,1], k = 3
// Output: 3
// Explanation:
// We must partition the array into k = 3 subarrays. The only valid partition is [1], [1], [1].
// Each subarray has sumArr = 1 and value = 1 × 2 / 2 = 1.
// The score of this partition is 1 + 1 + 1 = 3, which is the minimum possible score.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10^4
//     1 <= k <= nums.length 

import "fmt"
import "math/big"
import "math"

// Time Limit Exceeded 821 / 822 testcases passed
func minPartitionScore(nums []int, k int) int64 {
    n, inf := len(nums), int64(1 << 63 / 2)
    prefix := make([]int64, n + 1)
    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + int64(nums[i])    
    }
    dp, next := make([]int64, n + 1), make([]int64, n + 1)
    for i := 1; i <= n; i++ {
        dp[i] = inf
    }
    dp[0] = 0
    for p := 0; p < k; p++ {
        for i := 0; i <= n; i++ {
            next[i] = inf
        }
        for j := p + 1; j <= n; j++ {
            for i := p; i < j; i++ {
                lastSum := prefix[j] - prefix[i]
                lastScore := lastSum * (lastSum + 1) / 2
                score := dp[i] + lastScore
                next[j] = min(next[j], score)
            }
        }
        dp, next = next, dp 
    }
    return dp[n]
}

type Vec struct{ x, y int }

func (a Vec) sub(b Vec) Vec { return Vec{a.x - b.x, a.y - b.y} }
func (a Vec) dot(b Vec) int { return a.x*b.x + a.y*b.y }
func (a Vec) det(b Vec) int { return a.x*b.y - a.y*b.x } // 如果乘法会溢出，用 detCmp
func (a Vec) detCmp(b Vec) int {
    v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
    w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
    return v.Cmp(w)
}

func minPartitionScore1(nums []int, k int) int64 {
    n := len(nums)
    sum := make([]int, n+1)
    for i, x := range nums {
        sum[i+1] = sum[i] + x
    }
    f := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        f[i] = math.MaxInt / 2
    }
    for K := 1; K <= k; K++ {
        s := sum[K-1]
        q := []Vec{{s, f[K-1] + s*s - s}}
        for i := K; i <= n-(k-K); i++ {
            s = sum[i]
            p := Vec{-2 * s, 1}
            for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
                q = q[1:]
            }
            v := Vec{s, f[i] + s*s - s}
            f[i] = p.dot(q[0]) + s*s + s
            // 读者可以把 detCmp 改成 det 感受下这个算法的效率
            // 目前 det 也能过，可以试试 hack 一下
            for len(q) > 1 && q[len(q) - 1].sub(q[len(q) - 2]).detCmp(v.sub(q[len(q) - 1])) <= 0 {
                q = q[:len(q)-1]
            }
            q = append(q, v)
        }
    }
    return int64(f[n] / 2)
}

func main() {
    // Example 1:
    // Input: nums = [5,1,2,1], k = 2
    // Output: 25
    // Explanation:
    // We must partition the array into k = 2 subarrays. One optimal partition is [5] and [1, 2, 1].
    // The first subarray has sumArr = 5 and value = 5 × 6 / 2 = 15.
    // The second subarray has sumArr = 1 + 2 + 1 = 4 and value = 4 × 5 / 2 = 10.
    // The score of this partition is 15 + 10 = 25, which is the minimum possible score.
    fmt.Println(minPartitionScore([]int{5,1,2,1}, 2)) // 25
    // Example 2:
    // Input: nums = [1,2,3,4], k = 1
    // Output: 55
    // Explanation:
    // Since we must partition the array into k = 1 subarray, all elements belong to the same subarray: [1, 2, 3, 4].
    // This subarray has sumArr = 1 + 2 + 3 + 4 = 10 and value = 10 × 11 / 2 = 55.​​​​​​​
    // The score of this partition is 55, which is the minimum possible score.
    fmt.Println(minPartitionScore([]int{1,2,3,4}, 1)) // 55
    // Example 3:
    // Input: nums = [1,1,1], k = 3
    // Output: 3
    // Explanation:
    // We must partition the array into k = 3 subarrays. The only valid partition is [1], [1], [1].
    // Each subarray has sumArr = 1 and value = 1 × 2 / 2 = 1.
    // The score of this partition is 1 + 1 + 1 = 3, which is the minimum possible score.
    fmt.Println(minPartitionScore([]int{1,1,1}, 3)) // 3

    fmt.Println(minPartitionScore([]int{10000,10000,10000,10000,10000,10000,10000,10000}, 3)) // 1100040000
    fmt.Println(minPartitionScore([]int{1,2,3,4,5,6,7,8,9}, 2)) // 531
    fmt.Println(minPartitionScore([]int{9,8,7,6,5,4,3,2,1}, 2)) // 531

    fmt.Println(minPartitionScore1([]int{5,1,2,1}, 2)) // 25
    fmt.Println(minPartitionScore1([]int{1,2,3,4}, 1)) // 55
    fmt.Println(minPartitionScore1([]int{1,1,1}, 3)) // 3
    fmt.Println(minPartitionScore1([]int{10000,10000,10000,10000,10000,10000,10000,10000}, 3)) // 1100040000
    fmt.Println(minPartitionScore1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 531
    fmt.Println(minPartitionScore1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 531
}