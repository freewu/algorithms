package main

// 2069. Walking Robot Simulation II
// A width x height grid is on an XY-plane with the bottom-left cell at (0, 0) and the top-right cell at (width - 1, height - 1). 
// The grid is aligned with the four cardinal directions ("North", "East", "South", and "West"). 
// A robot is initially at cell (0, 0) facing direction "East".

// The robot can be instructed to move for a specific number of steps. 
// For each step, it does the following.
//     1. Attempts to move forward one cell in the direction it is facing.
//     2. If the cell the robot is moving to is out of bounds, 
//        the robot instead turns 90 degrees counterclockwise and retries the step.

// After the robot finishes moving the number of steps required, it stops and awaits the next instruction.

// Implement the Robot class:
//     Robot(int width, int height) 
//         Initializes the width x height grid with the robot at (0, 0) facing "East".
//     void step(int num) 
//         Instructs the robot to move forward num steps.
//     int[] getPos() 
//         Returns the current cell the robot is at, as an array of length 2, [x, y].
//     String getDir() 
//         Returns the current direction of the robot, "North", "East", "South", or "West".

// Example 1:
// example-1
// Input
// ["Robot", "step", "step", "getPos", "getDir", "step", "step", "step", "getPos", "getDir"]
// [[6, 3], [2], [2], [], [], [2], [1], [4], [], []]
// Output
// [null, null, null, [4, 0], "East", null, null, null, [1, 2], "West"]
// Explanation
// Robot robot = new Robot(6, 3); // Initialize the grid and the robot at (0, 0) facing East.
// robot.step(2);  // It moves two steps East to (2, 0), and faces East.
// robot.step(2);  // It moves two steps East to (4, 0), and faces East.
// robot.getPos(); // return [4, 0]
// robot.getDir(); // return "East"
// robot.step(2);  // It moves one step East to (5, 0), and faces East.
//                 // Moving the next step East would be out of bounds, so it turns and faces North.
//                 // Then, it moves one step North to (5, 1), and faces North.
// robot.step(1);  // It moves one step North to (5, 2), and faces North (not West).
// robot.step(4);  // Moving the next step North would be out of bounds, so it turns and faces West.
//                 // Then, it moves four steps West to (1, 2), and faces West.
// robot.getPos(); // return [1, 2]
// robot.getDir(); // return "West"

// Constraints:
//     2 <= width, height <= 100
//     1 <= num <= 10^5
//     At most 10^4 calls in total will be made to step, getPos, and getDir.

import "fmt"

type Robot struct {
    Height, Width, Equation int
    Position []int
    Direction string
}

func Constructor(width int, height int) Robot {
    return Robot{ height, width, width * 2 + height * 2 - 4, []int{0, 0}, "East" }
}

func (this *Robot) Step(num int) {
    num %= this.Equation
    if num == 0 { num = this.Equation }
    for i := 0; i < num; i++ {
        switch { // 到了边缘需要转向
        case this.Direction == "South" && this.Position[0] == 0 && this.Position[1] == 0:
            this.Direction = "East"
        case this.Direction == "North" && this.Position[0] == this.Width - 1 && this.Position[1] == this.Height - 1:
            this.Direction = "West"
        case this.Direction == "East" && this.Position[0] == this.Width - 1 && this.Position[1] == 0:
            this.Direction = "North"
        case this.Direction == "West" && this.Position[0] == 0 && this.Position[1] == this.Height - 1:
            this.Direction = "South"
        }
        switch this.Direction {
        case "East":  this.Position[0]++
        case "West":  this.Position[0]--
        case "North": this.Position[1]++
        case "South": this.Position[1]--
        }
    }
}

func (this *Robot) GetPos() []int {
    return this.Position
}

func (this *Robot) GetDir() string {
    return this.Direction
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Step(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */

func main() {
    // Robot robot = new Robot(6, 3); // Initialize the grid and the robot at (0, 0) facing East.
    obj := Constructor(6,3)
    fmt.Println(obj)
    // robot.step(2);  // It moves two steps East to (2, 0), and faces East.
    obj.Step(2) 
    fmt.Println(obj) // (2, 0, East)
    // robot.step(2);  // It moves two steps East to (4, 0), and faces East.
    obj.Step(2) 
    fmt.Println(obj) // (4, 0, East)
    // robot.getPos(); // return [4, 0]
    fmt.Println(obj.GetPos) // [4, 0]
    // robot.getDir(); // return "East"
    fmt.Println(obj.GetDir) // "East"
    // robot.step(2);  // It moves one step East to (5, 0), and faces East.
    //                 // Moving the next step East would be out of bounds, so it turns and faces North.
    //                 // Then, it moves one step North to (5, 1), and faces North.
    obj.Step(2) 
    fmt.Println(obj) // (5, 1, North)
    // robot.step(1);  // It moves one step North to (5, 2), and faces North (not West).
    obj.Step(1) 
    fmt.Println(obj) // (5, 2, North)
    // robot.step(4);  // Moving the next step North would be out of bounds, so it turns and faces West.
    //                 // Then, it moves four steps West to (1, 2), and faces West.
    obj.Step(4) 
    fmt.Println(obj) // (1, 2, West)
    // robot.getPos(); // return [1, 2]
    fmt.Println(obj.GetPos()) // [1, 2]
    // robot.getDir(); // return "West"
    fmt.Println(obj.GetDir()) // West
}