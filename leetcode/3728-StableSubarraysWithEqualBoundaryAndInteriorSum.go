package main

// 3728. Stable Subarrays With Equal Boundary and Interior Sum
// You are given an integer array capacity.

// A subarray capacity[l..r] is considered stable if:
//     1. Its length is at least 3.
//     2. The first and last elements are each equal to the sum of all elements strictly between them 
//        (i.e., capacity[l] = capacity[r] = capacity[l + 1] + capacity[l + 2] + ... + capacity[r - 1]).

// Return an integer denoting the number of stable subarrays.

// Example 1:
// Input: capacity = [9,3,3,3,9]
// Output: 2
// Explanation:
// [9,3,3,3,9] is stable because the first and last elements are both 9, and the sum of the elements strictly between them is 3 + 3 + 3 = 9.
// [3,3,3] is stable because the first and last elements are both 3, and the sum of the elements strictly between them is 3.

// Example 2:
// Input: capacity = [1,2,3,4,5]
// Output: 0
// Explanation:
// No subarray of length at least 3 has equal first and last elements, so the answer is 0.

// Example 3:
// Input: capacity = [-4,4,0,0,-8,-4]
// Output: 1
// Explanation:
// [-4,4,0,0,-8,-4] is stable because the first and last elements are both -4, and the sum of the elements strictly between them is 4 + 0 + 0 + (-8) = -4

// Constraints:
//     3 <= capacity.length <= 10^5
//     -10^9 <= capacity[i] <= 10^9

import "fmt"

func countStableSubarrays(capacity []int) int64 {
    type Pair struct{ val, sum int }
    count := map[Pair]int{}
    res, sum := 0, capacity[0] // 前缀和
    for i := 1; i < len(capacity); i++ {
        res += (count[Pair{capacity[i], sum}])
        count[Pair{capacity[i - 1], capacity[i - 1] + sum}]++   
        sum += capacity[i]
    }
    return int64(res)
}

func countStableSubarrays1(capacity []int) int64 {
    res, n := 0, len(capacity)
    sum := make([]int, n + 1)
    for i, v := range capacity {
        sum[i+1] = v + sum[i]
    }
    type Pair struct{ x, y int}
    mark := make(map[Pair]int, n + 1)
    for i := 1; i < n + 1; i++ {
        x := capacity[i-1]
        prev := sum[i] - x * 2
        if v, has := mark[Pair{x, prev}]; has {
            res += v
        }
        if i >= 2 {
            if capacity[i-1] == capacity[i-2] && prev == sum[i-1] - capacity[i-2] * 2 {
                res--
            }
        }
        p := Pair{x, sum[i]}
        mark[p]++
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: capacity = [9,3,3,3,9]
    // Output: 2
    // Explanation:
    // [9,3,3,3,9] is stable because the first and last elements are both 9, and the sum of the elements strictly between them is 3 + 3 + 3 = 9.
    // [3,3,3] is stable because the first and last elements are both 3, and the sum of the elements strictly between them is 3.
    fmt.Println(countStableSubarrays([]int{9,3,3,3,9})) // 2
    // Example 2:
    // Input: capacity = [1,2,3,4,5]
    // Output: 0
    // Explanation:
    // No subarray of length at least 3 has equal first and last elements, so the answer is 0.
    fmt.Println(countStableSubarrays([]int{1,2,3,4,5})) // 0
    // Example 3:
    // Input: capacity = [-4,4,0,0,-8,-4]
    // Output: 1
    // Explanation:
    // [-4,4,0,0,-8,-4] is stable because the first and last elements are both -4, and the sum of the elements strictly between them is 4 + 0 + 0 + (-8) = -4
    fmt.Println(countStableSubarrays([]int{-4,4,0,0,-8,-4})) // 1

    fmt.Println(countStableSubarrays([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(countStableSubarrays([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(countStableSubarrays1([]int{9,3,3,3,9})) // 2
    fmt.Println(countStableSubarrays1([]int{1,2,3,4,5})) // 0
    fmt.Println(countStableSubarrays1([]int{-4,4,0,0,-8,-4})) // 1
    fmt.Println(countStableSubarrays1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(countStableSubarrays1([]int{9,8,7,6,5,4,3,2,1})) // 0
}