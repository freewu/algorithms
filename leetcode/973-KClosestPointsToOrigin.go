package main

// 973. K Closest Points to Origin
// Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k, 
// return the k closest points to the origin (0, 0).

// The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).

// You may return the answer in any order. 
// The answer is guaranteed to be unique (except for the order that it is in).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/03/closestplane1.jpg" />
// Input: points = [[1,3],[-2,2]], k = 1
// Output: [[-2,2]]
// Explanation:
// The distance between (1, 3) and the origin is sqrt(10).
// The distance between (-2, 2) and the origin is sqrt(8).
// Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
// We only want the closest k = 1 points from the origin, so the answer is just [[-2,2]].

// Example 2:
// Input: points = [[3,3],[5,-1],[-2,4]], k = 2
// Output: [[3,3],[-2,4]]
// Explanation: The answer [[-2,4],[3,3]] would also be accepted.
 
// Constraints:
//     1 <= k <= points.length <= 10^4
//     -10^4 <= xi, yi <= 10^4

import "fmt"
import "sort"
import "math/rand"
import "time"

func kClosest(points [][]int, k int) [][]int {
    res, distanceMap, distances := [][]int{}, make(map[int][]int), []int{}
    for i := 0; i < len(points); i++ {
        distance := points[i][0] * points[i][0] + points[i][1] * points[i][1]
        if _, ok := distanceMap[distance]; !ok { // 去重的距离
            distances = append(distances, distance)
        }
        distanceMap[distance] = append(distanceMap[distance], i)
    }
    sort.Ints(distances)
    for i := 0; i < len(distances); i++ {
        distanceIndexes := distanceMap[distances[i]]
        for _, index := range distanceIndexes {
            if k == 0 {
                break
            }
            res = append(res, points[index])
            k--
        }
        if k == 0 {
            break
        }
    }
    return res
}

// quick select
func kClosest1(points [][]int, k int) [][]int {
    rand.Seed(time.Now().UnixNano())
    distance := func(p []int) int { return p[0]*p[0] + p[1]*p[1] }
    partition := func(points [][]int, p, q int) int {
        r := p + rand.Intn(q - p + 1) // choose a random location and swap to first
        points[p], points[r] = points[r], points[p]
        pivot := points[p]
        i, j := p-1, q+1
        for {
            i++; j--
            for distance(points[i]) < distance(pivot) { i++ }
            for distance(points[j]) > distance(pivot) { j-- }
            if i >= j { return j  }
            points[i], points[j] = points[j], points[i]
        }
    }
    quickselect := func (points [][]int, k int) {
        p, q := 0, len(points)-1
        for p <= q {
            m := partition(points, p, q)
            if p == q {
                return
            }
            if m >= k {
                q = m
            } else {
                p = m+1
            }
        }
    }
    quickselect(points, k)
    return points[:k]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/03/closestplane1.jpg" />
    // Input: points = [[1,3],[-2,2]], k = 1
    // Output: [[-2,2]]
    // Explanation:
    // The distance between (1, 3) and the origin is sqrt(10).
    // The distance between (-2, 2) and the origin is sqrt(8).
    // Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
    // We only want the closest k = 1 points from the origin, so the answer is just [[-2,2]].
    fmt.Println(kClosest([][]int{{1,3},{-2,2}}, 1)) // [[-2,2]]
    // Example 2:
    // Input: points = [[3,3],[5,-1],[-2,4]], k = 2
    // Output: [[3,3],[-2,4]]
    // Explanation: The answer [[-2,4],[3,3]] would also be accepted.
    fmt.Println(kClosest([][]int{{3,3},{5,-1},{-2,4}}, 2)) // [[3,3],[-2,4]]

    fmt.Println(kClosest1([][]int{{1,3},{-2,2}}, 1)) // [[-2,2]]
    fmt.Println(kClosest1([][]int{{3,3},{5,-1},{-2,4}}, 2)) // [[3,3],[-2,4]]
}