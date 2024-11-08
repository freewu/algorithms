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
}