package main

// 3724. Minimum Operations to Transform Array
// You are given two integer arrays nums1 of length n and nums2 of length n + 1.

// You want to transform nums1 into nums2 using the minimum number of operations.

// You may perform the following operations any number of times, each time choosing an index i:
//     1. Increase nums1[i] by 1.
//     2. Decrease nums1[i] by 1.
//     3. Append nums1[i] to the end of the array.

// Return the minimum number of operations required to transform nums1 into nums2.

// Example 1:
// Input: nums1 = [2,8], nums2 = [1,7,3]
// Output: 4
// Explanation:
// Step	i	Operation	nums1[i]	    Updated  nums1
// 1	    0	Append	        -	            [2, 8, 2]  
// 2	    0	Decrement	Decreases to 1	    [1, 8, 2]
// 3	    1	Decrement	Decreases to 7	    [1, 7, 2]
// 4	    2	Increment	Increases to 3	    [1, 7, 3]
// Thus, after 4 operations nums1 is transformed into nums2.

// Example 2:
// Input: nums1 = [1,3,6], nums2 = [2,4,5,3]
// Output: 4
// Explanation:
// Step	i	Operation	nums1[i]	        Updated nums1
// 1	    1	Append	        -	            [1, 3, 6, 3]
// 2	    0	Increment	Increases to 2	    [2, 3, 6, 3]
// 3	    1	Increment	Increases to 4	    [2, 4, 6, 3]
// 4	    2	Decrement	Decreases to 5	    [2, 4, 5, 3]
// Thus, after 4 operations nums1 is transformed into nums2.

// Example 3:
// Input: nums1 = [2], nums2 = [3,4]
// Output: 3
// Explanation:
// Step	i	Operation	nums1[i]	        Updated nums1
// 1	    0	Increment	Increases to 3	    [3]
// 2	    0	Append	        -	            [3, 3]
// 3	    1	Increment	Increases to 4	    [3, 4]  
// Thus, after 3 operations nums1 is transformed into nums2.

// Constraints:
//     1 <= n == nums1.length <= 10^5
//     nums2.length == n + 1
//     1 <= nums1[i], nums2[i] <= 10^5

import "fmt"

func minOperations(nums1 []int, nums2 []int) int64 {
    res, last, operCountForLast := 1, nums2[len(nums2)-1], 1 << 31
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := range nums1 {
        res += (abs(nums1[i] - nums2[i]))
        if operCountForLast != 0 { // isn't necessary, just to avoid extra operations
            from, to := min(nums1[i], nums2[i]), max(nums1[i], nums2[i])
            if from <= last && last <= to {
                operCountForLast = 0
            } else {
                operCountForLast = min( operCountForLast, min(abs(from-last), abs(to-last))) 
            }
        }
    }
    return int64(res + operCountForLast)
}

func minOperations1(nums1 []int, nums2 []int) int64 {
    res, extra, val := 0, 1 << 31, nums2[len(nums1)]
    for i := 0; i < len(nums1); i++ {
        a, b := min(nums1[i], nums2[i]), max(nums1[i], nums2[i])
        res += b - a
        if val >= a && val <= b {
            extra = 0
        } else if val < a {
            extra = min(extra, a - val)
        } else {
            extra = min(extra, val - b)
        }
    }
    return int64(res + extra + 1)
}

func main() {
    // Example 1:
    // Input: nums1 = [2,8], nums2 = [1,7,3]
    // Output: 4
    // Explanation:
    // Step	i	Operation	nums1[i]	    Updated  nums1
    // 1	    0	Append	        -	            [2, 8, 2]  
    // 2	    0	Decrement	Decreases to 1	    [1, 8, 2]
    // 3	    1	Decrement	Decreases to 7	    [1, 7, 2]
    // 4	    2	Increment	Increases to 3	    [1, 7, 3]
    // Thus, after 4 operations nums1 is transformed into nums2.
    fmt.Println(minOperations([]int{2,8}, []int{1,7,3})) // 4
    // Example 2:
    // Input: nums1 = [1,3,6], nums2 = [2,4,5,3]
    // Output: 4
    // Explanation:
    // Step	i	Operation	nums1[i]	        Updated nums1
    // 1	    1	Append	        -	            [1, 3, 6, 3]
    // 2	    0	Increment	Increases to 2	    [2, 3, 6, 3]
    // 3	    1	Increment	Increases to 4	    [2, 4, 6, 3]
    // 4	    2	Decrement	Decreases to 5	    [2, 4, 5, 3]
    // Thus, after 4 operations nums1 is transformed into nums2.
    fmt.Println(minOperations([]int{1,3,6}, []int{2,4,5,3})) // 4
    // Example 3:
    // Input: nums1 = [2], nums2 = [3,4]
    // Output: 3
    // Explanation:
    // Step	i	Operation	nums1[i]	        Updated nums1
    // 1	    0	Increment	Increases to 3	    [3]
    // 2	    0	Append	        -	            [3, 3]
    // 3	    1	Increment	Increases to 4	    [3, 4]  
    // Thus, after 3 operations nums1 is transformed into nums2.
    fmt.Println(minOperations([]int{2}, []int{3,4})) // 3 

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9}, []int{0,1,2,3,4,5,6,7,8,9})) // 2
    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1,0})) // 42
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1}, []int{0,1,2,3,4,5,6,7,8,9})) // 42
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1,0})) // 2

    fmt.Println(minOperations1([]int{2,8}, []int{1,7,3})) // 4
    fmt.Println(minOperations1([]int{1,3,6}, []int{2,4,5,3})) // 4
    fmt.Println(minOperations1([]int{2}, []int{3,4})) // 3 
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9}, []int{0,1,2,3,4,5,6,7,8,9})) // 2
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1,0})) // 42
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1}, []int{0,1,2,3,4,5,6,7,8,9})) // 42
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1,0})) // 2
}