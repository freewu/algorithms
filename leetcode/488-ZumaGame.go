package main

// 488. Zuma Game
// You are playing a variation of the game Zuma.

// In this variation of Zuma, there is a single row of colored balls on a board, 
// where each ball can be colored red 'R', yellow 'Y', blue 'B', green 'G', or white 'W'.
// You also have several colored balls in your hand.

// Your goal is to clear all of the balls from the board. On each turn:
//     Pick any ball from your hand and insert it in between two balls in the row or on either end of the row.
//     If there is a group of three or more consecutive balls of the same color, remove the group of balls from the board.
//     If this removal causes more groups of three or more of the same color to form, then continue removing each group until there are none left.
//     If there are no more balls on the board, then you win the game.
//     Repeat this process until you either win or do not have any more balls in your hand.

// Given a string board, representing the row of balls on the board, and a string hand, representing the balls in your hand, 
// return the minimum number of balls you have to insert to clear all the balls from the board. 
// If you cannot clear all the balls from the board using the balls in your hand, return -1.

// Example 1:
// Input: board = "WRRBBW", hand = "RB"
// Output: -1
// Explanation: It is impossible to clear all the balls. The best you can do is:
// - Insert 'R' so the board becomes WRRRBBW. WRRRBBW -> WBBW.
// - Insert 'B' so the board becomes WBBBW. WBBBW -> WW.
// There are still balls remaining on the board, and you are out of balls to insert.

// Example 2:
// Input: board = "WWRRBBWW", hand = "WRBRW"
// Output: 2
// Explanation: To make the board empty:
// - Insert 'R' so the board becomes WWRRRBBWW. WWRRRBBWW -> WWBBWW.
// - Insert 'B' so the board becomes WWBBBWW. WWBBBWW -> WWWW -> empty.
// 2 balls from your hand were needed to clear the board.

// Example 3:
// Input: board = "G", hand = "GGGGG"
// Output: 2
// Explanation: To make the board empty:
// - Insert 'G' so the board becomes GG.
// - Insert 'G' so the board becomes GGG. GGG -> empty.
// 2 balls from your hand were needed to clear the board.

// Constraints:
//     1 <= board.length <= 16
//     1 <= hand.length <= 5
//     board and hand consist of the characters 'R', 'Y', 'B', 'G', and 'W'.
//     The initial row of balls on the board will not have any groups of three or more consecutive balls of the same color.

import "fmt"
import "sort"
import "strings"

// 超出时间限制 55 / 57
// func findMinStep(board string, hand string) int {
//     q, mp, res := [][]string{{board, hand}}, make(map[string]bool), 0
//     var del3 func(str string) string
//     del3 = func(str string) string {
//         cnt := 1
//         for i := 1; i < len(str); i++ {
//             if str[i] == str[i-1] {
//                 cnt++
//             } else {
//                 if cnt >= 3 {
//                     return del3(str[0:i-cnt] + str[i:])
//                 }
//                 cnt = 1
//             }
//         }
//         if cnt >= 3 {
//             return str[0 : len(str)-cnt]
//         }
//         return str
//     }
//     for len(q) > 0 {
//         length := len(q)
//         res++
//         for length > 0 {
//             length--
//             cur := q[0]
//             q = q[1:]
//             curB, curH := cur[0], cur[1]
//             for i := 0; i < len(curB); i++ {
//                 for j := 0; j < len(curH); j++ {
//                     curB2 := del3(curB[0:i] + string(curH[j]) + curB[i:])
//                     curH2 := curH[0:j] + curH[j+1:]
//                     if len(curB2) == 0 {
//                         return res
//                     }
//                     if _, ok := mp[curB2 + curH2]; ok {
//                         continue
//                     }
//                     mp[curB2 + curH2] = true
//                     q = append(q, []string{curB2, curH2})
//                 }
//             }
//         }
//     }
//     return -1
// }

func findMinStep(board string, hand string) int {
    hand_slice := strings.Split(hand, "")
    sort.Strings(hand_slice)
    hand = strings.Join(hand_slice, "")
    q, visit, depth := [][]string{}, map[string]bool{}, 0
    q = append(q, []string{board, hand})
    key := fmt.Sprintf("%v_%v", board, hand)
    visit[key] = true
    var clean func (s string) string
    clean = func (s string) string {
        count := 1
        for i := 1; i < len(s); i++ {
            if s[i] == s[i-1] {
                count += 1
            } else {
                if count >= 3 {
                    return clean(s[:i-count] + s[i:])
                } else {
                    count = 1
                }
            }
        }
        if count < 3 {
            return s
        }
        return s[:len(s)-count]
    }
    for len(q) > 0 {
        depth++
        k := len(q)
        for k > 0 {
            k--
            cur_board := q[0][0]
            cur_hand := q[0][1]
            q = q[1:]
            for i := 0; i < len(cur_board)+1; i++ {
                for j := 0; j < len(cur_hand); j++ {
                    if j > 0 && cur_hand[j] == cur_hand[j-1] {
                        continue
                    }
                    if i > 0 && cur_board[i-1] == cur_hand[j] {
                        continue
                    }

                    choose := false
                    if i > 0 && i < len(cur_board) && cur_board[i-1] == cur_board[i] && cur_board[i] != cur_hand[j] {
                        choose = true
                    }
                    if i < len(cur_board) && cur_board[i] == cur_hand[j] {
                        choose = true
                    }
                    if choose {
                        new_board := clean(cur_board[:i] + string(cur_hand[j]) + cur_board[i:])
                        if new_board == "" {
                            return depth
                        }
                        new_hand := cur_hand[:j] + cur_hand[j+1:]
                        key := fmt.Sprintf("%v_%v", new_board, new_hand)
                        if visit[key] {
                            continue
                        }
                        visit[key] = true
                        q = append(q, []string{new_board, new_hand})
                    }
                }
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: board = "WRRBBW", hand = "RB"
    // Output: -1
    // Explanation: It is impossible to clear all the balls. The best you can do is:
    // - Insert 'R' so the board becomes WRRRBBW. WRRRBBW -> WBBW.
    // - Insert 'B' so the board becomes WBBBW. WBBBW -> WW.
    // There are still balls remaining on the board, and you are out of balls to insert.
    fmt.Println(findMinStep("WRRBBW","RB")) // -1
    // Example 2: 
    // Input: board = "WWRRBBWW", hand = "WRBRW"
    // Output: 2
    // Explanation: To make the board empty:
    // - Insert 'R' so the board becomes WWRRRBBWW. WWRRRBBWW -> WWBBWW.
    // - Insert 'B' so the board becomes WWBBBWW. WWBBBWW -> WWWW -> empty.
    // 2 balls from your hand were needed to clear the board.
    fmt.Println(findMinStep("WWRRBBWW","WRBRW")) // 2
    // Example 3:
    // Input: board = "G", hand = "GGGGG"
    // Output: 2
    // Explanation: To make the board empty:
    // - Insert 'G' so the board becomes GG.
    // - Insert 'G' so the board becomes GGG. GGG -> empty.
    // 2 balls from your hand were needed to clear the board.
    fmt.Println(findMinStep("G","GGGGG")) // 2
}