package main

// 2101. Detonate the Maximum Bombs
// You are given a list of bombs. 
// The range of a bomb is defined as the area where its effect can be felt. 
// This area is in the shape of a circle with the center as the location of the bomb.

// The bombs are represented by a 0-indexed 2D integer array bombs where bombs[i] = [xi, yi, ri]. 
// xi and yi denote the X-coordinate and Y-coordinate of the location of the ith bomb, 
// whereas ri denotes the radius of its range.

// You may choose to detonate a single bomb. 
// When a bomb is detonated, it will detonate all bombs that lie in its range. 
// These bombs will further detonate the bombs that lie in their ranges.

// Given the list of bombs, return the maximum number of bombs 
// that can be detonated if you are allowed to detonate only one bomb.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/11/06/desmos-eg-3.png" />
// Input: bombs = [[2,1,3],[6,1,4]]
// Output: 2
// Explanation:
// The above figure shows the positions and ranges of the 2 bombs.
// If we detonate the left bomb, the right bomb will not be affected.
// But if we detonate the right bomb, both bombs will be detonated.
// So the maximum bombs that can be detonated is max(1, 2) = 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/11/06/desmos-eg-2.png" />
// Input: bombs = [[1,1,5],[10,10,5]]
// Output: 1
// Explanation:
// Detonating either bomb will not detonate the other bomb, so the maximum number of bombs that can be detonated is 1.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/11/07/desmos-eg1.png" />
// Input: bombs = [[1,2,3],[2,3,1],[3,4,2],[4,5,3],[5,6,4]]
// Output: 5
// Explanation:
// The best bomb to detonate is bomb 0 because:
// - Bomb 0 detonates bombs 1 and 2. The red circle denotes the range of bomb 0.
// - Bomb 2 detonates bomb 3. The blue circle denotes the range of bomb 2.
// - Bomb 3 detonates bomb 4. The green circle denotes the range of bomb 3.
// Thus all 5 bombs are detonated.

// Constraints:
//     1 <= bombs.length <= 100
//     bombs[i].length == 3
//     1 <= xi, yi, ri <= 10^5

import "fmt"
import "math/bits"

// bfs
func maximumDetonation(bombs [][]int) int {
    res := 0
    for firstIndex, _ := range bombs {
        count, queue, visitedMap := 1,[]int{ firstIndex }, make(map[int]bool)
        visitedMap[firstIndex] = true
        for len(queue) > 0 {
            i := queue[0]
            queue = queue[1:]
            for j, _ := range bombs {
                if visitedMap[j] {
                    continue
                }
                xi, yi, ri := bombs[i][0], bombs[i][1], bombs[i][2]
                xj, yj := bombs[j][0], bombs[j][1]
                if (xi - xj) * (xi - xj) + (yi - yj) * (yi - yj) <= ri * ri {
                    queue = append(queue, j)
                    count++
                    visitedMap[j]=true
                }
            }
        }
        if res < count {
            res = count
        }
    }
    return res
}

func maximumDetonation1(bombs [][]int) (ans int) {
    n := len(bombs)
    f := make([]bitset, n)
    for i := range f {
        f[i] = newBitset(n)
        f[i].set(i)
    }
    for i, p := range bombs {
        x, y, r := p[0], p[1], p[2]
        for j, q := range bombs {
            dx := x - q[0]
            dy := y - q[1]
            if j != i && dx*dx+dy*dy <= r*r {
                f[i].set(j)
            }
        }
    }

    for k := range f {
        for i := range f {
            if f[i].has(k) {
                f[i].or(f[k])
            }
        }
    }

    for _, bs := range f {
        ans = max(ans, bs.onesCount())
    }
    return
}

const w = bits.UintSize

type bitset []uint

func newBitset(n int) bitset {
    return make(bitset, (n+w-1)/w)
}

func (b bitset) has(p int) bool {
    return b[p/w]&(1<<(p%w)) != 0
}

func (b bitset) set(p int) {
    b[p/w] |= 1 << (p % w)
}

func (b bitset) or(c bitset) {
    for i, x := range c {
        b[i] |= x
    }
}

func (b bitset) onesCount() (c int) {
    for _, x := range b {
        c += bits.OnesCount(x)
    }
    return
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/11/06/desmos-eg-3.png" />
    // Input: bombs = [[2,1,3],[6,1,4]]
    // Output: 2
    // Explanation:
    // The above figure shows the positions and ranges of the 2 bombs.
    // If we detonate the left bomb, the right bomb will not be affected.
    // But if we detonate the right bomb, both bombs will be detonated.
    // So the maximum bombs that can be detonated is max(1, 2) = 2.
    fmt.Println(maximumDetonation([][]int{{2,1,3},{6,1,4}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/11/06/desmos-eg-2.png" />
    // Input: bombs = [[1,1,5],[10,10,5]]
    // Output: 1
    // Explanation:
    // Detonating either bomb will not detonate the other bomb, so the maximum number of bombs that can be detonated is 1.
    fmt.Println(maximumDetonation([][]int{{1,1,5},{10,10,5}})) // 1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/11/07/desmos-eg1.png" />
    // Input: bombs = [[1,2,3],[2,3,1],[3,4,2],[4,5,3],[5,6,4]]
    // Output: 5
    // Explanation:
    // The best bomb to detonate is bomb 0 because:
    // - Bomb 0 detonates bombs 1 and 2. The red circle denotes the range of bomb 0.
    // - Bomb 2 detonates bomb 3. The blue circle denotes the range of bomb 2.
    // - Bomb 3 detonates bomb 4. The green circle denotes the range of bomb 3.
    // Thus all 5 bombs are detonated.
    fmt.Println(maximumDetonation([][]int{{1,2,3},{2,3,1},{3,4,2},{4,5,3},{5,6,4}})) // 5

    fmt.Println(maximumDetonation1([][]int{{2,1,3},{6,1,4}})) // 2
    fmt.Println(maximumDetonation1([][]int{{1,1,5},{10,10,5}})) // 1
    fmt.Println(maximumDetonation1([][]int{{1,2,3},{2,3,1},{3,4,2},{4,5,3},{5,6,4}})) // 5
}