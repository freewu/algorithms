package main

// 3698. Split Array With Minimum Difference
// You are given an integer array nums.

// Split the array into exactly two subarrays, left and right, such that left is strictly increasing and right is strictly decreasing.

// Return the minimum possible absolute difference between the sums of left and right. If no valid split exists, return -1.

// A subarray is a contiguous non-empty sequence of elements within an array.

// An array is said to be strictly increasing if each element is strictly greater than its previous one (if exists).

// An array is said to be strictly decreasing if each element is strictly smaller than its previous one (if exists).

// Example 1:
// Input: nums = [1,3,2]
// Output: 2
// Explanation:
// i	left	right	Validity	left sum	right sum	Absolute difference
// 0	[1]	[3, 2]	Yes	1	5	|1 - 5| = 4
// 1	[1, 3]	[2]	Yes	4	2	|4 - 2| = 2
// Thus, the minimum absolute difference is 2.

// Example 2:
// Input: nums = [1,2,4,3]
// Output: 4
// Explanation:
// i	left	right	Validity	left sum	right sum	Absolute difference
// 0	[1]	[2, 4, 3]	No	1	9	-
// 1	[1, 2]	[4, 3]	Yes	3	7	|3 - 7| = 4
// 2	[1, 2, 4]	[3]	Yes	7	3	|7 - 3| = 4
// Thus, the minimum absolute difference is 4.

// Example 3:
// Input: nums = [3,1,2]
// Output: -1
// Explanation:
// No valid split exists, so the answer is -1.

// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func splitArray(nums []int) int64 {
    n, prefix, i := len(nums), nums[0], 1
    for i < n && nums[i] > nums[i-1] {
        prefix += nums[i]
        i++
    }
    // 最长严格递减后缀
    suffix := nums[n - 1]
    j := n - 2
    for j >= 0 && nums[j] > nums[j+1] {
        suffix += nums[j]
        j--
    }
    if i - 1 < j { return -1 } // 情况一 [1,2,3,1,3,2]，无解
    diff := prefix - suffix
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    if i - 1 == j { return int64(abs(diff)) } // 情况二  1,2,3,3,2]，有唯一分割方案 [1,2,3]+[3,2]。 如果 i−1=j，返回 [0,i−1] 的元素和与 [i,n−1] 的元素和的绝对差。 
    return int64(min(abs(diff + nums[i - 1]), abs(diff - nums[i - 1]))) // 情况三，suffix 多算了一个 nums[i-1]，或者 prefix 多算了一个 nums[i - 1]
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2]
    // Output: 2
    // Explanation:
    // i	left	right	Validity	left sum	right sum	Absolute difference
    // 0	[1]	[3, 2]	Yes	1	5	|1 - 5| = 4
    // 1	[1, 3]	[2]	Yes	4	2	|4 - 2| = 2
    // Thus, the minimum absolute difference is 2.
    fmt.Println(splitArray([]int{1,3,2})) // 2
    // Example 2:
    // Input: nums = [1,2,4,3]
    // Output: 4
    // Explanation:
    // i	left	right	Validity	left sum	right sum	Absolute difference
    // 0	[1]	[2, 4, 3]	No	1	9	-
    // 1	[1, 2]	[4, 3]	Yes	3	7	|3 - 7| = 4
    // 2	[1, 2, 4]	[3]	Yes	7	3	|7 - 3| = 4
    // Thus, the minimum absolute difference is 4.
    fmt.Println(splitArray([]int{1,2,4,3})) // 4
    // Example 3:
    // Input: nums = [3,1,2]
    // Output: -1
    // Explanation:
    // No valid split exists, so the answer is -1.
    fmt.Println(splitArray([]int{3,1,2})) // -1

    fmt.Println(splitArray([]int{1,2,3,4,5,6,7,8,9})) // 27
    fmt.Println(splitArray([]int{9,8,7,6,5,4,3,2,1})) // 27
}