package main

// 864. Shortest Path to Get All Keys
// You are given an m x n grid grid where:
//     '.' is an empty cell.
//     '#' is a wall.
//     '@' is the starting point.
//     Lowercase letters represent keys.
//     Uppercase letters represent locks.

// You start at the starting point and one move consists of walking one space in one of the four cardinal directions. 
// You cannot walk outside the grid, or walk into a wall.

// If you walk over a key, you can pick it up and you cannot walk over a lock unless you have its corresponding key.
//     For some 1 <= k <= 6, there is exactly one lowercase and one uppercase letter of the first k letters of the English alphabet in the grid. 
//     This means that there is exactly one key for each lock, 
//     and one lock for each key; and also that the letters used to represent the keys 
//     and locks were chosen in the same order as the English alphabet.

// Return the lowest number of moves to acquire all keys. If it is impossible, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/23/lc-keys2.jpg" />
// Input: grid = ["@.a..","###.#","b.A.B"]
// Output: 8
// Explanation: Note that the goal is to obtain all the keys not to open all the locks.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/07/23/lc-key2.jpg" />
// Input: grid = ["@..aA","..B#.","....b"]
// Output: 6

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/07/23/lc-keys3.jpg" />
// Input: grid = ["@Aa"]
// Output: -1

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 30
//     grid[i][j] is either an English letter, '.', '#', or '@'. 
//     There is exactly one '@' in the grid.
//     The number of keys in the grid is in the range [1, 6].
//     Each key in the grid is unique.
//     Each key in the grid has a matching lock.

import "fmt"
import "unicode"

// bfs
func shortestPathAllKeys(grid []string) int {
    type Position struct {
        x,y int
        keys int
    }
    keysCount, lastQueue := 0, []Position{}
    m, n := len(grid), len(grid[0])
    isKey := func(x, y int) bool { return grid[x][y] >= 'a' && grid[x][y] <= 'k'; } // 判断是否是钥匙 小写为钥匙 大写是锁
    for i, row := range grid {
        for j, cell := range row {
            if cell == '@' { // 找到起点
                lastQueue = []Position{Position{x:i, y:j, keys:0}}
            } else if isKey(i,j) {
                keysCount++ // 钥匙个数
            }
        }
    }
    seen, queue, steps := make(map[Position]bool), make([]Position, 0), 0
    for len(lastQueue) > 0 {
        for _, pos := range lastQueue {
            x,y := pos.x, pos.y
            if pos.keys == ((1 << keysCount) - 1) { return steps - 1 } // check if all keys found 
            if x < 0 || x == m || y < 0 || y == n { continue } // out of bounds
            if grid[x][y] == '#' { continue } // Wall
            if seen[pos] { continue } // Already seen
            
            if grid[x][y] >= 'A' && grid[x][y] <= 'K' { // lock
                if pos.keys & (1 << int(grid[x][y] -'A')) == 0 { // 还未找到 钥匙
                    continue // don't have the key
                }
            }
            if isKey(x, y) { // key
                pos.keys |= 1 << int(grid[x][y] - 'a')
            }
            seen[pos] = true
            queue = append(queue, Position{x:x+1, y:y, keys:pos.keys}) // 右
            queue = append(queue, Position{x:x-1, y:y, keys:pos.keys}) // 左
            queue = append(queue, Position{x:x, y:y+1, keys:pos.keys}) // 下
            queue = append(queue, Position{x:x, y:y-1, keys:pos.keys}) // 上
        }
        lastQueue, queue = queue, lastQueue[:0]
        steps++
    }
    return -1 // can't find a solution
}

func shortestPathAllKeys1(grid []string) int {
    // 分层最短路问题, 图上的vertex不在对应grid上的格子,而是格子+一个状态(BFS)
    // 主要是开锁问题,那么状态是有无对应的钥匙
    // '.' 代表一个空房间
    // '#' 代表一堵墙
    // '@' 是起点
    // 小写字母代表钥匙
    // 大写字母代表锁
    // 钥匙的数目范围是 [1, 6]
    keyMask := 0
    m := len(grid)
    n := len(grid[0])
    mat := make([][]byte, m)

    keyCnt := 0
    starts := make([][2]int, 0) // 0:xy(x*n+y) 1:state
    // 将grid转为matrix,收集起点,收集钥匙
    for i, row := range grid {
        bs := []byte(row)
        mat[i] = bs

        for j, b := range bs {
            if unicode.IsLower(rune(b)) {
                keyMask |= 1 << (b - 'a')
                keyCnt++
            }
            if b == '@' { //题目并未明确说是一个起点, 是否可能多个起点?
                starts = append(starts, [2]int{i*n + j, 0})
            }
        }
    }

    mxState := 1 << keyCnt
    q := make([][2]int, m*n*mxState) // 0:xy(x*n+y) 1:state
    visit := make([][]bool, m*n)
    for i := range visit {
        visit[i] = make([]bool, mxState)
    }
    ql, qr := 0, 0
    for _, start := range starts {
        q[qr] = start
        qr++
        visit[start[0]][start[1]] = true
    }

    dir4 := []struct{ x, y int }{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
    // bfs
    level := 0
    for ql < qr {  // bug!! ql,qr写反
        level++ //  level从0开始,此时指向下一层

        size := qr - ql
        for a := 0; a < size; a++ {
            cur := q[ql]
            ql++   // bug!! ql忘了++

            xy, s := cur[0], cur[1]
            x, y := xy/n, xy%n
            for _, d := range dir4 {
                nx, ny, ns := d.x+x, d.y+y, s
                nxy := nx*n + ny
                if 0 <= nx && nx < m && 0 <= ny && ny < n && !visit[nxy][ns] { // 重大bug!!!! 对于bfs来说,一定防止走visit
                    nc := mat[nx][ny]
                    if nc == '#' || // 遇到墙
                        unicode.IsUpper(rune(nc)) && ns&(1<<(nc-'A')) == 0 { // 遇到锁,但没钥匙
                        continue // 重大bug!!!! 消灭一个尝试,进行下一个尝试是用continue
                    }
                    if unicode.IsLower(rune(nc)) { // 遇到钥匙,持有钥匙
                        ns |= 1 << (nc - 'a')
                        // 剪枝,收集到所有钥匙提前退出
                        if ns == keyMask {
                            return level  // bug!!不是level+1, 不是问的第几层,而是问的走多少步, 从0层到level层需要走level步,共level+1层
                        }
                    }
                    visit[nxy][ns] = true
                    q[qr] = [2]int{nxy, ns}
                    qr++
                }
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/07/23/lc-keys2.jpg" />
    // Input: grid = ["@.a..","###.#","b.A.B"]
    // Output: 8
    // Explanation: Note that the goal is to obtain all the keys not to open all the locks.
    fmt.Println(shortestPathAllKeys([]string{"@.a..","###.#","b.A.B"})) // 9
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/07/23/lc-key2.jpg" />
    // Input: grid = ["@..aA","..B#.","....b"]
    // Output: 6
    fmt.Println(shortestPathAllKeys([]string{"@..aA","..B#.","....b"})) /// 6
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/07/23/lc-keys3.jpg" />
    // Input: grid = ["@Aa"]
    // Output: -1
    fmt.Println(shortestPathAllKeys([]string{"@Aa"})) // -1

    fmt.Println(shortestPathAllKeys1([]string{"@.a..","###.#","b.A.B"})) // 9
    fmt.Println(shortestPathAllKeys1([]string{"@..aA","..B#.","....b"})) /// 6
    fmt.Println(shortestPathAllKeys1([]string{"@Aa"})) // -1
}