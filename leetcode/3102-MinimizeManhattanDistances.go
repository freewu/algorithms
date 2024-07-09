package main

// 3102. Minimize Manhattan Distances
// You are given a array points representing integer coordinates of some points on a 2D plane, where points[i] = [xi, yi].
// The distance between two points is defined as their Manhattan distance.
// Return the minimum possible value for maximum distance between any two points by removing exactly one point.

// Example 1:
// Input: points = [[3,10],[5,15],[10,2],[4,4]]
// Output: 12
// Explanation:
// The maximum distance after removing each point is the following:
// After removing the 0th point the maximum distance is between points (5, 15) and (10, 2), which is |5 - 10| + |15 - 2| = 18.
// After removing the 1st point the maximum distance is between points (3, 10) and (10, 2), which is |3 - 10| + |10 - 2| = 15.
// After removing the 2nd point the maximum distance is between points (5, 15) and (4, 4), which is |5 - 4| + |15 - 4| = 12.
// After removing the 3rd point the maximum distance is between points (5, 15) and (10, 2), which is |5 - 10| + |15 - 2| = 18.
// 12 is the minimum possible maximum distance between any two points after removing exactly one point.

// Example 2:
// Input: points = [[1,1],[1,1],[1,1]]
// Output: 0
// Explanation:
// Removing any of the points results in the maximum distance between any two points of 0.

// Constraints:
//     3 <= points.length <= 10^5
//     points[i].length == 2
//     1 <= points[i][0], points[i][1] <= 10^8

import "fmt"

func minimumDistance(points [][]int) int {
    mninf, mxinf := -1 << 32 - 1, 1 << 32 - 1
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    manhattan := func (points [][]int, i, j int) int {
        return abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
    }
    maxManhattanDistance := func (points [][]int, remove int) []int {
        n := len(points)
        maxSum, minSum, maxDiff, minDiff := mninf, mxinf, mninf, mxinf
        maxSumIndex, minSumIndex, maxDiffIndex, minDiffIndex := -1, -1, -1, -1
        for i := 0; i < n; i++ {
            if i != remove {
                sum, diff := points[i][0]+points[i][1], points[i][0]-points[i][1]
                if sum > maxSum { maxSumIndex, maxSum = i, sum }
                if sum < minSum { minSumIndex, minSum = i, sum }
                if diff > maxDiff { maxDiffIndex, maxDiff = i, diff }
                if diff < minDiff { minDiffIndex, minDiff = i, diff }
            }
        }
        if max(maxSum - minSum, maxDiff - minDiff) == maxSum - minSum {
            return []int{ maxSumIndex, minSumIndex }
        }
        return []int{ maxDiffIndex, minDiffIndex }
    }
    m := maxManhattanDistance(points, -1)
    m1, m2 :=  maxManhattanDistance(points, m[0]), maxManhattanDistance(points, m[1])
    return min(manhattan(points, m1[0], m1[1]), manhattan(points, m2[0], m2[1]))
}

func main() {
// Example 1:
// Input: points = [[3,10],[5,15],[10,2],[4,4]]
// Output: 12
// Explanation:
// The maximum distance after removing each point is the following:
// After removing the 0th point the maximum distance is between points (5, 15) and (10, 2), which is |5 - 10| + |15 - 2| = 18.
// After removing the 1st point the maximum distance is between points (3, 10) and (10, 2), which is |3 - 10| + |10 - 2| = 15.
// After removing the 2nd point the maximum distance is between points (5, 15) and (4, 4), which is |5 - 4| + |15 - 4| = 12.
// After removing the 3rd point the maximum distance is between points (5, 15) and (10, 2), which is |5 - 10| + |15 - 2| = 18.
// 12 is the minimum possible maximum distance between any two points after removing exactly one point.
fmt.Println(minimumDistance([][]int{{3,10},{5,15},{10,2},{4,4}})) // 12
// Example 2:
// Input: points = [[1,1],[1,1],[1,1]]
// Output: 0
// Explanation:
// Removing any of the points results in the maximum distance between any two points of 0.
fmt.Println(minimumDistance([][]int{{1,1},{1,1},{1,1}})) // 0
}