package main

// 1057. Campus Bikes
// On a campus represented on the X-Y plane, there are n workers and m bikes, with n <= m.

// You are given an array workers of length n where workers[i] = [xi, yi] is the position of the ith worker. 
// You are also given an array bikes of length m where bikes[j] = [xj, yj] is the position of the jth bike. 
// All the given positions are unique.

// Assign a bike to each worker. Among the available bikes and workers, 
// we choose the (workeri, bikej) pair with the shortest Manhattan distance between each other and assign the bike to that worker.

// If there are multiple (workeri, bikej) pairs with the same shortest Manhattan distance, we choose the pair with the smallest worker index. 
// If there are multiple ways to do that, we choose the pair with the smallest bike index. Repeat this process until there are no available workers.

// Return an array answer of length n, where answer[i] is the index (0-indexed) of the bike that the ith worker is assigned to.
// The Manhattan distance between two points p1 and p2 is Manhattan(p1, p2) = |p1.x - p2.x| + |p1.y - p2.y|.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/03/06/1261_example_1_v2.png" />
// Input: workers = [[0,0],[2,1]], bikes = [[1,2],[3,3]]
// Output: [1,0]
// Explanation: Worker 1 grabs Bike 0 as they are closest (without ties), and Worker 0 is assigned Bike 1. So the output is [1, 0].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/03/06/1261_example_2_v2.png" />
// Input: workers = [[0,0],[1,1],[2,0]], bikes = [[1,0],[2,2],[2,1]]
// Output: [0,2,1]
// Explanation: Worker 0 grabs Bike 0 at first. Worker 1 and Worker 2 share the same distance to Bike 2, thus Worker 1 is assigned to Bike 2, and Worker 2 will take Bike 1. So the output is [0,2,1].
 
// Constraints:
//     n == workers.length
//     m == bikes.length
//     1 <= n <= m <= 1000
//     workers[i].length == bikes[j].length == 2
//     0 <= xi, yi < 1000
//     0 <= xj, yj < 1000
//     All worker and bike locations are unique.

import "fmt"

func assignBikes(workers [][]int, bikes [][]int) []int {
    list := make([][3]int, 0,len(workers) * len(bikes)) 
    vW, vB := make([]bool, len(workers)), make([]bool, len(bikes))  // 工人已分配的标记, 车已分配的标记
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i, w := range workers { // 找到所有组合
        for j, b := range bikes {
            list = append(list, [3]int{i, j, abs(w[0] - b[0]) + abs(w[1] - b[1])})
        }
    }
    less := func (a,b [3]int) bool{
        if a[2]==b[2]{
            if a[0]==b[0]{
                return a[1]<b[1]
            }
            return a[0]<b[0]
        }
        return a[2]<b[2]
    }
    var quick func(list [][3]int, left,right int)
    quick = func(list [][3]int, left,right int) {
        l,r, p := left, right, (left + right) >> 1
        pv := list[p]
        for l < r {
            for l <= p && less(list[l], pv) {
                l++
            }
            if l <= p {
                list[p] = list[l]
                p = l
            }
            for r >= p && less(pv, list[r]) {
                r--
            }
            if r >= p {
                list[p] = list[r]
                p = r
            }
        }
        list[p] = pv
        if left < p {
            quick(list, left, p-1)
        }
        if right > p {
            quick(list, p+1, right)
        }
    }
    quick(list, 0, len(list)-1) // 快速排序
    res := make([]int, len(workers))
    for _,p := range list { // 依次取出即可
        if vW[p[0]] || vB[p[1]] {
            continue
        }
        res[p[0]] = p[1]
        vW[p[0]], vB[p[1]] = true, true
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/03/06/1261_example_1_v2.png" />
    // Input: workers = [[0,0],[2,1]], bikes = [[1,2],[3,3]]
    // Output: [1,0]
    // Explanation: Worker 1 grabs Bike 0 as they are closest (without ties), and Worker 0 is assigned Bike 1. So the output is [1, 0].
    fmt.Println(assignBikes([][]int{{0,0},{2,1}},[][]int{{1,2},{3,3}})) // [1,0]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/03/06/1261_example_2_v2.png" />
    // Input: workers = [[0,0],[1,1],[2,0]], bikes = [[1,0],[2,2],[2,1]]
    // Output: [0,2,1]
    // Explanation: Worker 0 grabs Bike 0 at first. Worker 1 and Worker 2 share the same distance to Bike 2, thus Worker 1 is assigned to Bike 2, and Worker 2 will take Bike 1. So the output is [0,2,1].
    fmt.Println(assignBikes([][]int{{0,0},{1,1},{2,0}},[][]int{{1,0},{2,2},{2,1}})) // [0,2,1]
}