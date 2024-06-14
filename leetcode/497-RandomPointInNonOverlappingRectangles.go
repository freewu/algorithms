package main

// 497. Random Point in Non-overlapping Rectangles
// You are given an array of non-overlapping axis-aligned rectangles rects where rects[i] = [ai, bi, xi, yi] indicates 
// that (ai, bi) is the bottom-left corner point of the ith rectangle and (xi, yi) is the top-right corner point of the ith rectangle. 
// Design an algorithm to pick a random integer point inside the space covered by one of the given rectangles.
// A point on the perimeter of a rectangle is included in the space covered by the rectangle.

// Any integer point inside the space covered by one of the given rectangles should be equally likely to be returned.
// Note that an integer point is a point that has integer coordinates.

// Implement the Solution class:
//     Solution(int[][] rects) Initializes the object with the given rectangles rects.
//     int[] pick() Returns a random integer point [u, v] inside the space covered by one of the given rectangles.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/24/lc-pickrandomrec.jpg" />
// Input
// ["Solution", "pick", "pick", "pick", "pick", "pick"]
// [[[[-2, -2, 1, 1], [2, 2, 4, 6]]], [], [], [], [], []]
// Output
// [null, [1, -2], [1, -1], [-1, -2], [-2, -2], [0, 0]]
// Explanation
// Solution solution = new Solution([[-2, -2, 1, 1], [2, 2, 4, 6]]);
// solution.pick(); // return [1, -2]
// solution.pick(); // return [1, -1]
// solution.pick(); // return [-1, -2]
// solution.pick(); // return [-2, -2]
// solution.pick(); // return [0, 0]
 
// Constraints:
//     1 <= rects.length <= 100
//     rects[i].length == 4
//     -10^9 <= ai < xi <= 10^9
//     -10^9 <= bi < yi <= 10^9
//     xi - ai <= 2000
//     yi - bi <= 2000
//     All the rectangles do not overlap.
//     At most 10^4 calls will be made to pick.

import "fmt"
import "math/rand"
import "sort"

// type Solution struct {
//     n, total int
//     rects [][]int
//     prefixsum []int
// }

// func Constructor(rects [][]int) Solution {
//     n := len(rects)
//     prefixsum := make([]int, n)
//     for i, rect := range rects {
//         length := rect[2] - rect[0] + 1
//         width := rect[3] - rect[1] + 1
//         area := length * width
//         prefixsum[i] = area
//         if i > 0 {
//             prefixsum[i] += prefixsum[i - 1]
//         }
//     }
//     return Solution{n, prefixsum[n - 1], rects, prefixsum}
// }

// func (this *Solution) Pick() []int {
//     point := random(1, this.total)
//     index := binarySearch(this.prefixsum, point)
//     rect := this.rects[index]
//     return []int{random(rect[0], rect[2]), random(rect[1], rect[3])}
// }

// func random(mn, mx int) int {
//     return mn + rand.Intn(mx - mn + 1)
// }

// func binarySearch(a []int, target int) int {
//     low, high := 0, len(a) - 1
//     for low < high {
//         mid := low + ((high - low) >> 1)
//         if a[mid] < target {
//             low = mid + 1
//         } else {
//             high = mid
//         }
//     }
//     return low
// }

type Solution struct {
    sum []int
    rects [][]int
}

func Constructor(rects [][]int) Solution {
    res := make([]int, len(rects) + 1)
    for i, v := range rects {
        a, b, x, y := v[0], v[1], v[2], v[3]
        res[i+1] = res[i] + (x - a + 1) * (y - b + 1)
    }
    return Solution{res, rects}
}

func (this *Solution) Pick() []int {
    n := rand.Intn(this.sum[len(this.sum)-1])
    pos := sort.SearchInts(this.sum, n + 1) - 1
    a, b, y := this.rects[pos][0], this.rects[pos][1], this.rects[pos][3]
    return []int{(n - this.sum[pos]) / (y - b + 1) + a, (n - this.sum[pos])  % (y - b + 1) + b}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */

func main() {
    // Solution solution = new Solution([[-2, -2, 1, 1], [2, 2, 4, 6]]);
    obj := Constructor([][]int{{-2, -2, 1, 1},{2, 2, 4, 6}})
    fmt.Println(obj)
    // solution.pick(); // return [1, -2]
    fmt.Println(obj.Pick()) //  [1, -2]
    // solution.pick(); // return [1, -1]
    fmt.Println(obj.Pick()) //  [1, -1]
    // solution.pick(); // return [-1, -2]
    fmt.Println(obj.Pick()) //  [-1, -2]
    // solution.pick(); // return [-2, -2]
    fmt.Println(obj.Pick()) //  [-2, -2]
    // solution.pick(); // return [0, 0]
    fmt.Println(obj.Pick()) //  [0, 0]
}