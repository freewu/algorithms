package main

// 1564. Put Boxes Into the Warehouse I
// You are given two arrays of positive integers, boxes and warehouse, representing the heights of some boxes of unit width and the heights of n rooms in a warehouse respectively. 
// The warehouse's rooms are labelled from 0 to n - 1 from left to right where warehouse[i] (0-indexed) is the height of the ith room.

// Boxes are put into the warehouse by the following rules:
//     Boxes cannot be stacked.
//     You can rearrange the insertion order of the boxes.
//     Boxes can only be pushed into the warehouse from left to right only.
//     If the height of some room in the warehouse is less than the height of a box, then that box and all other boxes behind it will be stopped before that room.

// Return the maximum number of boxes you can put into the warehouse.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/08/26/11.png" />
// Input: boxes = [4,3,4,1], warehouse = [5,3,3,4,1]
// Output: 3
// Explanation: 
// <img src="https://assets.leetcode.com/uploads/2020/08/26/12.png" />
// We can first put the box of height 1 in room 4. Then we can put the box of height 3 in either of the 3 rooms 1, 2, or 3. Lastly, we can put one box of height 4 in room 0.
// There is no way we can fit all 4 boxes in the warehouse.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/08/26/21.png" />
// Input: boxes = [1,2,2,3,4], warehouse = [3,4,1,2]
// Output: 3
// Explanation: 
// <img src="https://assets.leetcode.com/uploads/2020/08/26/22.png" />
// Notice that it's not possible to put the box of height 4 into the warehouse since it cannot pass the first room of height 3.
// Also, for the last two rooms, 2 and 3, only boxes of height 1 can fit.
// We can fit 3 boxes maximum as shown above. The yellow box can also be put in room 2 instead.
// Swapping the orange and green boxes is also valid, or swapping one of them with the red box.

// Example 3:
// Input: boxes = [1,2,3], warehouse = [1,2,3,4]
// Output: 1
// Explanation: Since the first room in the warehouse is of height 1, we can only put boxes of height 1.

// Constraints:
//     n == warehouse.length
//     1 <= boxes.length, warehouse.length <= 10^5
//     1 <= boxes[i], warehouse[i] <= 10^9

import "fmt"
import "sort"

func maxBoxesInWarehouse(boxes []int, warehouse []int) int {
    i, count := len(boxes) - 1, 0
    sort.Ints(boxes)
    for _, room := range warehouse {
        // 从最小到最大遍历方框
        // 丢弃不适合当前仓库的箱子
        for i >= 0 && boxes[i] > room {
            i--
        }
        if i == -1 {
            return count
        }
        count++
        i--
    }
    return count
}

func maxBoxesInWarehouse1(boxes []int, warehouse []int) int {
    // 对仓库房间的高度进行预处理，以获得可用高度
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < len(warehouse); i++ {
        warehouse[i] = min(warehouse[i-1], warehouse[i])
    }
    // 从小到大遍历方框
    sort.Ints(boxes)
    count := 0
    for i := len(warehouse) - 1; i >= 0; i-- {
        if count < len(boxes) && boxes[count] <= warehouse[i] { // 清点能放进当前库房的箱子
            count++
        }
    }
    return count
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/08/26/11.png" />
    // Input: boxes = [4,3,4,1], warehouse = [5,3,3,4,1]
    // Output: 3
    // Explanation: 
    // <img src="https://assets.leetcode.com/uploads/2020/08/26/12.png" />
    // We can first put the box of height 1 in room 4. Then we can put the box of height 3 in either of the 3 rooms 1, 2, or 3. Lastly, we can put one box of height 4 in room 0.
    // There is no way we can fit all 4 boxes in the warehouse.
    fmt.Println(maxBoxesInWarehouse([]int{4,3,4,1}, []int{5,3,3,4,1})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/08/26/21.png" />
    // Input: boxes = [1,2,2,3,4], warehouse = [3,4,1,2]
    // Output: 3
    // Explanation: 
    // <img src="https://assets.leetcode.com/uploads/2020/08/26/22.png" />
    // Notice that it's not possible to put the box of height 4 into the warehouse since it cannot pass the first room of height 3.
    // Also, for the last two rooms, 2 and 3, only boxes of height 1 can fit.
    // We can fit 3 boxes maximum as shown above. The yellow box can also be put in room 2 instead.
    // Swapping the orange and green boxes is also valid, or swapping one of them with the red box.
    fmt.Println(maxBoxesInWarehouse([]int{1,2,2,3,4}, []int{3,4,1,2})) // 3
    // Example 3:
    // Input: boxes = [1,2,3], warehouse = [1,2,3,4]
    // Output: 1
    // Explanation: Since the first room in the warehouse is of height 1, we can only put boxes of height 1.
    fmt.Println(maxBoxesInWarehouse([]int{1,2,3}, []int{1,2,3,4})) // 1

    fmt.Println(maxBoxesInWarehouse1([]int{4,3,4,1}, []int{5,3,3,4,1})) // 4
    fmt.Println(maxBoxesInWarehouse1([]int{1,2,2,3,4}, []int{3,4,1,2})) // 3
    fmt.Println(maxBoxesInWarehouse1([]int{1,2,3}, []int{1,2,3,4})) // 1
}