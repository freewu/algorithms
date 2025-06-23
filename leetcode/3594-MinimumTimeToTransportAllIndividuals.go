package main

// 3594. Minimum Time to Transport All Individuals
// You are given n individuals at a base camp who need to cross a river to reach a destination using a single boat. 
// The boat can carry at most k people at a time. 
// The trip is affected by environmental conditions that vary cyclically over m stages.

// Each stage j has a speed multiplier mul[j]:
//     1. If mul[j] > 1, the trip slows down.
//     2. If mul[j] < 1, the trip speeds up.

// Each individual i has a rowing strength represented by time[i], the time (in minutes) it takes them to cross alone in neutral conditions.

// Rules:
//     1. A group g departing at stage j takes time equal to the maximum time[i] among its members, multiplied by mul[j] minutes to reach the destination.
//     2. After the group crosses the river in time d, the stage advances by floor(d) % m steps.
//     3. If individuals are left behind, one person must return with the boat. 
//        Let r be the index of the returning person, the return takes time[r] × mul[current_stage], defined as return_time, and the stage advances by floor(return_time) % m.

// Return the minimum total time required to transport all individuals. If it is not possible to transport all individuals to the destination, return -1.

// Example 1:
// Input: n = 1, k = 1, m = 2, time = [5], mul = [1.0,1.3]
// Output: 5.00000
// Explanation:
// Individual 0 departs from stage 0, so crossing time = 5 × 1.00 = 5.00 minutes.
// All team members are now at the destination. Thus, the total time taken is 5.00 minutes.

// Example 2:
// Input: n = 3, k = 2, m = 3, time = [2,5,8], mul = [1.0,1.5,0.75]
// Output: 14.50000
// Explanation:
// The optimal strategy is:
// Send individuals 0 and 2 from the base camp to the destination from stage 0. The crossing time is max(2, 8) × mul[0] = 8 × 1.00 = 8.00 minutes. The stage advances by floor(8.00) % 3 = 2, so the next stage is (0 + 2) % 3 = 2.
// Individual 0 returns alone from the destination to the base camp from stage 2. The return time is 2 × mul[2] = 2 × 0.75 = 1.50 minutes. The stage advances by floor(1.50) % 3 = 1, so the next stage is (2 + 1) % 3 = 0.
// Send individuals 0 and 1 from the base camp to the destination from stage 0. The crossing time is max(2, 5) × mul[0] = 5 × 1.00 = 5.00 minutes. The stage advances by floor(5.00) % 3 = 2, so the final stage is (0 + 2) % 3 = 2.
// All team members are now at the destination. The total time taken is 8.00 + 1.50 + 5.00 = 14.50 minutes.

// Example 3:
// Input: n = 2, k = 1, m = 2, time = [10,10], mul = [2.0,2.0]
// Output: -1.00000
// Explanation:
// Since the boat can only carry one person at a time, it is impossible to transport both individuals as one must always return. Thus, the answer is -1.00.

// Constraints:
//     1 <= n == time.length <= 12
//     1 <= k <= 5
//     1 <= m <= 5
//     1 <= time[i] <= 100
//     m == mul.length
//     0.5 <= mul[i] <= 2.0

import "fmt"
import "math"
import "math/bits"
import "container/heap"

type Tuple struct {
    dis         float64
    stage, mask int
    state       uint8 // 状态机：0 未过河，1 已过河
}
type hp []Tuple
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(Tuple)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func minTime(n, k, m int, time []int, mul []float64) float64 {
    u := 1 << n
    // 计算每个 time 子集的最大值
    maxTime := make([]int, u)
    for i, t := range time {
        highBit := 1 << i
        for mask, mx := range maxTime[:highBit] {
            maxTime[highBit|mask] = max(mx, t)
        }
    }
    // 把 maxTime 中的大小大于 k 的集合改为 inf
    for i := range maxTime {
        if bits.OnesCount(uint(i)) > k {
            maxTime[i] = math.MaxInt
        }
    }
    dis := make([][][2]float64, m)
    for i := range dis {
        dis[i] = make([][2]float64, u)
        for j := range dis[i] {
            dis[i][j] = [2]float64{ math.MaxFloat64, math.MaxFloat64 }
        }
    }
    h := hp{}
    push := func(d float64, stage, mask int, state uint8) {
        if d < dis[stage][mask][state] {
            dis[stage][mask][state] = d
            heap.Push(&h, Tuple{d, stage, mask, state})
        }
    }
    push(0, 0, u-1, 0) // 起点
    for len(h) > 0 {
        top := heap.Pop(&h).(Tuple)
        d := top.dis
        stage := top.stage
        left := top.mask // 剩余没有过河的人
        state := top.state
        if left == 0 { // 所有人都过河了
            return d
        }
        if d > dis[stage][left][state] { continue }
        if state == 0 {
            // 枚举 sub 这群人坐一艘船
            for sub := left; sub > 0; sub = (sub - 1) & left {
                if maxTime[sub] != math.MaxInt {
                    cost := float64(maxTime[sub]) * mul[stage]
                    push(d+cost, (stage+int(cost))%m, left^sub, 1)
                }
            }
        } else {
            // 枚举回来的人
            for s, lb := u-1^left, 0; s > 0; s ^= lb {
                lb = s & -s
                cost := float64(maxTime[lb]) * mul[stage]
                push(d+cost, (stage+int(cost))%m, left^lb, 0)
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: n = 1, k = 1, m = 2, time = [5], mul = [1.0,1.3]
    // Output: 5.00000
    // Explanation:
    // Individual 0 departs from stage 0, so crossing time = 5 × 1.00 = 5.00 minutes.
    // All team members are now at the destination. Thus, the total time taken is 5.00 minutes.
    fmt.Println(minTime(1, 1, 2, []int{5}, []float64{1.0,1.3})) // 5.00000
    // Example 2:
    // Input: n = 3, k = 2, m = 3, time = [2,5,8], mul = [1.0,1.5,0.75]
    // Output: 14.50000
    // Explanation:
    // The optimal strategy is:
    // Send individuals 0 and 2 from the base camp to the destination from stage 0. The crossing time is max(2, 8) × mul[0] = 8 × 1.00 = 8.00 minutes. The stage advances by floor(8.00) % 3 = 2, so the next stage is (0 + 2) % 3 = 2.
    // Individual 0 returns alone from the destination to the base camp from stage 2. The return time is 2 × mul[2] = 2 × 0.75 = 1.50 minutes. The stage advances by floor(1.50) % 3 = 1, so the next stage is (2 + 1) % 3 = 0.
    // Send individuals 0 and 1 from the base camp to the destination from stage 0. The crossing time is max(2, 5) × mul[0] = 5 × 1.00 = 5.00 minutes. The stage advances by floor(5.00) % 3 = 2, so the final stage is (0 + 2) % 3 = 2.
    // All team members are now at the destination. The total time taken is 8.00 + 1.50 + 5.00 = 14.50 minutes.
    fmt.Println(minTime(3, 2, 3, []int{2,5,8}, []float64{1.0,1.5,0.75})) // 14.50000
    // Example 3:
    // Input: n = 2, k = 1, m = 2, time = [10,10], mul = [2.0,2.0]
    // Output: -1.00000
    // Explanation:
    // Since the boat can only carry one person at a time, it is impossible to transport both individuals as one must always return. Thus, the answer is -1.00.
    fmt.Println(minTime(2, 1, 2, []int{10,10}, []float64{2.0,2.0})) // -1.00000
}