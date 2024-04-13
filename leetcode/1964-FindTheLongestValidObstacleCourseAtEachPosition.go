package main

// 1964. Find the Longest Valid Obstacle Course at Each Position
// You want to build some obstacle courses. You are given a 0-indexed integer array obstacles of length n, 
// where obstacles[i] describes the height of the ith obstacle.
// For every index i between 0 and n - 1 (inclusive), find the length of the longest obstacle course in obstacles such that:
//     You choose any number of obstacles between 0 and i inclusive.
//     You must include the ith obstacle in the course.
//     You must put the chosen obstacles in the same order as they appear in obstacles.

// Every obstacle (except the first) is taller than or the same height as the obstacle immediately before it.
// Return an array ans of length n, where ans[i] is the length of the longest obstacle course for index i as described above.

// Example 1:
// Input: obstacles = [1,2,3,2]
// Output: [1,2,3,3]
// Explanation: The longest valid obstacle course at each position is:
// - i = 0: [1], [1] has length 1.
// - i = 1: [1,2], [1,2] has length 2.
// - i = 2: [1,2,3], [1,2,3] has length 3.
// - i = 3: [1,2,3,2], [1,2,2] has length 3.

// Example 2:
// Input: obstacles = [2,2,1]
// Output: [1,2,1]
// Explanation: The longest valid obstacle course at each position is:
// - i = 0: [2], [2] has length 1.
// - i = 1: [2,2], [2,2] has length 2.
// - i = 2: [2,2,1], [1] has length 1.

// Example 3:
// Input: obstacles = [3,1,5,6,4,2]
// Output: [1,1,2,3,2,2]
// Explanation: The longest valid obstacle course at each position is:
// - i = 0: [3], [3] has length 1.
// - i = 1: [3,1], [1] has length 1.
// - i = 2: [3,1,5], [3,5] has length 2. [1,5] is also valid.
// - i = 3: [3,1,5,6], [3,5,6] has length 3. [1,5,6] is also valid.
// - i = 4: [3,1,5,6,4], [3,4] has length 2. [1,4] is also valid.
// - i = 5: [3,1,5,6,4,2], [1,2] has length 2.
 
// Constraints:
//     n == obstacles.length
//     1 <= n <= 10^5
//     1 <= obstacles[i] <= 10^7

import "fmt"
import "sort"

// 二分
func longestObstacleCourseAtEachPosition(obstacles []int) []int {
    sub, res := []int{}, []int{}
    search := func (nums []int, target int) int {
        l, r := 0, len(nums)
        for l < r {
            m := l + (r-l)/2
            if nums[m] <= target {
                l = m + 1
            } else {
                r = m
            }
        }
        return l
    }
    for _, v := range obstacles {
        i := search(sub, v)
        if i == len(sub) {
            sub = append(sub, v)
        } else {
            sub[i] = v
        }
        res = append(res, i + 1)
    }
    return res
}

// dp
func longestObstacleCourseAtEachPosition1(obstacles []int) []int {
    if len(obstacles) == 0 {
        return nil
    }
    if len(obstacles) == 1 {
        return []int{1}
    }
    dp, res := []int{}, make([]int, len(obstacles))
    for i := 0; i < len(obstacles); i++ {
        p := sort.SearchInts(dp, obstacles[i] + 1)
        if p < len(dp) {
            dp[p] = obstacles[i]
        } else {
            dp = append(dp, obstacles[i])
        }
        res[i] = p + 1
    }
    return res
}

func main() {
    // Explanation: The longest valid obstacle course at each position is:
    // - i = 0: [1], [1] has length 1.
    // - i = 1: [1,2], [1,2] has length 2.
    // - i = 2: [1,2,3], [1,2,3] has length 3.
    // - i = 3: [1,2,3,2], [1,2,2] has length 3.
    fmt.Println(longestObstacleCourseAtEachPosition([]int{1,2,3,2})) // [1,2,3,3]
    // Explanation: The longest valid obstacle course at each position is:
    // - i = 0: [2], [2] has length 1.
    // - i = 1: [2,2], [2,2] has length 2.
    // - i = 2: [2,2,1], [1] has length 1.
    fmt.Println(longestObstacleCourseAtEachPosition([]int{2,2,1})) // [1,2,1]
    // Explanation: The longest valid obstacle course at each position is:
    // - i = 0: [3], [3] has length 1.
    // - i = 1: [3,1], [1] has length 1.
    // - i = 2: [3,1,5], [3,5] has length 2. [1,5] is also valid.
    // - i = 3: [3,1,5,6], [3,5,6] has length 3. [1,5,6] is also valid.
    // - i = 4: [3,1,5,6,4], [3,4] has length 2. [1,4] is also valid.
    // - i = 5: [3,1,5,6,4,2], [1,2] has length 2.
    fmt.Println(longestObstacleCourseAtEachPosition([]int{3,1,5,6,4,2})) // [1,1,2,3,2,2]

    fmt.Println(longestObstacleCourseAtEachPosition1([]int{1,2,3,2})) // [1,2,3,3]
    fmt.Println(longestObstacleCourseAtEachPosition1([]int{2,2,1})) // [1,2,1]
    fmt.Println(longestObstacleCourseAtEachPosition1([]int{3,1,5,6,4,2})) // [1,1,2,3,2,2]
}