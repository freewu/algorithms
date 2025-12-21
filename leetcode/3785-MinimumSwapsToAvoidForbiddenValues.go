package main

// 3785. Minimum Swaps to Avoid Forbidden Values
// You are given two integer arrays, nums and forbidden, each of length n.

// You may perform the following operation any number of times (including zero):

//     Choose two distinct indices i and j, and swap nums[i] with nums[j].

// Return the minimum number of swaps required such that, for every index i, the value of nums[i] is not equal to forbidden[i]. 
// If no amount of swaps can ensure that every index avoids its forbidden value, return -1.

// Example 1:
// Input: nums = [1,2,3], forbidden = [3,2,1]
// Output: 1
// Explanation:
// One optimal set of swaps:
// Select indices i = 0 and j = 1 in nums and swap them, resulting in nums = [2, 1, 3].
// After this swap, for every index i, nums[i] is not equal to forbidden[i].

// Example 2:
// Input: nums = [4,6,6,5], forbidden = [4,6,5,5]
// Output: 2
// Explanation:
// One optimal set of swaps:
// Select indices i = 0 and j = 2 in nums and swap them, resulting in nums = [6, 6, 4, 5].
// Select indices i = 1 and j = 3 in nums and swap them, resulting in nums = [6, 5, 4, 6].
// After these swaps, for every index i, nums[i] is not equal to forbidden[i].

// Example 3:
// Input: nums = [7,7], forbidden = [8,7]
// Output: -1
// Explanation:
// It is not possible to make nums[i] different from forbidden[i] for all indices.

// Example 4:
// Input: nums = [1,2], forbidden = [2,1]
// Output: 0
// Explanation:
// No swaps are required because nums[i] is already different from forbidden[i] for all indices, so the answer is 0.

// Constraints:
//     1 <= n == nums.length == forbidden.length <= 10^5
//     1 <= nums[i], forbidden[i] <= 10^9

import "fmt"

func minSwaps(nums, forbidden []int) int {
    n := len(nums)
    total := make(map[int]int)
    for _, v := range nums {
        total[v]++
    }
    count := make(map[int]int)
    k, mx := 0, 0
    for i, x := range forbidden {
        total[x]++
        if total[x] > n { return -1 }
        if x == nums[i] {
            k++
            count[x]++
            mx = max(mx, count[x])
        }
    }
    return max((k + 1)/2, mx)
}

func minSwaps1(nums []int, forbidden []int) int {
    freq := make(map[int]int)
    n := len(nums)
    for i := 0; i < n; i++ {
        if nums[i] == forbidden[i] {
            freq[nums[i]]++
        }
    }
    best, sum := 0, 0
    for _, v := range freq {
        if v > best {
            best = v
        }
        sum += v
    }
    if best * 2 <= sum { return (sum + 1) / 2 }
    res := sum - best
    remain := 2 * best - sum
    who := -1
    for k, v := range freq {
        if v == best {
            who = k
            break
        }
    }
    for i := 0; i < n && remain > 0; i++ {
        if nums[i] != forbidden[i] && forbidden[i] != who && nums[i] != who {
            remain--
            res++
        }
    }
    if remain > 0 { return -1 }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3], forbidden = [3,2,1]
    // Output: 1
    // Explanation:
    // One optimal set of swaps:
    // Select indices i = 0 and j = 1 in nums and swap them, resulting in nums = [2, 1, 3].
    // After this swap, for every index i, nums[i] is not equal to forbidden[i].
    fmt.Println(minSwaps([]int{1,2,3}, []int{3,2,1})) // 1
    // Example 2:
    // Input: nums = [4,6,6,5], forbidden = [4,6,5,5]
    // Output: 2
    // Explanation:
    // One optimal set of swaps:
    // Select indices i = 0 and j = 2 in nums and swap them, resulting in nums = [6, 6, 4, 5].
    // Select indices i = 1 and j = 3 in nums and swap them, resulting in nums = [6, 5, 4, 6].
    // After these swaps, for every index i, nums[i] is not equal to forbidden[i].
    fmt.Println(minSwaps([]int{4,6,6,5}, []int{4,6,5,5})) // 2
    // Example 3:
    // Input: nums = [7,7], forbidden = [8,7]
    // Output: -1
    // Explanation:
    // It is not possible to make nums[i] different from forbidden[i] for all indices.
    fmt.Println(minSwaps([]int{7,7}, []int{8,7})) // -1
    // Example 4:
    // Input: nums = [1,2], forbidden = [2,1]
    // Output: 0
    // Explanation:
    // No swaps are required because nums[i] is already different from forbidden[i] for all indices, so the answer is 0.
    fmt.Println(minSwaps([]int{1,2}, []int{2,1})) // 0

    fmt.Println(minSwaps([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 1
    fmt.Println(minSwaps([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(minSwaps([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(minSwaps([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 5

    fmt.Println(minSwaps1([]int{1,2,3}, []int{3,2,1})) // 1
    fmt.Println(minSwaps1([]int{4,6,6,5}, []int{4,6,5,5})) // 2
    fmt.Println(minSwaps1([]int{7,7}, []int{8,7})) // -1
    fmt.Println(minSwaps1([]int{1,2}, []int{2,1})) // 0
    fmt.Println(minSwaps1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 1
    fmt.Println(minSwaps1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(minSwaps1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(minSwaps1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 5
}