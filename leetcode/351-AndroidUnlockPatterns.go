package main

// 351. Android Unlock Patterns
// Android devices have a special lock screen with a 3 x 3 grid of dots. 
// Users can set an "unlock pattern" by connecting the dots in a specific sequence, forming a series of joined line segments where each segment's endpoints are two consecutive dots in the sequence. 
// A sequence of k dots is a valid unlock pattern if both of the following are true:
//     All the dots in the sequence are distinct.
//     If the line segment connecting two consecutive dots in the sequence passes through the center of any other dot, the other dot must have previously appeared in the sequence. No jumps through the center non-selected dots are allowed.
//         For example, connecting dots 2 and 9 without dots 5 or 6 appearing beforehand is valid because the line from dot 2 to dot 9 does not pass through the center of either dot 5 or 6.
//         However, connecting dots 1 and 3 without dot 2 appearing beforehand is invalid because the line from dot 1 to dot 3 passes through the center of dot 2.

// Here are some example valid and invalid unlock patterns:
//     <img src="https://assets.leetcode.com/uploads/2018/10/12/android-unlock.png" />

//     The 1st pattern [4,1,3,6] is invalid because the line connecting dots 1 and 3 pass through dot 2, but dot 2 did not previously appear in the sequence.
//     The 2nd pattern [4,1,9,2] is invalid because the line connecting dots 1 and 9 pass through dot 5, but dot 5 did not previously appear in the sequence.
//     The 3rd pattern [2,4,1,3,6] is valid because it follows the conditions. 
//         The line connecting dots 1 and 3 meets the condition because dot 2 previously appeared in the sequence.
//     The 4th pattern [6,5,4,1,9,2] is valid because it follows the conditions. 
//         The line connecting dots 1 and 9 meets the condition because dot 5 previously appeared in the sequence.

// Given two integers m and n, return the number of unique and valid unlock patterns of the Android grid lock screen that consist of at least m keys and at most n keys.
// Two unlock patterns are considered unique if there is a dot in one sequence that is not in the other, or the order of the dots is different.

// Example 1:
// Input: m = 1, n = 1
// Output: 9

// Example 2:
// Input: m = 1, n = 2
// Output: 65

// Constraints:
//     1 <= m, n <= 9

import "fmt"
import "math"

func numberOfPatterns(m int, n int) int {
    res := 0
    keyborad := [9][2]int{{0, 0}, {1, 0}, {2, 0}, {0, 1}, {1, 1}, {2, 1}, {0, 2}, {1, 2}, {2, 2}}
    nextAvailable := func (passed ...int) []int {
        history, x, y, available := make(map[int]bool, len(passed)), keyborad[passed[len(passed)-1]][0], keyborad[passed[len(passed)-1]][1], make([]int, 0)
        for _, pass := range passed {
            history[pass] = true
        }
        for index, position := range keyborad {
            if diffx, diffy := position[0]-x, position[1]-y; (((diffx%2 != 0) || int(math.Abs(float64(diffy))) != 2) && (int(math.Abs(float64(diffx))) != 2 || (diffy%2 != 0)) || history[(diffx/2+x)*3+y+(diffy/2)]) && !history[index] {
                available = append(available, index)
            }
        }
        return available
    }
    var findZip func(m int, n int, passed ...int) [][]int
    findZip = func (m int, n int, passed ...int) [][]int {
        zip := make([][]int, 0)
        if len(passed) <= n {
            if len(passed) >= m {
                zip = append(zip, passed)
            }
            if len(passed) < n {
                for _, num := range nextAvailable(passed...) {
                    zip = append(zip, findZip(m, n, append(passed, num)...)...)
                }
            }
        }
        return zip
    }
    for index := range keyborad {
        res += len(findZip(m, n, index))
    }
    return res
}

// 打表
func numberOfPatterns1(m int, n int) int {
    l := []int{ 9, 56, 320, 1624, 7152, 26016, 72912, 140704, 140704 }
    res := 0
    for i := m - 1; i < n; i++ {
        res += l[i]
    }
    return res
}

func main() {
    // Example 1:
    // Input: m = 1, n = 1
    // Output: 9
    fmt.Println(numberOfPatterns(1,1)) // 9
    // Example 2:
    // Input: m = 1, n = 2
    // Output: 65
    fmt.Println(numberOfPatterns(1,2)) // 65

    fmt.Println(numberOfPatterns1(1,1)) // 9
    fmt.Println(numberOfPatterns1(1,2)) // 65
}