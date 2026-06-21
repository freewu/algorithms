package main

// 3964. Minimum Lights to Illuminate a Road
// You are given an integer array lights of length n, representing positions 0 through n - 1 on a road.

// For each position i:
//     1. If lights[i] = v, where v > 0, there is a working bulb at position i that illuminates every position from max(0, i - v) to min(n - 1, i + v), inclusive.
//     2. If lights[i] = 0, there is no working bulb at position i.

// A position is visible if it is illuminated by at least one working bulb.

// You may install additional bulbs at any positions. 
// Each additional bulb installed at position j illuminates positions from max(0, j - 1) to min(n - 1, j + 1), inclusive.

// Return the minimum number of additional bulbs required to make every position on the road visible.

// Example 1:
// Input: lights = [0,0,0,0]
// Output: 2
// Explanation:
// One optimal placement is:
// Install an additional bulb at position 1, illuminating positions [0, 1, 2].
// Install an additional bulb at position 3, illuminating positions [2, 3].
// Therefore, the minimum number of additional bulbs required is 2.

// Example 2:
// Input: lights = [0,0,0,2,0]
// Output: 1
// Explanation:
// Since lights[3] = 2, the working bulb at position 3 illuminates positions [1, 2, 3, 4].
// Installing an additional bulb at position 1 illuminates positions [0, 1, 2], making every position visible.
// Therefore, the minimum number of additional bulbs required is 1.
 
// Constraints:
//     1 <= n == lights.length <= 10^5
//     0 <= lights[i] <= n

import "fmt"
import "slices"

func minLights(lights []int) int {
    res, frontier, n := 0, 0, len(lights)
    ranges := lights[:0]
    for i, v := range lights {
        if v != 0 {
            ranges = append(ranges, max(0, i - v) << 32 | min(n - 1, i + v))    
        }
    }
    calcAdditional := func (l, frontier int) int {
        return (max(l-frontier, 0) + 2) / 3
    }
    slices.Sort(ranges)
    for _, v := range ranges {
        l, r := v >> 32, v << 32 >> 32
        res += calcAdditional(l, frontier)
        frontier = max(frontier, r + 1)
    }
    return res + calcAdditional(n, frontier)
}

func minLights1(lights []int) int {
    index, val, n := 0, lights[0], len(lights)
    for i := 1; i < n; i++ {
        curr := lights[i]
        if curr == 0 {
            if i - index <= val {
                lights[i] = 1
            } else {
                index, val = i, lights[i] 
            }
        } else if curr > 0 {
            if curr >= val - i + index || i - index > val {
                index, val = i, lights[i] 
            }
        }  
    }
    index, val = n - 1, lights[n - 1]
    for i := n - 2; i >= 0; i-- {
        curr := lights[i]
        if curr == 0 {
            if index - i <= val {
                lights[i] = 1
            }else{
                index, val = i, lights[i] 
            }
        } else if curr > 0 {
            if curr >= val - index + i || index - i > val {
                index, val = i, lights[i] 
            }
        }  
    }
    res, i := 0, 0
    for i < n {
        if lights[i] > 0 {
            i++
            continue
        }
        res++
        k := 1
        if i + 1 <n && lights[i + 1] == 0 {
            k++
        }
        if i + 2 <n && lights[i + 2] == 0 {
            k++
        }
        i += k 
    }
    return res
}

func minLights2(lights []int) int {
    res, until, n := 0, -1, len(lights)
    on := make([]bool, n)
    for i := 0; i < n; i++ {
        if lights[i] != 0 {
            if i + lights[i] > until {
                until = i + lights[i]
            }
        }
        if until >= i {
            on[i] = true
        }
    }
    until = n
    for i := n - 1; i >= 0; i-- {
        if lights[i] != 0 {
            if i - lights[i] < until {
                until = i - lights[i]
            }
        }
        if until <= i {
            on[i] = true
        }
    }
    ok := func(x int) bool {
        return 0 <= x && x < n
    }
    up := func(x int) {
        for _, p := range []int{-1, 0, 1} {
            if ok(x + p) {
                on[x + p] = true
            }
        }
    }
    for i := 0; i < n; i++ {
        if on[i] {
            continue
        }
        if i == n - 1 {
            up(i)
        } else {
            up(i + 1)
        }
        res++
    }
    return res
}

func main() {
    // Example 1:
    // Input: lights = [0,0,0,0]
    // Output: 2
    // Explanation:
    // One optimal placement is:
    // Install an additional bulb at position 1, illuminating positions [0, 1, 2].
    // Install an additional bulb at position 3, illuminating positions [2, 3].
    // Therefore, the minimum number of additional bulbs required is 2.
    fmt.Println(minLights([]int{0,0,0,0})) // 2
    // Example 2:
    // Input: lights = [0,0,0,2,0]
    // Output: 1
    // Explanation:
    // Since lights[3] = 2, the working bulb at position 3 illuminates positions [1, 2, 3, 4].
    // Installing an additional bulb at position 1 illuminates positions [0, 1, 2], making every position visible.
    // Therefore, the minimum number of additional bulbs required is 1.
    fmt.Println(minLights([]int{0,0,0,2,0})) // 1 

    fmt.Println(minLights([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minLights([]int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(minLights([]int{0,1,1})) // 1

    fmt.Println(minLights1([]int{0,0,0,0})) // 2
    fmt.Println(minLights1([]int{0,0,0,2,0})) // 1 
    fmt.Println(minLights1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minLights1([]int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(minLights1([]int{0,1,1})) // 0

    fmt.Println(minLights2([]int{0,0,0,0})) // 2
    fmt.Println(minLights2([]int{0,0,0,2,0})) // 1 
    fmt.Println(minLights2([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minLights2([]int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(minLights2([]int{0,1,1})) // 0
}