package main

// 4011. Count Subarrays With Even Odd Ratio I
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
// Subarray    | Values        | Even Count | Odd Count | Ratio
// nums[0..0]	| [1]	        | 0	         | 1	     | 0 / 1
// nums[0..1]	| [1, 2]	    | 1	         | 1	     | 1 / 1
// nums[0..2]	| [1, 2, 1]     | 1	         | 2	     | 1 / 2
// nums[0..3]	| [1, 2, 1, 2]  | 2	         | 2	     | 2 / 2
// nums[1..2]	| [2, 1]	    | 1	         | 1	     | 1 / 1
// nums[2..2]	| [1]	        | 0	         | 1	     | 0 / 1
// nums[2..3]	| [1, 2]	    | 1	         | 1	     | 1 / 1
// Thus, the number of valid subarrays is 7.

// Example 2:
// Input: nums = [2,2,1], a = 2, b = 1
// Output: 3
// Explanation:
// The following are the valid subarrays:
// Subarray    | Values        | Even Count | Odd Count | Ratio
// nums[0..2]	| [2, 2, 1]	| 2	         | 1	     | 2 / 1
// nums[1..2]	| [2, 1]	| 1	         | 1	     | 1 / 1
// nums[2..2]	| [1]	    | 0	         | 1	     | 0 / 1    
// Thus, the number of valid subarrays is 3.

// Example 3:
// Input: nums = [2,2,2], a = 1, b = 1
// Output: 0
// Explanation:
// Every subarray contains 0 odd numbers, so no subarray is valid.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 1000
//     1 <= a, b <= 1000

import "fmt"
import "sort"

func countRatioSubarrays(nums []int, a int, b int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        even, odd := 0, 0
        for j := i; j < len(nums); j++ {
            if nums[j]%2 == 0 {
                even++
            } else {
                odd++
            }
            if odd > 0 && even * b <= a * odd {
                res++
            }
        }
    }
    return res
}

func countRatioSubarrays1(nums []int, a int, b int) int {
    n, even, odd := len(nums), 0, 0
    arr := make([]int, n + 1)
    for i, v := range nums {
        if v%2 == 0 {
            even++
        } else {
            odd++
        }
        arr[i+1] = b * even - a * odd
    }
    sorted := append([]int(nil), arr...)
    sort.Ints(sorted)
    bit := make([]int, n+2)
    res, inserted := 0, 0
    for _, v := range arr {
        index := sort.SearchInts(sorted, v) + 1
        less := 0
        for i := index - 1; i > 0; i -= i & -i {
            less += bit[i]
        }
        res += inserted - less
        for i := index; i < len(bit); i += i & -i {
            bit[i]++
        }
        inserted++
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1,2], a = 3, b = 2
    // Output: 7
    // Explanation:
    // The following are the valid subarrays:
    // Subarray	Values	Even Count	Odd Count	Ratio
    // nums[0..0]	[1]	0	1	0 / 1
    // nums[0..1]	[1, 2]	1	1	1 / 1
    // nums[0..2]	[1, 2, 1]	1	2	1 / 2
    // nums[0..3]	[1, 2, 1, 2]	2	2	2 / 2
    // nums[1..2]	[2, 1]	1	1	1 / 1
    // nums[2..2]	[1]	0	1	0 / 1
    // nums[2..3]	[1, 2]	1	1	1 / 1
    // Thus, the number of valid subarrays is 7.
    fmt.Println(countRatioSubarrays([]int{1,2,1,2}, 3, 2)) // 7
    // Example 2:
    // Input: nums = [2,2,1], a = 2, b = 1
    // Output: 3
    // Explanation:
    // The following are the valid subarrays:
    // Subarray	Values	Even Count	Odd Count	Ratio
    // nums[0..2]	[2, 2, 1]	2	1	2 / 1
    // nums[1..2]	[2, 1]	1	1	1 / 1
    // nums[2..2]	[1]	0	1	0 / 1
    // Thus, the number of valid subarrays is 3.
    fmt.Println(countRatioSubarrays([]int{2,2,1}, 2, 1)) // 3
    // Example 3:
    // Input: nums = [2,2,2], a = 1, b = 1
    // Output: 0
    // Explanation:
    // Every subarray contains 0 odd numbers, so no subarray is valid.
    fmt.Println(countRatioSubarrays([]int{2,2,2}, 1, 1)) // 0

    fmt.Println(countRatioSubarrays([]int{1,2,3,4,5,6,7,8,9}, 3, 2)) // 38  
    fmt.Println(countRatioSubarrays([]int{9,8,7,6,5,4,3,2,1}, 3, 2)) // 38

    fmt.Println(countRatioSubarrays([]int{1,2,3,4,5,6,7,8,9}, 1, 1)) // 35 
    fmt.Println(countRatioSubarrays([]int{9,8,7,6,5,4,3,2,1}, 1, 1)) // 35
    fmt.Println(countRatioSubarrays([]int{1,2,3,4,5,6,7,8,9}, 1, 1000)) // 5  
    fmt.Println(countRatioSubarrays([]int{9,8,7,6,5,4,3,2,1}, 1, 1000)) // 5
    fmt.Println(countRatioSubarrays([]int{1,2,3,4,5,6,7,8,9}, 1000, 1)) // 41 
    fmt.Println(countRatioSubarrays([]int{9,8,7,6,5,4,3,2,1}, 1000, 1)) // 41
    fmt.Println(countRatioSubarrays([]int{1,2,3,4,5,6,7,8,9}, 1000, 1000)) // 35
    fmt.Println(countRatioSubarrays([]int{9,8,7,6,5,4,3,2,1}, 1000, 1000)) // 35

    fmt.Println(countRatioSubarrays1([]int{1,2,1,2}, 3, 2)) // 7
    fmt.Println(countRatioSubarrays1([]int{2,2,1}, 2, 1)) // 3
    fmt.Println(countRatioSubarrays1([]int{2,2,2}, 1, 1)) // 0
    fmt.Println(countRatioSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 3, 2)) // 38  
    fmt.Println(countRatioSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 3, 2)) // 38
    fmt.Println(countRatioSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 1, 1)) // 35 
    fmt.Println(countRatioSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 1, 1)) // 35
    fmt.Println(countRatioSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 1, 1000)) // 5  
    fmt.Println(countRatioSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 1, 1000)) // 5
    fmt.Println(countRatioSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 1000, 1)) // 41 
    fmt.Println(countRatioSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 1000, 1)) // 41
    fmt.Println(countRatioSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 1000, 1000)) // 35
    fmt.Println(countRatioSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 1000, 1000)) // 35
}