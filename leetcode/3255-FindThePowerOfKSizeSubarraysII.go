package main

// 3255. Find the Power of K-Size Subarrays II
// You are given an array of integers nums of length n and a positive integer k.

// The power of an array is defined as:
//     Its maximum element if all of its elements are consecutive and sorted in ascending order.
//     -1 otherwise.

// You need to find the power of all subarrays of nums of size k.

// Return an integer array results of size n - k + 1, where results[i] is the power of nums[i..(i + k - 1)].

// Example 1:
// Input: nums = [1,2,3,4,3,2,5], k = 3
// Output: [3,4,-1,-1,-1]
// Explanation:
// There are 5 subarrays of nums of size 3:
// [1, 2, 3] with the maximum element 3.
// [2, 3, 4] with the maximum element 4.
// [3, 4, 3] whose elements are not consecutive.
// [4, 3, 2] whose elements are not sorted.
// [3, 2, 5] whose elements are not consecutive.

// Example 2:
// Input: nums = [2,2,2,2,2], k = 4
// Output: [-1,-1]

// Example 3:
// Input: nums = [3,2,3,2,3,2], k = 2
// Output: [-1,3,-1,3,-1]

// Constraints:
//     1 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^6
//     1 <= k <= n

import "fmt"

func resultsArray(nums []int, k int) []int {
    res, l := make([]int, 0, len(nums)), 0
    for i, v := range nums {
        if i > 0 && nums[i-1] + 1 == v {
            l++
        } else {
            l = 1
        }
        if l >= k { // 如果 所有 元素都是依次 连续 且 上升 的，那么能量值为 最大 的元素
            res = append(res, nums[i])
        } else if i >= k - 1 { // 否则为 -1
            res = append(res, -1)
        }
    }
    return res
}

func resultsArray1(nums []int, k int) []int {
    res, count := make([]int, len(nums) - k + 1), 0
    for i := range res {
        res[i] = -1
    }
    for i, v := range nums {
        if i == 0 || v == nums[i - 1] + 1 {
            count++
        } else {
            count = 1
        }
        if count >= k {
            res[i - k + 1] = v
        }
    }
    return res
}

func resultsArray2(nums []int, k int) []int {
    n, start, prev := len(nums), 0, nums[0]
    res := make([]int, n - k + 1) 
    for i := range res { // fill -1
        res[i] = -1 
    }
    for i := 1; i < n; i += 1 {
        if prev + 1 == nums[i] {
            prev = nums[i] 
            continue 
        }
        if ((i - start) >= k) {
            for j := i - k; j >= start; j-- {
                res[j] = nums[j + k - 1]
            }
        }
        prev, start = nums[i], i
    } 
    if (n - start) >= k {
        for j := n - k; j >= start; j-- {
            res[j] = nums[j + k - 1]
        }
    }
    return res 
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,3,2,5], k = 3
    // Output: [3,4,-1,-1,-1]
    // Explanation:
    // There are 5 subarrays of nums of size 3:
    // [1, 2, 3] with the maximum element 3.
    // [2, 3, 4] with the maximum element 4.
    // [3, 4, 3] whose elements are not consecutive.
    // [4, 3, 2] whose elements are not sorted.
    // [3, 2, 5] whose elements are not consecutive.
    fmt.Println(resultsArray([]int{1,2,3,4,3,2,5}, 3)) // [3,4,-1,-1,-1]
    // Example 2:
    // Input: nums = [2,2,2,2,2], k = 4
    // Output: [-1,-1]
    fmt.Println(resultsArray([]int{2,2,2,2,2}, 4)) // [-1,-1]
    // Example 3:
    // Input: nums = [3,2,3,2,3,2], k = 2
    // Output: [-1,3,-1,3,-1]
    fmt.Println(resultsArray([]int{3,2,3,2,3,2}, 2)) // [-1,3,-1,3,-1]

    fmt.Println(resultsArray1([]int{1,2,3,4,3,2,5}, 3)) // [3,4,-1,-1,-1]
    fmt.Println(resultsArray1([]int{2,2,2,2,2}, 4)) // [-1,-1]
    fmt.Println(resultsArray1([]int{3,2,3,2,3,2}, 2)) // [-1,3,-1,3,-1]

    fmt.Println(resultsArray2([]int{1,2,3,4,3,2,5}, 3)) // [3,4,-1,-1,-1]
    fmt.Println(resultsArray2([]int{2,2,2,2,2}, 4)) // [-1,-1]
    fmt.Println(resultsArray2([]int{3,2,3,2,3,2}, 2)) // [-1,3,-1,3,-1]
}