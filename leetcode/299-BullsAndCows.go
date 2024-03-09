package main 

// 299. Bulls and Cows
// You are playing the Bulls and Cows game with your friend.
// You write down a secret number and ask your friend to guess what the number is. 
// When your friend makes a guess, you provide a hint with the following info:
//         The number of "bulls", which are digits in the guess that are in the correct position.
//         The number of "cows", which are digits in the guess that are in your secret number but are located in the wrong position. Specifically, the non-bull digits in the guess that could be rearranged such that they become bulls.

// Given the secret number secret and your friend's guess guess, return the hint for your friend's guess.
// The hint should be formatted as "xAyB", where x is the number of bulls and y is the number of cows. Note that both secret and guess may contain duplicate digits.

// Example 1:
// Input: secret = "1807", guess = "7810"
// Output: "1A3B"
// Explanation: Bulls are connected with a '|' and cows are underlined:
// "1807"
//   |
// "7810"

// Example 2:
// Input: secret = "1123", guess = "0111"
// Output: "1A1B"
// Explanation: Bulls are connected with a '|' and cows are underlined:
// "1123"        "1123"
//   |      or     |
// "0111"        "0111"
// Note that only one of the two unmatched 1s is counted as a cow since the non-bull digits can only be rearranged to allow one 1 to be a bull.
 
// Constraints:
//         1 <= secret.length, guess.length <= 1000
//         secret.length == guess.length
//         secret and guess consist of digits only.

import "fmt"

// 计算下标一致并且对应下标的元素一致的个数，即 x
// secret 和 guess 分别去除 x 个公牛的元素,剩下 secret 和 guess 求共同的元素个数就是 y
// 把 x， y 转换成字符串，分别与 A 和 B 进行拼接返回结果
func getHint(secret string, guess string) string {
	x, y := 0, 0
	m := make(map[byte]int)
	var sa []byte
	n := len(secret)
	for i := 0; i < n; i++ {
        // 猜对的个数
		if secret[i] == guess[i] {
			x++
		} else {
			m[secret[i]] += 1
			sa = append(sa, guess[i])
		}
	}
	for _, v := range sa {
		if _, ok := m[v]; ok {
            // 去除 x 个公牛的元素,剩下 secret 和 guess 求共同的元素个数就是 y
			if m[v] > 1 {
				m[v] -= 1
			} else {
				delete(m, v)
			}
			y++
		}
	}
    return fmt.Sprintf("%dA%dB",x, y)
}

func getHint1(secret string, guess string) string {
	x, y := 0, 0
	ca_g := make([]int, 10, 10)
	ca_s := make([]int, 10, 10)
	length := len(secret)
	for i := 0; i < length; i++ {
		if guess[i] == secret[i] {
			x++
		} else {
			ca_g[guess[i]-'0']++
			ca_s[secret[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		if ca_g[i] >= ca_s[i] {
			y += ca_s[i]
		} else {
			y += ca_g[i]
		}
	}
	return fmt.Sprintf("%dA%dB", x, y)
}

func main() {
    fmt.Println(getHint("1807","7810")) // 1A3B
    fmt.Println(getHint("1123","0111")) // 1A1B

    fmt.Println(getHint1("1807","7810")) // 1A3B
    fmt.Println(getHint1("1123","0111")) // 1A1B
}