package main

// 871. Minimum Number of Refueling Stops
// A car travels from a starting position to a destination which is target miles east of the starting position.

// There are gas stations along the way. 
// The gas stations are represented as an array stations where stations[i] = [positioni, fueli] indicates 
// that the ith gas station is positioni miles east of the starting position and has fueli liters of gas.

// The car starts with an infinite tank of gas, which initially has startFuel liters of fuel in it. 
// It uses one liter of gas per one mile that it drives. 
// When the car reaches a gas station, it may stop and refuel, transferring all the gas from the station into the car.

// Return the minimum number of refueling stops the car must make in order to reach its destination. 
// If it cannot reach the destination, return -1.

// Note that if the car reaches a gas station with 0 fuel left, the car can still refuel there. 
// If the car reaches the destination with 0 fuel left, it is still considered to have arrived.

// Example 1:
// Input: target = 1, startFuel = 1, stations = []
// Output: 0
// Explanation: We can reach the target without refueling.

// Example 2:
// Input: target = 100, startFuel = 1, stations = [[10,100]]
// Output: -1
// Explanation: We can not reach the target (or even the first gas station).

// Example 3:
// Input: target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]
// Output: 2
// Explanation: We start with 10 liters of fuel.
// We drive to position 10, expending 10 liters of fuel.  We refuel from 0 liters to 60 liters of gas.
// Then, we drive from position 10 to position 60 (expending 50 liters of fuel),
// and refuel from 10 liters to 50 liters of gas.  We then drive to and reach the target.
// We made 2 refueling stops along the way, so we return 2.

// Constraints:
//     1 <= target, startFuel <= 10^9
//     0 <= stations.length <= 500
//     1 <= positioni < positioni+1 < target
//     1 <= fueli < 10^9

import "fmt"
import "sort"
import "container/heap"

// dp
func minRefuelStops(target int, startFuel int, stations [][]int) int {
    if startFuel > target { // 一箱油就够了 不需要中途加油
        return 0
    }
    dp := make([]int, len(stations)+1)
    dp[0] = startFuel
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range stations {
        location, capacity := v[0], v[1]
        for j := i + 1; j > 0; j-- {
            // can take?
            if location <= dp[j-1] {
                dp[j] = max(dp[j], dp[j-1]+capacity)
            }
        }
    }
    for i, v := range dp {
        if v >= target {
            return i
        }
    }
    return -1
}

func minRefuelStops1(target int, startFuel int, stations [][]int) int {
    res, cur := 0, startFuel
    hp := make(MaxHeap, 0)
    i := 0 // 当前加油站的index
    for cur < target {
        if i < len(stations) && cur >= stations[i][0] {
            heap.Push(&hp, stations[i][1])
            i++
        } else {
            if len(hp) == 0 {
                return -1
            }
            pop := heap.Pop(&hp).(int)
            res++
            cur += pop
        }
    }
    return res
}

type MaxHeap []int
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
func (h *MaxHeap) Peek() interface{} {
    if len(*h) > 0 {
        return (*h)[0]
    }
    return 0
}

func minRefuelStops2(target, startFuel int, stations [][]int) int {
    stations = append(stations, []int{target, 0})
    res, prePosition, curFuel := 0, 0, startFuel
    fuelHeap := &hp{}
    for _, station := range stations {
        position, fuel := station[0], station[1]
        curFuel -= position - prePosition       // 每行驶 1 英里用掉 1 升汽油
        for fuelHeap.Len() > 0 && curFuel < 0 { // 没油了
            curFuel += heap.Pop(fuelHeap).(int) // 选油量最多的油桶
            res++
        }
        if curFuel < 0 { // 无法到达
            return -1
        }
        heap.Push(fuelHeap, fuel) // 留着后面加油
        prePosition = position
    }
    return res
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func main() {
    // Example 1:
    // Input: target = 1, startFuel = 1, stations = []
    // Output: 0
    // Explanation: We can reach the target without refueling.
    fmt.Println(minRefuelStops(1,1,[][]int{})) // 0
    // Example 2:
    // Input: target = 100, startFuel = 1, stations = [[10,100]]
    // Output: -1
    // Explanation: We can not reach the target (or even the first gas station).
    fmt.Println(minRefuelStops(100,1,[][]int{{10,100}})) // -1
    // Example 3:
    // Input: target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]
    // Output: 2
    // Explanation: We start with 10 liters of fuel.
    // We drive to position 10, expending 10 liters of fuel.  We refuel from 0 liters to 60 liters of gas.
    // Then, we drive from position 10 to position 60 (expending 50 liters of fuel),
    // and refuel from 10 liters to 50 liters of gas.  We then drive to and reach the target.
    // We made 2 refueling stops along the way, so we return 2.
    fmt.Println(minRefuelStops(100,10,[][]int{{10,60},{20,30},{30,30},{60,40}})) // 2

    fmt.Println(minRefuelStops1(1,1,[][]int{})) // 0
    fmt.Println(minRefuelStops1(100,1,[][]int{{10,100}})) // -1
    fmt.Println(minRefuelStops1(100,10,[][]int{{10,60},{20,30},{30,30},{60,40}})) // 2

    fmt.Println(minRefuelStops2(1,1,[][]int{})) // 0
    fmt.Println(minRefuelStops2(100,1,[][]int{{10,100}})) // -1
    fmt.Println(minRefuelStops2(100,10,[][]int{{10,60},{20,30},{30,30},{60,40}})) // 2
}