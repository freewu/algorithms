package main

// 3134. Find the Median of the Uniqueness Array
// You are given an integer array nums. 
// The uniqueness array of nums is the sorted array that contains the number of distinct elements of all the subarrays of nums. 
// In other words, it is a sorted array consisting of distinct(nums[i..j]), for all 0 <= i <= j < nums.length.

// Here, distinct(nums[i..j]) denotes the number of distinct elements in the subarray that starts at index i and ends at index j.

// Return the median of the uniqueness array of nums.

// Note that the median of an array is defined as the middle element of the array when it is sorted in non-decreasing order. 
// If there are two choices for a median, the smaller of the two values is taken.

// Example 1:
// Input: nums = [1,2,3]
// Output: 1
// Explanation:
// The uniqueness array of nums is [distinct(nums[0..0]), distinct(nums[1..1]), distinct(nums[2..2]), distinct(nums[0..1]), distinct(nums[1..2]), distinct(nums[0..2])] which is equal to [1, 1, 1, 2, 2, 3]. The uniqueness array has a median of 1. Therefore, the answer is 1.

// Example 2:
// Input: nums = [3,4,3,4,5]
// Output: 2
// Explanation:
// The uniqueness array of nums is [1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3]. The uniqueness array has a median of 2. Therefore, the answer is 2.

// Example 3:
// Input: nums = [4,3,5,4]
// Output: 2
// Explanation:
// The uniqueness array of nums is [1, 1, 1, 1, 2, 2, 2, 3, 3, 3]. The uniqueness array has a median of 2. Therefore, the answer is 2.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"
import "sort"

func medianOfUniquenessArray(nums []int) int {
    n := len(nums)
    median := (int64(n) * int64(n + 1) / 2 + 1) / 2
    // 检测数组中不同元素数目小于等于 t 的连续子数组数目是否大于等于 median
    check := func(t int) bool {
        cnt := make(map[int]int)
        tot := int64(0)
        for i, j := 0, 0; i < n; i++ {
            cnt[nums[i]]++
            for len(cnt) > t {
                cnt[nums[j]]--
                if cnt[nums[j]] == 0 {
                    delete(cnt, nums[j])
                }
                j++
            }
            tot += int64(i - j + 1)
        }
        return tot >= median
    }
    res, low, high := 0, 1, n
    for low <= high {
        mid := (low + high) / 2
        if check(mid) {
            res, high= mid, mid - 1
        } else {
            low = mid + 1
        }
    }
    return res
}

func medianOfUniquenessArray1(nums []int) int {
    n := len(nums)
    k := (n*(n+1)/2 + 1) / 2
    return 1 + sort.Search(n-1, func(upper int) bool {
        upper++
        cnt, l, freq := 0, 0, map[int]int{}
        for r, in := range nums {
            freq[in]++ // 移入右端点
            for len(freq) > upper { // 窗口内元素过多
                out := nums[l]
                freq[out]-- // 移出左端点
                if freq[out] == 0 {
                    delete(freq, out)
                }
                l++
            }
            cnt += r - l + 1 // 右端点固定为 r 时，有 r-l+1 个合法左端点
            if cnt >= k {
                return true
            }
        }
        return false
    })
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 1
    // Explanation:
    // The uniqueness array of nums is [distinct(nums[0..0]), distinct(nums[1..1]), distinct(nums[2..2]), distinct(nums[0..1]), distinct(nums[1..2]), distinct(nums[0..2])] which is equal to [1, 1, 1, 2, 2, 3]. The uniqueness array has a median of 1. Therefore, the answer is 1.
    fmt.Println(medianOfUniquenessArray([]int{1,2,3})) // 1
    // Example 2:
    // Input: nums = [3,4,3,4,5]
    // Output: 2
    // Explanation:
    // The uniqueness array of nums is [1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3]. The uniqueness array has a median of 2. Therefore, the answer is 2.
    fmt.Println(medianOfUniquenessArray([]int{3,4,3,4,5})) // 2
    // Example 3:
    // Input: nums = [4,3,5,4]
    // Output: 2
    // Explanation:
    // The uniqueness array of nums is [1, 1, 1, 1, 2, 2, 2, 3, 3, 3]. The uniqueness array has a median of 2. Therefore, the answer is 2.
    fmt.Println(medianOfUniquenessArray([]int{4,3,5,4})) // 2

    fmt.Println(medianOfUniquenessArray1([]int{1,2,3})) // 1
    fmt.Println(medianOfUniquenessArray1([]int{3,4,3,4,5})) // 2
    fmt.Println(medianOfUniquenessArray1([]int{4,3,5,4})) // 2
}