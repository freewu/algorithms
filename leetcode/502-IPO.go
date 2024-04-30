package main

// 502. IPO
// Suppose LeetCode will start its IPO soon. 
// In order to sell a good price of its shares to Venture Capital, LeetCode would like to work on some projects to increase its capital before the IPO. 
// Since it has limited resources, it can only finish at most k distinct projects before the IPO. 
// Help LeetCode design the best way to maximize its total capital after finishing at most k distinct projects.

// You are given n projects where the ith project has a pure profit profits[i] and a minimum capital of capital[i] is needed to start it.

// Initially, you have w capital. When you finish a project, you will obtain its pure profit and the profit will be added to your total capital.

// Pick a list of at most k distinct projects from given projects to maximize your final capital, 
// and return the final maximized capital.

// The answer is guaranteed to fit in a 32-bit signed integer.

// Example 1:
// Input: k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
// Output: 4
// Explanation: Since your initial capital is 0, you can only start the project indexed 0.
// After finishing it you will obtain profit 1 and your capital becomes 1.
// With capital 1, you can either start the project indexed 1 or the project indexed 2.
// Since you can choose at most 2 projects, you need to finish the project indexed 2 to get the maximum capital.
// Therefore, output the final maximized capital, which is 0 + 1 + 3 = 4.

// Example 2:
// Input: k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
// Output: 6

// Constraints:
//     1 <= k <= 10^5
//     0 <= w <= 10^9
//     n == profits.length
//     n == capital.length
//     1 <= n <= 10^5
//     0 <= profits[i] <= 10^4
//     0 <= capital[i] <= 10^9

import "fmt"
import "container/heap"
import "sort"

type project struct {
    Cap int
    Pro int
}

type capheap []project

func (ch capheap) Len() int {
    return len(ch)
}

func (ch capheap) Less(i, j int) bool {
    return ch[i].Cap < ch[j].Cap
}

func (ch capheap) Swap(i, j int) {
    ch[i], ch[j] = ch[j], ch[i]
}

func (ch *capheap) Push(x interface{}) {
    *ch = append(*ch, x.(project))
}

func (ch *capheap) Pop() interface{} {
    l := len(*ch)
    item := (*ch)[l-1]
    (*ch) = (*ch)[:l-1]   
    return item
}



type pfheap []project

func (ch pfheap) Len() int {
    return len(ch)
}

func (ch pfheap) Less(i, j int) bool {
    return ch[j].Pro < ch[i].Pro
}

func (ch pfheap) Swap(i, j int) {
    ch[i], ch[j] = ch[j], ch[i]
}

func (ch *pfheap) Push(x interface{}) {
    *ch = append(*ch, x.(project))
}

func (ch *pfheap) Pop() interface{} {
    l := len(*ch)
    item := (*ch)[l-1]
    (*ch) = (*ch)[:l-1]
    return item
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
    capHeap, profHeap := &capheap{}, &pfheap{}
    heap.Init(capHeap)
    heap.Init(profHeap)
    for k, v := range capital {
        heap.Push(capHeap, project{
            Cap : v,
            Pro : profits[k],
        })
    }
    for k > 0 {
        k--
        for len(*capHeap) > 0 && (*capHeap)[0].Cap <= w {
            heap.Push(profHeap, heap.Pop(capHeap))
        }
        if len(*profHeap) == 0 {
            break
        }
        item := heap.Pop(profHeap).(project)
        w = w + item.Pro
        
    }
    return w
}

func findMaximizedCapital1(k int, w int, profits []int, capital []int) int {
    type project struct {
        p,c int
        isDone bool
    }
    pj := make([]project,len(profits))
    for i:=0;i<len(profits);i++{
        pj[i] = project{profits[i],capital[i],false}
    }
    if k > len(pj) {
        k = len(pj)
    }
    sort.Slice(pj,func(i,j int) bool {
        return pj[i].p > pj[j].p
    })
    start := 0
    for k > 0 {
        flag := false
        for i:=start;i<len(pj);i++{
            if w >= pj[i].c && !pj[i].isDone {
                w += pj[i].p 
                pj[i].isDone = true
                flag = true
                if i == start {
                    start++
                }
                k--
                break
            }
        }
        if !flag {
            break
        }
    }
    return w
}

func main() {
    // Example 1:
    // Input: k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
    // Output: 4
    // Explanation: Since your initial capital is 0, you can only start the project indexed 0.
    // After finishing it you will obtain profit 1 and your capital becomes 1.
    // With capital 1, you can either start the project indexed 1 or the project indexed 2.
    // Since you can choose at most 2 projects, you need to finish the project indexed 2 to get the maximum capital.
    // Therefore, output the final maximized capital, which is 0 + 1 + 3 = 4.
    fmt.Println(findMaximizedCapital(2,0,[]int{1,2,3}, []int{0,1,1})) // 4
    // Example 2:
    // Input: k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
    // Output: 6
    fmt.Println(findMaximizedCapital(3,0,[]int{1,2,3}, []int{0,1,2})) // 6
    
    fmt.Println(findMaximizedCapital1(2,0,[]int{1,2,3}, []int{0,1,1})) // 4
    fmt.Println(findMaximizedCapital1(3,0,[]int{1,2,3}, []int{0,1,2})) // 6
}