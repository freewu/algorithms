package main

// 1320. Minimum Distance to Type a Word Using Two Fingers

// <img src="https://assets.leetcode.com/uploads/2020/01/02/leetcode_keyboard.png" />

// You have a keyboard layout as shown above in the X-Y plane, 
// where each English uppercase letter is located at some coordinate.
//     For example, the letter 'A' is located at coordinate (0, 0), the letter 'B' is located at coordinate (0, 1), 
//     the letter 'P' is located at coordinate (2, 3) and the letter 'Z' is located at coordinate (4, 1).

// Given the string word, return the minimum total distance to type such string using only two fingers.

// The distance between coordinates (x1, y1) and (x2, y2) is |x1 - x2| + |y1 - y2|.

// Note that the initial positions of your two fingers are considered free so do not count towards your total distance, 
// also your two fingers do not have to start at the first letter or the first two letters.

// Example 1:
// Input: word = "CAKE"
// Output: 3
// Explanation: Using two fingers, one optimal way to type "CAKE" is: 
// Finger 1 on letter 'C' -> cost = 0 
// Finger 1 on letter 'A' -> cost = Distance from letter 'C' to letter 'A' = 2 
// Finger 2 on letter 'K' -> cost = 0 
// Finger 2 on letter 'E' -> cost = Distance from letter 'K' to letter 'E' = 1 
// Total distance = 3

// Example 2:
// Input: word = "HAPPY"
// Output: 6
// Explanation: Using two fingers, one optimal way to type "HAPPY" is:
// Finger 1 on letter 'H' -> cost = 0
// Finger 1 on letter 'A' -> cost = Distance from letter 'H' to letter 'A' = 2
// Finger 2 on letter 'P' -> cost = 0
// Finger 2 on letter 'P' -> cost = Distance from letter 'P' to letter 'P' = 0
// Finger 1 on letter 'Y' -> cost = Distance from letter 'A' to letter 'Y' = 4
// Total distance = 6

// Constraints:
//     2 <= word.length <= 300
//     word consists of uppercase English letters.

import "fmt"
import "slices"

func minimumDistance(word string) int {
    type Finger struct {
        Left, Right int
    }
    dp := map[Finger]int{}
    dp[Finger{26,26}] = 1
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    update := func (dp map[Finger]int, finger Finger, distance int) {
        if v, ok := dp[finger]; ok {
            dp[finger] = min(v, distance)
        } else {
            dp[finger] = distance
        }
    }
    getDistance := func(a, b int) int {
        if a == 26 || b == 26 { return 0 }
        row, col := a/6 - b/6, a % 6 - b % 6
        return abs(row)+abs(col)
    }
    for _, c := range word {
        now := int(c-'A')
        nextDp := map[Finger]int{}
        for finger, distance := range dp {
            newFinger := finger
            newFinger.Left = now
            update(nextDp, newFinger, distance + getDistance(finger.Left, now))
            newFinger = finger
            newFinger.Right = now
            update(nextDp, newFinger, distance + getDistance(finger.Right, now))
        }
        dp = nextDp
    }
    res := 1 << 31
    for _, distance := range dp {
        res = min(res, distance)
    }
    return res - 1
}

// 空间压缩,
// 核心思想, 在i位置时,两根手指一根必指在i-1的字符上,另外一根手指在其它字符上,所以在i位置,只需要1个额外状态!!: 另一跟手指志在哪里
func minimumDistance2(s string) int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    move := func(a, b int) int {
        x1, y1 := a / 6, a % 6
        x2, y2 := b / 6, b % 6
        return abs(x1 - x2) + abs(y1 - y2) // 横向移动距离 + 纵赂移动距离
    }
    n := len(s)
    // f[i][j] 在i位置上,j表示肯定一个手指指在i字符,另一根手指指在什么字符上.
    dp := make([][]int, n) // trick!! 因为确定一个手指指在i-1的字符位置上,所以可以省掉一个状态
    for i := range dp {
        dp[i] = make([]int, 26)
        for j := range dp[i] {
            dp[i][j] = 1 << 31 // 注意!! 必须初始化为无效值, 因为如果[0:i-1]不含有某些字符,那么另一根手指就不能随便指向了,应该是无效的
        }
    }
    for j := range dp[0] { // 在0位置上,1根手指指在0字符不消耗移动, 另外一个手指可以指在任意字符,也不消耗移动次数(是为了将来转移到i位置不消耗move次数,就好像原来就呆在这个位置)
        dp[0][j] = 0
    }
    for i := 1; i < n; i++ {
        cur, pre := int(s[i]-'A'), int(s[i-1]-'A')
        for j := 0; j < 26; j++ {
            // 在i-1字符的手指移动到 i字符,指向j的另一根手指不动
            dp[i][j] = dp[i-1][j] + move(cur, pre) // bug!! 枚举的另一根手指,但移动的是在i-1位置的手机,所以move(cur,pre)而不是move(cur,j)
        }
        // 指在i-1字符的手指不动, 另一根手指指向i+1字符 => 那么当前j就是i-1的状态了
        for k := 0; k < 26; k++ {
            dp[i][pre] = min(dp[i][pre], dp[i-1][k] + move(cur, k))
        }
    }
    return slices.Min(dp[n-1])
}

// 要点是如何定义合适使用第二根手指
// dfs(i,AChar,BChar)
// 在i位置,可以做的决定,1.使用两根手指的任意一个 2.第一次使用第二根手指(那么之前只能使用第一根手指)
func minimumDistance1(s string) int {
    coordinate := func(ch byte) (int, int) {
        idx := ch - 'A'
        return int(idx/6), int(idx%6) // 每行 6个
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    moveCnt := func(a, b byte) int {
        x1, y1 := coordinate(a)
        x2, y2 := coordinate(b)
        return abs(x1 - x2) + abs(y1 - y2) // byte uint8类型,如果相减会溢出
    }
    n := len(s)
    memo := make([][26][26][2]int, n) // [1,n-1]
    for i := range memo {
        for j := range memo[i] {
            for k := range memo[i][j] {
                memo[i][j][k] = [2]int{-1, -1}
            }
        }
    }
    var dfs func(i int, a, b byte, startB int) int
    dfs = func(i int, a, b byte, startB int) int {
        if i >= n { return 0 }
        cur := s[i]
        p := &memo[i][a-'A'][b-'A'][startB] // bug!! 题目给出的是字符串是由大写字母组成
        if *p != -1 {                       // 重大bug!! 开始memo没有初始化为-1,造成TLE,因为相同的字符串比如"ADDDDDDDDDD"会让之后的选择不增加移动次数,造成缓存判断失效从而重复计算
            return *p
        }
        res := moveCnt(a, cur) + dfs(i+1, cur, b, startB) // 使用A手指
        if startB == 1 { // 尝试使用B手指
            res = min(res, moveCnt(b, cur)+dfs(i+1, a, cur, 1))
        } else {
            res = min(res, 0+dfs(i+1, a, cur, 1)) // 第一次放置B手指没有一点次数
        }
        *p = res
        return res
    }
    return dfs(1, s[0], 'A', 0) // 开始时,必然使用A手指放在0位置字符上, b初始化为'a',防止memo越界
}

func main() {
    // Example 1:
    // Input: word = "CAKE"
    // Output: 3
    // Explanation: Using two fingers, one optimal way to type "CAKE" is: 
    // Finger 1 on letter 'C' -> cost = 0 
    // Finger 1 on letter 'A' -> cost = Distance from letter 'C' to letter 'A' = 2 
    // Finger 2 on letter 'K' -> cost = 0 
    // Finger 2 on letter 'E' -> cost = Distance from letter 'K' to letter 'E' = 1 
    // Total distance = 3 
    fmt.Println(minimumDistance("CAKE")) // 3
    // Example 2:
    // Input: word = "HAPPY"
    // Output: 6
    // Explanation: Using two fingers, one optimal way to type "HAPPY" is:
    // Finger 1 on letter 'H' -> cost = 0
    // Finger 1 on letter 'A' -> cost = Distance from letter 'H' to letter 'A' = 2
    // Finger 2 on letter 'P' -> cost = 0
    // Finger 2 on letter 'P' -> cost = Distance from letter 'P' to letter 'P' = 0
    // Finger 1 on letter 'Y' -> cost = Distance from letter 'A' to letter 'Y' = 4
    // Total distance = 6
    fmt.Println(minimumDistance("HAPPY")) // 6

    fmt.Println(minimumDistance1("CAKE")) // 3
    fmt.Println(minimumDistance1("HAPPY")) // 6

    fmt.Println(minimumDistance2("CAKE")) // 3
    fmt.Println(minimumDistance2("HAPPY")) // 6
}