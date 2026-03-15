package main

// 3873. Maximum Points Activated with One Addition
// You are given a 2D integer array points, where points[i] = [xi, yi] represents the coordinates of the ith point. 
// All coordinates in points are distinct.

// If a point is activated, then all points that have the same x-coordinate or y-coordinate become activated as well.

// Activation continues until no additional points can be activated.

// You may add one additional point at any integer coordinate (x, y) not already present in points. 
// Activation begins by activating this newly added point.

// Return an integer denoting the maximum number of points that can be activated, including the newly added point.

// Example 1:
// Input: points = [[1,1],[1,2],[2,2]]
// Output: 4
// Explanation:
// Adding and activating a point such as (1, 3) causes activations:
// (1, 3) shares x = 1 with (1, 1) and (1, 2) -> (1, 1) and (1, 2) become activated.
// (1, 2) shares y = 2 with (2, 2) -> (2, 2) becomes activated.
// Thus, the activated points are (1, 3), (1, 1), (1, 2), (2, 2), so 4 points in total. We can show this is the maximum activated.

// Example 2:
// Input: points = [[2,2],[1,1],[3,3]]
// Output: 3
// Explanation:
// Adding and activating a point such as (1, 2) causes activations:
// (1, 2) shares x = 1 with (1, 1) -> (1, 1) becomes activated.
// (1, 2) shares y = 2 with (2, 2) -> (2, 2) becomes activated.
// Thus, the activated points are (1, 2), (1, 1), (2, 2), so 3 points in total. We can show this is the maximum activated.

// Example 3:
// Input: points = [[2,3],[2,2],[1,1],[4,5]]
// Output: 4
// Explanation:
// Adding and activating a point such as (2, 1) causes activations:
// (2, 1) shares x = 2 with (2, 3) and (2, 2) -> (2, 3) and (2, 2) become activated.
// (2, 1) shares y = 1 with (1, 1) -> (1, 1) becomes activated.
// Thus, the activated points are (2, 1), (2, 3), (2, 2), (1, 1), so 4 points in total.

// Constraints:
//     1 <= points.length <= 10^5
//     points[i] = [xi, yi]
//     -10^9 <= xi, yi <= 10^9
//     points contains all distinct coordinates.

import "fmt"

func maxActivated(points [][]int) int {
    mp := map[int]int{} // 哈希表并查集
    var find func(int) int
    find = func(x int) int {
        fx, ok := mp[x]
        if !ok {
            mp[x] = x
            fx = x
        }
        if fx != x {
            mp[x] = find(fx)
            return mp[x]
        }
        return x
    }
    const offset int = 3e9
    for _, p := range points {
        mp[find(p[0])] = find(p[1] + offset)   
    }
    size := map[int]int{}
    for _, p := range points {
        size[find(p[0])]++ // 统计连通块的大小
    }
    mx1, mx2 := 0, 0
    for _, sz := range size {
        if sz > mx1 {
            mx2 = mx1
            mx1 = sz
        } else if sz > mx2 {
            mx2 = sz
        }
    }
    return mx1 + mx2 + 1
}

func maxActivated1(points [][]int) int {
    n := len(points)
    parents,size := make([]int, n), make([]int, n)
    for i := range parents {
        parents[i],  size[i] = i, 1
    }
    var find func(i int, parents []int) int
    find = func(i int, parents []int) int {
        if parents[i] != i {
            parents[i] = find(parents[i], parents)
        }
        return parents[i]
    }
    union := func(i,j int, parents, size []int) {
        x,y := find(i, parents), find(j, parents)
        if x != y {
            parents[x] = y
            size[y] += size[x]
        }
    }
    mX, mY := make(map[int]int), make(map[int]int)
    for i, point := range points {
        x,y := point[0], point[1]
        if ptr,ok := mX[x]; ok {
            union(ptr, i, parents, size)
        } else {
            mX[x] = i
        }
        if ptr,ok := mY[y]; ok {
            union(ptr, i, parents, size)
        } else {
            mY[y] = i
        }
    }
    v1,v2 := 0,0
    for i := range parents {
        if parents[i] == i {
            if v2 < size[i] {
                v1,v2 = v2, size[i]
            } else if v1 < size[i] {
                v1 = size[i]
            }
        }
    }
    return v1 + v2 + 1
}

func main() {
    // Example 1:
    // Input: points = [[1,1],[1,2],[2,2]]
    // Output: 4
    // Explanation:
    // Adding and activating a point such as (1, 3) causes activations:
    // (1, 3) shares x = 1 with (1, 1) and (1, 2) -> (1, 1) and (1, 2) become activated.
    // (1, 2) shares y = 2 with (2, 2) -> (2, 2) becomes activated.
    // Thus, the activated points are (1, 3), (1, 1), (1, 2), (2, 2), so 4 points in total. We can show this is the maximum activated.
    fmt.Println(maxActivated([][]int{{1,1},{1,2},{2,2}})) // 4
    // Example 2:
    // Input: points = [[2,2],[1,1],[3,3]]
    // Output: 3
    // Explanation:
    // Adding and activating a point such as (1, 2) causes activations:
    // (1, 2) shares x = 1 with (1, 1) -> (1, 1) becomes activated.
    // (1, 2) shares y = 2 with (2, 2) -> (2, 2) becomes activated.
    // Thus, the activated points are (1, 2), (1, 1), (2, 2), so 3 points in total. We can show this is the maximum activated.
    fmt.Println(maxActivated([][]int{{2,2},{1,1},{3,3}})) // 3
    // Example 3:
    // Input: points = [[2,3],[2,2],[1,1],[4,5]]
    // Output: 4
    // Explanation:
    // Adding and activating a point such as (2, 1) causes activations:
    // (2, 1) shares x = 2 with (2, 3) and (2, 2) -> (2, 3) and (2, 2) become activated.
    // (2, 1) shares y = 1 with (1, 1) -> (1, 1) becomes activated.
    // Thus, the activated points are (2, 1), (2, 3), (2, 2), (1, 1), so 4 points in total.
    fmt.Println(maxActivated([][]int{{2,3},{2,2},{1,1},{4,5}})) // 4

    fmt.Println(maxActivated1([][]int{{1,1},{1,2},{2,2}})) // 4
    fmt.Println(maxActivated1([][]int{{2,2},{1,1},{3,3}})) // 3
    fmt.Println(maxActivated1([][]int{{2,3},{2,2},{1,1},{4,5}})) // 4
}