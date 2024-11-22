package main

// 1968. Array With Elements Not Equal to Average of Neighbors
// You are given a 0-indexed array nums of distinct integers. 
// You want to rearrange the elements in the array such that every element in the rearranged array is not equal to the average of its neighbors.

// More formally, the rearranged array should have the property 
// such that for every i in the range 1 <= i < nums.length - 1, (nums[i-1] + nums[i+1]) / 2 is not equal to nums[i].

// Return any rearrangement of nums that meets the requirements.

// Example 1:
// Input: nums = [1,2,3,4,5]
// Output: [1,2,4,5,3]
// Explanation:
// When i=1, nums[i] = 2, and the average of its neighbors is (1+4) / 2 = 2.5.
// When i=2, nums[i] = 4, and the average of its neighbors is (2+5) / 2 = 3.5.
// When i=3, nums[i] = 5, and the average of its neighbors is (4+3) / 2 = 3.5.

// Example 2:
// Input: nums = [6,2,0,9,7]
// Output: [9,7,6,2,0]
// Explanation:
// When i=1, nums[i] = 7, and the average of its neighbors is (9+6) / 2 = 7.5.
// When i=2, nums[i] = 6, and the average of its neighbors is (7+2) / 2 = 4.5.
// When i=3, nums[i] = 2, and the average of its neighbors is (6+0) / 2 = 3.

// Constraints:
//     3 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^5

import "fmt"
import "sort"

func rearrangeArray(nums []int) []int {
    sort.Ints(nums)
    for i := 0; i < len(nums) - 1; i += 2 {
        nums[i], nums[i+1] = nums[i+1], nums[i]
    }
    return nums
}

func rearrangeArray1(nums []int) []int {
    flag := true
    for flag {
        flag = false
        for i := 1; i < len(nums) - 1; i++ {
            if (nums[i - 1] + nums[i + 1]) / 2 == nums[i] {
                nums[i - 1], nums[i] = nums[i], nums[i - 1]
                flag = true
            }
        }
    }
    return nums
}

func rearrangeArray2(nums []int) []int {
    for i, toggle := 1, false; i < len(nums)-1; i, toggle = i + 1, !toggle {
        if toggle {
            if nums[i] > nums[i-1] { nums[i-1], nums[i] = nums[i], nums[i-1] }
            if nums[i] > nums[i+1] { nums[i+1], nums[i] = nums[i], nums[i+1] }
        } else {
            if nums[i] < nums[i-1] { nums[i-1], nums[i] = nums[i], nums[i-1] }
            if nums[i] < nums[i+1] { nums[i+1], nums[i] = nums[i], nums[i+1] }
        }
    }
    return nums
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5]
    // Output: [1,2,4,5,3]
    // Explanation:
    // When i=1, nums[i] = 2, and the average of its neighbors is (1+4) / 2 = 2.5.
    // When i=2, nums[i] = 4, and the average of its neighbors is (2+5) / 2 = 3.5.
    // When i=3, nums[i] = 5, and the average of its neighbors is (4+3) / 2 = 3.5.
    fmt.Println(rearrangeArray([]int{1,2,3,4,5})) // [1,2,4,5,3]
    // Example 2:
    // Input: nums = [6,2,0,9,7]
    // Output: [9,7,6,2,0]
    // Explanation:
    // When i=1, nums[i] = 7, and the average of its neighbors is (9+6) / 2 = 7.5.
    // When i=2, nums[i] = 6, and the average of its neighbors is (7+2) / 2 = 4.5.
    // When i=3, nums[i] = 2, and the average of its neighbors is (6+0) / 2 = 3.
    fmt.Println(rearrangeArray([]int{6,2,0,9,7})) // [9,7,6,2,0]

    fmt.Println(rearrangeArray1([]int{1,2,3,4,5})) // [1,2,4,5,3]
    fmt.Println(rearrangeArray1([]int{6,2,0,9,7})) // [6 2 0 9 7]

    fmt.Println(rearrangeArray2([]int{1,2,3,4,5})) // [1 3 2 5 4]
    fmt.Println(rearrangeArray2([]int{6,2,0,9,7})) // [2 6 0 9 7]
}