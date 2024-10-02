package main

// 1499. Max Value of Equation
// You are given an array points containing the coordinates of points on a 2D plane, 
// sorted by the x-values, where points[i] = [xi, yi] such that xi < xj for all 1 <= i < j <= points.length. 
// You are also given an integer k.

// Return the maximum value of the equation yi + yj + |xi - xj| where |xi - xj| <= k and 1 <= i < j <= points.length.

// It is guaranteed that there exists at least one pair of points that satisfy the constraint |xi - xj| <= k.

// Example 1:
// Input: points = [[1,3],[2,0],[5,10],[6,-10]], k = 1
// Output: 4
// Explanation: The first two points satisfy the condition |xi - xj| <= 1 and if we calculate the equation we get 3 + 0 + |1 - 2| = 4. Third and fourth points also satisfy the condition and give a value of 10 + -10 + |5 - 6| = 1.
// No other pairs satisfy the condition, so we return the max of 4 and 1.

// Example 2:
// Input: points = [[0,0],[3,0],[9,2]], k = 3
// Output: 3
// Explanation: Only the first two points have an absolute difference of 3 or less in the x-values, and give the value of 0 + 0 + |0 - 3| = 3.

// Constraints:
//     2 <= points.length <= 10^5
//     points[i].length == 2
//     -10^8 <= xi, yi <= 10^8
//     0 <= k <= 2 * 10^8
//     xi < xj for all 1 <= i < j <= points.length
//     xi form a strictly increasing sequence.

import "fmt"
import "container/heap"

func findMaxValueOfEquation(points [][]int, k int) int {
    pq := &PriorityQueue{}
    heap.Init(pq)
    res := -1 << 31
    for _, v := range points {
        p2 := Point{v[0], v[1]}
        for pq.Len() != 0 {
            p1 := heap.Pop(pq).(Point)
            if p2.x - p1.x <= k {
                equation := p1.y + p2.y + p2.x - p1.x
                if res < equation {
                    res = equation
                }
                heap.Push(pq, p1)
                break
            }
        }
        heap.Push(pq, p2)
    }
    return res
}

type Point struct { x, y int }
type PriorityQueue []Point
func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].y - pq[i].x > pq[j].y - pq[j].x }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(Point))
}
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    x := old[n-1]
    *pq = old[0 : n-1]
    return x
}

func findMaxValueOfEquation1(points [][]int, k int) int {
    type Pair struct {
        X   int
        Y_X int // y-x
    }
    res, stack := -1 << 31, make([]Pair, 0)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range points {
        x, y := v[0], v[1]
        for len(stack) > 0 && x - stack[0].X > k {
            stack = stack[1:]
        }
        if len(stack) > 0 {
            res = max(res, x + y + stack[0].Y_X)
        }
        for len(stack) > 0 && stack[len(stack)-1].Y_X <= y-x {
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, Pair{X: x, Y_X: y - x})
    }
    return res
}

func findMaxValueOfEquation2(points [][]int, k int) int {
    res, queue := -1 << 31, [][]int{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, p := range points {
        for len(queue) > 0 && p[0] - queue[0][0] > k {
            queue = queue[1:]
        }
        if len(queue) > 0 {
            res = max(res, p[0] - queue[0][0] + p[1] + queue[0][1])
        }
        for len(queue) > 0 && queue[len(queue) - 1][1] - queue[len(queue)-1][0] < p[1] - p[0] {
            queue = queue[:len(queue) - 1]
        }
        queue = append(queue, p)
    }
    return res
}

func main() {
    // Example 1:
    // Input: points = [[1,3],[2,0],[5,10],[6,-10]], k = 1
    // Output: 4
    // Explanation: The first two points satisfy the condition |xi - xj| <= 1 and if we calculate the equation we get 3 + 0 + |1 - 2| = 4. Third and fourth points also satisfy the condition and give a value of 10 + -10 + |5 - 6| = 1.
    // No other pairs satisfy the condition, so we return the max of 4 and 1.
    fmt.Println(findMaxValueOfEquation([][]int{{1,3},{2,0},{5,10},{6,-10}}, 1)) // 4
    // Example 2:
    // Input: points = [[0,0],[3,0],[9,2]], k = 3
    // Output: 3
    // Explanation: Only the first two points have an absolute difference of 3 or less in the x-values, and give the value of 0 + 0 + |0 - 3| = 3.
    fmt.Println(findMaxValueOfEquation([][]int{{0,0},{3,0},{9,2}}, 3)) // 3

    fmt.Println(findMaxValueOfEquation([][]int{{-19,9},{-15,-19},{-5,-8}}, 10)) // -6

    fmt.Println(findMaxValueOfEquation1([][]int{{1,3},{2,0},{5,10},{6,-10}}, 1)) // 4
    fmt.Println(findMaxValueOfEquation1([][]int{{0,0},{3,0},{9,2}}, 3)) // 3
    fmt.Println(findMaxValueOfEquation1([][]int{{-19,9},{-15,-19},{-5,-8}}, 10)) // -6

    fmt.Println(findMaxValueOfEquation2([][]int{{1,3},{2,0},{5,10},{6,-10}}, 1)) // 4
    fmt.Println(findMaxValueOfEquation2([][]int{{0,0},{3,0},{9,2}}, 3)) // 3
    fmt.Println(findMaxValueOfEquation2([][]int{{-19,9},{-15,-19},{-5,-8}}, 10)) // -6
}