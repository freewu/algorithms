package main

// 447. Number of Boomerangs
// You are given n points in the plane that are all distinct, where points[i] = [xi, yi]. 
// A boomerang is a tuple of points (i, j, k) such that the distance between i and j equals the distance between i and k (the order of the tuple matters).

// Return the number of boomerangs.

// Example 1:
// Input: points = [[0,0],[1,0],[2,0]]
// Output: 2
// Explanation: The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]].

// Example 2:
// Input: points = [[1,1],[2,2],[3,3]]
// Output: 2

// Example 3:
// Input: points = [[1,1]]
// Output: 0
 
// Constraints:
//     n == points.length
//     1 <= n <= 500
//     points[i].length == 2
//     -10^4 <= xi, yi <= 10^4
//     All the points are unique.

import "fmt"

func numberOfBoomerangs(points [][]int) int {
    res := 0
    dis := func (pa, pb []int) int { // 求两点之间的距离
        return (pa[0]-pb[0])*(pa[0]-pb[0]) + (pa[1]-pb[1])*(pa[1]-pb[1])
    }
    for i := 0; i < len(points); i++ {
        record := make(map[int]int, len(points))
        for j := 0; j < len(points); j++ {
            if j != i {
                // 求出两两点之间的距离，然后把这些距离记录在 map 中，key 是距离，value 是这个距离出现了多少
                record[dis(points[i], points[j])]++
            }
        }
        // 遍历 map，把里面距离大于 2 的 key 都拿出来，value 对应的是个数，在这些个数里面任取 2 个点就是解
        // 利用排列组合，C n 2 就可以得到这个距离的结果，最后把这些排列组合的结果累加起来
        for _, v := range record {
            res += v * (v - 1)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: points = [[0,0],[1,0],[2,0]]
    // Output: 2
    // Explanation: The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]].
    fmt.Println(numberOfBoomerangs([][]int{{0,0},{1,0},{2,0}})) // 2
    // Example 2: 
    // Input: points = [[1,1],[2,2],[3,3]]
    // Output: 2
    fmt.Println(numberOfBoomerangs([][]int{{1,1},{2,2},{3,3}})) // 2
    // Example 3:
    // Input: points = [[1,1]]
    // Output: 0
    fmt.Println(numberOfBoomerangs([][]int{{1,1}})) // 0
}