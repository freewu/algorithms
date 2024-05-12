package main

// 353. Design Snake Game
// Design a Snake game that is played on a device with screen size height x width. 
// Play the game online if you are not familiar with the game.
// The snake is initially positioned at the top left corner (0, 0) with a length of 1 unit.
// You are given an array food where food[i] = (ri, ci) is the row and column position of a piece of food that the snake can eat. 
// When a snake eats a piece of food, its length and the game's score both increase by 1.
// Each piece of food appears one by one on the screen, meaning the second piece of food will not appear until the snake eats the first piece of food.
// When a piece of food appears on the screen, it is guaranteed that it will not appear on a block occupied by the snake.
// The game is over if the snake goes out of bounds (hits a wall) or if its head occupies a space that its body occupies after moving (i.e. a snake of length 4 cannot run into itself).

// Implement the SnakeGame class:
//     SnakeGame(int width, int height, int[][] food) Initializes the object with a screen of size height x width and the positions of the food.
//     int move(String direction) Returns the score of the game after applying one direction move by the snake. If the game is over, return -1.
    
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/13/snake.jpg" />
// Input
// ["SnakeGame", "move", "move", "move", "move", "move", "move"]
// [[3, 2, [[1, 2], [0, 1]]], ["R"], ["D"], ["R"], ["U"], ["L"], ["U"]]
// Output
// [null, 0, 0, 1, 1, 2, -1]
// Explanation
// SnakeGame snakeGame = new SnakeGame(3, 2, [[1, 2], [0, 1]]);
// snakeGame.move("R"); // return 0
// snakeGame.move("D"); // return 0
// snakeGame.move("R"); // return 1, snake eats the first piece of food. The second piece of food appears at (0, 1).
// snakeGame.move("U"); // return 1
// snakeGame.move("L"); // return 2, snake eats the second food. No more food appears.
// snakeGame.move("U"); // return -1, game over because snake collides with border
 
// Constraints:
//     1 <= width, height <= 10^4
//     1 <= food.length <= 50
//     food[i].length == 2
//     0 <= ri < height
//     0 <= ci < width
//     direction.length == 1
//     direction is 'U', 'D', 'L', or 'R'.
//     At most 104 calls will be made to move.

import "fmt"

// type SnakeGame struct {
//     snake [][]int
//     food [][]int
//     foodid int
//     width int
//     height int
// }

// var dirs = map[string][]int{
//     "U":{-1,0},
//     "D":{1,0},
//     "L":{0,-1},
//     "R":{0,1},
// }

// func Constructor(width int, height int, food [][]int) SnakeGame {
//     return SnakeGame{[][]int{{0,0}}, food, 0, width, height}
// }


// func (this *SnakeGame) Move(direction string) int {
//     dir := dirs[direction]
//     head := this.snake[len(this.snake)-1]
//     x, y := head[0], head[1]
//     nx, ny := x+dir[0], y+dir[1]
//     if nx < 0 || nx >= this.height || ny < 0 || ny >= this.width {
//         return -1
//     }
//     if this.foodid < len(this.food) && nx == this.food[this.foodid][0] && ny == this.food[this.foodid][1] {
//         this.snake = append(this.snake, []int{nx, ny})
//         this.foodid++
//     } else {
//         this.snake = this.snake[1:]
//         for _, p := range this.snake {
//             if nx == p[0] && ny == p[1] {
//                 return -1
//             }
//         }
//         this.snake = append(this.snake, []int{nx, ny})
//     }
//     return len(this.snake)-1
// }

type point struct {
    x, y int
}

type SnakeGame struct {
    width, height int
    q []point
    m map[point]bool
    foods []point
}

var delta = map[string]point{
    "U": {-1, 0},
    "D": {1, 0},
    "L": {0, -1},
    "R": {0, 1},
}

func Constructor(width int, height int, food [][]int) SnakeGame {
    q := []point{{0, 0}}
    m := map[point]bool{{0, 0}: true}
    foods := []point{}
    for _, f := range food {
        foods = append(foods, point{f[0], f[1]})
    }
    return SnakeGame{
        width: width,
        height: height,
        foods: foods,
        q: q,
        m: m,
    }
}

func (this *SnakeGame) Move(direction string) int {
    head := this.q[len(this.q)-1]
    next := point{head.x+delta[direction].x, head.y+delta[direction].y}
    if next.x < 0 || next.x >= this.height || next.y < 0 || next.y >= this.width {
        return -1
    }
    if next != this.q[0] && this.m[next] {
        return -1
    }
    if len(this.foods) > 0 && next == this.foods[0] {
        this.foods = this.foods[1:]
    } else {
        delete(this.m, this.q[0])
        this.q = this.q[1:]
    }
    this.q = append(this.q, next)
    this.m[next] = true
    return len(this.q)-1
}

/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */

func main() {
    // SnakeGame snakeGame = new SnakeGame(3, 2, [[1, 2], [0, 1]]);
    obj := Constructor(3,2,[][]int{{1,2}, {0,1}})
    fmt.Println(obj)
    // snakeGame.move("R"); // return 0
    fmt.Println(obj.Move("R")) // 0
    // snakeGame.move("D"); // return 0
    fmt.Println(obj.Move("D")) // 0
    // snakeGame.move("R"); // return 1, snake eats the first piece of food. The second piece of food appears at (0, 1).
    fmt.Println(obj.Move("R")) // 1
    // snakeGame.move("U"); // return 1
    fmt.Println(obj.Move("U")) // 1
    // snakeGame.move("L"); // return 2, snake eats the second food. No more food appears.
    fmt.Println(obj.Move("L")) // 2
    // snakeGame.move("U"); // return -1, game over because snake collides with border
    fmt.Println(obj.Move("U")) // -1
}