package main

// 1244. Design A Leaderboard
// Design a Leaderboard class, which has 3 functions:
//     addScore(playerId, score): 
//         Update the leaderboard by adding score to the given player's score. 
//         If there is no player with such id in the leaderboard, add him to the leaderboard with the given score.
//     top(K): 
//         Return the score sum of the top K players.
//     reset(playerId): 
//         Reset the score of the player with the given id to 0 (in other words erase it from the leaderboard). 
//         It is guaranteed that the player was added to the leaderboard before calling this function.

// Initially, the leaderboard is empty.

// Example 1:
// Input: 
// ["Leaderboard","addScore","addScore","addScore","addScore","addScore","top","reset","reset","addScore","top"]
// [[],[1,73],[2,56],[3,39],[4,51],[5,4],[1],[1],[2],[2,51],[3]]
// Output: 
// [null,null,null,null,null,null,73,null,null,null,141]
// Explanation: 
// Leaderboard leaderboard = new Leaderboard ();
// leaderboard.addScore(1,73);   // leaderboard = [[1,73]];
// leaderboard.addScore(2,56);   // leaderboard = [[1,73],[2,56]];
// leaderboard.addScore(3,39);   // leaderboard = [[1,73],[2,56],[3,39]];
// leaderboard.addScore(4,51);   // leaderboard = [[1,73],[2,56],[3,39],[4,51]];
// leaderboard.addScore(5,4);    // leaderboard = [[1,73],[2,56],[3,39],[4,51],[5,4]];
// leaderboard.top(1);           // returns 73;
// leaderboard.reset(1);         // leaderboard = [[2,56],[3,39],[4,51],[5,4]];
// leaderboard.reset(2);         // leaderboard = [[3,39],[4,51],[5,4]];
// leaderboard.addScore(2,51);   // leaderboard = [[2,51],[3,39],[4,51],[5,4]];
// leaderboard.top(3);           // returns 141 = 51 + 51 + 39;
 
// Constraints:
//     1 <= playerId, K <= 10000
//     It's guaranteed that K is less than or equal to the current number of players.
//     1 <= score <= 100
//     There will be at most 1000 function calls.

import "fmt"

type Leaderboard struct {
    board map[int]int // 用 map 记录选手对应的分数
}

func Constructor() Leaderboard {
    return Leaderboard{board:map[int]int{}}
}

func (this *Leaderboard) AddScore(playerId int, score int) {
    this.board[playerId] += score
}

func (this *Leaderboard) Top(k int) int {
    res, index, a := 0, 0, make([]int, len(this.board))
    for _, i := range this.board {
        a[index] = i
        index++
    }
    topK(a, k)
    for i := len(a) - k; i < len(a); i++ {
        res += a[i]
    }
    return res
}

func (this *Leaderboard) Reset(playerId int) {
    this.board[playerId] = 0
}

// 快速选择topk
func topK(a []int, k int) {
    l, h := 0, len(a)-1
    temp := a[l]
    for l < h {
        for h > l && a[h] >= temp {
            h--
        }
        a[h], a[l] = a[l], a[h]
        for l < h && a[l] < temp {
            l++
        }

        a[h], a[l] = a[l], a[h]
    }
    if l+k == len(a) {
        return
    }
    if l+k < len(a) {
        topK(a[l+1:], k)
        return
    }
    topK(a[:l], k+l-len(a))
}

/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */

func main() {
    // Leaderboard leaderboard = new Leaderboard ();
    obj := Constructor()
    fmt.Println(obj)
    // leaderboard.addScore(1,73);   // leaderboard = [[1,73]];
    obj.AddScore(1,73)
    fmt.Println(obj)
    // leaderboard.addScore(2,56);   // leaderboard = [[1,73],[2,56]];
    obj.AddScore(2,56)
    fmt.Println(obj)
    // leaderboard.addScore(3,39);   // leaderboard = [[1,73],[2,56],[3,39]];
    obj.AddScore(3,39)
    fmt.Println(obj)
    // leaderboard.addScore(4,51);   // leaderboard = [[1,73],[2,56],[3,39],[4,51]];
    obj.AddScore(4,51)
    fmt.Println(obj)
    // leaderboard.addScore(5,4);    // leaderboard = [[1,73],[2,56],[3,39],[4,51],[5,4]];
    obj.AddScore(15,4)
    fmt.Println(obj)
    // leaderboard.top(1);           // returns 73;
    fmt.Println(obj.Top(1)) // 73
    // leaderboard.reset(1);         // leaderboard = [[2,56],[3,39],[4,51],[5,4]];
    obj.Reset(1)
    fmt.Println(obj)
    // leaderboard.reset(2);         // leaderboard = [[3,39],[4,51],[5,4]];
    obj.Reset(2)
    fmt.Println(obj)
    // leaderboard.addScore(2,51);   // leaderboard = [[2,51],[3,39],[4,51],[5,4]];
    obj.AddScore(2,51)
    fmt.Println(obj)
    // leaderboard.top(3);           // returns 141 = 51 + 51 + 39;
    fmt.Println(obj.Top(3)) // 141
}