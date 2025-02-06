package main

// LCP 59. 搭桥过河
// 欢迎各位勇者来到力扣城，本次试炼主题为「搭桥过河」。

// 勇者面前有一段长度为 num 的河流，河流可以划分为若干河道。
// 每条河道上恰有一块浮木，wood[i] 记录了第 i 条河道上的浮木初始的覆盖范围。
//     1. 当且仅当浮木与相邻河道的浮木覆盖范围有重叠时，勇者才可以在两条浮木间移动
//     2. 勇者 仅能在岸上 通过花费一点「自然之力」，使任意一条浮木沿着河流移动一个单位距离

// 请问勇者跨越这条河流，最少需要花费多少「自然之力」。

// 示例 1：
// 输入： num = 10, wood = [[1,2],[4,7],[8,9]] 
// 输出： 3 
// 解释：如下图所示， 
// 将 [1,2] 浮木移动至 [3,4]，花费 2「自然之力」， 
// 将 [8,9] 浮木移动至 [7,8]，花费 1「自然之力」， 
// 此时勇者可以顺着 [3,4]->[4,7]->[7,8] 跨越河流，
// 因此，勇者最少需要花费 3 点「自然之力」跨越这条河流
// <img src="https://pic.leetcode-cn.com/1648196478-ophADL-wood%20(2).gif" /> 

// 示例 2：
// 输入： num = 10, wood = [[1,5],[1,1],[10,10],[6,7],[7,8]] 
// 输出： 10 
// 解释： 
// 将 [1,5] 浮木移动至 [2,6]，花费 1「自然之力」， 
// 将 [1,1] 浮木移动至 [6,6]，花费 5「自然之力」， 
// 将 [10,10] 浮木移动至 [6,6]，花费 4「自然之力」， 
// 此时勇者可以顺着 [2,6]->[6,6]->[6,6]->[6,7]->[7,8] 跨越河流， 
// 因此，勇者最少需要花费 10 点「自然之力」跨越这条河流

// 示例 3：
// 输入： num = 5, wood = [[1,2],[2,4]] 
// 输出： 0 
// 解释：勇者不需要移动浮木，仍可以跨越这条河流

// 提示:
//     1 <= num <= 10^9
//     1 <= wood.length <= 10^5
//     wood[i].length == 2
//     1 <= wood[i][0] <= wood[i][1] <= num

import "fmt"
import "container/heap"

// 定义一个大顶堆
type MaxHeap []int64
func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int64)) }
func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// 定义一个小顶堆
type MinHeap []int64
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int64)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func buildBridge(num int, wood [][]int) int64 {
    if len(wood) == 0 { return 0 }
    L, R := &MaxHeap{},  &MinHeap{}
    heap.Init(L)
    heap.Init(R)
    heap.Push(L, int64(wood[0][0]))
    heap.Push(R, int64(wood[0][0]))
    var biasL, biasR, res int64 = 0, 0, 0
    for i := 1; i < len(wood); i++ {
        biasL -= int64(wood[i][1] - wood[i][0])
        biasR += int64(wood[i-1][1] - wood[i-1][0])
        l0, r0, x := (*L)[0] + biasL, (*R)[0] + biasR, int64(wood[i][0])
        if x < l0 {
            res += l0 - x
            heap.Pop(L)
            heap.Push(L, x-biasL)
            heap.Push(L, x-biasL)
            heap.Push(R, l0-biasR)
        } else if x > r0 {
            res += x - r0
            heap.Pop(R)
            heap.Push(R, x-biasR)
            heap.Push(R, x-biasR)
            heap.Push(L, r0-biasL)
        } else {
            heap.Push(L, x-biasL)
            heap.Push(R, x-biasR)
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入： num = 10, wood = [[1,2],[4,7],[8,9]] 
    // 输出： 3 
    // 解释：如下图所示， 
    // 将 [1,2] 浮木移动至 [3,4]，花费 2「自然之力」， 
    // 将 [8,9] 浮木移动至 [7,8]，花费 1「自然之力」， 
    // 此时勇者可以顺着 [3,4]->[4,7]->[7,8] 跨越河流，
    // 因此，勇者最少需要花费 3 点「自然之力」跨越这条河流
    // <img src="https://pic.leetcode-cn.com/1648196478-ophADL-wood%20(2).gif" /> 
    fmt.Println(buildBridge(10, [][]int{{1,2},{4,7},{8,9}} )) // 3
    // 示例 2：
    // 输入： num = 10, wood = [[1,5],[1,1],[10,10],[6,7],[7,8]] 
    // 输出： 10 
    // 解释： 
    // 将 [1,5] 浮木移动至 [2,6]，花费 1「自然之力」， 
    // 将 [1,1] 浮木移动至 [6,6]，花费 5「自然之力」， 
    // 将 [10,10] 浮木移动至 [6,6]，花费 4「自然之力」， 
    // 此时勇者可以顺着 [2,6]->[6,6]->[6,6]->[6,7]->[7,8] 跨越河流， 
    // 因此，勇者最少需要花费 10 点「自然之力」跨越这条河流
    fmt.Println(buildBridge(10, [][]int{{1,5},{1,1},{10,10},{6,7},{7,8}} )) // 10
    // 示例 3：
    // 输入： num = 5, wood = [[1,2],[2,4]] 
    // 输出： 0 
    // 解释：勇者不需要移动浮木，仍可以跨越这条河流
    fmt.Println(buildBridge(5, [][]int{{1,2},{2,4}} )) // 0
}