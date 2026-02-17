package main

// 3846. Total Distance to Type a String Using One Finger
// There is a special keyboard where keys are arranged in a rectangular grid as follows.
// q	w	e	r	t	y	u	i	o	p
// a	s	d	f	g	h	j	k	l	 
// z	x	c	v	b	n	m	 	 	 

// You are given a string s that consists of lowercase English letters only. 
// Return an integer denoting the total distance to type s using only one finger. Your finger starts on the key 'a'.

// The distance between two keys at (r1, c1) and (r2, c2) is |r1 - r2| + |c1 - c2|.

// Example 1:
// Input: s = "hello"
// Output: 17
// Explanation:
// Your finger starts at 'a', which is at (1, 0).
// Move to 'h', which is at (1, 5). The distance is |1 - 1| + |0 - 5| = 5.
// Move to 'e', which is at (0, 2). The distance is |1 - 0| + |5 - 2| = 4.
// Move to 'l', which is at (1, 8). The distance is |0 - 1| + |2 - 8| = 7.
// Move to 'l', which is at (1, 8). The distance is |1 - 1| + |8 - 8| = 0.
// Move to 'o', which is at (0, 8). The distance is |1 - 0| + |8 - 8| = 1.
// Total distance is 5 + 4 + 7 + 0 + 1 = 17.

// Example 2:
// Input: s = "a"
// Output: 0
// Explanation:
// Your finger starts at 'a', which is at (1, 0).
// Move to 'a', which is at (1, 0). The distance is |1 - 1| + |0 - 0| = 0.
// Total distance is 0.
 
// Constraints:
//     1 <= s.length <= 10^4
//     s consists of lowercase English letters only.

import "fmt"

func totalDistance(s string) int {
    // 定义键盘布局，建立每个字符到坐标 (行, 列) 的映射
    keyboard := map[rune][2]int{
        'q': {0, 0}, 'w': {0, 1}, 'e': {0, 2}, 'r': {0, 3}, 't': {0, 4}, 'y': {0, 5}, 'u': {0, 6}, 'i': {0, 7}, 'o': {0, 8}, 'p': {0, 9},
        'a': {1, 0}, 's': {1, 1}, 'd': {1, 2}, 'f': {1, 3}, 'g': {1, 4}, 'h': {1, 5}, 'j': {1, 6}, 'k': {1, 7}, 'l': {1, 8},
        'z': {2, 0}, 'x': {2, 1}, 'c': {2, 2}, 'v': {2, 3}, 'b': {2, 4}, 'n': {2, 5}, 'm': {2, 6},
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    res, row, col := 0, keyboard['a'][0], keyboard['a'][1] // // 初始位置在 'a'
    for _, c := range s { // 遍历字符串中的每个字符
        row1, col1 := keyboard[c][0], keyboard[c][1] // 获取目标字符的坐标
        distance := abs(row - row1) + abs(col - col1) // 计算曼哈顿距离并累加
        res += distance
        row, col = row1, col1 // 更新当前位置为目标字符位置
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "hello"
    // Output: 17
    // Explanation:
    // Your finger starts at 'a', which is at (1, 0).
    // Move to 'h', which is at (1, 5). The distance is |1 - 1| + |0 - 5| = 5.
    // Move to 'e', which is at (0, 2). The distance is |1 - 0| + |5 - 2| = 4.
    // Move to 'l', which is at (1, 8). The distance is |0 - 1| + |2 - 8| = 7.
    // Move to 'l', which is at (1, 8). The distance is |1 - 1| + |8 - 8| = 0.
    // Move to 'o', which is at (0, 8). The distance is |1 - 0| + |8 - 8| = 1.
    // Total distance is 5 + 4 + 7 + 0 + 1 = 17.
    fmt.Println(totalDistance("hello")) // 17
    // Example 2:
    // Input: s = "a"
    // Output: 0
    // Explanation:
    // Your finger starts at 'a', which is at (1, 0).
    // Move to 'a', which is at (1, 0). The distance is |1 - 1| + |0 - 0| = 0.
    // Total distance is 0. 
    fmt.Println(totalDistance("a")) // 0

    fmt.Println(totalDistance("bluefrog")) // 30
    fmt.Println(totalDistance("leetcode")) // 37
    fmt.Println(totalDistance("freewu")) // 11
}
