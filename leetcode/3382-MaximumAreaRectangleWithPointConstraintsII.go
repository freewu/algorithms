package main

// 3382. Maximum Area Rectangle With Point Constraints II
// There are n points on an infinite plane. 
// You are given two integer arrays xCoord and yCoord where (xCoord[i], yCoord[i]) represents the coordinates of the ith point.

// Your task is to find the maximum area of a rectangle that:
//     1. Can be formed using four of these points as its corners.
//     2. Does not contain any other point inside or on its border.
//     3. Has its edges parallel to the axes.

// Return the maximum area that you can obtain or -1 if no such rectangle is possible.

// Example 1:
// Input: xCoord = [1,1,3,3], yCoord = [1,3,1,3]
// Output: 4
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/11/02/example1.png" />
// We can make a rectangle with these 4 points as corners and there is no other point that lies inside or on the border. Hence, the maximum possible area would be 4.

// Example 2:
// Input: xCoord = [1,1,3,3,2], yCoord = [1,3,1,3,2]
// Output: -1
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/11/02/example2.png" />
// There is only one rectangle possible is with points [1,1], [1,3], [3,1] and [3,3] but [2,2] will always lie inside it. Hence, returning -1.

// Example 3:
// Input: xCoord = [1,1,3,3,1,3], yCoord = [1,3,1,3,2,2]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/11/02/example3.png" />
// The maximum area rectangle is formed by the points [1,3], [1,2], [3,2], [3,3], which has an area of 2. Additionally, the points [1,1], [1,2], [3,1], [3,2] also form a valid rectangle with the same area.

// Constraints:
//     1 <= xCoord.length == yCoord.length <= 2 * 10^5
//     0 <= xCoord[i], yCoord[i] <= 8 * 10^7
//     All the given points are unique.

import "fmt"
import "sort"
import "maps"
import "slices"

// Time Limit Exceeded 592 / 593 
func maxRectangleArea(xCoord []int, yCoord []int) int64 {
    // 将点按 x 坐标排序
    points := make([][2]int, len(xCoord))
    for i := range xCoord {
        points[i] = [2]int{xCoord[i], yCoord[i]}
    }
    sort.Slice(points, func(i, j int) bool {
        if points[i][0] == points[j][0] {
            return points[i][1] < points[j][1]
        }
        return points[i][0] < points[j][0]
    })
    // 使用 map 存储每个 x 对应的 y 列表
    d := make(map[int][]int)
    for _, point := range points {
        x, y := point[0], point[1]
        d[x] = append(d[x], y)
    }
    // 存储每个 (x, y1) 对应的 (x, y2)
    up := make(map[[2]int][2]int)
    for x, ys := range d {
        for i := 0; i < len(ys)-1; i++ {
            y1, y2 := ys[i], ys[i+1]
            up[[2]int{x, y1}] = [2]int{x, y2}
        }
    }
    // 使用 map 存储 y 对应的 x
    yxmap := make(map[int]int)
    // 使用切片模拟 SortedList
    var sl []int
    res := int64(-1)
    // 遍历排序后的 x 坐标
    sortedX := make([]int, 0, len(d))
    for x := range d {
        sortedX = append(sortedX, x)
    }
    sort.Ints(sortedX)
    contains := func(sl []int, y int) bool { // 检查切片中是否包含某个元素
        pos := sort.SearchInts(sl, y)
        return pos < len(sl) && sl[pos] == y
    }
    insertSorted := func(sl []int, y int) []int { // 在有序切片中插入元素
        pos := sort.SearchInts(sl, y)
        sl = append(sl, 0)
        copy(sl[pos+1:], sl[pos:])
        sl[pos] = y
        return sl
    }
    remove := func(sl []int, pos int) []int { // 从切片中移除元素
        return append(sl[:pos], sl[pos+1:]...)
    }
    for _, x := range sortedX {
        ys := d[x]
        for i := 0; i < len(ys)-1; i++ {
            y1, y2 := ys[i], ys[i+1]
            if contains(sl, y1) {
                prevX := yxmap[y1]
                if val, ok := up[[2]int{prevX, y1}]; ok && val[1] == y2 {
                    area := int64(y2-y1) * int64(x-prevX)
                    if area > res {
                        res = area
                    }
                }
            }
        }
        // 维护 sl 列表
        for _, y := range ys {
            for len(sl) > 0 {
                pos := sort.SearchInts(sl, y)
                if pos == len(sl) || sl[pos] > y {
                    pos--
                }
                if pos >= 0 {
                    remy := sl[pos]
                    remx := yxmap[remy]
                    if val, ok := up[[2]int{remx, remy}]; ok && val[1] >= y {
                        sl = remove(sl, pos)
                    } else {
                        break
                    }
                } else {
                    break
                }
            }
        }
        for _, y := range ys[:len(ys)-1] {
            if !contains(sl, y) {
                sl = insertSorted(sl, y)
            }
            yxmap[y] = x
        }
    }
    return res
}


// 树状数组模板
type Fenwick []int

func (f Fenwick) add(i int) {
    for ; i < len(f); i += i & -i {
        f[i]++
    }
}

// [1,i] 中的元素和
func (f Fenwick) pre(i int) int {
    res := 0
    for ; i > 0; i &= i - 1 {
        res += f[i]
    }
    return res
}

// [l,r] 中的元素和
func (f Fenwick) query(l, r int) int {
    return f.pre(r) - f.pre(l-1)
}

func maxRectangleArea1(xCoord, yCoord []int) int64 {
    xMap, yMap := map[int][]int{}, map[int][]int{}  // 同一列的所有点的纵坐标, 同一行的所有点的横坐标
    for i, x := range xCoord {
        y := yCoord[i]
        xMap[x], yMap[y] = append(xMap[x], y), append(yMap[y], x)
    }
    // 预处理每个点的正下方的点
    type Pair struct{ x, y int }
    below := map[Pair]int{}
    for x, ys := range xMap {
        sort.Ints(ys)
        for i := 1; i < len(ys); i++ {
            below[Pair{x, ys[i]}] = ys[i-1]
        }
    }
    // 预处理每个点的正左边的点
    left := map[Pair]int{}
    for y, xs := range yMap {
        sort.Ints(xs)
        for i := 1; i < len(xs); i++ {
            left[Pair{xs[i], y}] = xs[i-1]
        }
    }
    // 离散化用
    xs := slices.Sorted(maps.Keys(xMap))
    ys := slices.Sorted(maps.Keys(yMap))
    // 收集询问：矩形区域（包括边界）的点的个数
    type Query struct{ x1, x2, y1, y2, area int }
    queries := []Query{}
    // 枚举 (x2,y2) 作为矩形的右上角
    for x2, listY := range xMap {
        for i := 1; i < len(listY); i++ {
            // 计算矩形左下角 (x1,y1)
            y2 := listY[i]
            x1, ok := left[Pair{x2, y2}]
            if !ok { continue }
            y1 := listY[i-1] // (x2,y2) 下面的点（矩形右下角）的纵坐标
            // 矩形右下角的左边的点的横坐标必须是 x1
            if x, ok := left[Pair{x2, y1}]; !ok || x != x1 { continue }
            // 矩形左上角的下边的点的纵坐标必须是 y1
            if y, ok := below[Pair{x1, y2}]; !ok || y != y1 { continue }
            queries = append(queries, Query{
                sort.SearchInts(xs, x1), // 离散化
                sort.SearchInts(xs, x2),
                sort.SearchInts(ys, y1),
                sort.SearchInts(ys, y2),
                (x2 - x1) * (y2 - y1),
            })
        }
    }
    // 离线询问
    type Data struct{ qid, sign, y1, y2 int }
    qs := make([][]Data, len(xs))
    for i, q := range queries {
        if q.x1 > 0 {
            qs[q.x1-1] = append(qs[q.x1-1], Data{i, -1, q.y1, q.y2})
        }
        qs[q.x2] = append(qs[q.x2], Data{i, 1, q.y1, q.y2})
    }
    // 回答询问
    arr := make([]int, len(queries))
    tree := make(Fenwick, len(ys)+1)
    for i, x := range xs {
        // 把横坐标为 x 的所有点都加到树状数组中
        for _, y := range xMap[x] {
            tree.add(sort.SearchInts(ys, y) + 1) // 离散化
        }
        for _, q := range qs[i] {
            // 查询横坐标 <= x（已满足）且纵坐标在 [y1,y2] 中的点的个数
            arr[q.qid] += q.sign * tree.query(q.y1+1, q.y2+1)
        }
    }
    res := -1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range arr {
        if v == 4 { // 矩形区域（包括边界）恰好有 4 个点
            res = max(res, queries[i].area)
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: xCoord = [1,1,3,3], yCoord = [1,3,1,3]
    // Output: 4
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/11/02/example1.png" />
    // We can make a rectangle with these 4 points as corners and there is no other point that lies inside or on the border. Hence, the maximum possible area would be 4.
    fmt.Println(maxRectangleArea([]int{1,1,3,3}, []int{1,3,1,3})) // 4
    // Example 2:
    // Input: xCoord = [1,1,3,3,2], yCoord = [1,3,1,3,2]
    // Output: -1
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/11/02/example2.png" />
    // There is only one rectangle possible is with points [1,1], [1,3], [3,1] and [3,3] but [2,2] will always lie inside it. Hence, returning -1.
    fmt.Println(maxRectangleArea([]int{1,1,3,3,2}, []int{1,3,1,3,2})) // -1
    // Example 3:
    // Input: xCoord = [1,1,3,3,1,3], yCoord = [1,3,1,3,2,2]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/11/02/example3.png" />
    // The maximum area rectangle is formed by the points [1,3], [1,2], [3,2], [3,3], which has an area of 2. Additionally, the points [1,1], [1,2], [3,1], [3,2] also form a valid rectangle with the same area.
    fmt.Println(maxRectangleArea([]int{1,1,3,3,1,3}, []int{1,3,1,3,2,2})) // 2

    fmt.Println(maxRectangleArea1([]int{1,1,3,3}, []int{1,3,1,3})) // 4
    fmt.Println(maxRectangleArea1([]int{1,1,3,3,2}, []int{1,3,1,3,2})) // -1
    fmt.Println(maxRectangleArea1([]int{1,1,3,3,1,3}, []int{1,3,1,3,2,2})) // 2
}