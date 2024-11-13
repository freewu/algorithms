package main

// 3346. Maximum Frequency of an Element After Performing Operations I
// You are given an integer array nums and two integers k and numOperations.

// You must perform an operation numOperations times on nums, where in each operation you:
//     1. Select an index i that was not selected in any previous operations.
//     2. Add an integer in the range [-k, k] to nums[i].

// Return the maximum possible frequency of any element in nums after performing the operations.

// Example 1:
// Input: nums = [1,4,5], k = 1, numOperations = 2
// Output: 2
// Explanation:
// We can achieve a maximum frequency of two by:
//     Adding 0 to nums[1]. nums becomes [1, 4, 5].
//     Adding -1 to nums[2]. nums becomes [1, 4, 4].

// Example 2:
// Input: nums = [5,11,20,20], k = 5, numOperations = 1
// Output: 2
// Explanation:
// We can achieve a maximum frequency of two by:
//     Adding 0 to nums[1].

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     0 <= k <= 10^5
//     0 <= numOperations <= nums.length

import "fmt"
import "sort"

func maxFrequency(nums []int, k int, numOperations int) int {
    res, n := 1, len(nums) // 这个数成为众数以后可以出现的最多的次数 (len(nums[i]-k,nums[i]+k), numOperations)
    sort.Ints(nums)
    count := make(map[int]int, n)
    for _, v := range nums {
        count[v]++
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    find1 := func (x int, nums []int) int { // 求第一个>=l的元素下标
        res, l, r := -1, 0, len(nums) - 1
        for l <= r {
            mid := (l + r + 1) / 2
            if nums[mid] >= x {
                res, r = mid, mid - 1
            } else {
                l = mid + 1
            }
        }
        return res
    }
    find2 := func(x int, nums []int) int { // <= r 的最后一个元素下标
        res, l, r := -1, 0, len(nums)-1
        for l <= r {
            mid := (l + r + 1) / 2
            if nums[mid] <= x {
                res, l = mid, mid + 1
            } else {
                r = mid - 1
            }
        }
        return res
    }
    // 求第一个 >=l 的元素下标 ，<=r 的最后一个元素下标
    for i := nums[0]; i <= nums[n-1]; i++ {
        l, r := i - k, i + k
        idxl, idxr := find1(l, nums), find2(r, nums)
        res = max(res, min(idxr - idxl + 1, numOperations + count[i])) // 查询区间[idxl,idxr]里i的出现次数
    }
    return res
}

func maxFrequency1(nums []int, k int, numOperations int) int {
    sort.Ints(nums)
    res, n := 0, len(nums)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, l, r := 0, 0, 0; i < n; {
        j := i
        for l < j && nums[j]-nums[l] > k  { l++ }
        r = max(r, i)
        for r < n && nums[r]-nums[j] <= k { r++ }
        for i < n && nums[i] == nums[j] { i++ }
        res = max(res, i - j + min(j - l + r - i, numOperations))
    }
    for l, r := 0, 0; r < n; r++ { // 所有的数，都变化的情况
        for l < r && nums[r]-nums[l] > 2 * k { l++ }
        res = max(res, min(numOperations, r - l + 1))
    }
    return res
}

func maxFrequency2(nums []int, k int, numOperations int) int {
    res, sum, mx := 0, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        mx = max(mx, v)
    }
    mx++
    count := make([]int, mx)
    for _, v := range nums {
        count[v]++
    }
    for i := 0; i < min(k, mx); i++ {
        sum += count[i]
    }
    for i := 0; i < mx; i++ {
        if i + k < mx { sum += count[i+k] }
        res = max(res, count[i] + min(numOperations, sum - count[i]))
        if i - k >= 0 {
            sum -= count[i - k]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,4,5], k = 1, numOperations = 2
    // Output: 2
    // Explanation:
    // We can achieve a maximum frequency of two by:
    //     Adding 0 to nums[1]. nums becomes [1, 4, 5].
    //     Adding -1 to nums[2]. nums becomes [1, 4, 4].
    fmt.Println(maxFrequency([]int{1,4,5}, 1, 2)) // 2
    // Example 2: 
    // Input: nums = [5,11,20,20], k = 5, numOperations = 1
    // Output: 2
    // Explanation:
    // We can achieve a maximum frequency of two by:
    //     Adding 0 to nums[1].
    fmt.Println(maxFrequency([]int{5,11,20,20}, 5, 1)) // 2

    fmt.Println(maxFrequency1([]int{1,4,5}, 1, 2)) // 2
    fmt.Println(maxFrequency1([]int{5,11,20,20}, 5, 1)) // 2

    fmt.Println(maxFrequency2([]int{1,4,5}, 1, 2)) // 2
    fmt.Println(maxFrequency2([]int{5,11,20,20}, 5, 1)) // 2
}