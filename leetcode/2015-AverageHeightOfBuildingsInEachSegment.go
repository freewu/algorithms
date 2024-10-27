package main

// 2015. Average Height of Buildings in Each Segment
// A perfectly straight street is represented by a number line. 
// The street has building(s) on it and is represented by a 2D integer array buildings, where buildings[i] = [starti, endi, heighti]. 
// This means that there is a building with heighti in the half-closed segment [starti, endi).

// You want to describe the heights of the buildings on the street with the minimum number of non-overlapping segments. 
// The street can be represented by the 2D integer array street where street[j] = [leftj, rightj, averagej] describes a half-closed segment [leftj, rightj) of the road where the average heights of the buildings in the segment is averagej.
//     For example, if buildings = [[1,5,2],[3,10,4]], the street could be represented by street = [[1,3,2],[3,5,3],[5,10,4]] because:
//     From 1 to 3, there is only the first building with an average height of 2 / 1 = 2.
//     From 3 to 5, both the first and the second building are there with an average height of (2+4) / 2 = 3.
//     From 5 to 10, there is only the second building with an average height of 4 / 1 = 4.

// Given buildings, return the 2D integer array street as described above (excluding any areas of the street where there are no buldings). 
// You may return the array in any order.

// The average of n elements is the sum of the n elements divided (integer division) by n.

// A half-closed segment [a, b) is the section of the number line between points a and b including point a and not including point b.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/09/21/image-20210921224001-2.png" />
// Input: buildings = [[1,4,2],[3,9,4]]
// Output: [[1,3,2],[3,4,3],[4,9,4]]
// Explanation:
// From 1 to 3, there is only the first building with an average height of 2 / 1 = 2.
// From 3 to 4, both the first and the second building are there with an average height of (2+4) / 2 = 3.
// From 4 to 9, there is only the second building with an average height of 4 / 1 = 4.

// Example 2:
// Input: buildings = [[1,3,2],[2,5,3],[2,8,3]]
// Output: [[1,3,2],[3,8,3]]
// Explanation:
// From 1 to 2, there is only the first building with an average height of 2 / 1 = 2.
// From 2 to 3, all three buildings are there with an average height of (2+3+3) / 3 = 2.
// From 3 to 5, both the second and the third building are there with an average height of (3+3) / 2 = 3.
// From 5 to 8, there is only the last building with an average height of 3 / 1 = 3.
// The average height from 1 to 3 is the same so we can group them into one segment.
// The average height from 3 to 8 is the same so we can group them into one segment.

// Example 3:
// Input: buildings = [[1,2,1],[5,6,1]]
// Output: [[1,2,1],[5,6,1]]
// Explanation:
// From 1 to 2, there is only the first building with an average height of 1 / 1 = 1.
// From 2 to 5, there are no buildings, so it is not included in the output.
// From 5 to 6, there is only the second building with an average height of 1 / 1 = 1.
// We cannot group the segments together because an empty space with no buildings seperates the segments.

// Constraints:
//     1 <= buildings.length <= 10^5
//     buildings[i].length == 3
//     0 <= starti < endi <= 10^8
//     1 <= heighti <= 10^5

import "fmt"
import "sort"

// 扫描线
func averageHeightOfBuildings(buildings [][]int) [][]int {
    arr, lines := [][]int{}, [][]int{}
    for _, v := range buildings {
        lines = append(lines, []int{ v[0], v[2], 0 })
        lines = append(lines, []int{ v[1], v[2], 1 })
    }
    sort.Slice(lines, func(i, j int) bool {
        if lines[i][0] == lines[j][0] { return lines[i][2] > lines[j][2] }
        return lines[i][0] < lines[j][0]
    })
    count, last, height := 0, 0, 0
    for i := 0; i < len(lines); {
        pos, h, edge := lines[i][0], lines[i][1], lines[i][2]
        t, j := 1, i + 1
        for ; j < len(lines) && lines[j][0] == pos && lines[j][2] == edge; j++ {
            h += lines[j][1] // 合并相同坐标，相同类型的高度
            t++
        }
        i = j
        if count == 0 {
            height = h
            count += t
            last = pos
            continue
        }
        newHeight, newCount := 0, 0
        if edge == 1 {
            newHeight, newCount = height - h, count - t
        } else {
            newHeight, newCount = height + h , count + t
        }
        if last != pos && (newCount == 0 || height / count != newHeight / newCount) {
            arr = append(arr, []int{ last, pos, height / count })
            last = pos
        }
        height, count = newHeight, newCount
    }
    res := [][]int{}
    for i := 0; i < len(arr); i++ {
        insert := arr[i]
        if len(res) > 0 && res[len(res) - 1][1] == insert[0] && res[len(res) - 1][2]== insert[2] {
            res[len(res) - 1][1] = insert[1]
        } else {
            res = append(res, insert)
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/09/21/image-20210921224001-2.png" />
    // Input: buildings = [[1,4,2],[3,9,4]]
    // Output: [[1,3,2],[3,4,3],[4,9,4]]
    // Explanation:
    // From 1 to 3, there is only the first building with an average height of 2 / 1 = 2.
    // From 3 to 4, both the first and the second building are there with an average height of (2+4) / 2 = 3.
    // From 4 to 9, there is only the second building with an average height of 4 / 1 = 4.
    fmt.Println(averageHeightOfBuildings([][]int{{1,4,2},{3,9,4}})) // [[1,3,2],[3,4,3],[4,9,4]]
    // Example 2:
    // Input: buildings = [[1,3,2],[2,5,3],[2,8,3]]
    // Output: [[1,3,2],[3,8,3]]
    // Explanation:
    // From 1 to 2, there is only the first building with an average height of 2 / 1 = 2.
    // From 2 to 3, all three buildings are there with an average height of (2+3+3) / 3 = 2.
    // From 3 to 5, both the second and the third building are there with an average height of (3+3) / 2 = 3.
    // From 5 to 8, there is only the last building with an average height of 3 / 1 = 3.
    // The average height from 1 to 3 is the same so we can group them into one segment.
    // The average height from 3 to 8 is the same so we can group them into one segment.
    fmt.Println(averageHeightOfBuildings([][]int{{1,3,2},{2,5,3},{2,8,3}})) // [[1,3,2],[3,8,3]]
    // Example 3:
    // Input: buildings = [[1,2,1],[5,6,1]]
    // Output: [[1,2,1],[5,6,1]]
    // Explanation:
    // From 1 to 2, there is only the first building with an average height of 1 / 1 = 1.
    // From 2 to 5, there are no buildings, so it is not included in the output.
    // From 5 to 6, there is only the second building with an average height of 1 / 1 = 1.
    // We cannot group the segments together because an empty space with no buildings seperates the segments.
    fmt.Println(averageHeightOfBuildings([][]int{{1,2,1},{5,6,1}})) // [[1,2,1],[5,6,1]]
}