package main

// 1606. Find Servers That Handled Most Number of Requests
// You have k servers numbered from 0 to k-1 that are being used to handle multiple requests simultaneously. 
// Each server has infinite computational capacity but cannot handle more than one request at a time. 
// The requests are assigned to servers according to a specific algorithm:
//     1. The ith (0-indexed) request arrives.
//     2. If all servers are busy, the request is dropped (not handled at all).
//     3. If the (i % k)th server is available, assign the request to that server.
//     4. Otherwise, assign the request to the next available server 
//        (wrapping around the list of servers and starting from 0 if necessary). 
//        For example, if the ith server is busy, try to assign the request to the (i+1)th server, then the (i+2)th server, and so on.

// You are given a strictly increasing array arrival of positive integers, 
// where arrival[i] represents the arrival time of the ith request, 
// and another array load, where load[i] represents the load of the ith request (the time it takes to complete). 
// Your goal is to find the busiest server(s). 
// A server is considered busiest if it handled the most number of requests successfully among all the servers.

// Return a list containing the IDs (0-indexed) of the busiest server(s). 
// You may return the IDs in any order.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/08/load-1.png" />
// Input: k = 3, arrival = [1,2,3,4,5], load = [5,2,3,3,3] 
// Output: [1] 
// Explanation: 
// All of the servers start out available.
// The first 3 requests are handled by the first 3 servers in order.
// Request 3 comes in. Server 0 is busy, so it's assigned to the next available server, which is 1.
// Request 4 comes in. It cannot be handled since all servers are busy, so it is dropped.
// Servers 0 and 2 handled one request each, while server 1 handled two requests. Hence server 1 is the busiest server.

// Example 2:
// Input: k = 3, arrival = [1,2,3,4], load = [1,2,1,2]
// Output: [0]
// Explanation: 
// The first 3 requests are handled by first 3 servers.
// Request 3 comes in. It is handled by server 0 since the server is available.
// Server 0 handled two requests, while servers 1 and 2 handled one request each. Hence server 0 is the busiest server.

// Example 3:
// Input: k = 3, arrival = [1,2,3], load = [10,12,11]
// Output: [0,1,2]
// Explanation: Each server handles a single request, so they are all considered the busiest.

// Constraints:
//     1 <= k <= 10^5
//     1 <= arrival.length, load.length <= 10^5
//     arrival.length == load.length
//     1 <= arrival[i], load[i] <= 10^9
//     arrival is strictly increasing.

import "fmt"
import "container/heap"

type ListHeap [][]int
func (h ListHeap) Len() int{return len(h)}
func (h ListHeap) Less(i, j int) bool{return h[i][0] < h[j][0]}
func (h ListHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i]}
func (h *ListHeap) Pop() interface{} {
    n := len(*h)
    v := (*h)[n - 1]
    *h = (*h)[0 : n - 1]
    return v
}
func (h *ListHeap) Push(v interface{}) {
    *h = append(*h, v.([]int))
}

type IntHeap []int
func (h IntHeap) Len() int{return len(h)}
func (h IntHeap) Less(i, j int) bool{return h[i] < h[j]}
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i]}
func (h *IntHeap) Pop() interface{} {
    n := len(*h)
    v := (*h)[n - 1]
    *h = (*h)[0 : n - 1]
    return v
}
func (h *IntHeap) Push(v interface{}) {
    *h = append(*h, v.(int))
}

func busiestServers(k int, arrival []int, load []int) []int {
    used, available, cnts := &ListHeap{}, &IntHeap{}, make([]int, k)
    for i := 0; i < k; i++ {
        heap.Push(available, i)
    }
    for i := 0; i < len(arrival); i++ {
        arr, duration := arrival[i], load[i]
        for used.Len() > 0 && (*used)[0][0] <= arr {
            cur := heap.Pop(used).([]int)
            heap.Push(available, i + ((cur[1] - i) % k + k) % k)
        }
        if available.Len() > 0 {
            idx := heap.Pop(available).(int) % k
            cnts[idx]++
            heap.Push(used, []int{arr + duration, idx})
        }
    }
    res, mx := []int{}, 0
    for i := 0; i < k; i++ {
        if cnts[i] > mx {
            mx = cnts[i]
            res = []int{ i }
        } else if cnts[i] == mx {
            res = append(res, i)
        }
    }
    return res
}

type ServerPair struct{ id, wt int } // wt: 在busy堆是结束时间, 在idle时重新设置的序号
type ServerHeap []ServerPair

func (h ServerHeap)  Len() int           { return len(h) }
func (h ServerHeap)  Less(i, j int) bool { return h[i].wt < h[j].wt }
func (h ServerHeap)  Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ServerHeap) Push(v any)         { *h = append(*h, v.(ServerPair)) }
func (h *ServerHeap) Pop() (v any)       { a := *h; v, *h = a[len(a)-1], a[:len(a)-1]; return v }

func busiestServers1(k int, arrival []int, load []int) []int {
    // 要维护一个busy堆,用来维护繁忙的服务器, 另外需要一个idle队列维护
    // 方法一: idle使用BSTTree, 是最直白的,当有空闲服务时,取 >=i%k 第一个; 取不到的话,就取最小的
    // 方法二: 使用堆(trick,难懂)
    // idle堆的入堆的序列号: i+((id-i)%k+k)%k), 也就是使用一个比i大(或相等)且和id同余的数 最为服务器的新id
    // 相当于将id修改到[i,i+k)区间了, 如果id比i%k小, 那么它就在区间的后半部分,也就是本轮不会被选到(一轮k个服务器,[0,k) [k,2k)...)
    // 如果id比i%k大(或相等(,那么它还是在这一轮能被选到)
    // 关键点! 因为 newId是和id同余的,所以它本身和i的大小没有联系, 它只会存在于 id, id+k, id+2k...这些固定的位置上,所以当i编号时,它的位置不变!!
    // 关键点2 因为一轮的任务能将后面解锁的同余服务器全部消耗掉,所以不会出现遗留被扔到下一轮,这个服务器如果被放到本轮,那么在本轮一定会被消耗掉,所以堆弹出最小值即可
    busy, idle := ServerHeap{}, make(ServerHeap, k)
    for i := 0; i < k; i++ {
        idle[i] = ServerPair{i, i}
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, count, mx := []int{}, make([]int, k), -1
    for i, start := range arrival {
        end := start + load[i]
        for busy.Len() > 0 && busy[0].wt <= start {
            cur := heap.Pop(&busy).(ServerPair)
            nxId := i + ((cur.id - i) % k + k) % k // 找到一个nxId,和id同余,并且被偏移到了[i:i+k)区间,偏移需要避免负数
            heap.Push(&idle, ServerPair{cur.id, nxId})
        }
        if idle.Len() > 0 {
            cur := heap.Pop(&idle).(ServerPair)
            count[cur.id]++
            mx = max(mx, count[cur.id])
            cur.wt = end
            heap.Push(&busy, cur)
        }
    }
    for i, v := range count {
        if v == mx {
            res = append(res, i)
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/08/load-1.png" />
    // Input: k = 3, arrival = [1,2,3,4,5], load = [5,2,3,3,3] 
    // Output: [1] 
    // Explanation: 
    // All of the servers start out available.
    // The first 3 requests are handled by the first 3 servers in order.
    // Request 3 comes in. Server 0 is busy, so it's assigned to the next available server, which is 1.
    // Request 4 comes in. It cannot be handled since all servers are busy, so it is dropped.
    // Servers 0 and 2 handled one request each, while server 1 handled two requests. Hence server 1 is the busiest server.
    fmt.Println(busiestServers(3, []int{1,2,3,4,5}, []int{5,2,3,3,3})) // [1] 
    // Example 2:
    // Input: k = 3, arrival = [1,2,3,4], load = [1,2,1,2]
    // Output: [0]
    // Explanation: 
    // The first 3 requests are handled by first 3 servers.
    // Request 3 comes in. It is handled by server 0 since the server is available.
    // Server 0 handled two requests, while servers 1 and 2 handled one request each. Hence server 0 is the busiest server.
    fmt.Println(busiestServers(3, []int{1,2,3,4}, []int{1,2,1,2})) // [0]
    // Example 3:
    // Input: k = 3, arrival = [1,2,3], load = [10,12,11]
    // Output: [0,1,2]
    // Explanation: Each server handles a single request, so they are all considered the busiest.
    fmt.Println(busiestServers(3, []int{1,2,3}, []int{10,12,11})) // [0,1,2]

    fmt.Println(busiestServers1(3, []int{1,2,3,4,5}, []int{5,2,3,3,3})) // [1] 
    fmt.Println(busiestServers1(3, []int{1,2,3,4}, []int{1,2,1,2})) // [0]
    fmt.Println(busiestServers1(3, []int{1,2,3}, []int{10,12,11})) // [0,1,2]
}