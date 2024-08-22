package main

// 853. Car Fleet
// There are n cars at given miles away from the starting mile 0, traveling to reach the mile target.

// You are given two integer array position and speed, both of length n, 
// where position[i] is the starting mile of the ith car and speed[i] is the speed of the ith car in miles per hour.

// A car cannot pass another car, but it can catch up and then travel next to it at the speed of the slower car.

// A car fleet is a car or cars driving next to each other. 
// The speed of the car fleet is the minimum speed of any car in the fleet.

// If a car catches up to a car fleet at the mile target, it will still be considered as part of the car fleet.

// Return the number of car fleets that will arrive at the destination.

// Example 1:
// Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
// Output: 3
// Explanation:
// The cars starting at 10 (speed 2) and 8 (speed 4) become a fleet, meeting each other at 12. The fleet forms at target.
// The car starting at 0 (speed 1) does not catch up to any other car, so it is a fleet by itself.
// The cars starting at 5 (speed 1) and 3 (speed 3) become a fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches target.

// Example 2:
// Input: target = 10, position = [3], speed = [3]
// Output: 1
// Explanation:
// There is only one car, hence there is only one fleet.

// Example 3:
// Input: target = 100, position = [0,2,4], speed = [4,2,1]
// Output: 1
// Explanation:
// The cars starting at 0 (speed 4) and 2 (speed 2) become a fleet, meeting each other at 4. The car starting at 4 (speed 1) travels to 5.
// Then, the fleet at 4 (speed 2) and the car at position 5 (speed 1) become one fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches target.

// Constraints:
//     n == position.length == speed.length
//     1 <= n <= 10^5
//     0 < target <= 10^6
//     0 <= position[i] < target
//     All the values of position are unique.
//     0 < speed[i] <= 10^6

import "fmt"
import "sort"
import "slices"

// stack 
func carFleet(target int, position []int, speed []int) int {
    n := len(position)
    cars := make([][]int, n)
    for i, p := range position {
        cars[i] = []int{p, speed[i]}
    }
    sort.Slice(cars, func(i, j int) bool { 
        return cars[i][0] < cars[j][0] 
    })
    stack := []float64{}
    for _, c := range cars {
        t := float64(target - c[0]) / float64(c[1])
        // previous car will get to dest before cur (t)
        // pop so it will arrive with (t)
        for len(stack) > 0 && t >= stack[len(stack)-1] {
            stack = stack[:len(stack)-1] // pop
        }
        stack = append(stack, t) // push
    }
    return len(stack) // cars arriving together
}

func carFleet1(target int, position []int, speed []int) int {
    n := len(position)
    cars := make([][]float64, n)
    for i, p := range position {
        cars[i] = []float64{ float64(p), float64(target-p) / float64(speed[i]) }
    }
    sort.Slice(cars, func(i, j int) bool { return cars[i][0] < cars[j][0] })
    cur, fleet := 0.0, 0
    for i := n-1; i >= 0; i-- {
        if cars[i][1] > cur {
            cur = cars[i][1]
            fleet++
        }
    }
    return fleet
}

func carFleet2(target int, position []int, speed []int) int {
    // 排序 + 单调栈
    // 按照pos排序,后面如果能追上前面的,就会变为前面的(被拦住) => 转换为计算costTime是否小于前面更方便
    // 注意!! 题目要求的是"车队的数量" 而不是 "车队中车的数量"
    n := len(position)
    type pair struct {
        pos  int
        time float64
    }
    arr := make([]pair, 0, n)
    for i, pos := range position {
        p := pair{pos, float64(target-pos) / float64(speed[i])}
        arr = append(arr, p)
    }
    slices.SortFunc(arr, func(a, b pair) int {
        return a.pos - b.pos
    })

    // 如何计算车队数量
    // 从前往后不好处理, 因为 1号车追不上2号, 但最终有个耗时极大的end,可以将前面的全部拦住
    // 方法一:逆序思维: 可以倒序遍历,设当前有n个车队,发现当前车比前一辆耗费时间长,则融合车队-1 (从后往前好处理,因为当前处理的就是拦路虎)

    // 方法二: 单调栈, 栈中都是未处理的数据, 后面如果遇到一个时间更长的,它就是瓶颈,弹出所有比它时间短的(融合为它), 然后它处于未处理状态,如栈
    // 最终栈中的元素是所有不能被后面的车拦住的车,
    stack := make([]float64, 0, n) // 栈中塞的是costTime
    for _, cur := range arr {
        for len(stack) > 0 && cur.time >= stack[len(stack)-1] { 
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, cur.time)
    }
    return len(stack)
}

func main() {
    // Example 1:
    // Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
    // Output: 3
    // Explanation:
    // The cars starting at 10 (speed 2) and 8 (speed 4) become a fleet, meeting each other at 12. The fleet forms at target.
    // The car starting at 0 (speed 1) does not catch up to any other car, so it is a fleet by itself.
    // The cars starting at 5 (speed 1) and 3 (speed 3) become a fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches target.
    fmt.Println(carFleet(12, []int{10,8,0,5,3}, []int{2,4,1,1,3})) // 3
    // Example 2:
    // Input: target = 10, position = [3], speed = [3]
    // Output: 1
    // Explanation:
    // There is only one car, hence there is only one fleet.
    fmt.Println(carFleet(10, []int{3}, []int{3})) // 1
    // Example 3:
    // Input: target = 100, position = [0,2,4], speed = [4,2,1]
    // Output: 1
    // Explanation:
    // The cars starting at 0 (speed 4) and 2 (speed 2) become a fleet, meeting each other at 4. The car starting at 4 (speed 1) travels to 5.
    // Then, the fleet at 4 (speed 2) and the car at position 5 (speed 1) become one fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches target.
    fmt.Println(carFleet(100, []int{0,2,4}, []int{4,2,1})) // 1

    fmt.Println(carFleet1(12, []int{10,8,0,5,3}, []int{2,4,1,1,3})) // 3
    fmt.Println(carFleet1(10, []int{3}, []int{3})) // 1
    fmt.Println(carFleet1(100, []int{0,2,4}, []int{4,2,1})) // 1

    fmt.Println(carFleet2(12, []int{10,8,0,5,3}, []int{2,4,1,1,3})) // 3
    fmt.Println(carFleet2(10, []int{3}, []int{3})) // 1
    fmt.Println(carFleet2(100, []int{0,2,4}, []int{4,2,1})) // 1
}