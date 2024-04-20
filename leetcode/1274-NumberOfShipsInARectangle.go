package main

// 1274. Number of Ships in a Rectangle
// (This problem is an interactive problem.)
// Each ship is located at an integer point on the sea represented by a cartesian plane,
// and each integer point may contain at most 1 ship.

// You have a function Sea.hasShips(topRight, bottomLeft) which takes two points as arguments 
// and returns true If there is at least one ship in the rectangle represented by the two points, 
// including on the boundary.

// Given two points: the top right and bottom left corners of a rectangle, r
// eturn the number of ships present in that rectangle. 
// It is guaranteed that there are at most 10 ships in that rectangle.

// Submissions making more than 400 calls to hasShips will be judged Wrong Answer. 
// Also, any solutions that attempt to circumvent the judge will result in disqualification.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/07/26/1445_example_1.PNG" />
// Input: 
// ships = [[1,1],[2,2],[3,3],[5,5]], topRight = [4,4], bottomLeft = [0,0]
// Output: 3
// Explanation: From [0,0] to [4,4] we can count 3 ships within the range.

// Example 2:
// Input: ans = [[1,1],[2,2],[3,3]], topRight = [1000,1000], bottomLeft = [0,0]
// Output: 3
 
// Constraints:
//     On the input ships is only given to initialize the map internally. You must solve this problem "blindfolded". In other words, you must find the answer using the given hasShips API, without knowing the ships position.
//     0 <= bottomLeft[0] <= topRight[0] <= 1000
//     0 <= bottomLeft[1] <= topRight[1] <= 1000
//     topRight != bottomLeft

import "fmt"

type Sea struct {
}

func (this Sea) hasShips(topRight,  bottomLeft []int) bool {
    return true
}

/**
 * // This is Sea's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Sea struct {
 *     func hasShips(topRight, bottomLeft []int) bool {}
 * }
 */

// 四分搜索
func countShips(sea Sea, tr, bl []int) int {
    if tr[0] < bl[0] || tr[1] < bl[1] || !sea.hasShips(tr, bl) {
        return 0
    }
    if tr[0] == bl[0] && tr[1]==bl[1] {
        return 1
    }
    hx, hy, lx, ly := tr[0], tr[1], bl[0], bl[1]
    mx, my := (hx + lx) >> 1, (hy + ly) >> 1
    return countShips(sea, tr, []int{mx + 1, my + 1}) + countShips(sea, []int{mx, my}, bl) + countShips(sea, []int{mx, hy}, []int{lx, my + 1}) + countShips(sea, []int{hx, my}, []int{mx + 1, ly})
}

func main() {

}