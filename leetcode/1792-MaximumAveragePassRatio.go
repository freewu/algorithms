package main

// 1792. Maximum Average Pass Ratio
// There is a school that has classes of students and each class will be having a final exam. 
// You are given a 2D integer array classes, where classes[i] = [passi, totali]. 
// You know beforehand that in the ith class, there are totali total students, but only passi number of students will pass the exam.

// You are also given an integer extraStudents. 
// There are another extraStudents brilliant students that are guaranteed to pass the exam of any class they are assigned to. 
// You want to assign each of the extraStudents students to a class in a way that maximizes the average pass ratio across all the classes.

// The pass ratio of a class is equal to the number of students of the class that will pass the exam divided by the total number of students of the class. 
// The average pass ratio is the sum of pass ratios of all the classes divided by the number of the classes.

// Return the maximum possible average pass ratio after assigning the extraStudents students. 
// Answers within 10^-5 of the actual answer will be accepted.

// Example 1:
// Input: classes = [[1,2],[3,5],[2,2]], extraStudents = 2
// Output: 0.78333
// Explanation: You can assign the two extra students to the first class. The average pass ratio will be equal to (3/4 + 3/5 + 2/2) / 3 = 0.78333.

// Example 2:
// Input: classes = [[2,4],[3,9],[4,5],[2,10]], extraStudents = 4
// Output: 0.53485

// Constraints:
//     1 <= classes.length <= 10^5
//     classes[i].length == 2
//     1 <= passi <= totali <= 10^5
//     1 <= extraStudents <= 10^5

import "fmt"
import "container/heap"

type MaxHeap [][2]int

func (h MaxHeap) Len() int      { return len(h) }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i int, j int) bool {
    return (float64(h[i][0] + 1) / float64(h[i][1] + 1) - float64(h[i][0]) / float64(h[i][1])) > 
           (float64(h[j][0] + 1) / float64(h[j][1] + 1) - float64(h[j][0]) / float64(h[j][1]))
}
func (h *MaxHeap) Push(a interface{}) { *h = append(*h, a.([2]int)) }
func (h *MaxHeap) Pop() interface{} {
    l := len(*h)
    res := (*h)[l - 1]
    *h = (*h)[:l - 1]
    return res
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
    mxh := &MaxHeap{}
    for _, v := range classes {
        heap.Push(mxh, [2]int{ v[0], v[1] })
    }
    for extraStudents > 0 {
        v := heap.Pop(mxh).([2]int)
        extraStudents--
        v[0]++
        v[1]++
        heap.Push(mxh, v)
    }
    sum := float64(0)
    for _, v := range *mxh {
        sum += float64(v[0]) / float64(v[1])
    }
    return sum / float64(len(classes))
}

var g [][]int
func maxAverageRatio1(classes [][]int, extraStudents int) float64 {
    g = classes
    h := Heap{}
    for i := range classes {
        h.Push(i)
    }
    for i := 0; i < extraStudents; i++ {
        v := h.Pop()
        classes[v][1]++
        classes[v][0]++
        h.Push(v)
    }
    sum := 0.
    for _, v := range classes {
        sum += float64(v[0]) / float64(v[1])
    }
    return sum / float64(len(classes))
}


type Heap struct {
    data []int
}

func (h *Heap) Len() int {
    return len(h.data)
}

func (h *Heap) Less(i, j int) bool {
    a1, b1 := calc(h.data[i])
    a2, b2 := calc(h.data[j])
    return a1 * b2 > b1 * a2
}

func calc(i int) (int, int) {
    a, b := g[i][0], g[i][1]
    return (b-a), (b+1)*b
}

func (h *Heap) Swap(i, j int) {
    h.data[i], h.data[j] = h.data[j], h.data[i]
}

/*
func (h *Heap) Push(x interface{}) {
    h.data = append(h.data, x.(int))
}

func (h *Heap) Pop() interface{} {
    l := h.Len()
    v := h.data[l-1]
    h.data = h.data[:l-1]
    return v
}
*/

func (h *Heap) Push(x int) {
    h.data = append(h.data, x)
    h.filterUp(h.Len() - 1)
}

func (h *Heap) Top() int {
    return h.data[0]
}

func (h *Heap) Pop() int {
    l := h.Len()
    h.Swap(0, l-1)
    v := h.data[l-1]
    h.data = h.data[:l-1]
    h.filterDown(0)
    return v
}

func (h *Heap) filterUp(t int) {
    for t > 0 {
        p := (t - 1) / 2
        if h.Less(t, p) {
            h.Swap(t, p)
            t = p
        } else {
            break
        }
    }
}

func (h *Heap) filterDown(t int) {
    l := h.Len()
    left := 1
    for left < l {
        if left+1 < l && h.Less(left+1, left) {
            left += 1
        }
        if h.Less(left, t) {
            h.Swap(left, t)
            t = left
            left = t*2 + 1
        } else {
            break
        }
    }
}

func main() {
    // Example 1:
    // Input: classes = [[1,2],[3,5],[2,2]], extraStudents = 2
    // Output: 0.78333
    // Explanation: You can assign the two extra students to the first class. The average pass ratio will be equal to (3/4 + 3/5 + 2/2) / 3 = 0.78333.
    fmt.Println(maxAverageRatio([][]int{{1,2},{3,5},{2,2}}, 2)) // 0.78333
    // Example 2:
    // Input: classes = [[2,4],[3,9],[4,5],[2,10]], extraStudents = 4
    // Output: 0.53485
    fmt.Println(maxAverageRatio([][]int{{2,4},{3,9},{4,5},{2,10}}, 4)) // 0.53485

    fmt.Println(maxAverageRatio1([][]int{{1,2},{3,5},{2,2}}, 2)) // 0.78333
    fmt.Println(maxAverageRatio1([][]int{{2,4},{3,9},{4,5},{2,10}}, 4)) // 0.53485
}