package main

import "fmt"
import "time"

// 131. Palindrome Partitioning
// Given a string s, partition s such that every substring of the partition is a palindrome.
// Return all possible palindrome partitioning of s.

// Example 1:
// Input: s = "aab"
// Output: [["a","a","b"],["aa","b"]]

// Example 2:
// Input: s = "a"
// Output: [["a"]]

// Constraints:
// 		1 <= s.length <= 16
// 		s contains only lowercase English letters.

// DFS 递归求解
func partition(s string) [][]string {
	result := [][]string{}
	size := len(s)
	if size == 0 {
		return result
	}
	current := make([]string, 0, size)
	dfs(s, 0, current, &result)
	return result
}

func dfs(s string, idx int, cur []string, result *[][]string) {
	start, end := idx, len(s)
	if start == end {
		temp := make([]string, len(cur))
		copy(temp, cur)
		*result = append(*result, temp)
		return
	}
	for i := start; i < end; i++ {
		// 只处理回文的情况
		if isPal(s, start, i) {
			dfs(s, i+1, append(cur, s[start:i+1]), result)
		}
	}
}

// 判断是否是回文
func isPal(str string, s, e int) bool {
	for s < e {
		if str[s] != str[e] {
			return false
		}
		s++
		e--
	}
	return true
}

// best solution
func partition1(s string) [][]string {

    check := func(i, j int) bool {
        for i < j {
            if s[i] != s[j] {
                return false
            }
            i++
            j--
        }

        return true
    }

    var path []string
    var res [][]string 
    var dfs func(start int) 
    dfs = func(start int) {
        if start >= len(s) {
            tmp := make([]string, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        }

        for i := start; i < len(s); i++ {
            if check(start, i) {
                path = append(path, s[start:i+1])
                dfs(i + 1)
                path = path[:len(path) - 1]
            }
        }
    }

    dfs(0)
    return res
}

func main() {
	start := time.Now() // 获取当前时间
	fmt.Println(partition("aab")) // [[a a b] [aa b]]
	fmt.Println(partition("a")) // [[a]]
	fmt.Printf("ladderLength use : %v \r\n",time.Since(start))

	start = time.Now() // 获取当前时间
	fmt.Println(partition1("aab")) // [[a a b] [aa b]]
	fmt.Println(partition1("a")) // [[a]]
	fmt.Printf("ladderLength use : %v \r\n",time.Since(start))
}