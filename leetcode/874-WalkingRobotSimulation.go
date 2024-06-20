package main

// 874. Walking Robot Simulation
// A robot on an infinite XY-plane starts at point (0, 0) facing north. 
// The robot can receive a sequence of these three possible types of commands:
//     -2: Turn left 90 degrees.
//     -1: Turn right 90 degrees.
//     1 <= k <= 9: Move forward k units, one unit at a time.

// Some of the grid squares are obstacles. The ith obstacle is at grid point obstacles[i] = (xi, yi).
// If the robot runs into an obstacle, then it will instead stay in its current location and move on to the next command.

// Return the maximum Euclidean distance that the robot ever gets from the origin squared (i.e. if the distance is 5, return 25).

// Note:
//     North means +Y direction.
//     East means +X direction.
//     South means -Y direction.
//     West means -X direction.
//     There can be obstacle in [0,0].
 
// Example 1:
// Input: commands = [4,-1,3], obstacles = []
// Output: 25
// Explanation: The robot starts at (0, 0):
// 1. Move north 4 units to (0, 4).
// 2. Turn right.
// 3. Move east 3 units to (3, 4).
// The furthest point the robot ever gets from the origin is (3, 4), which squared is 32 + 42 = 25 units away.

// Example 2:
// Input: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
// Output: 65
// Explanation: The robot starts at (0, 0):
// 1. Move north 4 units to (0, 4).
// 2. Turn right.
// 3. Move east 1 unit and get blocked by the obstacle at (2, 4), robot is at (1, 4).
// 4. Turn left.
// 5. Move north 4 units to (1, 8).
// The furthest point the robot ever gets from the origin is (1, 8), which squared is 12 + 82 = 65 units away.

// Example 3:
// Input: commands = [6,-1,-1,6], obstacles = []
// Output: 36
// Explanation: The robot starts at (0, 0):
// 1. Move north 6 units to (0, 6).
// 2. Turn right.
// 3. Turn right.
// 4. Move south 6 units to (0, 0).
// The furthest point the robot ever gets from the origin is (0, 6), which squared is 62 = 36 units away.
 
// Constraints:
//     1 <= commands.length <= 10^4
//     commands[i] is either -2, -1, or an integer in the range [1, 9].
//     0 <= obstacles.length <= 10^4
//     -3 * 10^4 <= xi, yi <= 3 * 10^4
//     The answer is guaranteed to be less than 2^31.

import "fmt"

func robotSim(commands []int, obstacles [][]int) int {
    type point struct {
        x, y int
    }
    dx, dy := [4]int{0, 1, 0, -1}, [4]int{1, 0, -1, 0}
    set := make(map[point]bool)
    for _, v := range obstacles {
        set[point{v[0], v[1]}] = true
    }
    res, x, y, direction := 0, 0, 0, 0
    for _, c := range commands {
        if c == -1 { // 向右转 90 度
            direction = (direction + 1) % 4
        } else if c == -2 { // 向左转 90 度
            direction = (direction - 1 + 4) % 4
        } else {
            for i := 0; i < c; i++ {
                nx, ny := x + dx[direction], y + dy[direction]
                if set[point{nx, ny}] { // 机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，并继续执行下一个命令
                    break
                } else {
                    x, y = nx, ny
                    ed := x * x + y * y // Euclidean Distance
                    if ed > res {
                        res = ed
                    }
                }
            }
        }
    }
    return res
}

func robotSim1(commands []int, obstacles [][]int) int {
    calcHash := func (x, y int) int { return (x + 30000) * 60000 + y + 30000 }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    obstaclesMap := map[int][]int{}
    dx, dy := []int{0, 1, 0, -1}, []int{1, 0, -1, 0} // N, E, S, W
    res, x, y, direction := 0, 0, 0, 0
    for _, v := range obstacles {
        curHash := calcHash(v[0], v[1])
        obstaclesMap[curHash] = v
    }
    for _, v := range commands {
        if v == -1 { // 向右转 90 度
            direction = (direction + 1) % 4
        } else if v == -2 { // 向左转 90 度
            direction = (direction - 1 + 4) % 4
        } else {
            for i := 0; i < v; i++ {
                nextX, nextY := x + dx[direction], y + dy[direction]
                tmp := calcHash(nextX, nextY)
                if _, ok := obstaclesMap[tmp]; ok {
                    break
                }
                x, y = nextX, nextY
                res = max(res, x * x + y * y)
            }
        }
    }
    return res
}

func robotSim2(commands []int, obstaclesIn [][]int) int {
    type point struct {
        x, y int
    }
    obstacles := map[point]bool{}
    for _, location := range obstaclesIn {
        obstacles[point{location[0], location[1]}] = true
    }
    currPos := point{0, 0}
    res, dirX, dirY := 0,0, 1
    for _, command := range commands {
        switch (command) {
            case -2: // 向左转 90 度
                if dirY != 0 {
                    dirX, dirY = dirY * -1, 0
                } else {
                    dirY, dirX = dirX, 0
                }
            case -1: // 向右转 90 度
                if dirY != 0 {
                    dirX, dirY = dirY, 0
                } else {
                    dirY, dirX = dirX * -1, 0
                }
            default:
                for i := 0; i < command; i++ {
                    newPos := point{currPos.x + dirX, currPos.y + dirY}
                    if _, ok := obstacles[newPos]; ok { // 机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，并继续执行下一个命令
                        break
                    }
                    currPos = newPos
                }
        }

        currDist := currPos.x * currPos.x + currPos.y * currPos.y
        if currDist > res {
            res = currDist
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: commands = [4,-1,3], obstacles = []
    // Output: 25
    // Explanation: The robot starts at (0, 0):
    // 1. Move north 4 units to (0, 4).
    // 2. Turn right.
    // 3. Move east 3 units to (3, 4).
    // The furthest point the robot ever gets from the origin is (3, 4), which squared is 32 + 42 = 25 units away.
    fmt.Println(robotSim([]int{4,-1,3},[][]int{})) // 25
    // Example 2:
    // Input: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
    // Output: 65
    // Explanation: The robot starts at (0, 0):
    // 1. Move north 4 units to (0, 4).
    // 2. Turn right.
    // 3. Move east 1 unit and get blocked by the obstacle at (2, 4), robot is at (1, 4).
    // 4. Turn left.
    // 5. Move north 4 units to (1, 8).
    // The furthest point the robot ever gets from the origin is (1, 8), which squared is 12 + 82 = 65 units away.
    fmt.Println(robotSim([]int{4,-1,4,-2,4},[][]int{{2,4}})) // 65
    // Example 3:
    // Input: commands = [6,-1,-1,6], obstacles = []
    // Output: 36
    // Explanation: The robot starts at (0, 0):
    // 1. Move north 6 units to (0, 6).
    // 2. Turn right.
    // 3. Turn right.
    // 4. Move south 6 units to (0, 0).
    // The furthest point the robot ever gets from the origin is (0, 6), which squared is 62 = 36 units away.
    fmt.Println(robotSim([]int{6,-1,-1,6},[][]int{})) // 36

    fmt.Println(robotSim1([]int{4,-1,3},[][]int{})) // 25
    fmt.Println(robotSim1([]int{4,-1,4,-2,4},[][]int{{2,4}})) // 65
    fmt.Println(robotSim1([]int{6,-1,-1,6},[][]int{})) // 36

    fmt.Println(robotSim2([]int{4,-1,3},[][]int{})) // 25
    fmt.Println(robotSim2([]int{4,-1,4,-2,4},[][]int{{2,4}})) // 65
    fmt.Println(robotSim2([]int{6,-1,-1,6},[][]int{})) // 36
}