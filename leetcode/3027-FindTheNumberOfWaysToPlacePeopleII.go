package main

// 3027. Find the Number of Ways to Place People II
// You are given a 2D array points of size n x 2 representing integer coordinates of some points on a 2D-plane, 
// where points[i] = [xi, yi].

// We define the right direction as positive x-axis (increasing x-coordinate) 
// and the left direction as negative x-axis (decreasing x-coordinate). 
// Similarly, we define the up direction as positive y-axis (increasing y-coordinate) 
// and the down direction as negative y-axis (decreasing y-coordinate)

// You have to place n people, including Alice and Bob, at these points such that there is exactly one person at every point. 
// Alice wants to be alone with Bob, so Alice will build a rectangular fence with Alice's position as the upper left corner 
// and Bob's position as the lower right corner of the fence (Note that the fence might not enclose any area, i.e. it can be a line). 
// If any person other than Alice and Bob is either inside the fence or on the fence, Alice will be sad.

// Return the number of pairs of points where you can place Alice and Bob, 
// such that Alice does not become sad on building the fence.

// Note that Alice can only build a fence with Alice's position as the upper left corner, and Bob's position as the lower right corner. 
// For example, Alice cannot build either of the fences in the picture below with four corners (1, 1), (1, 3), (3, 1), and (3, 3), because:
//     With Alice at (3, 3) and Bob at (1, 1), Alice's position is not the upper left corner and Bob's position is not the lower right corner of the fence.
//     With Alice at (1, 3) and Bob at (1, 1), Bob's position is not the lower right corner of the fence.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/01/04/example1alicebob.png" />
// Input: points = [[1,1],[2,2],[3,3]]
// Output: 0
// Explanation: There is no way to place Alice and Bob such that Alice can build a fence with Alice's position as the upper left corner and Bob's position as the lower right corner. Hence we return 0. 

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2024/02/04/example2alicebob.png" />
// Input: points = [[6,2],[4,4],[2,6]]
// Output: 2
// Explanation: There are two ways to place Alice and Bob such that Alice will not be sad:
// - Place Alice at (4, 4) and Bob at (6, 2).
// - Place Alice at (2, 6) and Bob at (4, 4).
// You cannot place Alice at (2, 6) and Bob at (6, 2) because the person at (4, 4) will be inside the fence.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2024/02/04/example4alicebob.png" />
// Input: points = [[3,1],[1,3],[1,1]]
// Output: 2
// Explanation: There are two ways to place Alice and Bob such that Alice will not be sad:
// - Place Alice at (1, 1) and Bob at (3, 1).
// - Place Alice at (1, 3) and Bob at (1, 1).
// You cannot place Alice at (1, 3) and Bob at (3, 1) because the person at (1, 1) will be on the fence.
// Note that it does not matter if the fence encloses any area, the first and second fences in the image are valid.

// Constraints:
//     2 <= n <= 1000
//     points[i].length == 2
//     -109 <= points[i][0], points[i][1] <= 10^9
//     All points[i] are distinct.

import "fmt"
import "sort"

func numberOfPairs(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        if points[i][1] > points[j][1]  { return true }
        if points[i][1] == points[j][1] { return points[i][0] < points[j][0]}
        return false
    })
    res := 0
    for i := 0; i < len(points); i++ {
        mn, mx := points[i][0], 1 << 31
        for j := i + 1; j < len(points); j++ {
            if points[j][0] >= mn && points[j][0] < mx {
                res++
                mx = points[j][0]
            }
        }
    }
    return res
}

func numberOfPairs1(points [][]int) int {
    res, n := 0, len(points)
    sort.Slice(points, func(i, j int) bool {
        if points[i][0] == points[j][0] { return points[i][1] < points[j][1] }
        return points[i][0] > points[j][0]
    })
    for i := 1; i < n; i++ {
        y1, mx := points[i][1], -1 << 31
        for j := i - 1; j >= 0; j-- {
            y2 := points[j][1]
            if y2 > y1 { continue }
            if y2 > mx {
                res++
                mx = y2
                if mx == y1 { break }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/01/04/example1alicebob.png" />
    // Input: points = [[1,1],[2,2],[3,3]]
    // Output: 0
    // Explanation: There is no way to place Alice and Bob such that Alice can build a fence with Alice's position as the upper left corner and Bob's position as the lower right corner. Hence we return 0. 
    fmt.Println(numberOfPairs([][]int{{1,1},{2,2},{3,3}})) // 0
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2024/02/04/example2alicebob.png" />
    // Input: points = [[6,2],[4,4],[2,6]]
    // Output: 2
    // Explanation: There are two ways to place Alice and Bob such that Alice will not be sad:
    // - Place Alice at (4, 4) and Bob at (6, 2).
    // - Place Alice at (2, 6) and Bob at (4, 4).
    // You cannot place Alice at (2, 6) and Bob at (6, 2) because the person at (4, 4) will be inside the fence.
    fmt.Println(numberOfPairs([][]int{{6,2},{4,4},{2,6}})) // 2
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2024/02/04/example4alicebob.png" />
    // Input: points = [[3,1],[1,3],[1,1]]
    // Output: 2
    // Explanation: There are two ways to place Alice and Bob such that Alice will not be sad:
    // - Place Alice at (1, 1) and Bob at (3, 1).
    // - Place Alice at (1, 3) and Bob at (1, 1).
    // You cannot place Alice at (1, 3) and Bob at (3, 1) because the person at (1, 1) will be on the fence.
    // Note that it does not matter if the fence encloses any area, the first and second fences in the image are valid.
    fmt.Println(numberOfPairs([][]int{{3,1},{1,3},{1,1}})) // 2

    fmt.Println(numberOfPairs1([][]int{{1,1},{2,2},{3,3}})) // 0
    fmt.Println(numberOfPairs1([][]int{{6,2},{4,4},{2,6}})) // 2
    fmt.Println(numberOfPairs1([][]int{{3,1},{1,3},{1,1}})) // 2
}