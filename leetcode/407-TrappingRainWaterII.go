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

// func trapRainWater(height [][]int) int {
//     res, m, n := 0, len(height), len(height[0])
//     visited := make([][]bool, m)
//     for i := range visited {
//         visited[i] = make([]bool, n)
//     }
//     heap := binaryheap.NewWith(func(a, b interface{}) int {
//         priorityA := a.([3]int)
//         priorityB := b.([3]int)
//         return utils.IntComparator(priorityA[2], priorityB[2])
//     })
//     for i := range height {
//         heap.Push([3]int{i, 0, height[i][0]})
//         heap.Push([3]int{i, n - 1, height[i][n-1]})
//     }
//     for i := range height[0] {
//         heap.Push([3]int{0, i, height[0][i]})
//         heap.Push([3]int{m - 1, i, height[m-1][i]})
//     }
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
//     for !heap.Empty() {
//         top, _ := heap.Pop()
//         cell := top.([3]int)
//         visited[cell[0]][cell[1]] = true
//         for _, dir := range dirs {
//             x := cell[0] + dir[0]
//             y := cell[1] + dir[1]
//             if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] {
//                 continue
//             }
//             visited[x][y] = true
//             res += max(cell[2]-height[x][y], 0)
//             heap.Push([3]int{x, y, max(height[x][y], cell[2])})
//         }
//     }
//     return res
// }

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
}