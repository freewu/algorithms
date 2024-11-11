package main

// 1834. Single-Threaded CPU
// You are given n​​​​​​ tasks labeled from 0 to n - 1 represented by a 2D integer array tasks, 
// where tasks[i] = [enqueueTimei, processingTimei] means that the i​​​​​​th​​​​ task will be available to process at enqueueTimei 
// and will take processingTimei to finish processing.

// You have a single-threaded CPU that can process at most one task at a time and will act in the following way:
//     1. If the CPU is idle and there are no available tasks to process, the CPU remains idle.
//     2. If the CPU is idle and there are available tasks, the CPU will choose the one with the shortest processing time. 
//        If multiple tasks have the same shortest processing time, it will choose the task with the smallest index.
//     3. Once a task is started, the CPU will process the entire task without stopping.
//     4. The CPU can finish a task then start a new one instantly.

// Return the order in which the CPU will process the tasks.

// Example 1:
// Input: tasks = [[1,2],[2,4],[3,2],[4,1]]
// Output: [0,2,3,1]
// Explanation: The events go as follows: 
// - At time = 1, task 0 is available to process. Available tasks = {0}.
// - Also at time = 1, the idle CPU starts processing task 0. Available tasks = {}.
// - At time = 2, task 1 is available to process. Available tasks = {1}.
// - At time = 3, task 2 is available to process. Available tasks = {1, 2}.
// - Also at time = 3, the CPU finishes task 0 and starts processing task 2 as it is the shortest. Available tasks = {1}.
// - At time = 4, task 3 is available to process. Available tasks = {1, 3}.
// - At time = 5, the CPU finishes task 2 and starts processing task 3 as it is the shortest. Available tasks = {1}.
// - At time = 6, the CPU finishes task 3 and starts processing task 1. Available tasks = {}.
// - At time = 10, the CPU finishes task 1 and becomes idle.

// Example 2:
// Input: tasks = [[7,10],[7,12],[7,5],[7,4],[7,2]]
// Output: [4,3,2,0,1]
// Explanation: The events go as follows:
// - At time = 7, all the tasks become available. Available tasks = {0,1,2,3,4}.
// - Also at time = 7, the idle CPU starts processing task 4. Available tasks = {0,1,2,3}.
// - At time = 9, the CPU finishes task 4 and starts processing task 3. Available tasks = {0,1,2}.
// - At time = 13, the CPU finishes task 3 and starts processing task 2. Available tasks = {0,1}.
// - At time = 18, the CPU finishes task 2 and starts processing task 0. Available tasks = {1}.
// - At time = 28, the CPU finishes task 0 and starts processing task 1. Available tasks = {}.
// - At time = 40, the CPU finishes task 1 and becomes idle.

// Constraints:
//     tasks.length == n
//     1 <= n <= 10^5
//     1 <= enqueueTimei, processingTimei <= 10^9

import "fmt"
import "sort"
import "container/heap"
// import "github.com/emirpasic/gods/v2/queues/priorityqueue"

// type task struct {
//     enqueueTime, processingTime, index int
// }

// func getOrder(tasks [][]int) []int {
//     list := make([]task, len(tasks))
//     for i, ints := range tasks {
//         list[i] = task{ints[0], ints[1], i}
//     }
//     slices.SortFunc(list, func(a, b task) int {
//         return a.enqueueTime - b.enqueueTime
//     })
//     pq := priorityqueue.NewWith[task](func(x, y task) int {
//         if x.processingTime != y.processingTime {
//             return x.processingTime - y.processingTime
//         } else {
//             return x.index - y.index
//         }
//     })
//     res := make([]int, 0, len(tasks))
//     for time := 0; len(list) > 0 || !pq.Empty(); {
//         if pq.Empty() {
//             time = max(time, list[0].enqueueTime)
//         }
//         for len(list) > 0 && list[0].enqueueTime <= time {
//             pq.Enqueue(list[0])
//             list = list[1:]
//         }
//         value, _ := pq.Dequeue()
//         res = append(res, value.index)
//         time += value.processingTime
//     }
//     return res
// }

type Vertex struct {
    Position, Time int
}

type PriorityQueue []Vertex

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
    if pq[i].Time == pq[j].Time { return pq[i].Position < pq[j].Position }
    return pq[i].Time < pq[j].Time
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any) { *pq = append(*pq, x.(Vertex)) }
func (pq *PriorityQueue) Pop() any {
    n := len(*pq)
    res := (*pq)[n-1]
    *pq = (*pq)[:n-1]
    return res
}

func getOrder(tasks [][]int) []int {
    for i := range tasks {
        tasks[i] = append(tasks[i], i)
    }
    sort.Slice(tasks, func(i, j int) bool {
        return tasks[i][0] < tasks[j][0]
    })
    res, pq := []int{}, PriorityQueue{}
    index, time := 0, tasks[0][0]
    for index < len(tasks) || len(pq) > 0 {
        for index < len(tasks) && time >= tasks[index][0] {
            heap.Push(&pq, Vertex{ Position:tasks[index][2], Time: tasks[index][1] })
            index++
        }
        if len(pq) > 0 {
            v := heap.Pop(&pq).(Vertex)
            time += v.Time
            res = append(res,  v.Position)
        } else {
            time = tasks[index][0]
        }
    }
    return res
}

type Task struct {
	Index, Enqueue,Processing int
}

type MinHeap []Task

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
    // Sort by processing time, then by index if processing times are equal
    if h[i].Processing == h[j].Processing { return h[i].Index < h[j].Index }
    return h[i].Processing < h[j].Processing
}
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Task)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    res := old[len(old) - 1]
    *h = old[:len(old) - 1]
    return res
}

func getOrder1(tasks [][]int) []int {
    arr := []Task{}
    for i, task := range tasks {
        arr = append(arr, Task{ Index: i, Enqueue: task[0], Processing: task[1], })
    }
    sort.Slice(arr, func(i, j int) bool { // Sort tasks by their enqueue time
        return arr[i].Enqueue < arr[j].Enqueue
    })
    mnh := &MinHeap{}
    heap.Init(mnh)
    res, now, index := []int{}, 0, 0
    for len(res) < len(arr) {
        for index < len(arr) && arr[index].Enqueue <= now { // Add all tasks that have become available by the current time
            heap.Push(mnh, arr[index])
            index++
        }
        if mnh.Len() == 0 { // If no tasks are available, jump to the next available task's enqueue time
            now = arr[index].Enqueue
        } else { // Process the task with the shortest processing time
            task := heap.Pop(mnh).(Task)
            res = append(res, task.Index)
            now += task.Processing
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: tasks = [[1,2],[2,4],[3,2],[4,1]]
    // Output: [0,2,3,1]
    // Explanation: The events go as follows: 
    // - At time = 1, task 0 is available to process. Available tasks = {0}.
    // - Also at time = 1, the idle CPU starts processing task 0. Available tasks = {}.
    // - At time = 2, task 1 is available to process. Available tasks = {1}.
    // - At time = 3, task 2 is available to process. Available tasks = {1, 2}.
    // - Also at time = 3, the CPU finishes task 0 and starts processing task 2 as it is the shortest. Available tasks = {1}.
    // - At time = 4, task 3 is available to process. Available tasks = {1, 3}.
    // - At time = 5, the CPU finishes task 2 and starts processing task 3 as it is the shortest. Available tasks = {1}.
    // - At time = 6, the CPU finishes task 3 and starts processing task 1. Available tasks = {}.
    // - At time = 10, the CPU finishes task 1 and becomes idle.
    fmt.Println(getOrder([][]int{{1,2},{2,4},{3,2},{4,1}})) // [0,2,3,1]
    // Example 2:
    // Input: tasks = [[7,10],[7,12],[7,5],[7,4],[7,2]]
    // Output: [4,3,2,0,1]
    // Explanation: The events go as follows:
    // - At time = 7, all the tasks become available. Available tasks = {0,1,2,3,4}.
    // - Also at time = 7, the idle CPU starts processing task 4. Available tasks = {0,1,2,3}.
    // - At time = 9, the CPU finishes task 4 and starts processing task 3. Available tasks = {0,1,2}.
    // - At time = 13, the CPU finishes task 3 and starts processing task 2. Available tasks = {0,1}.
    // - At time = 18, the CPU finishes task 2 and starts processing task 0. Available tasks = {1}.
    // - At time = 28, the CPU finishes task 0 and starts processing task 1. Available tasks = {}.
    // - At time = 40, the CPU finishes task 1 and becomes idle.
    fmt.Println(getOrder([][]int{{7,10},{7,12},{7,5},{7,4},{7,2}})) // [4,3,2,0,1]

    fmt.Println(getOrder1([][]int{{1,2},{2,4},{3,2},{4,1}})) // [0,2,3,1]
    fmt.Println(getOrder1([][]int{{7,10},{7,12},{7,5},{7,4},{7,2}})) // [4,3,2,0,1]
}