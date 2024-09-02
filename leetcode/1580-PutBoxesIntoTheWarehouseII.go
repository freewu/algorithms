package main

// 1580. Put Boxes Into the Warehouse II
// You are given two arrays of positive integers, boxes and warehouse, representing the heights of some boxes of unit width and the heights of n rooms in a warehouse respectively. 
// The warehouse's rooms are labeled from 0 to n - 1 from left to right where warehouse[i] (0-indexed) is the height of the ith room.

// Boxes are put into the warehouse by the following rules:
//     Boxes cannot be stacked.
//     You can rearrange the insertion order of the boxes.
//     Boxes can be pushed into the warehouse from either side (left or right)
//     If the height of some room in the warehouse is less than the height of a box, then that box and all other boxes behind it will be stopped before that room.

// Return the maximum number of boxes you can put into the warehouse.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/08/30/22.png" />
// Input: boxes = [1,2,2,3,4], warehouse = [3,4,1,2]
// Output: 4
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2020/08/30/22-1.png" />
// We can store the boxes in the following order:
// 1- Put the yellow box in room 2 from either the left or right side.
// 2- Put the orange box in room 3 from the right side.
// 3- Put the green box in room 1 from the left side.
// 4- Put the red box in room 0 from the left side.
// Notice that there are other valid ways to put 4 boxes such as swapping the red and green boxes or the red and orange boxes.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/08/30/22-2.png" />
// Input: boxes = [3,5,5,2], warehouse = [2,1,3,4,5]
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2020/08/30/22-3.png" />
// It is not possible to put the two boxes of height 5 in the warehouse since there's only 1 room of height >= 5.
// Other valid solutions are to put the green box in room 2 or to put the orange box first in room 2 before putting the green and red boxes.

// Constraints:
//     n == warehouse.length
//     1 <= boxes.length, warehouse.length <= 10^5
//     1 <= boxes[i], warehouse[i] <= 10^9

import "fmt"
import "sort"

func maxBoxesInWarehouse(boxes []int, warehouse []int) int {
    n, m, leftValid, rightValid, res := len(warehouse), len(boxes), 1 << 31, 1 << 31, 0
    validWarehouse := make([]int, n)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        leftValid = min(leftValid, warehouse[i])
        validWarehouse[i] = max(validWarehouse[i], leftValid)
    }
    for i := n - 1; i >= 0; i-- {
        rightValid = min(rightValid, warehouse[i])
        validWarehouse[i] = max(validWarehouse[i], rightValid)
    }
    sort.Ints(boxes)
    sort.Ints(validWarehouse)
    for i, j := 0, 0; i < m && j < n; i++ {
        for j < n && boxes[i] > validWarehouse[j] { j++}
        if j < n && boxes[i] <= validWarehouse[j] {
            res++
            j++
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/08/30/22.png" />
    // Input: boxes = [1,2,2,3,4], warehouse = [3,4,1,2]
    // Output: 4
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2020/08/30/22-1.png" />
    // We can store the boxes in the following order:
    // 1- Put the yellow box in room 2 from either the left or right side.
    // 2- Put the orange box in room 3 from the right side.
    // 3- Put the green box in room 1 from the left side.
    // 4- Put the red box in room 0 from the left side.
    // Notice that there are other valid ways to put 4 boxes such as swapping the red and green boxes or the red and orange boxes.
    fmt.Println(maxBoxesInWarehouse([]int{1,2,2,3,4}, []int{3,4,1,2})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/08/30/22-2.png" />
    // Input: boxes = [3,5,5,2], warehouse = [2,1,3,4,5]
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2020/08/30/22-3.png" />
    // It is not possible to put the two boxes of height 5 in the warehouse since there's only 1 room of height >= 5.
    // Other valid solutions are to put the green box in room 2 or to put the orange box first in room 2 before putting the green and red boxes.
    fmt.Println(maxBoxesInWarehouse([]int{3,5,5,2}, []int{2,1,3,4,5})) // 3
}