package main

// 面试题 16.22. Langtons Ant LCCI
// An ant is sitting on an infinite grid of white and black squares. 
// It initially faces right. All squares are white initially.

// At each step, it does the following:
//     (1) At a white square, flip the color of the square, turn 90 degrees right (clockwise), and move forward one unit.
//     (2) At a black square, flip the color of the square, turn 90 degrees left (counter-clockwise), and move forward one unit.

// Write a program to simulate the first K moves that the ant makes and print the final board as a grid.

// The grid should be represented as an array of strings, where each element represents one row in the grid. 
// The black square is represented as 'X', and the white square is represented as '_', 
// the square which is occupied by the ant is represented as 'L', 'U', 'R', 'D', 
// which means the left, up, right and down orientations respectively. 
// You only need to return the minimum matrix that is able to contain all squares that are passed through by the ant.

// Example 1:
// Input: 0
// Output: ["R"]

// Example 2:
// Input: 2
// Output: [ "_X", "LX" ]

// Example 3:
// Input: 5
// Output: ["_U", "X_", "XX" ]

// Note:
//     K <= 100000

import "fmt"

func printKMoves(K int) []string {
    m, start, x, y := 10000, 5000, 5000, 5000 // 用来压缩 map，起始 x, y 都为5000防止出现负值，会导致map压缩无效
    move := [4][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} // 一定要设置成顺时针【右下左上】
    orientation := 0 // 开始朝向右
    black := make(map[int]bool) // 使用black来记录变黑的位置
    xMin, xMax, yMin, yMax := start, start, start, start
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= K; i++ {
        if k := m * x + y; black[k] {
            // 逆时针
            orientation--
            if orientation == -1 {
                orientation += 4
            }
            delete(black, k)
        } else {
            black[x*m+y] = true
            orientation++
        }
        orientation %= 4
        x += move[orientation][0]
        y += move[orientation][1]
        xMin, xMax = min(xMin, x), max(xMax, x)
        yMin, yMax = min(yMin, y), max(yMax, y)
    }
    // 默认全白
    grid := make([][]byte, yMax - yMin+1)
    for i := range grid {
        grid[i] = make([]byte, xMax-xMin+1)
        for j := range grid[i] {
            grid[i][j] = '_'
        }
    }
    // 根据black来涂黑
    for key := range black {
        _x, _y := key/m, key%m
        grid[yMax-_y][_x-xMin] = 'X'
    }
    // 最后一个点的朝向
    orientationBytes:= []byte{'R', 'D', 'L', 'U'}
    grid[yMax-y][x-xMin] = orientationBytes[orientation]
    res := make([]string, len(grid))
    for i := range grid {
        res[i] = string(grid[i])
    }
    return res
}

func printKMoves1(K int) []string {
    x, y, x1, x2, y1, y2, d := 0,0,0,0,0,0,1
    directions :=[]int{ -1, 0, 1, 0, -1 }
    black := map[[2]int]bool{}
    mp := map[int]byte{ 0: 'U', 1: 'R', 2: 'D', 3: 'L' }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for K > 0 {
        if _, ok := black[[2]int{x,y}]; !ok {
            black[[2]int{x,y}] = true
            d = (d + 1) % 4
            x, y = x + directions[d], y + directions[d + 1]
        } else {
            delete(black,[2]int{x, y})
            d = (d + 3) % 4
            x, y = x + directions[d], y + directions[d + 1]
        }
        x1, x2, y1, y2 = min(x1, x), max(x2, x), min(y1, y), max(y2, y)
        K--
    }
    n := x2 - x1 + 1
    res, arr :=  make([]string, n), make([][]byte, n)
    for i := 0 ; i < n; i++ {
        arr[i] = make([]byte, y2 - y1 + 1)
    }
    for i := x1; i <= x2; i++ {
        for j := y1; j <= y2; j++ {
            if !black[[2]int{i,j}] {
                arr[i - x1][j - y1] = '_'
            } else {
                arr[i - x1][j - y1] = 'X'
            }
        }
    }
    arr[x - x1][y - y1] = mp[d]
    for i := 0; i < n; i++ {
        res[i] = string(arr[i])
    }
    return res
}

func main() {
    // Example 1:
    // Input: 0
    // Output: ["R"]
    fmt.Println(printKMoves(0)) // ["R"]
    // Example 2:
    // Input: 2
    // Output: [ "_X", "LX" ]
    fmt.Println(printKMoves(2)) // ["_X", "LX"]
    // Example 3:
    // Input: 5
    // Output: ["_U", "X_", "XX" ]
    fmt.Println(printKMoves(5)) // ["_U", "X_", "XX" ]
    
    fmt.Println(printKMoves(1)) // [X D]
    fmt.Println(printKMoves(8)) // [_XX XLX XX_]
    fmt.Println(printKMoves(1024)) // [___________XX____ ____________XX___ _________X_XX_X__ __XX____XX_X__X__ _X___XXXXX_X_X___ XXX_____XX_X_____ X_X_X__XX__XX_X__ ____XXX_X_XX_X_X_ _____XX_X_X____X_ ___X_X_X_XX___X__ _X_X_X_X_X___X___ X__X_XX__X_XXX___ X_XX_XX___XXXXX__ _XX_X_X__XXXXXX__ __XXX__X_X_X_____ ____X_XXX_X__X___ ___X___X__X__XL_X ___X__________XXX ____X______XX__X_ _____XXXXXX__XX__]
    // fmt.Println(printKMoves(99999)) // 
    // fmt.Println(printKMoves(100000)) // 

    fmt.Println(printKMoves1(0)) // ["R"]
    fmt.Println(printKMoves1(2)) // ["_X", "LX"]
    fmt.Println(printKMoves1(5)) // ["_U", "X_", "XX" ]
    fmt.Println(printKMoves1(1)) // [X D]
    fmt.Println(printKMoves1(8)) // [_XX XLX XX_]
    fmt.Println(printKMoves1(1024)) // [___________XX____ ____________XX___ _________X_XX_X__ __XX____XX_X__X__ _X___XXXXX_X_X___ XXX_____XX_X_____ X_X_X__XX__XX_X__ ____XXX_X_XX_X_X_ _____XX_X_X____X_ ___X_X_X_XX___X__ _X_X_X_X_X___X___ X__X_XX__X_XXX___ X_XX_XX___XXXXX__ _XX_X_X__XXXXXX__ __XXX__X_X_X_____ ____X_XXX_X__X___ ___X___X__X__XL_X ___X__________XXX ____X______XX__X_ _____XXXXXX__XX__]
    // fmt.Println(printKMoves(99999)) // 
    // fmt.Println(printKMoves(100000)) // 
}