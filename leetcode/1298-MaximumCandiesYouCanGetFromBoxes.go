package main

// 1298. Maximum Candies You Can Get from Boxes
// You have n boxes labeled from 0 to n - 1. 
// You are given four arrays: status, candies, keys, and containedBoxes where:
//     status[i] is 1 if the ith box is open and 0 if the ith box is closed,
//     candies[i] is the number of candies in the ith box,
//     keys[i] is a list of the labels of the boxes you can open after opening the ith box.
//     containedBoxes[i] is a list of the boxes you found inside the ith box.

// You are given an integer array initialBoxes that contains the labels of the boxes you initially have. 
// You can take all the candies in any open box and you can use the keys in it to open new boxes and you also can use the boxes you find in it.

// Return the maximum number of candies you can get following the rules above.

// Example 1:
// Input: status = [1,0,1,0], candies = [7,5,4,100], keys = [[],[],[1],[]], containedBoxes = [[1,2],[3],[],[]], initialBoxes = [0]
// Output: 16
// Explanation: You will be initially given box 0. You will find 7 candies in it and boxes 1 and 2.
// Box 1 is closed and you do not have a key for it so you will open box 2. You will find 4 candies and a key to box 1 in box 2.
// In box 1, you will find 5 candies and box 3 but you will not find a key to box 3 so box 3 will remain closed.
// Total number of candies collected = 7 + 4 + 5 = 16 candy.

// Example 2:
// Input: status = [1,0,0,0,0,0], candies = [1,1,1,1,1,1], keys = [[1,2,3,4,5],[],[],[],[],[]], containedBoxes = [[1,2,3,4,5],[],[],[],[],[]], initialBoxes = [0]
// Output: 6
// Explanation: You have initially box 0. Opening it you can find boxes 1,2,3,4 and 5 and their keys.
// The total number of candies will be 6.

// Constraints:
//     n == status.length == candies.length == keys.length == containedBoxes.length
//     1 <= n <= 1000
//     status[i] is either 0 or 1.
//     1 <= candies[i] <= 1000
//     0 <= keys[i].length <= n
//     0 <= keys[i][j] < n
//     All values of keys[i] are unique.
//     0 <= containedBoxes[i].length <= n
//     0 <= containedBoxes[i][j] < n
//     All values of containedBoxes[i] are unique.
//     Each box is contained in one box at most.
//     0 <= initialBoxes.length <= n
//     0 <= initialBoxes[i] < n

import "fmt"

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
    n := len(status)
    got, queue, used := make([]bool, n, n), make([]int, 0, 0), make([]bool, n, n)
    for _, i := range initialBoxes {
        got[i] = true
        if status[i] == 1 {
            queue = append(queue, i)
            used[i] = true
        }
    }
    res := 0
    for len(queue) != 0 {
        cur := queue[0]
        queue = queue[1:]
        res += candies[cur]
        for _, i := range keys[cur] {
            status[i] = 1
            if got[i] && !used[i] {
                queue = append(queue, i)
                used[i] = true
            }
        }
        for _, i := range containedBoxes[cur] {
            got[i] = true
            if status[i] == 1 && !used[i] {
                queue = append(queue, i)
                used[i] = true
            }
        }
    }
    return res
}

func maxCandies1(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
    res, n := 0, len(status)
    queue, mp, visited := []int{}, make([]bool, n), make([]bool, n)
    for _, v := range initialBoxes {
        queue = append(queue, v)
        mp[v] = true
    }
    for len(queue) > 0 {
        e := queue[0]
        queue = queue[1:]
        if !visited[e] && status[e] == 1 { // 没拿走 box 的物品，且可以打开
            res += candies[e]
            visited[e] = true
            queue = append(queue, containedBoxes[e]...)
            for _, k := range keys[e] {
                status[k] = 1
                if !visited[k] && mp[k] {
                    queue = append(queue, k)
                }
            }
        } else {
            mp[e] = true
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: status = [1,0,1,0], candies = [7,5,4,100], keys = [[],[],[1],[]], containedBoxes = [[1,2],[3],[],[]], initialBoxes = [0]
    // Output: 16
    // Explanation: You will be initially given box 0. You will find 7 candies in it and boxes 1 and 2.
    // Box 1 is closed and you do not have a key for it so you will open box 2. You will find 4 candies and a key to box 1 in box 2.
    // In box 1, you will find 5 candies and box 3 but you will not find a key to box 3 so box 3 will remain closed.
    // Total number of candies collected = 7 + 4 + 5 = 16 candy.
    fmt.Println(maxCandies([]int{1,0,1,0},[]int{7,5,4,100}, [][]int{{},{},{1},{}},[][]int{{1,2}, {3}, {},{}}, []int{0})) // 16
    // Example 2:
    // Input: status = [1,0,0,0,0,0], candies = [1,1,1,1,1,1], keys = [[1,2,3,4,5],[],[],[],[],[]], containedBoxes = [[1,2,3,4,5],[],[],[],[],[]], initialBoxes = [0]
    // Output: 6
    // Explanation: You have initially box 0. Opening it you can find boxes 1,2,3,4 and 5 and their keys.
    // The total number of candies will be 6.
    fmt.Println(maxCandies([]int{1,0,0,0,0,0},[]int{1,1,1,1,1,1}, [][]int{{1,2,3,4,5},{},{},{},{},{}},[][]int{{1,2,3,4,5},{},{},{},{},{}}, []int{0})) // 6

    fmt.Println(maxCandies1([]int{1,0,1,0},[]int{7,5,4,100}, [][]int{{},{},{1},{}},[][]int{{1,2}, {3}, {},{}}, []int{0})) // 16
    fmt.Println(maxCandies1([]int{1,0,0,0,0,0},[]int{1,1,1,1,1,1}, [][]int{{1,2,3,4,5},{},{},{},{},{}},[][]int{{1,2,3,4,5},{},{},{},{},{}}, []int{0})) // 6
}