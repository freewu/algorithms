package main

// 976. Largest Perimeter Triangle
// Given an integer array nums, return the largest perimeter of a triangle with a non-zero area, formed from three of these lengths. 
// If it is impossible to form any triangle of a non-zero area, return 0.

// Example 1:
// Input: nums = [2,1,2]
// Output: 5
// Explanation: You can form a triangle with three side lengths: 1, 2, and 2.

// Example 2:
// Input: nums = [1,2,1,10]
// Output: 0
// Explanation: 
// You cannot use the side lengths 1, 1, and 2 to form a triangle.
// You cannot use the side lengths 1, 1, and 10 to form a triangle.
// You cannot use the side lengths 1, 2, and 10 to form a triangle.
// As we cannot use any three side lengths to form a triangle of non-zero area, we return 0.
 
// Constraints:
//     3 <= nums.length <= 10^4
//     1 <= nums[i] <= 10^6

import "fmt"
import "sort"

func largestPerimeter(nums []int) int {
    sort.Ints(nums)
    abs := func (i int) int { if i >= 0 { return i; }; return i * -1; }
    for i := len(nums) -1; i >= 2; i-- {
        // 判断是否能组成三角形
        // 两边之和大于第三边，两边之差小于第三边
        if nums[i] + nums[i-1] > nums[i- 2] && abs(nums[i] - nums[i-1]) < nums[i- 2] {
            // 因为排序后 从大到小判断，所有能形成三角形一定是最大边长的
            return nums[i] + nums[i -1] + nums[i -2]
        }
    }
    return 0
}

func largestPerimeter1(nums []int) int {
    sort.Ints(nums)
    a := 0
    for i := len(nums) - 1; i > 1; i-- {
        a = nums[i-1] + nums[i-2]
        // 两边之和大于第三边
        if nums[i] < a {
            return nums[i] + a
        }
    }
    return 0
}

func largestPerimeter2(nums []int) int {
    res := 0
    if len(nums) < 3 { return 0 }
    sort.Ints(nums)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for len(nums) >= 3 {
        if nums[len(nums) - 1] >= nums[len(nums) - 2] + nums[len(nums) - 3] {
            nums = nums[:len(nums) - 1]
            continue
        }
        res = max(res, nums[len(nums) - 1] + nums[len(nums) - 2] + nums[len(nums) - 3])
        nums = nums[:len(nums) - 1]

    }
    return res
}

func main() {
    // Input: nums = [2,1,2]
    // Output: 5
    // Explanation: You can form a triangle with three side lengths: 1, 2, and 2.
    fmt.Println(largestPerimeter([]int{2,1,2})) // 5
    // You cannot use the side lengths 1, 1, and 2 to form a triangle.
    // You cannot use the side lengths 1, 1, and 10 to form a triangle.
    // You cannot use the side lengths 1, 2, and 10 to form a triangle.
    // As we cannot use any three side lengths to form a triangle of non-zero area, we return 0.
    fmt.Println(largestPerimeter([]int{1,2,1,10})) // 0

    fmt.Println(largestPerimeter([]int{1,2,3,4,5,6,7,8,9})) // 24
    fmt.Println(largestPerimeter([]int{9,8,7,6,5,4,3,2,1})) // 24

    fmt.Println(largestPerimeter1([]int{2,1,2})) // 5
    fmt.Println(largestPerimeter1([]int{1,2,1,10})) // 0
    fmt.Println(largestPerimeter1([]int{1,2,3,4,5,6,7,8,9})) // 24
    fmt.Println(largestPerimeter1([]int{9,8,7,6,5,4,3,2,1})) // 24

    fmt.Println(largestPerimeter2([]int{2,1,2})) // 5
    fmt.Println(largestPerimeter2([]int{1,2,1,10})) // 0
    fmt.Println(largestPerimeter2([]int{1,2,3,4,5,6,7,8,9})) // 24
    fmt.Println(largestPerimeter2([]int{9,8,7,6,5,4,3,2,1})) // 24
}