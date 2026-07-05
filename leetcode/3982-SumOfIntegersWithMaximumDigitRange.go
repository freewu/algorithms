package main 

// 3982. Sum of Integers with Maximum Digit Range
// You are given an integer array nums.

// The digit range of an integer is defined as the difference between its largest digit and smallest digit.

// For example, the digit range of 5724 is 7 - 2 = 5.

// Return the sum of all integers in nums whose digit range is equal to the maximum digit range among all integers in the array.

// Example 1:
// Input: nums = [5724,111,350]
// Output: 6074
// Explanation:
// i | nums[i] | Largest | Smallest | Digit Range
// 0 | 5724    | 7       | 2       | 5
// 1 | 111     | 1       | 1       | 0
// 2 | 350     | 5       | 0       | 5
// The maximum digit range is 5. The integers with this digit range are 5724 and 350, so the answer is 5724 + 350 = 6074.

// Example 2:
// Input: nums = [90,900]
// Output: 990
// Explanation:
// i | nums[i] | Largest | Smallest | Digit Range
// 0 | 90      | 9       | 9       | 0
// 1 | 900     | 9       | 0       | 9
// The maximum digit range is 9. Both integers have this digit range, so the answer is 90 + 900 = 990.

// Constraints:
//     1 <= nums.length <= 100
//     10 <= nums[i] <= 10^5

import "fmt"

func maxDigitRange(nums []int) int {
    res, maxRange := 0, 0
    for _, v := range nums {
        mn, mx := 9, 0
        for i := v; i > 0; i /= 10 {
            d := i % 10
            mn = min(mn, d)
            mx = max(mx, d)
        }
        r := mx - mn
        if r > maxRange {
            maxRange = r
            res = v // 重新累加
        } else if r == maxRange {
            res += v
        }
    }
    return res
}

func maxDigitRange1(nums []int) int {
    res, count := -1 << 61, make(map[int]int)
    for i, v := range nums {
        l, r := 1 << 61, -1 << 61
        for v > 0 {
            l = min(l, v % 10)
            r = max(r, v % 10)
            v /= 10
        }
        count[r - l] += nums[i]
    }
    for k := range count {
        res = max(res, k)
    }
    return count[res]
}

func main() {
    // Example 1:
    // Input: nums = [5724,111,350]
    // Output: 6074
    // Explanation:
    // i | nums[i] | Largest | Smallest | Digit Range
    // 0 | 5724    | 7       | 2       | 5
    // 1 | 111     | 1       | 1       | 0
    // 2 | 350     | 5       | 0       | 5
    // The maximum digit range is 5. The integers with this digit range are 5724 and 350, so the answer is 5724 + 350 = 6074.
    fmt.Println(maxDigitRange([]int{ 5724,111,350 })) // 6074
    // Example 2:
    // Input: nums = [90,900]
    // Output: 990
    // Explanation:
    // i | nums[i] | Largest | Smallest | Digit Range
    // 0 | 90      | 9       | 9       | 0
    // 1 | 900     | 9       | 0       | 9
    // The maximum digit range is 9. Both integers have this digit range, so the answer is 90 + 900 = 990.
    fmt.Println(maxDigitRange([]int{ 90,900 })) // 990

    fmt.Println(maxDigitRange([]int{ 1,2,3,4,5,6,7,8,9 })) // 45
    fmt.Println(maxDigitRange([]int{ 9,8,7,6,5,4,3,2,1 })) // 45

    fmt.Println(maxDigitRange1([]int{ 5724,111,350 })) // 6074
    fmt.Println(maxDigitRange1([]int{ 90,900 })) // 990
    fmt.Println(maxDigitRange1([]int{ 1,2,3,4,5,6,7,8,9 })) // 45
    fmt.Println(maxDigitRange1([]int{ 9,8,7,6,5,4,3,2,1 })) // 45
}