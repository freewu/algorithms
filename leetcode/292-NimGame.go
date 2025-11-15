package main

// 292. Nim Game
// You are playing the following Nim Game with your friend:
//     Initially, there is a heap of stones on the table.
//     You and your friend will alternate taking turns, and you go first.
//     On each turn, the person whose turn it is will remove 1 to 3 stones from the heap.
//     The one who removes the last stone is the winner.

// Given n, the number of stones in the heap, 
// return true if you can win the game assuming both you and your friend play optimally, otherwise return false.

// Example 1:
// Input: n = 4
// Output: false
// Explanation: These are the possible outcomes:
// 1. You remove 1 stone. Your friend removes 3 stones, including the last stone. Your friend wins.
// 2. You remove 2 stones. Your friend removes 2 stones, including the last stone. Your friend wins.
// 3. You remove 3 stones. Your friend removes the last stone. Your friend wins.
// In all outcomes, your friend wins.

// Example 2:
// Input: n = 1
// Output: true

// Example 3:
// Input: n = 2
// Output: true
 
// Constraints:
//     1 <= n <= 2^31 - 1

import "fmt"

// 如果Nim游戏中的规则稍微变动一下,每次最多只能取K个,怎么处理?
// 方法是将每堆石子数mod (k+1)
//     1、【K+1】-【K+1】-【K+1】-【K+1】-【K+1】-【n】，非平衡态
//     2、【K+1】-【K+1】-【K+1】-【K+1】-【K+1】，平衡态（每一堆个数亦或为0）
// 游戏人I能够在非平衡取子游戏中取胜，而游戏人II能够在平衡的取子游戏中取胜
func canWinNim(n int) bool {
    // 1-3 都可以一次性取完 所以 开始皆胜利
    if(n < 4) {
        return true
    }
    if n % 4 != 0 { // 非平衡态
        return true
    }
    // 后者保证每轮取 k + 1 (4) 就可以了
    // A 1  B 3
    // A 2  B 2
    // A 3  B 1 
    return false // 平衡态 后者必败 
}

func main() {
    // Example 1:
    // Input: n = 4
    // Output: false
    // Explanation: These are the possible outcomes:
    // 1. You remove 1 stone. Your friend removes 3 stones, including the last stone. Your friend wins.
    // 2. You remove 2 stones. Your friend removes 2 stones, including the last stone. Your friend wins.
    // 3. You remove 3 stones. Your friend removes the last stone. Your friend wins.
    // In all outcomes, your friend wins.
    fmt.Println(canWinNim(4)) // false
    // Example 2:
    // Input: n = 1
    // Output: true
    fmt.Println(canWinNim(1)) // true
    // Example 3:
    // Input: n = 2
    // Output: true
    fmt.Println(canWinNim(2)) // true
	
	fmt.Println(canWinNim(8)) // false
	fmt.Println(canWinNim(9)) // true
    fmt.Println(canWinNim(64)) // false
    fmt.Println(canWinNim(99)) // true
    fmt.Println(canWinNim(100)) // false
	fmt.Println(canWinNim(999)) // true
	fmt.Println(canWinNim(1024)) // false
    fmt.Println(canWinNim(1_000_000_007)) // true
    fmt.Println(canWinNim(1 << 31 - 1)) // true
}