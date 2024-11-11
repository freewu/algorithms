package main

// 2345. Finding the Number of Visible Mountains
// You are given a 0-indexed 2D integer array peaks where peaks[i] = [xi, yi] states that mountain i has a peak at coordinates (xi, yi). 
// A mountain can be described as a right-angled isosceles triangle, with its base along the x-axis and a right angle at its peak. 
// More formally, the gradients of ascending and descending the mountain are 1 and -1 respectively.

// A mountain is considered visible if its peak does not lie within another mountain (including the border of other mountains).

// Return the number of visible mountains.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/07/19/ex1.png" />
// Input: peaks = [[2,2],[6,3],[5,4]]
// Output: 2
// Explanation: The diagram above shows the mountains.
// - Mountain 0 is visible since its peak does not lie within another mountain or its sides.
// - Mountain 1 is not visible since its peak lies within the side of mountain 2.
// - Mountain 2 is visible since its peak does not lie within another mountain or its sides.
// There are 2 mountains that are visible.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/07/19/ex2new1.png" />
// Input: peaks = [[1,3],[1,3]]
// Output: 0
// Explanation: The diagram above shows the mountains (they completely overlap).
// Both mountains are not visible since their peaks lie within each other.

// Constraints:
//     1 <= peaks.length <= 10^5
//     peaks[i].length == 2
//     1 <= xi, yi <= 10^5

import "fmt"
import "sort"

func visibleMountains(peaks [][]int) int {
    arr, count := [][2]int{}, make(map[int]int)
    for _, peak := range peaks {
        arr = append(arr, [2]int{ peak[0] - peak[1], peak[1] + peak[0] })
        count[(peak[0] - peak[1]) * 100001 + peak[0] + peak[1]]++
    }
    sort.Slice(arr, func(i, j int) bool {
        if arr[i][0] == arr[j][0] { return arr[i][1] > arr[j][1] }
        return arr[i][0] < arr[j][0]
    })
    res, mx := 0, -100001
    for _, v := range arr {
        r := v[1]
        if r <= mx { continue }
        mx = r
        if count[v[0] * 100001 + v[1]] < 2 {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/07/19/ex1.png" />
    // Input: peaks = [[2,2],[6,3],[5,4]]
    // Output: 2
    // Explanation: The diagram above shows the mountains.
    // - Mountain 0 is visible since its peak does not lie within another mountain or its sides.
    // - Mountain 1 is not visible since its peak lies within the side of mountain 2.
    // - Mountain 2 is visible since its peak does not lie within another mountain or its sides.
    // There are 2 mountains that are visible.
    fmt.Println(visibleMountains([][]int{{2,2},{6,3},{5,4}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/07/19/ex2new1.png" />
    // Input: peaks = [[1,3],[1,3]]
    // Output: 0
    // Explanation: The diagram above shows the mountains (they completely overlap).
    // Both mountains are not visible since their peaks lie within each other.
    fmt.Println(visibleMountains([][]int{{1,3},{1,3}})) // 0
}