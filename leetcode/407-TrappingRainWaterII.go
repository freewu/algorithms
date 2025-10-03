package main

// 407. Trapping Rain Water II
// Given an m x n integer matrix heightMap representing the height of each unit cell in a 2D elevation map, 
// return the volume of water it can trap after raining.

// Example 1:
// <img src= "https://assets.leetcode.com/uploads/2021/04/08/trap1-3d.jpg" />
// Input: heightMap = [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]
// Output: 4
// Explanation: After the rain, water is trapped between the blocks.
// We have two small ponds 1 and 3 units trapped.
// The total volume of water trapped is 4.

// Example 2:
// <img src= "https://assets.leetcode.com/uploads/2021/04/08/trap2-3d.jpg" />
// Input: heightMap = [[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]
// Output: 10
 
// Constraints:
//     m == heightMap.length
//     n == heightMap[i].length
//     1 <= m, n <= 200
//     0 <= heightMap[i][j] <= 2 * 10^4

import "fmt"
import "sync"
import "container/heap"

var (
    psPool = sync.Pool{
        New: func() interface{} {
            return make([]int, 100*1024)
        },
    }
    poolsPool = sync.Pool{
        New: func() interface{} {
            return make([][]int, 10*1024)
        },
    }
)

func trapRainWater(heightMap [][]int) int {
    lenRow := len(heightMap) // 行的长度
    if lenRow == 0 {
        return 0
    }
    lenCol := len(heightMap[0]) // 列的长度
    if lenCol == 0 {
        return 0
    }
    l := lenRow * lenCol
    var ps []int
    var pools [][]int
    if l < 4*1024 {
        ps = make([]int, l)
        pools = make([][]int, lenRow)
    } else {
        ps = psPool.Get().([]int)
        if len(ps) < l {
            ps = make([]int, l)
        }
        defer psPool.Put(ps)
        pools = poolsPool.Get().([][]int)
        if len(pools) < lenRow {
            pools = make([][]int, lenRow)
        }
        defer poolsPool.Put(pools)
    }

    for i, j := 0, 0; i < lenRow; i, j = i+1, j+lenCol {
        pools[i] = ps[j : j+lenCol]
    }

    // 1. 向右下角收敛; 第一行、最后一行、最后一列不需要处理
    pools[0] = heightMap[0]
    for i := 1; i < lenRow-1; i++ {
        line := heightMap[i]
        upPools, curPools := pools[i-1], pools[i]
        curPools[0], curPools[lenCol-1] = line[0], line[lenCol-1] // 开头和结尾都是既定的值

        // 1.1 从左往右
        for j := 1; j < lenCol-1; j++ {
            upPool, leftPool := upPools[j], curPools[j-1]
            curPoint := line[j]

            if leftPool > upPool {
                leftPool = upPool // minPool
            }
            if leftPool > curPoint {
                curPools[j] = leftPool
            } else {
                curPools[j] = curPoint
            }
        }
        // 1.2 从右往左
        for j := lenCol - 2; j >= 0; j-- {
            rightPool := curPools[j+1]
            curPoint := line[j]

            if rightPool < curPools[j] {
                if rightPool > curPoint {
                    curPools[j] = rightPool
                } else {
                    curPools[j] = curPoint
                }
            }
        }
    }

    // 2. 向左上角回溯, 并同时收集 pool 存储量; 第一行、最后一行、最后一列不需要处理
    pools[lenRow-1] = heightMap[lenRow-1]
    for i := lenRow - 2; i > 0; i-- {
        line := heightMap[i]
        lowPools, curPools := pools[i+1], pools[i]

        // 1.1 从左往右
        for j := 1; j < lenCol-1; j++ {
            lowPool, leftPool := lowPools[j], curPools[j-1]
            curPoint := line[j]

            if leftPool > lowPool {
                leftPool = lowPool // minPool
            }
            if leftPool < curPools[j] {
                if leftPool > curPoint {
                    curPools[j] = leftPool
                } else {
                    curPools[j] = curPoint
                }
            }
        }
        // 1.2 从右往左
        for j := lenCol - 2; j >= 0; j-- {
            lowPool, rightPool := lowPools[j], curPools[j+1]
            curPoint, lowPoint := line[j], heightMap[i+1][j]

            if rightPool < curPools[j] {
                if rightPool > curPoint {
                    curPools[j] = rightPool
                } else {
                    curPools[j] = curPoint
                }
            }
            curPool := curPools[j]
            if lowPool > curPool && lowPool > lowPoint {
                // 此时需要回溯
                if curPool < lowPoint {
                    curPool = lowPoint
                }
                lowPools[j] = curPool
                // backtracking(heightMap, pools, i+1, j)
                // i++
                // break
                if i >= 1 {
                    backtracking(heightMap[i:], pools[i:], 1, j)
                    i++
                    break
                }
            }
        }
    }

    sum := 0
    for i := 1; i < lenRow-1; i++ {
        line := heightMap[i]
        curPools := pools[i]
        for j := 1; j < lenCol-1; j++ {
            if curPools[j] < line[j] {
                // log.Printf("i:%d,j:%d", i, j)
                continue
            }
            sum += curPools[j] - line[j]
        }
    }
    return sum
}

// 回溯，根据上下左右4个方向走，如果可以走就递归往前走
func backtracking(heightMap, pools [][]int, x, y int) {
    if x == 0 || y == 0 || x == len(heightMap)-1 || y == len(heightMap[0])-1 {
        return
    }

    cur := pools[x][y]
    up := pools[x-1][y]
    upPoint := heightMap[x-1][y]
    if up > cur && upPoint < up {
        if upPoint < cur {
            pools[x-1][y] = cur
        } else {
            pools[x-1][y] = upPoint
        }
        backtracking(heightMap, pools, x-1, y)
    }
    low := pools[x+1][y]
    lowPoint := heightMap[x+1][y]
    if low > cur && lowPoint < low {
        if lowPoint < cur {
            pools[x+1][y] = cur
        } else {
            pools[x+1][y] = lowPoint
        }
        backtracking(heightMap, pools, x+1, y)
    }
    left := pools[x][y-1]
    leftPoint := heightMap[x][y-1]
    if left > cur && leftPoint < left {
        if leftPoint < cur {
            pools[x][y-1] = cur
        } else {
            pools[x][y-1] = leftPoint
        }
        backtracking(heightMap, pools, x, y-1)
    }
    right := pools[x][y+1]
    rightPoint := heightMap[x][y+1]
    if right > cur && rightPoint < right {
        if rightPoint < cur {
            pools[x][y+1] = cur
        } else {
            pools[x][y+1] = rightPoint
        }
        backtracking(heightMap, pools, x, y+1)
    }
}

// Structure pour notre Priority Queue
type Cell struct {
    height int
    row    int
    col    int
}

// PriorityQueue implémente heap.Interface
type PriorityQueue []Cell

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].height < pq[j].height }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { item := x.(Cell); *pq = append(*pq, item) }
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func trapRainWater1(heightMap [][]int) int {
    if len(heightMap) <= 2 { return 0 }
    res, m, n := 0, len(heightMap), len(heightMap[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    pq := &PriorityQueue{}
    heap.Init(pq)
    // Ajout des bords quand la queue
    // On passe à la fois dans la première ligne et la dernière
    for i := 0; i < n; i++ {
        heap.Push(pq, Cell{ height: heightMap[0][i], row: 0, col: i })
        heap.Push(pq, Cell{ height: heightMap[m-1][i], row: m - 1, col: i })
        visited[0][i] = true
        visited[m-1][i] = true
    }
    // Pareille pour les colonnes
    // On passe à la fois dans la première colonne et la dernière
    for i := 1; i < m-1; i++ {
        heap.Push(pq, Cell{height: heightMap[i][0], row: i, col: 0})
        heap.Push(pq, Cell{height: heightMap[i][n-1], row: i, col: n - 1})
        visited[i][0] = true
        visited[i][n-1] = true
    }
    directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for pq.Len() > 0 { // Parcoure depuis les bords de l'intérieur
        cell := heap.Pop(pq).(Cell)
        // Vérifier les 4 cellules adjacentes
        for _, direction := range directions {
            newRow, newCol := cell.row+direction[0], cell.col+direction[1]
            // Vérifier si la nouvelle position est valide et non visitée
            if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && !visited[newRow][newCol] {
                visited[newRow][newCol] = true
                // Calcule l'eau piégée dans cette cellule
                if cell.height > heightMap[newRow][newCol] {
                    res += cell.height - heightMap[newRow][newCol]
                    // Add la cellule avec la nouvelle valeur (après remplissage d'eau)
                    heap.Push(pq, Cell{height: cell.height, row: newRow, col: newCol})
                } else {
                    // Add la cellule avec sa hauteur originale
                    heap.Push(pq, Cell{height: heightMap[newRow][newCol], row: newRow, col: newCol})
                }
            }
        }
    }
    return res
}


func trapRainWater2(heightMap [][]int) int {
    if len(heightMap) == 0 { return 0 }
    // 广度遍历，从四周向中间扩散
    // 先把水填到最高, 当前节点水小于四周时，代表水会溢出，则四周设为当前直接填水值
    res, maxHeight, m, n := 0, 0, len(heightMap), len(heightMap[0])
    for _, arrv := range heightMap {
        for _, v := range arrv {
            maxHeight = max(maxHeight, v)
        }
    }
    // 预填水
    waterMap := make([][]int, m)
    for i := range waterMap {
        tmp := make([]int, n)
        for j := range tmp {
            tmp[j] = maxHeight
        }
        waterMap[i] = tmp
    }
    type Pair struct {x, y int}
    queue := []Pair{}
    for i, arrv := range heightMap {
        for j, v := range arrv {
            if (i == 0 || i == m - 1 || j == 0 || j == n - 1) && v < waterMap[i][j] { // 边界检测
                waterMap[i][j] = v
                queue = append(queue, Pair{i, j})
            }
        }
    }
    for len(queue) > 0 {
        p := queue[0]
        queue = queue[1:]
        pPlace := []int{-1, 0, 1, 0, -1}
        x, y := p.x, p.y
        for i := 0; i < 4; i++ {
            nx, ny := x + pPlace[i], y + pPlace[i+1]
            if 0 <= nx && nx < m && 0 <= ny && ny < n && waterMap[x][y] < waterMap[nx][ny] && waterMap[nx][ny] > heightMap[nx][ny] {
                waterMap[nx][ny] = max(waterMap[x][y], heightMap[nx][ny])
                queue = append(queue, Pair{nx, ny})
            }
        }
    }
    for i, arrv := range waterMap {
        for j, v := range arrv {
            res += v - heightMap[i][j]
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src= "https://assets.leetcode.com/uploads/2021/04/08/trap1-3d.jpg" />
    // Input: heightMap = [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]
    // Output: 4
    // Explanation: After the rain, water is trapped between the blocks.
    // We have two small ponds 1 and 3 units trapped.
    // The total volume of water trapped is 4.
    fmt.Println(trapRainWater([][]int{{1,4,3,1,3,2},{3,2,1,3,2,4},{2,3,3,2,3,1}})) // 4
    // Example 2:
    // <img src= "https://assets.leetcode.com/uploads/2021/04/08/trap2-3d.jpg" />
    // Input: heightMap = [[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]
    // Output: 10
    fmt.Println(trapRainWater([][]int{{3,3,3,3,3},{3,2,2,2,3},{3,2,1,2,3},{3,2,2,2,3},{3,3,3,3,3}})) // 10

    fmt.Println(trapRainWater1([][]int{{1,4,3,1,3,2},{3,2,1,3,2,4},{2,3,3,2,3,1}})) // 4
    fmt.Println(trapRainWater1([][]int{{3,3,3,3,3},{3,2,2,2,3},{3,2,1,2,3},{3,2,2,2,3},{3,3,3,3,3}})) // 10

    fmt.Println(trapRainWater2([][]int{{1,4,3,1,3,2},{3,2,1,3,2,4},{2,3,3,2,3,1}})) // 4
    fmt.Println(trapRainWater2([][]int{{3,3,3,3,3},{3,2,2,2,3},{3,2,1,2,3},{3,2,2,2,3},{3,3,3,3,3}})) // 10
}