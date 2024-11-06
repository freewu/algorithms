package main

// 2158. Amount of New Area Painted Each Day
// There is a long and thin painting that can be represented by a number line. 
// You are given a 0-indexed 2D integer array paint of length n, where paint[i] = [starti, endi]. 
// This means that on the ith day you need to paint the area between starti and endi.

// Painting the same area multiple times will create an uneven painting so you only want to paint each area of the painting at most once.

// Return an integer array worklog of length n, where worklog[i] is the amount of new area that you painted on the ith day.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/02/01/screenshot-2022-02-01-at-17-16-16-diagram-drawio-diagrams-net.png" />
// Input: paint = [[1,4],[4,7],[5,8]]
// Output: [3,3,1]
// Explanation:
// On day 0, paint everything between 1 and 4.
// The amount of new area painted on day 0 is 4 - 1 = 3.
// On day 1, paint everything between 4 and 7.
// The amount of new area painted on day 1 is 7 - 4 = 3.
// On day 2, paint everything between 7 and 8.
// Everything between 5 and 7 was already painted on day 1.
// The amount of new area painted on day 2 is 8 - 7 = 1. 

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/02/01/screenshot-2022-02-01-at-17-17-45-diagram-drawio-diagrams-net.png" />
// Input: paint = [[1,4],[5,8],[4,7]]
// Output: [3,3,1]
// Explanation:
// On day 0, paint everything between 1 and 4.
// The amount of new area painted on day 0 is 4 - 1 = 3.
// On day 1, paint everything between 5 and 8.
// The amount of new area painted on day 1 is 8 - 5 = 3.
// On day 2, paint everything between 4 and 5.
// Everything between 5 and 7 was already painted on day 1.
// The amount of new area painted on day 2 is 5 - 4 = 1. 

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2022/02/01/screenshot-2022-02-01-at-17-19-49-diagram-drawio-diagrams-net.png" />
// Input: paint = [[1,5],[2,4]]
// Output: [4,0]
// Explanation:
// On day 0, paint everything between 1 and 5.
// The amount of new area painted on day 0 is 5 - 1 = 4.
// On day 1, paint nothing because everything between 2 and 4 was already painted on day 0.
// The amount of new area painted on day 1 is 0.

// Constraints:
//     1 <= paint.length <= 10^5
//     paint[i].length == 2
//     0 <= starti < endi <= 5 * 10^4

import "fmt"

type SegmentTree struct {
    tree []int
}

func (s *SegmentTree) update(index, left, right, low, high int) {
    if left > high || right < low || s.tree[index] == right - left + 1 { return }
    if left >= low && right <= high {
        s.tree[index] = right - left + 1
        return
    }
    mid := left + (right - left) >> 1
    s.update(index << 1, left, mid, low, high)
    s.update(index << 1 | 1, mid + 1, right, low, high)
    s.tree[index] = s.tree[index << 1] + s.tree[index << 1 | 1]
}

// 线段树
func amountPainted(paint [][]int) []int {
    n := 50001
    s, res := &SegmentTree{make([]int, 4 * n)}, make([]int, len(paint))
    for i := range paint {
        cur := s.tree[1]
        s.update(1, 0, n - 1, paint[i][0], paint[i][1]-1)
        res[i] = s.tree[1] - cur
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/02/01/screenshot-2022-02-01-at-17-16-16-diagram-drawio-diagrams-net.png" />
    // Input: paint = [[1,4],[4,7],[5,8]]
    // Output: [3,3,1]
    // Explanation:
    // On day 0, paint everything between 1 and 4.
    // The amount of new area painted on day 0 is 4 - 1 = 3.
    // On day 1, paint everything between 4 and 7.
    // The amount of new area painted on day 1 is 7 - 4 = 3.
    // On day 2, paint everything between 7 and 8.
    // Everything between 5 and 7 was already painted on day 1.
    // The amount of new area painted on day 2 is 8 - 7 = 1. 
    fmt.Println(amountPainted([][]int{{1,4},{4,7},{5,8}})) // [3,3,1]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/02/01/screenshot-2022-02-01-at-17-17-45-diagram-drawio-diagrams-net.png" />
    // Input: paint = [[1,4],[5,8],[4,7]]
    // Output: [3,3,1]
    // Explanation:
    // On day 0, paint everything between 1 and 4.
    // The amount of new area painted on day 0 is 4 - 1 = 3.
    // On day 1, paint everything between 5 and 8.
    // The amount of new area painted on day 1 is 8 - 5 = 3.
    // On day 2, paint everything between 4 and 5.
    // Everything between 5 and 7 was already painted on day 1.
    // The amount of new area painted on day 2 is 5 - 4 = 1. 
    fmt.Println(amountPainted([][]int{{1,4},{5,8},{4,7}})) // [3,3,1]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2022/02/01/screenshot-2022-02-01-at-17-19-49-diagram-drawio-diagrams-net.png" />
    // Input: paint = [[1,5],[2,4]]
    // Output: [4,0]
    // Explanation:
    // On day 0, paint everything between 1 and 5.
    // The amount of new area painted on day 0 is 5 - 1 = 4.
    // On day 1, paint nothing because everything between 2 and 4 was already painted on day 0.
    // The amount of new area painted on day 1 is 0.
    fmt.Println(amountPainted([][]int{{1,5},{2,4}})) // [4,0]
}