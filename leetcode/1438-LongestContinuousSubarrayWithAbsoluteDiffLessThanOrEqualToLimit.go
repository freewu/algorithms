package main

// 1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
// Given an array of integers nums and an integer limit, 
// return the size of the longest non-empty subarray such 
// that the absolute difference between any two elements of this subarray is less than or equal to limit.

// Example 1:
// Input: nums = [8,2,4,7], limit = 4
// Output: 2 
// Explanation: All subarrays are: 
// [8] with maximum absolute diff |8-8| = 0 <= 4.
// [8,2] with maximum absolute diff |8-2| = 6 > 4. 
// [8,2,4] with maximum absolute diff |8-2| = 6 > 4.
// [8,2,4,7] with maximum absolute diff |8-2| = 6 > 4.
// [2] with maximum absolute diff |2-2| = 0 <= 4.
// [2,4] with maximum absolute diff |2-4| = 2 <= 4.
// [2,4,7] with maximum absolute diff |2-7| = 5 > 4.
// [4] with maximum absolute diff |4-4| = 0 <= 4.
// [4,7] with maximum absolute diff |4-7| = 3 <= 4.
// [7] with maximum absolute diff |7-7| = 0 <= 4. 
// Therefore, the size of the longest subarray is 2.

// Example 2:
// Input: nums = [10,1,2,4,7,2], limit = 5
// Output: 4 
// Explanation: The subarray [2,4,7,2] is the longest since the maximum absolute diff is |2-7| = 5 <= 5.

// Example 3:
// Input: nums = [4,2,2,2,4,4,2,2], limit = 0
// Output: 3
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     0 <= limit <= 10^9

import "fmt"

func longestSubarray(nums []int, limit int) int {
    n, maxd, mind := 0, []int{}, []int{}
    for i := 0; i < len(nums); i++ {
        for len(maxd) != 0 && nums[i] > maxd[len(maxd) - 1] { 
            maxd = maxd[:len(maxd) - 1] 
        }
        for len(mind) != 0 && nums[i] < mind[len(mind) - 1] { 
            mind = mind[:len(mind) - 1] 
        }
        maxd = append(maxd, nums[i])
        mind = append(mind, nums[i])
        if maxd[0] - mind[0] > limit {
            if maxd[0] == nums[n] { 
                maxd = maxd[1:] 
            }
            if mind[0] == nums[n] { 
                mind = mind[1:] 
            }
            n++
        }
    }
    return len(nums) - n
}

func main() {
    // Example 1:
    // Input: nums = [8,2,4,7], limit = 4
    // Output: 2 
    // Explanation: All subarrays are: 
    // [8] with maximum absolute diff |8-8| = 0 <= 4.
    // [8,2] with maximum absolute diff |8-2| = 6 > 4. 
    // [8,2,4] with maximum absolute diff |8-2| = 6 > 4.
    // [8,2,4,7] with maximum absolute diff |8-2| = 6 > 4.
    // [2] with maximum absolute diff |2-2| = 0 <= 4.
    // [2,4] with maximum absolute diff |2-4| = 2 <= 4.
    // [2,4,7] with maximum absolute diff |2-7| = 5 > 4.
    // [4] with maximum absolute diff |4-4| = 0 <= 4.
    // [4,7] with maximum absolute diff |4-7| = 3 <= 4.
    // [7] with maximum absolute diff |7-7| = 0 <= 4. 
    // Therefore, the size of the longest subarray is 2.
    fmt.Println(longestSubarray([]int{8,2,4,7}, 4)) // 2
    // Example 2:
    // Input: nums = [10,1,2,4,7,2], limit = 5
    // Output: 4 
    // Explanation: The subarray [2,4,7,2] is the longest since the maximum absolute diff is |2-7| = 5 <= 5.
    fmt.Println(longestSubarray([]int{10,1,2,4,7,2}, 5)) // 4
    // Example 3:
    // Input: nums = [4,2,2,2,4,4,2,2], limit = 0
    // Output: 3
    fmt.Println(longestSubarray([]int{4,2,2,2,4,4,2,2}, 0)) // 3
}