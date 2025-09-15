package main

// 3678. Smallest Absent Positive Greater Than Average
// You are given an integer array nums.

// Return the smallest absent positive integer in nums such that it is strictly greater than the average of all elements in nums.

// The average of an array is defined as the sum of all its elements divided by the number of elements.

// Example 1:
// Input: nums = [3,5]
// Output: 6
// Explanation:
// The average of nums is (3 + 5) / 2 = 8 / 2 = 4.
// The smallest absent positive integer greater than 4 is 6.

// Example 2:
// Input: nums = [-1,1,2]
// Output: 3
// Explanation:
// ​​​​​​​The average of nums is (-1 + 1 + 2) / 3 = 2 / 3 = 0.667.
// The smallest absent positive integer greater than 0.667 is 3.

// Example 3:
// Input: nums = [4,-1]
// Output: 2
// Explanation:
// The average of nums is (4 + (-1)) / 2 = 3 / 2 = 1.50.
// The smallest absent positive integer greater than 1.50 is 2.

// Constraints:
//     1 <= nums.length <= 100
//     -100 <= nums[i] <= 100​​​​​​​

import "fmt"

func smallestAbsent(nums []int) int {
    sum, res := 0, 1
    for _, v := range nums {
        sum += v
    }
    avg := float64(sum) / float64(len(nums))
    for {
        if float64(res) > avg {
            found := false
            for _, v := range nums {
                if v == res {
                    found = true
                    break
                }
            }
            if !found {
                return res
            }
        }
        res++
    }
}

func main() {
    // Example 1:
    // Input: nums = [3,5]
    // Output: 6
    // Explanation:
    // The average of nums is (3 + 5) / 2 = 8 / 2 = 4.
    // The smallest absent positive integer greater than 4 is 6.
    fmt.Println(smallestAbsent([]int{3,5})) // 6
    // Example 2:
    // Input: nums = [-1,1,2]
    // Output: 3
    // Explanation:
    // ​​​​​​​The average of nums is (-1 + 1 + 2) / 3 = 2 / 3 = 0.667.
    // The smallest absent positive integer greater than 0.667 is 3.
    fmt.Println(smallestAbsent([]int{-1,1,2})) // 3
    // Example 3:
    // Input: nums = [4,-1]
    // Output: 2
    // Explanation:
    // The average of nums is (4 + (-1)) / 2 = 3 / 2 = 1.50.
    // The smallest absent positive integer greater than 1.50 is 2.
    fmt.Println(smallestAbsent([]int{4,-1})) // 2

    fmt.Println(smallestAbsent([]int{1,2,3,4,5,6,7,8,9})) // 10
    fmt.Println(smallestAbsent([]int{9,8,7,6,5,4,3,2,1})) // 10
}