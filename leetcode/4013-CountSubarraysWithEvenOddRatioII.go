package main

// 4013. Count Subarrays With Even Odd Ratio II
// You are given an integer array nums and two integers a and b.

// For a subarray, let:
//     1. x be the number of even elements.
//     2. y be the number of odd elements.

// The ratio of even to odd numbers in a subarray is defined as x / y, where the ratio is compared by its exact rational value.

// A subarray is considered valid if:
//     1. y > 0, and
//     2. x / y <= a / b.
    
// Return the number of valid subarrays in nums.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,2,1,2], a = 3, b = 2
// Output: 7
// Explanation:
// The following are the valid subarrays:
// Subarray   | Values         | Even Count | Odd Count | Ratio
// nums[0..0]	[1]	            | 0	         | 1         | 0 / 1
// nums[0..1]	[1, 2]	        | 1	         | 1         | 1 / 1
// nums[0..2]	[1, 2, 1]	    | 1	         | 2         | 1 / 2
// nums[0..3]	[1, 2, 1, 2]	| 2	         | 2         | 2 / 2
// nums[1..2]	[2, 1]	        | 1          | 1         | 1 / 1
// nums[2..2]	[1]	            | 0	         | 1         | 0 / 1
// nums[2..3]	[1, 2]	        | 1          | 1         | 1 / 1
// Thus, the number of valid subarrays is 7.

// Example 2:
// Input: nums = [2,2,1], a = 2, b = 1
// Output: 3
// Explanation:
// The following are the valid subarrays:
// Subarray   | Values         | Even Count | Odd Count | Ratio
// nums[0..2]	[2, 2, 1]	    | 2	         | 1         | 2 / 1
// nums[1..2]	[2, 1]	        | 1	         | 1         | 1 / 1
// nums[2..2]	[1]	            | 0	         | 1         | 0 / 1
// Thus, the number of valid subarrays is 3.

// Example 3:
// Input: nums = [2,2,2], a = 1, b = 1
// Output: 0
// Explanation:
// Every subarray contains 0 odd numbers, so no subarray is valid.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= a, b <= 10^9

import "fmt"
import "sort"
import "slices"

type Fenwick []int

func (t Fenwick) add(i int) {
    for ; i < len(t); i += i & -i {
        t[i]++
    }
}

func (t Fenwick) pre(i int) (res int) {
    for ; i > 0; i &= i - 1 {
        res += t[i]
    }
    return
}

// 值域树状数组
func countRatioSubarrays(nums []int, a, b int) int64 {
    sum := make([]int, len(nums) + 1)
    value := [2]int{-b, a}
    for i, v := range nums {
        sum[i+1] = sum[i] + value[v&1] // 偶数视作 -b，奇数视作 a
    }
    // sum 排序去重
    sorted := slices.Clone(sum)
    slices.Sort(sorted)
    sorted = slices.Compact(sorted)
    t := make(Fenwick, len(sorted)+1)
    res := 0
    for _, s := range sum {
        s = sort.SearchInts(sorted, s) + 1 // 离散化（从 1 开始）
        res += t.pre(s) // 计算在 s 左边有多少个 <= s 的数
        t.add(s)
    }
    return int64(res)
}

func countRatioSubarrays1(nums []int, a, b int) int64 {
    sum := make([]int, len(nums)+1) 
    value := [2]int{-b, a}
    for i, v := range nums {
        sum[i+1] = sum[i] + value[v&1] // 偶数视作 -b，奇数视作 a
    }
    var mergeCount func(sum []int) int
    mergeCount = func(sum []int) int { // 归并排序
        n := len(sum)
        if n <= 1 {
            return 0
        }
        left, right := slices.Clone(sum[:n/2]), slices.Clone(sum[n/2:])
        res := mergeCount(left) + mergeCount(right) // left 和 right 各自的合法数对个数
        l, r := 0, 0
        for i := range sum {
            // 计算一个在 left 中，另一个在 right 中的合法数对个数
            if l < len(left) && (r == len(right) || left[l] <= right[r]) {
                sum[i] = left[l]
                l++
            } else {
                res += l // left[:l] 中的数都 <= right[r]，这有 l 个
                sum[i] = right[r]
                r++
            }
        }
        return res
    }
    return int64(mergeCount(sum))
}

func countRatioSubarrays2(nums []int, a int, b int) int64 {
    res, n := int64(0), len(nums)
    prefix := make([]int,n+1)
    for i, v := range nums {
        if v % 2 == 0 {
            prefix[i+1] = prefix[i] + b
        } else {
            prefix[i+1] = prefix[i] - a
        }
    }
    sorted := append([]int{}, prefix...)
    sort.Ints(sorted)
    bit := make([]int, n + 2)
    add := func(i int) {
        for ; i <= n+1; i += i & -i {
            bit[i]++
        }
    }
    query := func(i int) int {
        s := 0
        for ; i > 0; i -= i & -i {
            s += bit[i]
        }
        return s
    }
    for k, p := range prefix {
        r := sort.SearchInts(sorted, p) + 1
        res += int64(k - query(r-1)) 
        add(r)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1,2], a = 3, b = 2
    // Output: 7
    // Explanation:
    // The following are the valid subarrays:
    // Subarray   | Values         | Even Count | Odd Count | Ratio
    // nums[0..0]	[1]	            | 0	         | 1         | 0 / 1
    // nums[0..1]	[1, 2]	        | 1	         | 1         | 1 / 1
    // nums[0..2]	[1, 2, 1]	    | 1	         | 2         | 1 / 2
    // nums[0..3]	[1, 2, 1, 2]	| 2	         | 2         | 2 / 2
    // nums[1..2]	[2, 1]	        | 1          | 1         | 1 / 1
    // nums[2..2]	[1]	            | 0	         | 1         | 0 / 1
    // nums[2..3]	[1, 2]	        | 1          | 1         | 1 / 1
    // Thus, the number of valid subarrays is 7.
    fmt.Println(countRatioSubarrays([]int{1,2,1,2}, 3, 2)) // 7
    // Example 2:
    // Input: nums = [2,2,1], a = 2, b = 1
    // Output: 3
    // Explanation:
    // The following are the valid subarrays:
    // Subarray   | Values         | Even Count | Odd Count | Ratio
    // nums[0..2]	[2, 2, 1]	    | 2	         | 1         | 2 / 1
    // nums[1..2]	[2, 1]	        | 1	         | 1         | 1 / 1
    // nums[2..2]	[1]	            | 0	         | 1         | 0 / 1
    // Thus, the number of valid subarrays is 3.
    fmt.Println(countRatioSubarrays([]int{2,2,1}, 2, 1)) // 3
    // Example 3:
    // Input: nums = [2,2,2], a = 1, b = 1
    // Output: 0
    // Explanation:
    // Every subarray contains 0 odd numbers, so no subarray is valid.
    fmt.Println(countRatioSubarrays([]int{2,2,2}, 1, 1)) // 0

    fmt.Println(countRatioSubarrays([]int{1,2,3,4,5,6,7,8,9}, 1, 1)) // 35 
    fmt.Println(countRatioSubarrays([]int{9,8,7,6,5,4,3,2,1}, 1, 1)) // 35
    fmt.Println(countRatioSubarrays([]int{1,2,3,4,5,6,7,8,9}, 1, 1_000_000_000)) // 5  
    fmt.Println(countRatioSubarrays([]int{9,8,7,6,5,4,3,2,1}, 1, 1_000_000_000)) // 5
    fmt.Println(countRatioSubarrays([]int{1,2,3,4,5,6,7,8,9}, 1_000_000_000, 1)) // 41 
    fmt.Println(countRatioSubarrays([]int{9,8,7,6,5,4,3,2,1}, 1_000_000_000000, 1)) // 41
    fmt.Println(countRatioSubarrays([]int{1,2,3,4,5,6,7,8,9}, 1_000_000_000, 1_000_000_000)) // 35
    fmt.Println(countRatioSubarrays([]int{9,8,7,6,5,4,3,2,1}, 1_000_000_000, 1_000_000_000)) // 35

    fmt.Println(countRatioSubarrays1([]int{1,2,1,2}, 3, 2)) // 7
    fmt.Println(countRatioSubarrays1([]int{2,2,1}, 2, 1)) // 3
    fmt.Println(countRatioSubarrays1([]int{2,2,2}, 1, 1)) // 0
    fmt.Println(countRatioSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 1, 1)) // 35 
    fmt.Println(countRatioSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 1, 1)) // 35
    fmt.Println(countRatioSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 1, 1_000_000_000)) // 5  
    fmt.Println(countRatioSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 1, 1_000_000_000)) // 5
    fmt.Println(countRatioSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 1_000_000_000, 1)) // 41 
    fmt.Println(countRatioSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 1_000_000_000000, 1)) // 41
    fmt.Println(countRatioSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 1_000_000_000, 1_000_000_000)) // 35
    fmt.Println(countRatioSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 1_000_000_000, 1_000_000_000)) // 35

    fmt.Println(countRatioSubarrays2([]int{1,2,1,2}, 3, 2)) // 7
    fmt.Println(countRatioSubarrays2([]int{2,2,1}, 2, 1)) // 3
    fmt.Println(countRatioSubarrays2([]int{2,2,2}, 1, 1)) // 0
    fmt.Println(countRatioSubarrays2([]int{1,2,3,4,5,6,7,8,9}, 1, 1)) // 35 
    fmt.Println(countRatioSubarrays2([]int{9,8,7,6,5,4,3,2,1}, 1, 1)) // 35
    fmt.Println(countRatioSubarrays2([]int{1,2,3,4,5,6,7,8,9}, 1, 1_000_000_000)) // 5  
    fmt.Println(countRatioSubarrays2([]int{9,8,7,6,5,4,3,2,1}, 1, 1_000_000_000)) // 5
    fmt.Println(countRatioSubarrays2([]int{1,2,3,4,5,6,7,8,9}, 1_000_000_000, 1)) // 41 
    fmt.Println(countRatioSubarrays2([]int{9,8,7,6,5,4,3,2,1}, 1_000_000_000000, 1)) // 41
    fmt.Println(countRatioSubarrays2([]int{1,2,3,4,5,6,7,8,9}, 1_000_000_000, 1_000_000_000)) // 35
    fmt.Println(countRatioSubarrays2([]int{9,8,7,6,5,4,3,2,1}, 1_000_000_000, 1_000_000_000)) // 35
}