package main

// 1041. Robot Bounded In Circle
// On an infinite plane, a robot initially stands at (0, 0) and faces north. Note that:
//     The north direction is the positive direction of the y-axis.
//     The south direction is the negative direction of the y-axis.
//     The east direction is the positive direction of the x-axis.
//     The west direction is the negative direction of the x-axis.

// The robot can receive one of three instructions:
//     "G": go straight 1 unit.
//     "L": turn 90 degrees to the left (i.e., anti-clockwise direction).
//     "R": turn 90 degrees to the right (i.e., clockwise direction).
//     The robot performs the instructions given in order, and repeats them forever.

// Return true if and only if there exists a circle in the plane such that the robot never leaves the circle.

// Example 1:
// Input: instructions = "GGLLGG"
// Output: true
// Explanation: The robot is initially at (0, 0) facing the north direction.
// "G": move one step. Position: (0, 1). Direction: North.
// "G": move one step. Position: (0, 2). Direction: North.
// "L": turn 90 degrees anti-clockwise. Position: (0, 2). Direction: West.
// "L": turn 90 degrees anti-clockwise. Position: (0, 2). Direction: South.
// "G": move one step. Position: (0, 1). Direction: South.
// "G": move one step. Position: (0, 0). Direction: South.
// Repeating the instructions, the robot goes into the cycle: (0, 0) --> (0, 1) --> (0, 2) --> (0, 1) --> (0, 0).
// Based on that, we return true.

// Example 2:
// Input: instructions = "GG"
// Output: false
// Explanation: The robot is initially at (0, 0) facing the north direction.
// "G": move one step. Position: (0, 1). Direction: North.
// "G": move one step. Position: (0, 2). Direction: North.
// Repeating the instructions, keeps advancing in the north direction and does not go into cycles.
// Based on that, we return false.

// Example 3:
// Input: instructions = "GL"
// Output: true
// Explanation: The robot is initially at (0, 0) facing the north direction.
// "G": move one step. Position: (0, 1). Direction: North.
// "L": turn 90 degrees anti-clockwise. Position: (0, 1). Direction: West.
// "G": move one step. Position: (-1, 1). Direction: West.
// "L": turn 90 degrees anti-clockwise. Position: (-1, 1). Direction: South.
// "G": move one step. Position: (-1, 0). Direction: South.
// "L": turn 90 degrees anti-clockwise. Position: (-1, 0). Direction: East.
// "G": move one step. Position: (0, 0). Direction: East.
// "L": turn 90 degrees anti-clockwise. Position: (0, 0). Direction: North.
// Repeating the instructions, the robot goes into the cycle: (0, 0) --> (0, 1) --> (-1, 1) --> (-1, 0) --> (0, 0).
// Based on that, we return true.
 
// Constraints:
//     1 <= instructions.length <= 100
//     instructions[i] is 'G', 'L' or, 'R'.

import "fmt"

func isRobotBounded(instructions string) bool {
    x, y, degree := 0, 0, 0
    addSubtractXAndY := []int{1, 1, -1, -1}

    for _, i := range instructions {
        if i == 'R' {
            degree = (degree + 1) % 4
        } else if i == 'L' {
            degree = (degree + 3) % 4
        } else {
            if degree == 0 || degree == 2 {
                y += addSubtractXAndY[degree]
            } else {
                x += addSubtractXAndY[degree]
            }
        }
    }
    return x == 0 && y == 0 || degree != 0
}

func isRobotBounded1(instructions string) bool {
    x, y, degree := 0, 0, 0
    addSubtractXAndY := []int{1, 1, -1, -1}

    for _, c := range instructions {
        switch c {
        // "R": turn 90 degrees to the right (i.e., clockwise direction).
        case 'R':
            degree = (degree + 1) % 4
        // "L": turn 90 degrees to the left (i.e., anti-clockwise direction).
        case 'L': 
            degree = (degree + 3) % 4
        // "G": go straight 1 unit.
        case 'G':
            if degree == 0 || degree == 2 {
                y += addSubtractXAndY[degree]
            } else {
                x += addSubtractXAndY[degree]
            }
        }
    }
    return x == 0 && y == 0 || degree != 0
}

func isRobotBounded2(instructions string) bool {
    judgeSite := func (instructions string, s map[int]int, o *int) bool {
        for _, v := range instructions {
            switch v {
            case 'G':
                s[*o] += 1
            case 'L':
                *o++
                *o %= 4
            case 'R':
                *o += 3
                *o %= 4
            }
        }
        if s[0]-s[2] == 0 && s[1]-s[3] == 0 {
            return false
        }
        return true
    }
    s, o, i := map[int]int{0: 0, 1: 0, 2: 0, 3: 0}, 0, 0
    for i < 4 && judgeSite(instructions, s, &o) {
        i++
    }
    if i < 4 {
        return true
    }
    return false
}

func main() {
    // Explanation: The robot is initially at (0, 0) facing the north direction.
    // "G": move one step. Position: (0, 1). Direction: North.
    // "G": move one step. Position: (0, 2). Direction: North.
    // "L": turn 90 degrees anti-clockwise. Position: (0, 2). Direction: West.
    // "L": turn 90 degrees anti-clockwise. Position: (0, 2). Direction: South.
    // "G": move one step. Position: (0, 1). Direction: South.
    // "G": move one step. Position: (0, 0). Direction: South.
    // Repeating the instructions, the robot goes into the cycle: (0, 0) --> (0, 1) --> (0, 2) --> (0, 1) --> (0, 0).
    // Based on that, we return true.
    fmt.Println(isRobotBounded("GGLLGG")) // true

    // Explanation: The robot is initially at (0, 0) facing the north direction.
    // "G": move one step. Position: (0, 1). Direction: North.
    // "G": move one step. Position: (0, 2). Direction: North.
    // Repeating the instructions, keeps advancing in the north direction and does not go into cycles.
    // Based on that, we return false.
    fmt.Println(isRobotBounded("GG")) // false

    // Example 3:
    // Input: instructions = "GL"
    // Output: true
    // Explanation: The robot is initially at (0, 0) facing the north direction.
    // "G": move one step. Position: (0, 1). Direction: North.
    // "L": turn 90 degrees anti-clockwise. Position: (0, 1). Direction: West.
    // "G": move one step. Position: (-1, 1). Direction: West.
    // "L": turn 90 degrees anti-clockwise. Position: (-1, 1). Direction: South.
    // "G": move one step. Position: (-1, 0). Direction: South.
    // "L": turn 90 degrees anti-clockwise. Position: (-1, 0). Direction: East.
    // "G": move one step. Position: (0, 0). Direction: East.
    // "L": turn 90 degrees anti-clockwise. Position: (0, 0). Direction: North.
    // Repeating the instructions, the robot goes into the cycle: (0, 0) --> (0, 1) --> (-1, 1) --> (-1, 0) --> (0, 0).
    // Based on that, we return true.
    fmt.Println(isRobotBounded("GL")) // true

    fmt.Println(isRobotBounded1("GGLLGG")) // true
    fmt.Println(isRobotBounded1("GG")) // false
    fmt.Println(isRobotBounded1("GL")) // true

    fmt.Println(isRobotBounded2("GGLLGG")) // true
    fmt.Println(isRobotBounded2("GG")) // false
    fmt.Println(isRobotBounded2("GL")) // true
}