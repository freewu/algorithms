package main

// 908. Smallest Range I
// You are given an integer array nums and an integer k.

// In one operation, you can choose any index i where 0 <= i < nums.length and change nums[i] to nums[i] + x where x is an integer from the range [-k, k]. 
// You can apply this operation at most once for each index i.

// The score of nums is the difference between the maximum and minimum elements in nums.

// Return the minimum score of nums after applying the mentioned operation at most once for each index in it.

// Example 1:
// Input: nums = [1], k = 0
// Output: 0
// Explanation: The score is max(nums) - min(nums) = 1 - 1 = 0.

// Example 2:
// Input: nums = [0,10], k = 2
// Output: 6
// Explanation: Change nums to be [2, 8]. The score is max(nums) - min(nums) = 8 - 2 = 6.

// Example 3:
// Input: nums = [1,3,6], k = 3
// Output: 0
// Explanation: Change nums to be [4, 4, 4]. The score is max(nums) - min(nums) = 4 - 4 = 0.

// Constraints:
//     1 <= nums.length <= 10^4
//     0 <= nums[i] <= 10^4
//     0 <= k <= 10^4

import "fmt"

func smallestRangeI(nums []int, k int) int {
    getMinAndMax := func (arr []int) (int, int) { // 获取最大值 & 最小值 
        mn, mx := arr[0], arr[0]
        for i := 1; i < len(arr); i++ {
            if arr[i] < mn {
                mn = arr[i]
            } else if arr[i] > mx {
                mx = arr[i]
            }
        }
        return mn, mx
    }
    mn, mx := getMinAndMax(nums)
    res := mx - mn - 2 * k
    if res < 0 {
        return 0
    }
    return res
}

func smallestRangeI1(nums []int, k int) int {
    mn, mx := nums[0], nums[0]
    for _, v := range nums[1:] {
        if v < mn {
            mn = v
        } else if v > mx {
            mx = v
        }
    }
    res := mx - mn - 2*k
    if res > 0 {
        return res
    }
    return 0
}

func main() {
    // Example 1:
    // Input: nums = [1], k = 0
    // Output: 0
    // Explanation: The score is max(nums) - min(nums) = 1 - 1 = 0.
    fmt.Println(smallestRangeI([]int{1}, 0)) // 0
    // Example 2:
    // Input: nums = [0,10], k = 2
    // Output: 6
    // Explanation: Change nums to be [2, 8]. The score is max(nums) - min(nums) = 8 - 2 = 6.
    fmt.Println(smallestRangeI([]int{0,10}, 2)) // 6
    // Example 3:
    // Input: nums = [1,3,6], k = 3
    // Output: 0
    // Explanation: Change nums to be [4, 4, 4]. The score is max(nums) - min(nums) = 4 - 4 = 0.
    fmt.Println(smallestRangeI([]int{1,3,6}, 3)) // 0

    fmt.Println(smallestRangeI1([]int{1}, 0)) // 0
    fmt.Println(smallestRangeI1([]int{0,10}, 2)) // 6
    fmt.Println(smallestRangeI1([]int{1,3,6}, 3)) // 0
}