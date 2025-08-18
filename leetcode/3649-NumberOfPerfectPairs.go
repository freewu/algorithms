package main

// 3649. Number of Perfect Pairs
// You are given an integer array nums.

// A pair of indices (i, j) is called perfect if the following conditions are satisfied:
//     1. i < j
//     2. Let a = nums[i], b = nums[j]. Then:
//         2.1 min(|a - b|, |a + b|) <= min(|a|, |b|)
//         2.3 max(|a - b|, |a + b|) >= max(|a|, |b|)

// Return the number of distinct perfect pairs.

// Note: The absolute value |x| refers to the non-negative value of x.

// Example 1:
// Input: nums = [0,1,2,3]
// Output: 2
// Explanation:
// There are 2 perfect pairs:
// (i, j)	(a, b)	min(|a − b|, |a + b|)	min(|a|, |b|)	max(|a − b|, |a + b|)	max(|a|, |b|)
// (1, 2)	(1, 2)	min(|1 − 2|, |1 + 2|) = 1	1	max(|1 − 2|, |1 + 2|) = 3	2
// (2, 3)	(2, 3)	min(|2 − 3|, |2 + 3|) = 1	2	max(|2 − 3|, |2 + 3|) = 5	3

// Example 2:
// Input: nums = [-3,2,-1,4]
// Output: 4
// Explanation:
// There are 4 perfect pairs:
// (i, j)	(a, b)	min(|a − b|, |a + b|)	min(|a|, |b|)	max(|a − b|, |a + b|)	max(|a|, |b|)
// (0, 1)	(-3, 2)	min(|-3 - 2|, |-3 + 2|) = 1	2	max(|-3 - 2|, |-3 + 2|) = 5	3
// (0, 3)	(-3, 4)	min(|-3 - 4|, |-3 + 4|) = 1	3	max(|-3 - 4|, |-3 + 4|) = 7	4
// (1, 2)	(2, -1)	min(|2 - (-1)|, |2 + (-1)|) = 1	1	max(|2 - (-1)|, |2 + (-1)|) = 3	2
// (1, 3)	(2, 4)	min(|2 - 4|, |2 + 4|) = 2	2	max(|2 - 4|, |2 + 4|) = 6	4

// Example 3:
// Input: nums = [1,10,100,1000]
// Output: 0
// Explanation:
// There are no perfect pairs. Thus, the answer is 0.

// Constraints:
//     2 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"
import "sort"

func perfectPairs(nums []int) int64 {
    res, curr, n := 0, 0, len(nums)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < n; i++ {
        nums[i] = abs(nums[i])
    }
    sort.Ints(nums)
    for i := 0; i < n; i++ {
        curr = nums[i]
        index := sort.SearchInts(nums[i + 1:], curr * 2 + 1)
        if i + 1 + index == n {
            index--
        }
        if nums[index+ 1 + i] > curr*2 {
            index--
        }
        res += (index + 1)
    }
    return int64(res)
}

func perfectPairs1(nums []int) int64 {
    res, right, n := 0, 0, len(nums)
    arr := make([]int, n)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i, v := range nums {
        arr[i] = abs(v)
    }
    sort.Ints(arr)
    for left := 0; left < n; left++ {
        target := 2 * arr[left]
        for right < n && arr[right] <= target {
            right++
        }
        v := right - 1 - left
        if v > 0 {
            res += v
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [0,1,2,3]
    // Output: 2
    // Explanation:
    // There are 2 perfect pairs:
    // (i, j)	(a, b)	min(|a − b|, |a + b|)	min(|a|, |b|)	max(|a − b|, |a + b|)	max(|a|, |b|)
    // (1, 2)	(1, 2)	min(|1 − 2|, |1 + 2|) = 1	1	max(|1 − 2|, |1 + 2|) = 3	2
    // (2, 3)	(2, 3)	min(|2 − 3|, |2 + 3|) = 1	2	max(|2 − 3|, |2 + 3|) = 5	3
    fmt.Println(perfectPairs([]int{0,1,2,3})) // 2
    // Example 2:
    // Input: nums = [-3,2,-1,4]
    // Output: 4
    // Explanation:
    // There are 4 perfect pairs:
    // (i, j)	(a, b)	min(|a − b|, |a + b|)	min(|a|, |b|)	max(|a − b|, |a + b|)	max(|a|, |b|)
    // (0, 1)	(-3, 2)	min(|-3 - 2|, |-3 + 2|) = 1	2	max(|-3 - 2|, |-3 + 2|) = 5	3
    // (0, 3)	(-3, 4)	min(|-3 - 4|, |-3 + 4|) = 1	3	max(|-3 - 4|, |-3 + 4|) = 7	4
    // (1, 2)	(2, -1)	min(|2 - (-1)|, |2 + (-1)|) = 1	1	max(|2 - (-1)|, |2 + (-1)|) = 3	2
    // (1, 3)	(2, 4)	min(|2 - 4|, |2 + 4|) = 2	2	max(|2 - 4|, |2 + 4|) = 6	4
    fmt.Println(perfectPairs([]int{-3,2,-1,4})) // 4
    // Example 3:
    // Input: nums = [1,10,100,1000]
    // Output: 0
    // Explanation:
    // There are no perfect pairs. Thus, the answer is 0.
    fmt.Println(perfectPairs([]int{1,10,100,1000})) // 0

    fmt.Println(perfectPairs([]int{1,2,3,4,5,6,7,8,9})) // 20
    fmt.Println(perfectPairs([]int{9,8,7,6,5,4,3,2,1})) // 20

    fmt.Println(perfectPairs1([]int{0,1,2,3})) // 2
    fmt.Println(perfectPairs1([]int{-3,2,-1,4})) // 4
    fmt.Println(perfectPairs1([]int{1,10,100,1000})) // 0
    fmt.Println(perfectPairs1([]int{1,2,3,4,5,6,7,8,9})) // 20
    fmt.Println(perfectPairs1([]int{9,8,7,6,5,4,3,2,1})) // 20
}
