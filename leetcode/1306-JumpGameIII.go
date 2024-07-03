package main

// 1306. Jump Game III
// Given an array of non-negative integers arr, you are initially positioned at start index of the array. 
// When you are at index i, you can jump to i + arr[i] or i - arr[i], check if you can reach any index with value 0.

// Notice that you can not jump outside of the array at any time.

// Example 1:
// Input: arr = [4,2,3,0,3,1,2], start = 5
// Output: true
// Explanation: 
// All possible ways to reach at index 3 with value 0 are: 
// index 5 -> index 4 -> index 1 -> index 3 
// index 5 -> index 6 -> index 4 -> index 1 -> index 3 

// Example 2:
// Input: arr = [4,2,3,0,3,1,2], start = 0
// Output: true 
// Explanation: 
// One possible way to reach at index 3 with value 0 is: 
// index 0 -> index 4 -> index 1 -> index 3

// Example 3:
// Input: arr = [3,0,2,1,2], start = 2
// Output: false
// Explanation: There is no way to reach at index 1 with value 0.

// Constraints:
//     1 <= arr.length <= 5 * 10^4
//     0 <= arr[i] < arr.length
//     0 <= start < arr.length

import "fmt"

func canReach(arr []int, start int) bool {
    visited := make([]bool, len(arr))
    var dfs func (pos int) bool
    dfs = func (pos int) bool {
        if arr[pos] == 0 {
            return true
        }
        flag1, flag2 := false, false
        visited[pos] = true
        next := pos + arr[pos]
        if next < len(arr) && !visited[next] {
            flag1 = dfs(next)
        }
        next = pos - arr[pos]
        if next >= 0 && !visited[next] {
            flag2 = dfs(next)
        }
        return flag1 || flag2
    }
    return dfs(start)    
}

func canReach1(arr []int, start int) bool {
    var dfs func(arr []int, pos int, visited map[int]bool) bool
    dfs = func(arr []int, pos int, visited map[int]bool) bool {
        if pos < 0 || pos >= len(arr) { return false } // 越界
        if _, ok := visited[pos]; ok {  return false } // 已到达过
        if arr[pos] == 0 { return true } // 到达 0
        visited[pos] = true
        return dfs(arr, pos + arr[pos], visited) || dfs(arr, pos - arr[pos], visited)
    }
    return dfs(arr, start, make(map[int]bool))
}

// bfs
func canReach2(arr []int, start int) bool {
    queue, n := []int{ start }, len(arr)
    visited := make([]bool, n)
    visited[start] = true
    for len(queue) > 0 {
        index := queue[0]
        queue = queue[1:]
        if arr[index] == 0 { return true } // 到达 0
        for _, nextIndex := range []int{ index + arr[index], index - arr[index] } { // { i + arr[i], i - arr[i] }
            if 0 <= nextIndex && nextIndex < n && visited[nextIndex] == false { // 未越界或未到达过加入到队列中
                visited[nextIndex] = true
                queue = append(queue, nextIndex)
            }
        }
    }
    return false
}

type Queue struct {
    myQueue []int
    head int    // queue的头在myQueue中的索引
    notEmpty bool
}

func (q *Queue) Push(a int) {
    q.myQueue = append(q.myQueue, a)
    q.notEmpty = true
}
func (q * Queue) Top() int {
    if (q.notEmpty) {
        return q.myQueue[q.head]
    }
    return 0
}
func (q *Queue) Pop() {
    q.head++
    if (q.head >= len(q.myQueue)) {
        q.head = len(q.myQueue)
        q.notEmpty = false
    }
}

func canReach3(arr []int, start int) bool {
    n := len(arr)
    q := Queue{}
    q.Push(start)
    for q.notEmpty {
        a := q.Top()
        if (arr[a] == -1) { // 表示arr[a]已经处理过了
            q.Pop()
            continue
        }
        if (arr[a] == 0) {
            return true
        }
        if (a + arr[a] < n && arr[a + arr[a]] != -1) { // 未越界&未访问过
            q.Push(a + arr[a])
        }
        if (a - arr[a] >= 0 && arr[a - arr[a]] != -1) {  // 未越界&未访问过
            q.Push(a - arr[a])
        }
        arr[a] = -1 // 表示这个位置已经跳到过了
        q.Pop()
    }
    return false
}

func main() {
    // Example 1:
    // Input: arr = [4,2,3,0,3,1,2], start = 5
    // Output: true
    // Explanation: 
    // All possible ways to reach at index 3 with value 0 are: 
    // index 5 -> index 4 -> index 1 -> index 3 
    // index 5 -> index 6 -> index 4 -> index 1 -> index 3 
    fmt.Println(canReach([]int{4,2,3,0,3,1,2}, 5)) // true
    // Example 2:
    // Input: arr = [4,2,3,0,3,1,2], start = 0
    // Output: true 
    // Explanation: 
    // One possible way to reach at index 3 with value 0 is: 
    // index 0 -> index 4 -> index 1 -> index 3
    fmt.Println(canReach([]int{4,2,3,0,3,1,2}, 0)) // true
    // Example 3:
    // Input: arr = [3,0,2,1,2], start = 2
    // Output: false
    // Explanation: There is no way to reach at index 1 with value 0.
    fmt.Println(canReach([]int{3,0,2,1,2}, 0)) // false

    fmt.Println(canReach1([]int{4,2,3,0,3,1,2}, 5)) // true
    fmt.Println(canReach1([]int{4,2,3,0,3,1,2}, 0)) // true
    fmt.Println(canReach1([]int{3,0,2,1,2}, 0)) // false

    fmt.Println(canReach2([]int{4,2,3,0,3,1,2}, 5)) // true
    fmt.Println(canReach2([]int{4,2,3,0,3,1,2}, 0)) // true
    fmt.Println(canReach2([]int{3,0,2,1,2}, 0)) // false

    fmt.Println(canReach3([]int{4,2,3,0,3,1,2}, 5)) // true
    fmt.Println(canReach3([]int{4,2,3,0,3,1,2}, 0)) // true
    fmt.Println(canReach3([]int{3,0,2,1,2}, 0)) // false
}