package main

// 1642. Furthest Building You Can Reach
// You are given an integer array heights representing the heights of buildings, some bricks, and some ladders.
// You start your journey from building 0 and move to the next building by possibly using bricks or ladders.
// While moving from building i to building i+1 (0-indexed),
// If the current building's height is greater than or equal to the next building's height, you do not need a ladder or bricks.
// If the current building's height is less than the next building's height, you can either use one ladder or (h[i+1] - h[i]) bricks.
// Return the furthest building index (0-indexed) you can reach if you use the given ladders and bricks optimally.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/27/q4.gif"/>
// Input: heights = [4,2,7,6,9,14,12], bricks = 5, ladders = 1
// Output: 4
// Explanation: Starting at building 0, you can follow these steps:
// - Go to building 1 without using ladders nor bricks since 4 >= 2.
// - Go to building 2 using 5 bricks. You must use either bricks or ladders because 2 < 7.
// - Go to building 3 without using ladders nor bricks since 7 >= 6.
// - Go to building 4 using your only ladder. You must use either bricks or ladders because 6 < 9.
// It is impossible to go beyond building 4 because you do not have any more bricks or ladders.

// Example 2:
// Input: heights = [4,12,2,7,3,18,20,3,19], bricks = 10, ladders = 2
// Output: 7

// Example 3:
// Input: heights = [14,3,19,3], bricks = 17, ladders = 0
// Output: 3
 
// Constraints:
//         1 <= heights.length <= 10^5
//         1 <= heights[i] <= 10^6
//         0 <= bricks <= 10^9
//         0 <= ladders <= heights.length

import "fmt"

import (
	"container/heap"
)
import "sort"

type HeightDiffPQ []int

func (pq HeightDiffPQ) Len() int            { return len(pq) }
func (pq HeightDiffPQ) Less(i, j int) bool  { return pq[i] < pq[j] }
func (pq HeightDiffPQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *HeightDiffPQ) Push(x interface{}) { *pq = append(*pq, x.(int)) }
func (pq *HeightDiffPQ) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return x
}

func furthestBuilding(heights []int, bricks int, ladder int) int {
	usedLadder := &HeightDiffPQ{}
	for i := 1; i < len(heights); i++ {
        // 高度一样不需要使用 bricks 或 ladder
		needbricks := heights[i] - heights[i-1]
		if needbricks < 0 {
			continue
		}
        // 维护一个长度为梯子个数的最小堆
		if ladder > 0 {
			heap.Push(usedLadder, needbricks)
			ladder--
		} else {
            // 当队列中元素超过梯子个数，便将队首最小值出队，出队的这个楼与楼的差距用砖头填补
			if len(*usedLadder) > 0 && needbricks > (*usedLadder)[0] {
				needbricks, (*usedLadder)[0] = (*usedLadder)[0], needbricks
				heap.Fix(usedLadder, 0)
			}
            // 所有砖头用完了，即是可以到达的最远楼号
			if bricks -= needbricks; bricks < 0 {
				return i - 1
			}
		}
	}
	return len(heights) - 1
}

// best solution
func furthestBuilding1(heights []int, bricks int, ladders int) int {
	q := hp{}
	n := len(heights)
	for i, a := range heights[:n-1] {
		b := heights[i+1]
		d := b - a
		if d > 0 {
			heap.Push(&q, d)
			if q.Len() > ladders {
				bricks -= heap.Pop(&q).(int)
				if bricks < 0 {
					return i
				}
			}
		}
	}
	return n - 1
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func main() {
    fmt.Println(furthestBuilding([]int{4,2,7,6,9,14,12},5,1)) // 4
    fmt.Println(furthestBuilding([]int{4,12,2,7,3,18,20,3,19},10,2)) // 7
    fmt.Println(furthestBuilding([]int{14,3,19,3},17,0)) // 3

    fmt.Println(furthestBuilding1([]int{4,2,7,6,9,14,12},5,1)) // 4
    fmt.Println(furthestBuilding1([]int{4,12,2,7,3,18,20,3,19},10,2)) // 7
    fmt.Println(furthestBuilding1([]int{14,3,19,3},17,0)) // 3
}