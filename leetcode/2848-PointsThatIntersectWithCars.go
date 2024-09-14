package main

// 2848. Points That Intersect With Cars
// You are given a 0-indexed 2D integer array nums representing the coordinates of the cars parking on a number line. 
// For any index i, nums[i] = [starti, endi] where starti is the starting point of the ith car and endi is the ending point of the ith car.

// Return the number of integer points on the line that are covered with any part of a car.

// Example 1:
// Input: nums = [[3,6],[1,5],[4,7]]
// Output: 7
// Explanation: All the points from 1 to 7 intersect at least one car, therefore the answer would be 7.

// Example 2:
// Input: nums = [[1,3],[5,8]]
// Output: 7
// Explanation: Points intersecting at least one car are 1, 2, 3, 5, 6, 7, 8. There are a total of 7 points, therefore the answer would be 7.

// Constraints:
//     1 <= nums.length <= 100
//     nums[i].length == 2
//     1 <= starti <= endi <= 100

import "fmt"

func numberOfPoints(nums [][]int) int {
    mp := [102]int{}
    for _, v := range nums {
        mp[v[0]]++ // 起点
        mp[v[1] + 1]-- // 到了终点
    }
    res, cars := 0, 0
    for _, c := range mp {
        cars += c // 累加车的数量(到终点的车会减去)
        if cars > 0 { // 和其他车有相交
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [[3,6],[1,5],[4,7]]
    // Output: 7
    // Explanation: All the points from 1 to 7 intersect at least one car, therefore the answer would be 7.
    fmt.Println(numberOfPoints([][]int{{3,6},{1,5},{4,7}})) // 7
    // Example 2:
    // Input: nums = [[1,3],[5,8]]
    // Output: 7
    // Explanation: Points intersecting at least one car are 1, 2, 3, 5, 6, 7, 8. There are a total of 7 points, therefore the answer would be 7.
    fmt.Println(numberOfPoints([][]int{{1,3},{5,8}})) // 7
}