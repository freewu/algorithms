package main

// 3763. Maximum Total Sum with Threshold Constraints
// You are given two integer arrays nums and threshold, both of length n.

// Starting at step = 1, you perform the following repeatedly:
//     1. Choose an unused index i such that threshold[i] <= step.
//         1.1 If no such index exists, the process ends.
//     2. Add nums[i] to your running total.
//     3. Mark index i as used and increment step by 1.

// Return the maximum total sum you can obtain by choosing indices optimally.

// Example 1:
// Input: nums = [1,10,4,2,1,6], threshold = [5,1,5,5,2,2]
// Output: 17
// Explanation:
// At step = 1, choose i = 1 since threshold[1] <= step. The total sum becomes 10. Mark index 1.
// At step = 2, choose i = 4 since threshold[4] <= step. The total sum becomes 11. Mark index 4.
// At step = 3, choose i = 5 since threshold[5] <= step. The total sum becomes 17. Mark index 5.
// At step = 4, we cannot choose indices 0, 2, or 3 because their thresholds are > 4, so we end the process.

// Example 2:
// Input: nums = [4,1,5,2,3], threshold = [3,3,2,3,3]
// Output: 0
// Explanation:
// At step = 1 there is no index i with threshold[i] <= 1, so the process ends immediately. Thus, the total sum is 0.

// Example 3:
// Input: nums = [2,6,10,13], threshold = [2,1,1,1]
// Output: 31
// Explanation:
// At step = 1, choose i = 3 since threshold[3] <= step. The total sum becomes 13. Mark index 3.
// At step = 2, choose i = 2 since threshold[2] <= step. The total sum becomes 23. Mark index 2.
// At step = 3, choose i = 1 since threshold[1] <= step. The total sum becomes 29. Mark index 1.
// At step = 4, choose i = 0 since threshold[0] <= step. The total sum becomes 31. Mark index 0.
// After step = 4 all indices have been chosen, so the process ends.

// Constraints:
//     n == nums.length == threshold.length
//     1 <= n <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= threshold[i] <= n

import "fmt"
import "sort"

// 超出时间限制 973 / 986 
func maxSum(nums []int, threshold []int) int64 {
    res, n := 0, len(nums)
    type Pair struct {  // 构造结构体切片，存储 threshold 和对应的 -nums （用于从小到大排序后取最大值）
        threshold int
        neg int
    }
    pairs := make([]Pair, n)
    for i := 0; i < n; i++ {
        pairs[i] = Pair{ threshold: threshold[i], neg: -nums[i], }
    }
    // 按threshold从小到大排序
    // 如果threshold相同，按negNum从小到大排序（等价于nums从大到小）
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if pairs[j].threshold < pairs[i].threshold || (pairs[j].threshold == pairs[i].threshold && pairs[j].neg < pairs[i].neg) {
                pairs[i], pairs[j] = pairs[j], pairs[i]
            }
        }
    }
    for i := 0; i < n; i++ {
        if pairs[i].threshold > i + 1 { // 当前是第i+1个元素（1-based），如果阈值不满足则返回当前和
            return int64(res)
        }
        res -= pairs[i].neg  // 累加-nums[i]的相反数（即nums[i]）
    }
    return int64(res)
}

func maxSum1(nums []int, threshold []int) int64 {
    res, n := 0, len(nums)
    type Pair struct {  // 构造结构体切片，存储 threshold 和对应的 -nums （用于从小到大排序后取最大值）
        threshold int
        neg int
    }
    pairs := make([]Pair, n)
    for i := 0; i < n; i++ {
        pairs[i] = Pair{ threshold: threshold[i], neg: -nums[i], }
    }
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].threshold < pairs[j].threshold || (pairs[i].threshold == pairs[j].threshold && pairs[i].neg < pairs[j].neg)
    })
    for i := 0; i < n; i++ {
        if pairs[i].threshold > i + 1 { // 当前是第i+1个元素（1-based），如果阈值不满足则返回当前和
            return int64(res)
        }
        res -= pairs[i].neg  // 累加-nums[i]的相反数（即nums[i]）
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,10,4,2,1,6], threshold = [5,1,5,5,2,2]
    // Output: 17
    // Explanation:
    // At step = 1, choose i = 1 since threshold[1] <= step. The total sum becomes 10. Mark index 1.
    // At step = 2, choose i = 4 since threshold[4] <= step. The total sum becomes 11. Mark index 4.
    // At step = 3, choose i = 5 since threshold[5] <= step. The total sum becomes 17. Mark index 5.
    // At step = 4, we cannot choose indices 0, 2, or 3 because their thresholds are > 4, so we end the process.
    fmt.Println(maxSum([]int{1,10,4,2,1,6}, []int{5,1,5,5,2,2})) // 17
    // Example 2:
    // Input: nums = [4,1,5,2,3], threshold = [3,3,2,3,3]
    // Output: 0
    // Explanation:
    // At step = 1 there is no index i with threshold[i] <= 1, so the process ends immediately. Thus, the total sum is 0.
    fmt.Println(maxSum([]int{4,1,5,2,3}, []int{3,3,2,3,3})) // 0
    // Example 3:
    // Input: nums = [2,6,10,13], threshold = [2,1,1,1]
    // Output: 31
    // Explanation:
    // At step = 1, choose i = 3 since threshold[3] <= step. The total sum becomes 13. Mark index 3.
    // At step = 2, choose i = 2 since threshold[2] <= step. The total sum becomes 23. Mark index 2.
    // At step = 3, choose i = 1 since threshold[1] <= step. The total sum becomes 29. Mark index 1.
    // At step = 4, choose i = 0 since threshold[0] <= step. The total sum becomes 31. Mark index 0.
    // After step = 4 all indices have been chosen, so the process ends.
    fmt.Println(maxSum([]int{2,6,10,13}, []int{2,1,1,1})) // 31

    fmt.Println(maxSum([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 45
    fmt.Println(maxSum([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxSum([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxSum([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(maxSum1([]int{1,10,4,2,1,6}, []int{5,1,5,5,2,2})) // 17
    fmt.Println(maxSum1([]int{4,1,5,2,3}, []int{3,3,2,3,3})) // 0
    fmt.Println(maxSum1([]int{2,6,10,13}, []int{2,1,1,1})) // 31
    fmt.Println(maxSum1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 45
    fmt.Println(maxSum1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxSum1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxSum1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 45
}