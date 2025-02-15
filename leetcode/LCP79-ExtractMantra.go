package main

// LCP 79. 提取咒文
// 随着兽群逐渐远去，一座大升降机缓缓的从地下升到了远征队面前。
// 借由这台升降机，他们将能够到达地底的永恒至森。 
// 在升降机的操作台上，是一个由魔法符号组成的矩阵，为了便于辨识，我们用小写字母来表示。 
// matrix[i][j] 表示矩阵第 i 行 j 列的字母。该矩阵上有一个提取装置，可以对所在位置的字母提取。 
// 提取装置初始位于矩阵的左上角 [0,0]，可以通过每次操作移动到上、下、左、右相邻的 1 格位置中。
// 提取装置每次移动或每次提取均记为一次操作。

// 远征队需要按照顺序，从矩阵中逐一取出字母以组成 mantra，才能够成功的启动升降机。
// 请返回他们 最少 需要消耗的操作次数。
// 如果无法完成提取，返回 -1。

// 注意：
//     提取装置可对同一位置的字母重复提取，每次提取一个
//     提取字母时，需按词语顺序依次提取

// 示例 1：
// 输入：matrix = ["sd","ep"], mantra = "speed"
// 输出：10
// 解释：如下图所示矩阵 
// <img src="https://pic.leetcode-cn.com/1646288670-OTlvAl-%E7%9F%A9%E9%98%B5%20(2).gif" />

// 示例 2：
// 输入：matrix = ["abc","daf","geg"]， mantra = "sad"
// 输出：-1
// 解释：矩阵中不存在 s ，无法提取词语

// 提示：
//     0 < matrix.length, matrix[i].length <= 100
//     0 < mantra.length <= 100
//     matrix 和 mantra 仅由小写字母组成

import "fmt"
import "container/list"
import "container/heap"

func extractMantra(matrix []string, mantra string) int {
    type State struct { x, y, steps, index int }
    rows, cols, n := len(matrix), len(matrix[0]), len(mantra)
    // BFS 队列
    queue := list.New()
    queue.PushBack(State{ 0, 0, 0, 0 })
    // 访问记录
    visited := make(map[[3]int]bool)
    visited[[3]int{0, 0, 0}] = true
    directions := []struct{ dx, dy int } { {-1, 0}, {1, 0}, {0, -1}, {0, 1}, }
    for queue.Len() > 0 {
        curr := queue.Remove(queue.Front()).(State)
        // 如果找到 mantra 的最后一个字符
        if curr.index == n - 1 && matrix[curr.x][curr.y] == mantra[curr.index] {
            return curr.steps + 1 // 最后一个字符的选择也算一次操作
        }
        // 当前字符匹配并选择（记一次操作）
        if matrix[curr.x][curr.y] == mantra[curr.index] {
            if !visited[[3]int{curr.x, curr.y, curr.index + 1}] {
                queue.PushBack(State{curr.x, curr.y, curr.steps + 1, curr.index + 1})
                visited[[3]int{curr.x, curr.y, curr.index + 1}] = true
            }
        }
        // 尝试四个方向移动（每次移动算一次操作）
        for _, dir := range directions {
            nx, ny := curr.x + dir.dx, curr.y + dir.dy
            if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
                if !visited[[3]int{ nx, ny, curr.index }] {
                    queue.PushBack(State{ nx, ny, curr.steps + 1, curr.index })
                    visited[[3]int{ nx, ny, curr.index }] = true
                }
            }
        }
    }
    return -1 // 无法完成提取
}

type Item struct {
    x, y, z, d int
}
type PriorityQueue []Item
func (this PriorityQueue) Len() int            {  return len(this) }
func (this PriorityQueue) Less(i,j int) bool   { return this[i].d < this[j].d }
func (this PriorityQueue) Swap(i,j int)        { this[i], this[j] = this[j], this[i] }
func (this *PriorityQueue) Push(v interface{}) { *this=append(*this,v.(Item)) }
func (this *PriorityQueue) Pop() interface{} {
    old := *this
    v := old[len(old) - 1]
    *this = old[:len(old) - 1]
    return v
}

func extractMantra1(matrix []string, mantra string) int {
    pq := PriorityQueue{}
    res, m, n := 1 << 31, len(matrix), len(matrix[0])
    directions := []int{ 0, 1, 0, -1, 0, 0 }
    visited := make([][][]int, m)
    for i := 0; i < m; i++ {
        visited[i] = make([][]int,n)
        for j := 0; j < n; j++ {
            visited[i][j] = make([]int, len(mantra) + 1)
            for k := 0; k <= len(mantra); k++ {
                visited[i][j][k] = 1 << 31
            }
        }
    }
    if mantra[0] == matrix[0][0] {
        heap.Push(&pq, Item{ 0,0,1,0 })
        visited[0][0][1] = 0
    } else {
        heap.Push(&pq, Item{ 0,0,0,0 })
        visited[0][0][0] = 0
    }
    for len(pq) > 0 {
        t := heap.Pop(&pq).(Item)
        if t.d > visited[t.x][t.y][t.z] || t.z == len(mantra) { continue }
        for dir := 0; dir < 5; dir++ {
            x, y := t.x + directions[dir], t.y + directions[dir + 1]
            if x < 0 || y < 0 || x >= m || y >= n { continue }
            z, d := t.z, t.d
            if matrix[x][y] == mantra[z] {
                z++
            }
            if dir < 4 {
                d++
            }
            if visited[x][y][z] > d {
                visited[x][y][z] = d
                heap.Push(&pq, Item{ x,y,z,d })
            }
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            res = min(res, visited[i][j][len(mantra)])
        }
    }
    if res == 1 << 31 {
        return -1
    }
    return res + len(mantra)
}

func main() {
    // 示例 1：
    // 输入：matrix = ["sd","ep"], mantra = "speed"
    // 输出：10
    // 解释：如下图所示矩阵 
    // <img src="https://pic.leetcode-cn.com/1646288670-OTlvAl-%E7%9F%A9%E9%98%B5%20(2).gif" />
    fmt.Println(extractMantra([]string{"sd","ep"}, "speed")) // 10
    // 示例 2：
    // 输入：matrix = ["abc","daf","geg"]， mantra = "sad"
    // 输出：-1
    // 解释：矩阵中不存在 s ，无法提取词语
    fmt.Println(extractMantra([]string{"abc","daf","geg"}, "sad")) // -1

    fmt.Println(extractMantra([]string{"blue","frog","leet", "code"}, "free")) // 7

    fmt.Println(extractMantra1([]string{"sd","ep"}, "speed")) // 10
    fmt.Println(extractMantra1([]string{"abc","daf","geg"}, "sad")) // -1
    fmt.Println(extractMantra1([]string{"blue","frog","leet", "code"}, "free")) // 7
}