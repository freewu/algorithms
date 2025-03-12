package main

// 3443. Maximum Manhattan Distance After K Changes
// You are given a string s consisting of the characters 'N', 'S', 'E', and 'W', 
// where s[i] indicates movements in an infinite grid:
//     'N' : Move north by 1 unit.
//     'S' : Move south by 1 unit.
//     'E' : Move east by 1 unit.
//     'W' : Move west by 1 unit.

// Initially, you are at the origin (0, 0). 
// You can change at most k characters to any of the four directions.

// Find the maximum Manhattan distance from the origin that can be achieved at any time while performing the movements in order.

// The Manhattan Distance between two cells (xi, yi) and (xj, yj) is |xi - xj| + |yi - yj|.

// Example 1:
// Input: s = "NWSE", k = 1
// Output: 3
// Explanation:
// Change s[2] from 'S' to 'N'. The string s becomes "NWNE".
// Movement	Position (x, y)	Manhattan Distance	Maximum
// s[0] == 'N'	(0, 1)	0 + 1 = 1	1
// s[1] == 'W'	(-1, 1)	1 + 1 = 2	2
// s[2] == 'N'	(-1, 2)	1 + 2 = 3	3
// s[3] == 'E'	(0, 2)	0 + 2 = 2	3
// The maximum Manhattan distance from the origin that can be achieved is 3. Hence, 3 is the output.

// Example 2:
// Input: s = "NSWWEW", k = 3
// Output: 6
// Explanation:
// Change s[1] from 'S' to 'N', and s[4] from 'E' to 'W'. The string s becomes "NNWWWW".
// The maximum Manhattan distance from the origin that can be achieved is 6. Hence, 6 is the output.

// Constraints:
//     1 <= s.length <= 10^5
//     0 <= k <= s.length
//     s consists of only 'N', 'S', 'E', and 'W'.

import "fmt"

func maxDistance(s string, k int) int {
    res, directions := 0, [][2]byte{ {'N', 'E'}, {'N', 'W'}, {'S', 'E'}, {'S', 'W'} }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, d := range directions {
        rem, count := k, 0
        for i := range s {
            if s[i] == d[0] || s[i] == d[1] {
                count++ // This move is already going in our target direction
            } else {
                // This move is going in the opposite direction
                if rem > 0 {
                    // We can convert this opposing move to our target direction
                    rem--
                    count++
                } else {
                    // Can't convert, this move works against our target direction
                    count--
                }
            }
            res = max(res, count)
        }
    }
    return res
}

func maxDistance1(s string, k int) int {
    res, count := 0, ['X']int{}
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range s {
        count[v]++
        left := k
        helper := func(a, b int) int {
            d := min(a, b, left)
            left -= d
            return abs(a - b) + 2 * d
        }
        res = max(res, helper(count['N'], count['S']) + helper(count['E'], count['W']))
    }
    return res
}

func maxDistance2(s string, k int) int {
    res, v, h, north, east := 0, 0, 0, 0, 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, c := range s {
        if c == 'N' {
            if north < 0 {
                v++
            }
            north++
        }
        if c == 'S' {
            if north > 0 {
                v++
            }
            north--
        }
        if c == 'E' {
            if east < 0 {
                h++
            }
            east++
        }
        if c == 'W' {
            if east > 0 {
                h++
            }
            east--
        }
        curr := abs(north) + abs(east)
        t := v + h
        if t > k {
            t = k
        }
        curr += 2 * t
        if curr > res {
            res = curr
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "NWSE", k = 1
    // Output: 3
    // Explanation:
    // Change s[2] from 'S' to 'N'. The string s becomes "NWNE".
    // Movement	Position (x, y)	Manhattan Distance	Maximum
    // s[0] == 'N'	(0, 1)	0 + 1 = 1	1
    // s[1] == 'W'	(-1, 1)	1 + 1 = 2	2
    // s[2] == 'N'	(-1, 2)	1 + 2 = 3	3
    // s[3] == 'E'	(0, 2)	0 + 2 = 2	3
    // The maximum Manhattan distance from the origin that can be achieved is 3. Hence, 3 is the output.
    fmt.Println(maxDistance("NWSE", 1)) // 3
    // Example 2:
    // Input: s = "NSWWEW", k = 3
    // Output: 6
    // Explanation:
    // Change s[1] from 'S' to 'N', and s[4] from 'E' to 'W'. The string s becomes "NNWWWW".
    // The maximum Manhattan distance from the origin that can be achieved is 6. Hence, 6 is the output.
    fmt.Println(maxDistance("NSWWEW", 3)) // 6

    fmt.Println(maxDistance1("NWSE", 1)) // 3
    fmt.Println(maxDistance1("NSWWEW", 3)) // 6

    fmt.Println(maxDistance2("NWSE", 1)) // 3
    fmt.Println(maxDistance2("NSWWEW", 3)) // 6
}