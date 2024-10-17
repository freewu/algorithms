package main

// 1515. Best Position for a Service Centre 
// A delivery company wants to build a new service center in a new city. 
// The company knows the positions of all the customers in this city on a 2D-Map 
// and wants to build the new center in a position such that the sum of the euclidean distances to all customers is minimum.

// Given an array positions where positions[i] = [xi, yi] is the position of the ith customer on the map, 
// return the minimum sum of the euclidean distances to all customers.

// In other words, you need to choose the position of the service center [xcentre, ycentre] 
// such that the following formula is minimized:
// <img src="https://assets.leetcode.com/uploads/2020/06/25/q4_edited.jpg" />

// Answers within 10^-5 of the actual value will be accepted.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/06/25/q4_e1.jpg" />
// Input: positions = [[0,1],[1,0],[1,2],[2,1]]
// Output: 4.00000
// Explanation: As shown, you can see that choosing [xcentre, ycentre] = [1, 1] will make the distance to each customer = 1, the sum of all distances is 4 which is the minimum possible we can achieve.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/06/25/q4_e3.jpg" />
// Input: positions = [[1,1],[3,3]]
// Output: 2.82843
// Explanation: The minimum possible sum of distances = sqrt(2) + sqrt(2) = 2.82843

// Constraints:
//     1 <= positions.length <= 50
//     positions[i].length == 2
//     0 <= xi, yi <= 100

import "fmt"
import "math"

func getMinDistSum(positions [][]int) float64 {
    dist := func(x, y float64) float64 {
        res := float64(0)
        for _, pos := range positions {
            x0, y0 := float64(pos[0]), float64(pos[1])
            res += math.Sqrt((x-x0) * (x-x0) + (y-y0) * (y-y0))
        }
        return res
    }
    res, x0, y0 := math.MaxFloat64, 0.0, 0.0
    left, right, top, bottom := 0.0, 100.0, 100.0, 0.0
    stop := 0.00001
    for delta := 10.; delta > stop; delta /= 10. {
        for x := left; x <= right; x += delta {
            for y := bottom; y <= top; y += delta {
                d := dist(x, y)
                if d < res {
                    x0, y0, res= x, y, d
                }
            }
        }
        left, right, top, bottom = x0 - delta, x0 + delta, y0 + delta, y0 - delta
    }
    return  res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/06/25/q4_e1.jpg" />
    // Input: positions = [[0,1],[1,0],[1,2],[2,1]]
    // Output: 4.00000
    // Explanation: As shown, you can see that choosing [xcentre, ycentre] = [1, 1] will make the distance to each customer = 1, the sum of all distances is 4 which is the minimum possible we can achieve.
    fmt.Println(getMinDistSum([][]int{{0,1},{1,0},{1,2},{2,1}})) // 4.00000
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/06/25/q4_e3.jpg" />
    // Input: positions = [[1,1],[3,3]]
    // Output: 2.82843
    // Explanation: The minimum possible sum of distances = sqrt(2) + sqrt(2) = 2.82843
    fmt.Println(getMinDistSum([][]int{{1,1},{3,3}})) // 2.82843
}