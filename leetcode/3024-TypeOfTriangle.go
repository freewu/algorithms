package main

// 3024. Type of Triangle
// You are given a 0-indexed integer array nums of size 3 which can form the sides of a triangle.
//     1. A triangle is called equilateral if it has all sides of equal length.
//     2. A triangle is called isosceles if it has exactly two sides of equal length.
//     3. A triangle is called scalene if all its sides are of different lengths.

// Return a string representing the type of triangle that can be formed or "none" if it cannot form a triangle.

// Example 1:
// Input: nums = [3,3,3]
// Output: "equilateral"
// Explanation: Since all the sides are of equal length, therefore, it will form an equilateral triangle.

// Example 2:
// Input: nums = [3,4,5]
// Output: "scalene"
// Explanation: 
// nums[0] + nums[1] = 3 + 4 = 7, which is greater than nums[2] = 5.
// nums[0] + nums[2] = 3 + 5 = 8, which is greater than nums[1] = 4.
// nums[1] + nums[2] = 4 + 5 = 9, which is greater than nums[0] = 3. 
// Since the sum of the two sides is greater than the third side for all three cases, therefore, it can form a triangle.
// As all the sides are of different lengths, it will form a scalene triangle.

// Constraints:
//     nums.length == 3
//     1 <= nums[i] <= 100

import "fmt"
import "sort"

func triangleType(nums []int) string {
    sort.Ints(nums)
    if nums[0] + nums[1] <= nums[2] { return "none"  }
    if nums[0] == nums[1] && nums[1] == nums[2] { return "equilateral" }
    if nums[0] == nums[1] || nums[1] == nums[2] { return "isosceles"   }
    return "scalene"
}

func triangleType1(nums []int) string {
    a, b, c := nums[0], nums[1], nums[2]
    if a + b <= c || b + c <= a || a + c <= b { return "none" }
    if a == b && b == c { return "equilateral" }
    if a == b || b == c || c == a { return "isosceles" }
    return "scalene"
}

func main() {
    // Example 1:
    // Input: nums = [3,3,3]
    // Output: "equilateral"
    // Explanation: Since all the sides are of equal length, therefore, it will form an equilateral triangle.
    fmt.Println(triangleType([]int{3,3,3})) // "equilateral"
    // Example 2:
    // Input: nums = [3,4,5]
    // Output: "scalene"
    // Explanation: 
    // nums[0] + nums[1] = 3 + 4 = 7, which is greater than nums[2] = 5.
    // nums[0] + nums[2] = 3 + 5 = 8, which is greater than nums[1] = 4.
    // nums[1] + nums[2] = 4 + 5 = 9, which is greater than nums[0] = 3. 
    // Since the sum of the two sides is greater than the third side for all three cases, therefore, it can form a triangle.
    // As all the sides are of different lengths, it will form a scalene triangle.
    fmt.Println(triangleType([]int{3,4,5})) // "scalene"

    fmt.Println(triangleType([]int{1,2,3})) // "none"
    fmt.Println(triangleType([]int{1,100,100})) // "isosceles"
    fmt.Println(triangleType([]int{1,2,100})) // "none"

    fmt.Println(triangleType1([]int{3,3,3})) // "equilateral"
    fmt.Println(triangleType1([]int{3,4,5})) // "scalene"
    fmt.Println(triangleType1([]int{1,2,3})) // "none"
    fmt.Println(triangleType1([]int{1,100,100})) // "isosceles"
    fmt.Println(triangleType1([]int{1,2,100})) // "none"
}