package main

// 587. Erect the Fence
// You are given an array trees where trees[i] = [xi, yi] represents the location of a tree in the garden.

// Fence the entire garden using the minimum length of rope, as it is expensive. 
// The garden is well-fenced only if all the trees are enclosed.

// Return the coordinates of trees that are exactly located on the fence perimeter. 
// You may return the answer in any order.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/24/erect2-plane.jpg" />
// Input: trees = [[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]
// Output: [[1,1],[2,0],[4,2],[3,3],[2,4]]
// Explanation: All the trees will be on the perimeter of the fence except the tree at [2, 2], which will be inside the fence.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/24/erect1-plane.jpg" />
// Input: trees = [[1,2],[2,2],[4,2]]
// Output: [[4,2],[2,2],[1,2]]
// Explanation: The fence forms a line that passes through all the trees.

// Constraints:
//     1 <= trees.length <= 3000
//     trees[i].length == 2
//     0 <= xi, yi <= 100
//     All the given positions are unique.

import "fmt"
import "sort"

// Jarvis 算法
func outerTrees(trees [][]int) [][]int {
    res, leftMost, n := [][]int{}, 0, len(trees)
    if n < 4 {
        return trees
    }
    for i, tree := range trees {
        if tree[0] < trees[leftMost][0] || (tree[0] == trees[leftMost][0] && tree[1] < trees[leftMost][1]) {
            leftMost = i
        }
    }
    cross := func (p, q, r []int) int {
        return (q[0]-p[0])*(r[1]-q[1]) - (q[1]-p[1])*(r[0]-q[0])
    }
    visited, p := make([]bool, n), leftMost
    for {
        q := (p + 1) % n
        for r, tree := range trees {
            if cross(trees[p], trees[q], tree) < 0 { // 如果 r 在 pq 的右侧，则 q = r
                q = r
            }
        }
        for i, b := range visited { // 是否存在点 i, 使得 p q i 在同一条直线上
            if !b && i != p && i != q && cross(trees[p], trees[q], trees[i]) == 0 {
                res = append(res, trees[i])
                visited[i] = true
            }
        }
        if !visited[q] {
            res = append(res, trees[q])
            visited[q] = true
        }
        p = q
        if p == leftMost {
            break
        }
    }
    return res
}

// Graham 算法
func outerTrees1(trees [][]int) [][]int {
    n, bottom := len(trees), 0
    if n < 4 {
        return trees
    }
    for i, tree := range trees { // 找到 y 最小的点 bottom
        if tree[1] < trees[bottom][1] {
            bottom = i
        }
    }
    cross := func (p, q, r []int) int {
        return (q[0]-p[0])*(r[1]-q[1]) - (q[1]-p[1])*(r[0]-q[0])
    }
    distance := func (p, q []int) int {
        return (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
    }
    trees[bottom], trees[0] = trees[0], trees[bottom]
    tr := trees[1:]
    sort.Slice(tr, func(i, j int) bool { // 以 bottom 原点，按照极坐标的角度大小进行排序
        a, b := tr[i], tr[j]
        diff := cross(trees[0], a, b)
        return diff > 0 || diff == 0 && distance(trees[0], a) < distance(trees[0], b)
    })
    // 对于凸包最后且在同一条直线的元素按照距离从大到小进行排序
    r := n - 1
    for r >= 0 && cross(trees[0], trees[n-1], trees[r]) == 0 {
        r--
    }
    for l, h := r+1, n-1; l < h; l++ {
        trees[l], trees[h] = trees[h], trees[l]
        h--
    }
    stack := []int{0, 1}
    for i := 2; i < n; i++ {
        // 如果当前元素与栈顶的两个元素构成的向量顺时针旋转，则弹出栈顶元素
        for len(stack) > 1 && cross(trees[stack[len(stack)-2]], trees[stack[len(stack)-1]], trees[i]) < 0 {
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    res := make([][]int, len(stack))
    for i, index := range stack {
        res[i] = trees[index]
    }
    return res
}

// Andrew 算法 单调链
func outerTrees2(trees [][]int) [][]int {
    n := len(trees)
    if n < 4 {
        return trees
    }
    cross := func (p, q, r []int) int {
        return (q[0]-p[0])*(r[1]-q[1]) - (q[1]-p[1])*(r[0]-q[0])
    }
    // 按照 x 从小到大排序，如果 x 相同，则按照 y 从小到大排序
    sort.Slice(trees, func(i, j int) bool { a, b := trees[i], trees[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })
    hull := []int{0} // hull[0] 需要入栈两次，不标记
    used := make([]bool, n)
    // 求凸包的下半部分
    for i := 1; i < n; i++ {
        for len(hull) > 1 && cross(trees[hull[len(hull)-2]], trees[hull[len(hull)-1]], trees[i]) < 0 {
            used[hull[len(hull)-1]] = false
            hull = hull[:len(hull)-1]
        }
        used[i] = true
        hull = append(hull, i)
    }
    // 求凸包的上半部分
    m := len(hull)
    for i := n - 2; i >= 0; i-- {
        if !used[i] {
            for len(hull) > m && cross(trees[hull[len(hull)-2]], trees[hull[len(hull)-1]], trees[i]) < 0 {
                used[hull[len(hull)-1]] = false
                hull = hull[:len(hull)-1]
            }
            used[i] = true
            hull = append(hull, i)
        }
    }
    // hull[0] 同时参与凸包的上半部分检测，因此需去掉重复的 hull[0]
    hull = hull[:len(hull)-1]
    res := make([][]int, len(hull))
    for i, index := range hull {
        res[i] = trees[index]
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/24/erect2-plane.jpg" />
    // Input: trees = [[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]
    // Output: [[1,1],[2,0],[4,2],[3,3],[2,4]]
    // Explanation: All the trees will be on the perimeter of the fence except the tree at [2, 2], which will be inside the fence.
    fmt.Println(outerTrees([][]int{{1,1},{2,2},{2,0},{2,4},{3,3},{4,2}})) // [[1,1],[2,0],[4,2],[3,3],[2,4]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/04/24/erect1-plane.jpg" />
    // Input: trees = [[1,2],[2,2],[4,2]]
    // Output: [[4,2],[2,2],[1,2]]
    // Explanation: The fence forms a line that passes through all the trees.
    fmt.Println(outerTrees([][]int{{1,2},{2,2},{4,2}})) // [[4,2],[2,2],[1,2]]

    fmt.Println(outerTrees1([][]int{{1,1},{2,2},{2,0},{2,4},{3,3},{4,2}})) // [[1,1],[2,0],[4,2],[3,3],[2,4]]
    fmt.Println(outerTrees1([][]int{{1,2},{2,2},{4,2}})) // [[4,2],[2,2],[1,2]]

    fmt.Println(outerTrees2([][]int{{1,1},{2,2},{2,0},{2,4},{3,3},{4,2}})) // [[1,1],[2,0],[4,2],[3,3],[2,4]]
    fmt.Println(outerTrees2([][]int{{1,2},{2,2},{4,2}})) // [[4,2],[2,2],[1,2]]
}