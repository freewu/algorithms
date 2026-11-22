package main

// 4120. Count Dominant Elements
// You are given an integer array nums.

// An element in nums is called dominant if it is greater than or equal to the average of its immediate neighbors.

// The immediate neighbors of nums[i] are nums[i - 1] and nums[i + 1], when those indices exist.

// A single element with no neighbors is considered dominant.

// Return the number of dominant elements in nums.

// Example 1:
// Input: nums = [1,10,2]
// Output: 1
// Explanation:
//     For nums[0] = 1, its only neighbor is 10. Since 1 < 10, it is not dominant.
//     For nums[1] = 10, its neighbors are 1 and 2. Their average is (1 + 2) / 2 = 1.5. Since 10 > 1.5, it is dominant.
//     For nums[2] = 2, its only neighbor is 10. Since 2 < 10, it is not dominant.
// Therefore, the answer is 1.

// Example 2:
// Input: nums = [6,4,2]
// Output: 2
// Explanation:
//     For nums[0] = 6, its only neighbor is 4. Since 6 > 4, it is dominant.
//     For nums[1] = 4, its neighbors are 6 and 2. Their average is (6 + 2) / 2 = 4. Since 4 = 4, it is dominant.
//     For nums[2] = 2, its only neighbor is 4. Since 2 < 4, it is not dominant.
// Therefore, the answer is 2.

// Example 3:
// Input: nums = [7]
// Output: 1
// Explanation:
//     For nums[0] = 7, it has no neighbors, so it is dominant.
// Therefore, the answer is 1.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"

func countDominantElements(nums []int) int {
    res, n := 0, len(nums)
    for i, v := range nums {
       sum, count := 0, 0
        if i > 0 {
            sum += nums[i - 1]
            count++
        }
        if i+1 < n {
            sum += nums[i + 1]
            count++
        }
        if count * v >= sum {
            res++
        }
    }
    return res
}

func countDominantElements1(nums []int) int {
    res := 0
    for i := range nums {
        sum, dom := 0.0, 0.0
        if i > 0 {
            sum += float64(nums[i - 1])
            dom += 1.0
        }
        if i < len(nums) - 1 {
            sum += float64(nums[i + 1])
            dom += 1.0
        }
        if dom == 0.0 || sum / dom <= float64(nums[i]) {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,10,2]
    // Output: 1
    // Explanation:
    //     For nums[0] = 1, its only neighbor is 10. Since 1 < 10, it is not dominant.
    //     For nums[1] = 10, its neighbors are 1 and 2. Their average is (1 + 2) / 2 = 1.5. Since 10 > 1.5, it is dominant.
    //     For nums[2] = 2, its only neighbor is 10. Since 2 < 10, it is not dominant.
    // Therefore, the answer is 1.
    fmt.Println(countDominantElements([]int{1,10,2})) // 1
    // Example 2:
    // Input: nums = [6,4,2]
    // Output: 2
    // Explanation:
    //     For nums[0] = 6, its only neighbor is 4. Since 6 > 4, it is dominant.
    //     For nums[1] = 4, its neighbors are 6 and 2. Their average is (6 + 2) / 2 = 4. Since 4 = 4, it is dominant.
    //     For nums[2] = 2, its only neighbor is 4. Since 2 < 4, it is not dominant.
    // Therefore, the answer is 2.
    fmt.Println(countDominantElements([]int{6,4,2})) // 2
    // Example 3:
    // Input: nums = [7]
    // Output: 1
    // Explanation:
    //     For nums[0] = 7, it has no neighbors, so it is dominant.
    // Therefore, the answer is 1.
    fmt.Println(countDominantElements([]int{7})) // 1

    fmt.Println(countDominantElements([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(countDominantElements([]int{9,8,7,6,5,4,3,2,1})) // 8

    fmt.Println(countDominantElements1([]int{1,10,2})) // 1
    fmt.Println(countDominantElements1([]int{6,4,2})) // 2
    fmt.Println(countDominantElements1([]int{7})) // 1
    fmt.Println(countDominantElements1([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(countDominantElements1([]int{9,8,7,6,5,4,3,2,1})) // 8
}