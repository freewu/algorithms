package main

// 2613. Beautiful Pairs
// You are given two 0-indexed integer arrays nums1 and nums2 of the same length. 
// A pair of indices (i,j) is called beautiful if|nums1[i] - nums1[j]| + |nums2[i] - nums2[j]| is the smallest amongst all possible indices pairs where i < j.

// Return the beautiful pair. 
// In the case that there are multiple beautiful pairs, return the lexicographically smallest pair.

// Note that
//     |x| denotes the absolute value of x.
//     A pair of indices (i1, j1) is lexicographically smaller than (i2, j2) if i1 < i2 or i1 == i2 and j1 < j2.

// Example 1:
// Input: nums1 = [1,2,3,2,4], nums2 = [2,3,1,2,3]
// Output: [0,3]
// Explanation: Consider index 0 and index 3. The value of |nums1[i]-nums1[j]| + |nums2[i]-nums2[j]| is 1, which is the smallest value we can achieve.

// Example 2:
// Input: nums1 = [1,2,4,3,2,5], nums2 = [1,4,2,3,5,1]
// Output: [1,4]
// Explanation: Consider index 1 and index 4. The value of |nums1[i]-nums1[j]| + |nums2[i]-nums2[j]| is 1, which is the smallest value we can achieve.

// Constraints:
//     2 <= nums1.length, nums2.length <= 10^5
//     nums1.length == nums2.length
//     0 <= nums1i <= nums1.length
//     0 <= nums2i <= nums2.length

import "fmt"
import "sort"

func beautifulPair(nums1 []int, nums2 []int) []int {
    n := len(nums1)
    pl := map[[2]int][]int{}
    for i := 0; i < n; i++ {
        k := [2]int{nums1[i], nums2[i]}
        pl[k] = append(pl[k], i)
    }
    points := [][3]int{}
    for i := 0; i < n; i++ {
        k := [2]int{nums2[i], nums1[i]}
        if len(pl[k]) > 1 { return []int{pl[k][0], pl[k][1]} }
        points = append(points, [3]int{nums1[i], nums2[i], i})
    }
    sort.Slice(points, func(i, j int) bool { 
        return points[i][0] < points[j][0] 
    })
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    distance := func (x1, y1, x2, y2 int) int { return abs(x1-x2) + abs(y1-y2) }
    var dfs func(l, r int) [3]int
    dfs = func(l, r int) [3]int {
        if l >= r { return [3]int{ 1 << 31, -1, -1 } }
        m := (l + r) >> 1
        x := points[m][0]
        t1, t2 := dfs(l, m), dfs(m + 1, r)
        if t1[0] > t2[0] || (t1[0] == t2[0] && (t1[1] > t2[1] || (t1[1] == t2[1] && t1[2] > t2[2]))) {
            t1 = t2
        }
        t := [][3]int{}
        for i := l; i <= r; i++ {
            if abs(points[i][0] - x) <= t1[0] {
                t = append(t, points[i])
            }
        }
        sort.Slice(t, func(i, j int) bool { 
            return t[i][1] < t[j][1]
        })
        for i := 0; i < len(t); i++ {
            for j := i + 1; j < len(t); j++ {
                if t[j][1]-t[i][1] > t1[0] { break }
                pi, pj := min(t[i][2], t[j][2]), max(t[i][2], t[j][2])
                d := distance(t[i][0], t[i][1], t[j][0], t[j][1])
                if d < t1[0] || (d == t1[0] && (pi < t1[1] || (pi == t1[1] && pj < t1[2]))) {
                    t1 = [3]int{d, pi, pj}
                }
            }
        }
        return t1
    }
    res := dfs(0, n-1)
    return []int{ res[1], res[2] }
}

func main() {
    // Example 1:
    // Input: nums1 = [1,2,3,2,4], nums2 = [2,3,1,2,3]
    // Output: [0,3]
    // Explanation: Consider index 0 and index 3. The value of |nums1[i]-nums1[j]| + |nums2[i]-nums2[j]| is 1, which is the smallest value we can achieve.
    fmt.Println(beautifulPair([]int{1,2,3,2,4}, []int{2,3,1,2,3})) // [0,3]
    // Example 2:
    // Input: nums1 = [1,2,4,3,2,5], nums2 = [1,4,2,3,5,1]
    // Output: [1,4]
    // Explanation: Consider index 1 and index 4. The value of |nums1[i]-nums1[j]| + |nums2[i]-nums2[j]| is 1, which is the smallest value we can achieve.
    fmt.Println(beautifulPair([]int{1,2,4,3,2,5}, []int{1,4,2,3,5,1})) // [1,4]
}