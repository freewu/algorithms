package main

// 2534. Time Taken to Cross the Door
// There are n persons numbered from 0 to n - 1 and a door. 
// Each person can enter or exit through the door once, taking one second.

// You are given a non-decreasing integer array arrival of size n, 
// where arrival[i] is the arrival time of the ith person at the door. 
// You are also given an array state of size n, where state[i] is 0 if person i wants to enter through the door or 1 if they want to exit through the door.

// If two or more persons want to use the door at the same time, they follow the following rules:
//     1. If the door was not used in the previous second, then the person who wants to exit goes first.
//     2. If the door was used in the previous second for entering, the person who wants to enter goes first.
//     3. If the door was used in the previous second for exiting, the person who wants to exit goes first.
//     4. If multiple persons want to go in the same direction, the person with the smallest index goes first.

// Return an array answer of size n where answer[i] is the second at which the ith person crosses the door.

// Note that:
//     Only one person can cross the door at each second.
//     A person may arrive at the door and wait without entering or exiting to follow the mentioned rules.

// Example 1:
// Input: arrival = [0,1,1,2,4], state = [0,1,0,0,1]
// Output: [0,3,1,2,4]
// Explanation: At each second we have the following:
// - At t = 0: Person 0 is the only one who wants to enter, so they just enter through the door.
// - At t = 1: Person 1 wants to exit, and person 2 wants to enter. Since the door was used the previous second for entering, person 2 enters.
// - At t = 2: Person 1 still wants to exit, and person 3 wants to enter. Since the door was used the previous second for entering, person 3 enters.
// - At t = 3: Person 1 is the only one who wants to exit, so they just exit through the door.
// - At t = 4: Person 4 is the only one who wants to exit, so they just exit through the door.

// Example 2:
// Input: arrival = [0,0,0], state = [1,0,1]
// Output: [0,2,1]
// Explanation: At each second we have the following:
// - At t = 0: Person 1 wants to enter while persons 0 and 2 want to exit. Since the door was not used in the previous second, the persons who want to exit get to go first. Since person 0 has a smaller index, they exit first.
// - At t = 1: Person 1 wants to enter, and person 2 wants to exit. Since the door was used in the previous second for exiting, person 2 exits.
// - At t = 2: Person 1 is the only one who wants to enter, so they just enter through the door.

// Constraints:
//     n == arrival.length == state.length
//     1 <= n <= 10^5
//     0 <= arrival[i] <= n
//     arrival is sorted in non-decreasing order.
//     state[i] is either 0 or 1.

import "fmt"
import "container/heap"

type PriorityQueue []int // val: person index

func (q PriorityQueue) Len() int { return len(q) }
func (q PriorityQueue) Less(i, j int) bool { return q[i] < q[j] }
func (q PriorityQueue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q *PriorityQueue) IsEmpty() bool { return q.Len() == 0 }
func (q *PriorityQueue) Push(x interface{}) { *q = append(*q, x.(int)) }
func (q *PriorityQueue) Pop() interface{} { // Pop 移除并返回堆顶元素 (即堆中最小的元素)。
    n := len(*q)
    x := (*q)[n-1] // 最后一个元素
    *q = (*q)[0 : n-1]
    return x
}

// arrival的索引是person idx， arrival[i]是人到达的时间
func timeTaken(arrival []int, state []int) []int {
    const (
        Enter = 0
        Exit  = 1
    )
    n := len(arrival)
    res := make([]int, n)
    queue := [2]*PriorityQueue{}
    queue[Enter], queue[Exit] = &PriorityQueue{}, &PriorityQueue{}
    heap.Init(queue[Enter])
    heap.Init(queue[Exit])
    
    i, curState, curTime, preState, preTime := 0, -1, 0, -1, -1 // 上一秒通过门的状态和时间，一开始无人通过, 都初始化为-1
    for i < n || !queue[Enter].IsEmpty() || !queue[Exit].IsEmpty() {
        for i < n && arrival[i] <= curTime { // 比如arrival = [0,0,0], 即curTime=0时刻同时进来3个人，3个人都入队
            heap.Push(queue[state[i]], i)
            i++
        }
        if queue[Exit].IsEmpty() && queue[Enter].IsEmpty() { // 这一秒curTime没有人，curTime直接重置到下一人来的时间
            curTime = arrival[i]
            continue
        }
        if preTime < 0 || curTime - preTime > 1 { // 确定当前应该通过的状态curState 上一秒无人通过，Exit优先
            if !queue[Exit].IsEmpty() {
                curState = Exit
            } else {
                curState = Enter
            }
        } else { // 上一秒有人通过，则跟随上一秒的状态preState
            if !queue[preState].IsEmpty() {
                curState = preState
            } else {
                curState = preState ^ 1
            }
        }
        // 更新preTime，preState和curTime
        res[heap.Pop(queue[curState]).(int)] = curTime
        preTime, preState = curTime, curState
        curTime++
    }
    return res
}

func timeTaken1(arrival []int, state []int) []int {
    n := len(arrival)
    res, queue := make([]int, n), [2][]int{}
    t, i, start := 0, 0, 1
    for i < n || len(queue[0]) > 0 || len(queue[1]) > 0 {
        for i < n && arrival[i] <= t {
            queue[state[i]] = append(queue[state[i]], i)
            i++
        }
        if len(queue[0]) > 0 && len(queue[1]) > 0 {
            res[queue[start][0]] = t
            queue[start] = queue[start][1:]
        } else if len(queue[0]) > 0 || len(queue[1]) > 0 {
            if len(queue[0]) == 0 {
                start = 1
            } else {
                start = 0
            }
            res[queue[start][0]] = t
            queue[start] = queue[start][1:]
        } else {
            start = 1
        }
        t++
    }
    return res
}

func main() {
    // Example 1:
    // Input: arrival = [0,1,1,2,4], state = [0,1,0,0,1]
    // Output: [0,3,1,2,4]
    // Explanation: At each second we have the following:
    // - At t = 0: Person 0 is the only one who wants to enter, so they just enter through the door.
    // - At t = 1: Person 1 wants to exit, and person 2 wants to enter. Since the door was used the previous second for entering, person 2 enters.
    // - At t = 2: Person 1 still wants to exit, and person 3 wants to enter. Since the door was used the previous second for entering, person 3 enters.
    // - At t = 3: Person 1 is the only one who wants to exit, so they just exit through the door.
    // - At t = 4: Person 4 is the only one who wants to exit, so they just exit through the door.
    fmt.Println(timeTaken([]int{0,1,1,2,4}, []int{0,1,0,0,1})) // [0,3,1,2,4]
    // Example 2:
    // Input: arrival = [0,0,0], state = [1,0,1]
    // Output: [0,2,1]
    // Explanation: At each second we have the following:
    // - At t = 0: Person 1 wants to enter while persons 0 and 2 want to exit. Since the door was not used in the previous second, the persons who want to exit get to go first. Since person 0 has a smaller index, they exit first.
    // - At t = 1: Person 1 wants to enter, and person 2 wants to exit. Since the door was used in the previous second for exiting, person 2 exits.
    // - At t = 2: Person 1 is the only one who wants to enter, so they just enter through the door.
    fmt.Println(timeTaken([]int{0,0,0}, []int{1,0,1})) // [0,2,1]

    fmt.Println(timeTaken1([]int{0,1,1,2,4}, []int{0,1,0,0,1})) // [0,3,1,2,4]
    fmt.Println(timeTaken1([]int{0,0,0}, []int{1,0,1})) // [0,2,1]
}