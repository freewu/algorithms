package main

// 2975. Maximum Square Area by Removing Fences From a Field
// There is a large (m - 1) x (n - 1) rectangular field with corners at (1, 1) 
// and (m, n) containing some horizontal and vertical fences given in arrays hFences and vFences respectively.

// Horizontal fences are from the coordinates (hFences[i], 1) to (hFences[i], n) 
// and vertical fences are from the coordinates (1, vFences[i]) to (m, vFences[i]).

// Return the maximum area of a square field that can be formed by removing some fences (possibly none) or -1 if it is impossible to make a square field.

// Since the answer may be large, return it modulo 10^9 + 7.

// Note: The field is surrounded by two horizontal fences from the coordinates (1, 1) to (1, n) and (m, 1) to (m, n) 
// and two vertical fences from the coordinates (1, 1) to (m, 1) and (1, n) to (m, n). 
// These fences cannot be removed.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/11/05/screenshot-from-2023-11-05-22-40-25.png" />
// Input: m = 4, n = 3, hFences = [2,3], vFences = [2]
// Output: 4
// Explanation: Removing the horizontal fence at 2 and the vertical fence at 2 will give a square field of area 4.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/11/22/maxsquareareaexample1.png" />
// Input: m = 6, n = 7, hFences = [2], vFences = [4]
// Output: -1
// Explanation: It can be proved that there is no way to create a square field by removing fences.

// Constraints:
//     3 <= m, n <= 10^9
//     1 <= hFences.length, vFences.length <= 600
//     1 < hFences[i] < m
//     1 < vFences[i] < n
//     hFences and vFences are unique.

import "fmt"
import "sort"

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
    hFences = append([]int{1}, append(hFences, m)...)
    vFences = append([]int{1}, append(vFences, n)...)
    sort.Ints(hFences)
    sort.Ints(vFences)
    mp := make(map[int]bool)
    for i := 0; i < len(hFences)-1; i++ {
        for j := i + 1; j < len(hFences); j++ {
            mp[hFences[j] - hFences[i]] = true
        }
    }
    res := -1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(vFences) - 1; i++ {
        for j := i + 1; j < len(vFences); j++ {
            diff := vFences[j] - vFences[i]
            if mp[diff] {
                res = max(res, (diff * diff))
            }
        }
    }
    return res % 1_000_000_007
}

func maximizeSquareArea1(m int, n int, hFences []int, vFences []int) int {
    // 最左边坐标为 1 不为 0
    hFences, vFences = append(hFences, 1), append(vFences, 1)
    hFences, vFences = append(hFences, m), append(vFences, n)
    sort.Ints(hFences)
    sort.Ints(vFences)
    mx := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    // 选「最大」，倒序枚举右下顶点、正序枚举左上顶点必定更优
    for i := len(hFences) - 1; i > -1; i-- {
        x := hFences[i]
        for j := len(vFences) - 1; j > -1; j-- {
            y := vFences[j]
            for ph, pv := 0, 0; ph < i && pv < j && min(x - hFences[ph], y - vFences[pv]) > mx; {
                if x - hFences[ph] == y - vFences[pv] {
                    mx = max(mx, x - hFences[ph])
                    break
                } else if x - hFences[ph] < y - vFences[pv] {
                    pv++
                } else {
                    ph++
                }
            }
        }
    }
    if mx > 0 { return mx * mx % 1_000_000_007 }
    return -1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/11/05/screenshot-from-2023-11-05-22-40-25.png" />
    // Input: m = 4, n = 3, hFences = [2,3], vFences = [2]
    // Output: 4
    // Explanation: Removing the horizontal fence at 2 and the vertical fence at 2 will give a square field of area 4.
    fmt.Println(maximizeSquareArea(4, 3, []int{2,3}, []int{2})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/11/22/maxsquareareaexample1.png" />
    // Input: m = 6, n = 7, hFences = [2], vFences = [4]
    // Output: -1
    // Explanation: It can be proved that there is no way to create a square field by removing fences.
    fmt.Println(maximizeSquareArea(6, 7, []int{2}, []int{4})) // -1

    fmt.Println(maximizeSquareArea(6, 7, []int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 64
    fmt.Println(maximizeSquareArea(6, 7, []int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 64
    fmt.Println(maximizeSquareArea(6, 7, []int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 64
    fmt.Println(maximizeSquareArea(6, 7, []int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 64

    fmt.Println(maximizeSquareArea1(4, 3, []int{2,3}, []int{2})) // 4
    fmt.Println(maximizeSquareArea1(6, 7, []int{2}, []int{4})) // -1
    fmt.Println(maximizeSquareArea1(6, 7, []int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 64
    fmt.Println(maximizeSquareArea1(6, 7, []int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 64
    fmt.Println(maximizeSquareArea1(6, 7, []int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 64
    fmt.Println(maximizeSquareArea1(6, 7, []int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 64
}