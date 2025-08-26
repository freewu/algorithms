package main

// 3661. Maximum Walls Destroyed by Robots
// There is an endless straight line populated with some robots and walls. 
// You are given integer arrays robots, distance, and walls:
//     1. robots[i] is the position of the ith robot.
//     2. distance[i] is the maximum distance the ith robot's bullet can travel.
//     3. walls[j] is the position of the jth wall.

// Every robot has one bullet that can either fire to the left or the right at most distance[i] meters.

// A bullet destroys every wall in its path that lies within its range. 
// Robots are fixed obstacles: if a bullet hits another robot before reaching a wall, it immediately stops at that robot and cannot continue.

// Return the maximum number of unique walls that can be destroyed by the robots.

// Notes:
//     1. A wall and a robot may share the same position; the wall can be destroyed by the robot at that position.
//     2. Robots are not destroyed by bullets.

// Example 1:
// Input: robots = [4], distance = [3], walls = [1,10]
// Output: 1
// Explanation:
// robots[0] = 4 fires left with distance[0] = 3, covering [1, 4] and destroys walls[0] = 1.
// Thus, the answer is 1.

// Example 2:
// Input: robots = [10,2], distance = [5,1], walls = [5,2,7]
// Output: 3
// Explanation:
// robots[0] = 10 fires left with distance[0] = 5, covering [5, 10] and destroys walls[0] = 5 and walls[2] = 7.
// robots[1] = 2 fires left with distance[1] = 1, covering [1, 2] and destroys walls[1] = 2.
// Thus, the answer is 3.

// Example 3:
// Input: robots = [1,2], distance = [100,1], walls = [10]
// Output: 0
// Explanation:
// In this example, only robots[0] can reach the wall, but its shot to the right is blocked by robots[1]; thus the answer is 0.

// Constraints:
//     1 <= robots.length == distance.length <= 10^5
//     1 <= walls.length <= 10^5
//     1 <= robots[i], walls[j] <= 10^9
//     1 <= distance[i] <= 10^5
//     All values in robots are unique
//     All values in walls are unique

import "fmt"
import "slices"
import "sort"

func maxWalls(robots []int, distance []int, walls []int) int {
    n := len(robots)
    type Pair struct{ x, dis int }
    arr := make([]Pair, n + 2)
    for i, x := range robots {
        arr[i] = Pair{ x, distance[i] }
    }
    arr[n + 1].x = 1 << 31 // 哨兵
    slices.SortFunc(arr, func(a, b Pair) int { return a.x - b.x })
    slices.Sort(walls)
    memo := make([][2]int, n + 1)
    for i := range memo {
        memo[i] = [2]int{-1, -1}
    }
    var dfs func(i, j int) int 
    dfs = func(i, j int) int {
        if i == 0 { return 0 }
        p := &memo[i][j]
        if *p != -1 { return *p }
        // 往左射，墙的坐标范围为 [leftX, a[i].x]
        leftX := max(arr[i].x - arr[i].dis, arr[i-1].x + 1) // +1 表示不能射到左边那个机器人
        left := sort.SearchInts(walls, leftX)
        cur := sort.SearchInts(walls, arr[i].x + 1)
        res := dfs(i-1, 0) + cur - left // 下标在 [left, cur-1] 中的墙都能摧毁
        // 往右射，墙的坐标范围为 [a[i].x, rightX]
        x2 := arr[i+1].x
        if j == 0 { // 右边那个机器人往左射
            x2 -= arr[i+1].dis
        }
        rightX := min(arr[i].x + arr[i].dis, x2-1) // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
        right := sort.SearchInts(walls, rightX + 1)
        cur = sort.SearchInts(walls, arr[i].x)
        res = max(res, dfs(i-1, 1) + right - cur) // 下标在 [cur, right-1] 中的墙都能摧毁
        *p = res
        return res
    }
    return dfs(n, 1)
}

func maxWalls1(robots []int, distance []int, walls []int) int {
    n, left, right, rightWall := len(robots), 0, 0, 0
    type Data struct {  pos, distance int }
    datas := make([]Data, n, n + 1)
    for i := 0; i < n; i++ {
        datas[i] = Data{robots[i], distance[i]}
    }
    slices.SortFunc(datas, func(a, b Data) int { return a.pos - b.pos })
    slices.Sort(walls)
    datas = append(datas, Data{ pos: 1 << 31 })
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, d := range datas[:len(robots)] {
        left1, left2, overlapping := left, right, false
        for len(walls) > 0 && walls[0] <= d.pos {
            if walls[0]+d.distance >= d.pos {
                left1++
                if walls[0] > rightWall {
                    left2++
                }
            }
            overlapping = walls[0] == d.pos
            walls = walls[1:]
        }
        right = max(left, right)
        if overlapping {
            right++
        }
        for _, wall := range walls {
            if wall < datas[i+1].pos && d.pos+d.distance >= wall {
                right++
                rightWall = wall
            } else {
                break
            }
        }
        left = max(left1, left2)
    }
    return max(left, right)
}

func main() {
    // Example 1:
    // Input: robots = [4], distance = [3], walls = [1,10]
    // Output: 1
    // Explanation:
    // robots[0] = 4 fires left with distance[0] = 3, covering [1, 4] and destroys walls[0] = 1.
    // Thus, the answer is 1.
    fmt.Println(maxWalls([]int{4}, []int{3}, []int{1,10})) // 1
    // Example 2:
    // Input: robots = [10,2], distance = [5,1], walls = [5,2,7]
    // Output: 3
    // Explanation:
    // robots[0] = 10 fires left with distance[0] = 5, covering [5, 10] and destroys walls[0] = 5 and walls[2] = 7.
    // robots[1] = 2 fires left with distance[1] = 1, covering [1, 2] and destroys walls[1] = 2.
    // Thus, the answer is 3.
    fmt.Println(maxWalls([]int{10,2}, []int{5,1}, []int{5,2,7})) // 3
    // Example 3:
    // Input: robots = [1,2], distance = [100,1], walls = [10]
    // Output: 0
    // Explanation:
    // In this example, only robots[0] can reach the wall, but its shot to the right is blocked by robots[1]; thus the answer is 0.
    fmt.Println(maxWalls([]int{1,2}, []int{100,1}, []int{10})) // 0

    fmt.Println(maxWalls1([]int{4}, []int{3}, []int{1,10})) // 1
    fmt.Println(maxWalls1([]int{10,2}, []int{5,1}, []int{5,2,7})) // 3
    fmt.Println(maxWalls1([]int{1,2}, []int{100,1}, []int{10})) // 0
}
