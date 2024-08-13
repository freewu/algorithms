package main

// 3245. Alternating Groups III
// There are some red and blue tiles arranged circularly. 
// You are given an array of integers colors and a 2D integers array queries.

// The color of tile i is represented by colors[i]:
//     colors[i] == 0 means that tile i is red.
//     colors[i] == 1 means that tile i is blue.

// An alternating group is a contiguous subset of tiles in the circle with alternating colors 
// (each tile in the group except the first and last one has a different color from its adjacent tiles in the group).

// You have to process queries of two types:
//     queries[i] = [1, sizei], determine the count of alternating groups with size sizei.
//     queries[i] = [2, indexi, colori], change colors[indexi] to colori.

// Return an array answer containing the results of the queries of the first type in order.
// Note that since colors represents a circle, the first and the last tiles are considered to be next to each other.

// Example 1:
// Input: colors = [0,1,1,0,1], queries = [[2,1,0],[1,4]]
// Output: [2]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-14-44.png">
// First query:
// Change colors[1] to 0.
// <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-20-25.png">
// Second query:
// Count of the alternating groups with size 4:
// <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-24-12.png">

// Example 2:
// Input: colors = [0,0,1,0,1,1], queries = [[1,3],[2,3,0],[1,5]]
// Output: [2,0]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-35-50.png">
// First query:
// Count of the alternating groups with size 3:
// <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-36-40.png">
// Second query: colors will not change.
// Third query: There is no alternating group with size 5.

// Constraints:
//     4 <= colors.length <= 5 * 10^4
//     0 <= colors[i] <= 1
//     1 <= queries.length <= 5 * 10^4
//     queries[i][0] == 1 or queries[i][0] == 2
//     For all i that:
//     queries[i][0] == 1: queries[i].length == 2, 3 <= queries[i][1] <= colors.length - 1
//     queries[i][0] == 2: queries[i].length == 3, 0 <= queries[i][1] <= colors.length - 1, 0 <= queries[i][2] <= 1

import "fmt"

type FenwickTree [][2]int

// op=1，添加一个 size
// op=-1，移除一个 size
func (this FenwickTree) update(size, op int) {
    for i := len(this) - size; i < len(this); i += i & -i {
        this[i][0] += op
        this[i][1] += op * size
    }
}

// 返回 >= size 的元素个数，元素和
func (this FenwickTree) query(size int) (int, int) {
    cnt, sum := 0, 0
    for i := len(this) - size; i > 0; i &= i - 1 {
        cnt += this[i][0]
        sum += this[i][1]
    }
    return cnt, sum 
}

func numberOfAlternatingGroups(colors[]int, queries [][]int) []int {
    n := len(colors)
    set := redblacktree.New[int, struct{}]()
    res, ft := []int{}, make(FenwickTree, n+1)

    // op=1，添加一个结束位置 i
    // op=-1，移除一个结束位置 i
    update := func(i, op int) {
        prev, ok := set.Floor(i)
        if !ok {
            prev = set.Right()
        }
        pre := prev.Key
        next, ok := set.Ceiling(i)
        if !ok {
            next = set.Left()
        }
        nxt := next.Key
        ft.update((nxt-pre+n-1)%n+1, -op) // 移除/添加旧长度
        ft.update((i-pre+n)%n, op)
        ft.update((nxt-i+n)%n, op) // 添加/移除新长度
    }
    // 添加一个结束位置 i
    add := func(i int) {
        if set.Empty() {
            ft.update(n, 1)
        } else {
            update(i, 1)
        }
        set.Put(i, struct{}{})
    }
    // 移除一个结束位置 i
    del := func(i int) {
        set.Remove(i)
        if set.Empty() {
            ft.update(n, -1)
        } else {
            update(i, -1)
        }
    }

    for i, c := range colors {
        if c == colors[(i+1) % n] {
            add(i) // i 是一个结束位置
        }
    }
    for _, q := range queries {
        if q[0] == 1 {
            if set.Empty() {
                res = append(res, n) // 每个长为 size 的子数组都符合要求
            } else {
                cnt, sum := ft.query(q[1])
                res = append(res, sum - cnt * (q[1]-1))
            }
        } else {
            i := q[1]
            if colors[i] == q[2] { // 无影响
                continue
            }
            pre, nxt := (i-1+n)%n, (i+1)%n
            // 修改前，先去掉结束位置
            if colors[pre] == colors[i] {
                del(pre)
            }
            if colors[i] == colors[nxt] {
                del(i)
            }
            colors[i] ^= 1
            // 修改后，添加新的结束位置
            if colors[pre] == colors[i] {
                add(pre)
            }
            if colors[i] == colors[nxt] {
                add(i)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: colors = [0,1,1,0,1], queries = [[2,1,0],[1,4]]
    // Output: [2]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-14-44.png">
    // First query:
    // Change colors[1] to 0.
    // <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-20-25.png">
    // Second query:
    // Count of the alternating groups with size 4:
    // <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-24-12.png">
    fmt.Println(numberOfAlternatingGroups([]int{0,1,1,0,1},[][]int{{2,1,0},{1,4}})) // [2]
    // Example 2:
    // Input: colors = [0,0,1,0,1,1], queries = [[1,3],[2,3,0],[1,5]]
    // Output: [2,0]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-35-50.png">
    // First query:
    // Count of the alternating groups with size 3:
    // <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-from-2024-06-03-20-36-40.png">
    // Second query: colors will not change.
    // Third query: There is no alternating group with size 5.
    fmt.Println(numberOfAlternatingGroups([]int{0,0,1,0,1,1},[][]int{{1,3},{2,3,0},{1,5}})) // [2,0]
}