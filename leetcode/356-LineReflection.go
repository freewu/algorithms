package main

// 356. Line Reflection
// Given n points on a 2D plane, find if there is such a line parallel to the y-axis that reflects the given points symmetrically.

// In other words, answer whether or not if there exists a line that after reflecting all points over the given line, 
// the original points' set is the same as the reflected ones.

// Note that there can be repeated points.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/04/23/356_example_1.PNG" />
// Input: points = [[1,1],[-1,1]]
// Output: true
// Explanation: We can choose the line x = 0.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/04/23/356_example_2.PNG" />
// Input: points = [[1,1],[-1,-1]]
// Output: false
// Explanation: We can't choose a line.

// Constraints:
//     n == points.length
//     1 <= n <= 10^4
//     -10^8 <= points[i][j] <= 10^8

// Follow up: Could you do better than O(n2)?

import "fmt"
import "sort"

func isReflected(points [][]int) bool {
    minX, maxX := points[0][0],points[0][0] // minX：最左边的横坐标，maxX：最右边的横坐标
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(points); i++ {
        minX = min(minX, points[i][0])
        maxX = max(maxX, points[i][0])
    }
    mid := float32(maxX + minX) / 2 // 对称中轴线
    pointsMap := make(map[[2]int]int) // pointsMap对点进行计数，后续用于去重，
    yMap := make(map[int]int) // yMap用于去重后对纵坐标计数
    sumX := float32(0)
    for i, v := range points{
        if  pointsMap[[2]int{v[0], v[1]}]  == 0 { // 如果不是重复点才计算
            sumX += mid - float32(points[i][0])
        }
        if  pointsMap[[2]int{v[0], v[1]}]  == 0 && float32(points[i][0]) != mid { // 如果不是重复点， 且点的横坐标不为mid
            yMap[points[i][1]]++ 
        }
        pointsMap[[2]int{v[0], v[1]}]++
        
    }
    if sumX != float32(0) { // 如果可以轴对称，则sumX必为0
        return false
    }
    for i := 0; i < len(points); i++ {
        if yMap[points[i][1]] % 2 != 0 { // 有一个纵坐标的出现次数是奇数次就返回false
            return false
        } 
    }
    return true 
}

func isReflected1(points [][]int) bool {
    hash := make(map[int]map[int]struct{}, 0)
    for _, point := range points {
        _, ok := hash[point[1]]
        if !ok {
            hash[point[1]] = make(map[int]struct{})
        }
        hash[point[1]][point[0]]=struct{}{}
    }
    first, mirror, nodes := false, float64(0), make([]int, 0, 1024)
    for _, subHash := range hash {
        nodes = nodes[:0]
        for n := range subHash {
            nodes = append(nodes, n)
        }
        sort.Slice(nodes, func(i, j int) bool {
            return nodes[i] < nodes[j]
        })
        i, j, mid := 0, len(nodes) - 1, float64(0)
        for i <= j {
            mid = float64(nodes[j] - nodes[i]) / float64(2) + float64(nodes[i])
            if first == true && mirror != mid {
                return false
            } 
            if first == false {
                mirror = mid
                first = true
            } 
            i++
            j--
        }
    }
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/04/23/356_example_1.PNG" />
    // Input: points = [[1,1],[-1,1]]
    // Output: true
    // Explanation: We can choose the line x = 0.
    fmt.Println(isReflected([][]int{{1,1},{-1,1}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/04/23/356_example_2.PNG" />
    // Input: points = [[1,1],[-1,-1]]
    // Output: false
    // Explanation: We can't choose a line.
    fmt.Println(isReflected([][]int{{1,1},{-1,-1}})) // false

    fmt.Println(isReflected1([][]int{{1,1},{-1,1}})) // true
    fmt.Println(isReflected1([][]int{{1,1},{-1,-1}})) // false
}