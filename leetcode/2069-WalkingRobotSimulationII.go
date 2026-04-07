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

type Robot1 struct {
    Height, Width, Equation int
    Position []int
    Direction string
}

func Constructor1(width int, height int) Robot1 {
    return Robot1{ height, width, width * 2 + height * 2 - 4, []int{0, 0}, "East" }
}

func (this *Robot1) Step(num int) {
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

func (this *Robot1) GetPos() []int {
    return this.Position
}

func (this *Robot1) GetDir() string {   
    return this.Direction
}


type Robot struct {
    x int
    y int
    dir int
    width int
    height int
    step int
}

func Constructor(width int, height int) Robot {
    return Robot{
        x: 0,
        y: 0,
        dir: 0,
        width: width,
        height: height,
    }
}
func (this *Robot) Step(num int) {
    this.step += num
}

func (this *Robot) Solve()  {
    dirs := [][]int{ {1,0},{0,1},{-1,0},{0,-1} }
    num := this.step
    this.step = 0
    num = (num - 1) % ((this.width + this.height - 2) * 2) + 1
    for num > 0 {
        curx := this.x + dirs[this.dir][0] * num
        cury := this.y + dirs[this.dir][1] * num  
        if curx >= this.width {
            num -= this.width - 1 - this.x
            this.x = this.width - 1
            this.dir = (this.dir+1) % 4
        } else if curx < 0 {
            num -= this.x
            this.x = 0
            this.dir = (this.dir + 1) % 4
        } else if cury >= this.height {
            num -= this.height - 1 - this.y
            this.y=this.height-1
            this.dir = (this.dir + 1) % 4
        } else if cury < 0 {
            num -= this.y
            this.y = 0
            this.dir = (this.dir + 1) % 4 
        } else {
            num = 0
            this.x = curx
            this.y = cury
        }
    }
}

func (this *Robot) GetPos() []int {
    if this.step > 0 {
        this.Solve()
    }
    return []int{ this.x, this.y }
}

func (this *Robot) GetDir() string {
    if this.step > 0 {
        this.Solve()
    }
    mp := []string{ "East", "North", "West", "South" }
    return mp[this.dir]
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


    // Robot robot = new Robot(6, 3); // Initialize the grid and the robot at (0, 0) facing East.
    obj11 := Constructor1(6,3)
    fmt.Println(obj11)
    // robot.step(2);  // It moves two steps East to (2, 0), and faces East.
    obj11.Step(2) 
    fmt.Println(obj11) // (2, 0, East)
    // robot.step(2);  // It moves two steps East to (4, 0), and faces East.
    obj11.Step(2) 
    fmt.Println(obj11) // (4, 0, East)
    // robot.getPos(); // return [4, 0]
    fmt.Println(obj11.GetPos) // [4, 0]
    // robot.getDir(); // return "East"
    fmt.Println(obj11.GetDir) // "East"
    // robot.step(2);  // It moves one step East to (5, 0), and faces East.
    //                 // Moving the next step East would be out of bounds, so it turns and faces North.
    //                 // Then, it moves one step North to (5, 1), and faces North.
    obj11.Step(2) 
    fmt.Println(obj11) // (5, 1, North)
    // robot.step(1);  // It moves one step North to (5, 2), and faces North (not West).
    obj11.Step(1) 
    fmt.Println(obj11) // (5, 2, North)
    // robot.step(4);  // Moving the next step North would be out of bounds, so it turns and faces West.
    //                 // Then, it moves four steps West to (1, 2), and faces West.
    obj11.Step(4) 
    fmt.Println(obj11) // (1, 2, West)
    // robot.getPos(); // return [1, 2]
    fmt.Println(obj11.GetPos()) // [1, 2]
    // robot.getDir(); // return "West"
    fmt.Println(obj11.GetDir()) // West
}