package main

// 2617. Minimum Number of Visited Cells in a Grid
// You are given a 0-indexed m x n integer matrix grid. Your initial position is at the top-left cell (0, 0).
// Starting from the cell (i, j), you can move to one of the following cells:
//     Cells (i, k) with j < k <= grid[i][j] + j (rightward movement), or
//     Cells (k, j) with i < k <= grid[i][j] + i (downward movement).

// Return the minimum number of cells you need to visit to reach the bottom-right cell (m - 1, n - 1). 
// If there is no valid path, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/01/25/ex1.png" />
// Input: grid = [[3,4,2,1],[4,2,3,1],[2,1,0,0],[2,4,0,0]]
// Output: 4
// Explanation: The image above shows one of the paths that visits exactly 4 cells.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/01/25/ex2.png" />
// Input: grid = [[3,4,2,1],[4,2,1,1],[2,1,1,0],[3,4,1,0]]
// Output: 3
// Explanation: The image above shows one of the paths that visits exactly 3 cells.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2023/01/26/ex3.png" />
// Input: grid = [[2,1,0],[1,0,0]]
// Output: -1
// Explanation: It can be proven that no path exists.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 10^5
//     1 <= m * n <= 10^5
//     0 <= grid[i][j] < m * n
//     grid[m - 1][n - 1] == 0

import "fmt"
import "sort"
import "container/heap"

// dp
func minimumVisitedCells(grid [][]int) int {
    m,n:=len(grid),len(grid[0])
    dp := make([][]int,m)
    for i:=0;i<m;i++ {
        dp[i]=make([]int,n)
        for j:=0;j<n;j++ {
            dp[i][j]=-1
        }
    }
    dp[m-1][n-1]=1
    k:=0
    for i:=m-2;i>=0;i-- {
        k = grid[i][n-1]+i
        if k>=m-1 {
            dp[i][n-1] = 2
        }else {
            minV:=-1
            for i2:=i+1;i2<=min(m-1,k);i2++ {
                if dp[i2][n-1]>0 {
                    if minV==-1 {
                        minV=dp[i2][n-1]
                    }else {
                        minV=min(minV,dp[i2][n-1])
                    }
                }
            } 
            if minV!=-1 {
                 dp[i][n-1]=minV+1
            }
        }
    }
    for j:=n-2;j>=0;j-- {
        k = grid[m-1][j]+j
        if k>=n-1 {
            dp[m-1][j] = 2
        }else {
              minV := -1
            for j2:=j+1;j2<=min(n-1,k);j2++ {
                if dp[m-1][j2]>0 {
                    if minV==-1 {
                        minV=dp[m-1][j2]
                    }else {
                        minV=min(minV,dp[m-1][j2])
                    }
                }
            }
              if minV!=-1 {
                 dp[m-1][j]=minV+1
            }
        }
    }
    for i:=m-2;i>=0;i-- {
        for j:=n-2;j>=0;j-- {
            minV := -1
            for j2:=j+1;j2<=min(n-1,grid[i][j]+j);j2++ {
                if dp[i][j2]>0 {
                    if minV==-1 {
                        minV=dp[i][j2]
                    }else {
                        minV=min(minV,dp[i][j2])
                    }
                }
            }
            for i2:=i+1;i2<=min(m-1,grid[i][j]+i);i2++ {
                if dp[i2][j]>0 {
                    if minV==-1 {
                        minV=dp[i2][j]
                    }else {
                        minV=min(minV,dp[i2][j])
                    }
                }
            }
            if minV!=-1 {
                 dp[i][j]=minV+1
            }
           
        }
    }
    return dp[0][0]
}

// dp
func minimumVisitedCells1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[m-1][n-1] = 1
	colAsc := make([][]int, n)
	for i := m - 1; i >= 0; i-- {
		var asc []int
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				asc = append(asc, j)
				colAsc[j] = append(colAsc[j], i)
				continue
			}
			p1 := sort.Search(len(asc), func(_i int) bool {
				return asc[_i] <= j+grid[i][j]
			})
			p2 := sort.Search(len(colAsc[j]), func(_i int) bool {
				return colAsc[j][_i] <= i+grid[i][j]
			})
			if p1 == len(asc) && p2 == len(colAsc[j]) {
				continue
			}
			var p int
			if p2 == len(colAsc[j]) {
				p = dp[i][asc[p1]]
			} else if p1 == len(asc) {
				p = dp[colAsc[j][p2]][j]
			} else {
				p = min(dp[i][asc[p1]], dp[colAsc[j][p2]][j])
			}
			dp[i][j] = p + 1
			for len(asc) > 0 && dp[i][asc[len(asc)-1]] >= dp[i][j] {
				asc = asc[:len(asc)-1]
			}
			asc = append(asc, j)
			for len(colAsc[j]) > 0 && dp[colAsc[j][len(colAsc[j])-1]][j] >= dp[i][j] {
				colAsc[j] = colAsc[j][:len(colAsc[j])-1]
			}
			colAsc[j] = append(colAsc[j], i)
			//fmt.Printf("dp[%d][%d] = %d, %v, %v\n", i, j, dp[i][j], asc, colAsc[j])
		}
	}
	return dp[0][0]
}

// PriorityQueue
func minimumVisitedCells2(grid [][]int) int {
    m, n := len(grid), len(grid[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[0][0] = 1
	row := make([]PriorityQueue, m)
    col := make([]PriorityQueue, n)

	update := func(x *int, y int) {
		if *x == -1 || y < *x {
			*x = y
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for len(row[i]) > 0 && row[i][0].second + grid[i][row[i][0].second] < j {
				heap.Pop(&row[i])
			}
			if len(row[i]) > 0 {
				update(&dist[i][j], dist[i][row[i][0].second] + 1)
			}

			for len(col[j]) > 0 && col[j][0].second + grid[col[j][0].second][j] < i {
				heap.Pop(&col[j])
			}
			if len(col[j]) > 0 {
				update(&dist[i][j], dist[col[j][0].second][j] + 1)
			}
			if dist[i][j] != -1 {
				heap.Push(&row[i], Pair{dist[i][j], j})
				heap.Push(&col[j], Pair{dist[i][j], i})
			}
		}
	}
	return dist[m - 1][n - 1]
}

type Pair struct {
    first int
    second int
}

type PriorityQueue []Pair

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].first < pq[j].first
}

func (pq *PriorityQueue) Push(x any) {
    *pq = append(*pq, x.(Pair))
}

func (pq *PriorityQueue) Pop() any {
    n := len(*pq)
    x := (*pq)[n - 1]
    *pq = (*pq)[:n - 1]
    return x
}

func main() {
    // Explanation: The image above shows one of the paths that visits exactly 4 cells.
    fmt.Println(minimumVisitedCells(
        [][]int{
            []int{3,4,2,1},
            []int{4,2,3,1},
            []int{2,1,0,0},
            []int{2,4,0,0},
        },
    )) // 4

    // Explanation: The image above shows one of the paths that visits exactly 3 cells.
    fmt.Println(minimumVisitedCells(
        [][]int{
            []int{3,4,2,1},
            []int{4,2,1,1},
            []int{2,1,1,0},
            []int{3,4,1,0},
        },
    )) // 3

    // Explanation: The image above shows one of the paths that visits exactly 3 cells.
    fmt.Println(minimumVisitedCells(
        [][]int{
            []int{2,1,0},
            []int{1,0,0},
        },
    )) // -1

    fmt.Println(minimumVisitedCells1(
        [][]int{
            []int{3,4,2,1},
            []int{4,2,3,1},
            []int{2,1,0,0},
            []int{2,4,0,0},
        },
    )) // 4
    fmt.Println(minimumVisitedCells1(
        [][]int{
            []int{3,4,2,1},
            []int{4,2,1,1},
            []int{2,1,1,0},
            []int{3,4,1,0},
        },
    )) // 3
    fmt.Println(minimumVisitedCells1(
        [][]int{
            []int{2,1,0},
            []int{1,0,0},
        },
    )) // -1

    fmt.Println(minimumVisitedCells2(
        [][]int{
            []int{3,4,2,1},
            []int{4,2,3,1},
            []int{2,1,0,0},
            []int{2,4,0,0},
        },
    )) // 4
    fmt.Println(minimumVisitedCells2(
        [][]int{
            []int{3,4,2,1},
            []int{4,2,1,1},
            []int{2,1,1,0},
            []int{3,4,1,0},
        },
    )) // 3
    fmt.Println(minimumVisitedCells2(
        [][]int{
            []int{2,1,0},
            []int{1,0,0},
        },
    )) // -1

}