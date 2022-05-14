package main

import "fmt"

/**
212. Word Search II

Given an m x n board of characters and a list of strings words, return all words on the board.
Each word must be constructed from letters of sequentially adjacent cells,
where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once in a word.

Example 1:

	Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
	Output: ["eat","oath"]

Example 2:

	Input: board = [["a","b"],["c","d"]], words = ["abcb"]
	Output: []


Constraints:

	m == board.length
	n == board[i].length
	1 <= m, n <= 12
	board[i][j] is a lowercase English letter.
	1 <= words.length <= 3 * 10^4
	1 <= words[i].length <= 10
	words[i] consists of lowercase English letters.
	All the strings of words are unique.

DFS
在第 79 题的基础上增加了一个 word 数组，要求找出所有出现在地图中的单词
*/

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	for _, v := range words {
		if exist(board, v) {
			res = append(res, v)
		}
	}
	return res
}

// these is 79 solution
var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i, v := range board {
		for j := range v {
			if searchWord(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func searchWord(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
	if index == len(word)-1 {
		return board[x][y] == word[index]
	}
	if board[x][y] == word[index] {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if isInBoard(board, nx, ny) && !visited[nx][ny] && searchWord(board, visited, word, index+1, nx, ny) {
				return true
			}
		}
		visited[x][y] = false
	}
	return false
}

// Best solution
func findWordsBest(board [][]byte, words []string) []string {
	m := len(board)
	n := len(board[0])
	/*
			注意到m和n不超过12。每个word的长度不超过10
		 	直观解法（应该会超时）：
				用一个map保存words中的所有字母
		 		直接在board中搜索出长度10以内的所有解，时间复杂度 O(m * n * 4^10) ，如果某个解也在map中，追加到结果数组
		 	优化：
			 	改用前缀树来保存words中的所有字符，这样在borad中搜索时，就可以结合前缀树进行剪枝：即无解的提前返回。
	*/
	t := &Trie{}
	for _, word := range words {
		t.Add(word)
	}

	ret := make([]string, 0)
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	var dfs func(x int, y int, node *Trie)
	dfs = func(x int, y int, node *Trie) {
		if node == nil || node.count == 0 { //无解，剪枝
			return
		}

		// 防止board[x][y]被重复使用
		if used[x][y] == true {
			return
		}
		used[x][y] = true

		if node.isEnd == true { //有阶段解
			ret = append(ret, node.word)
			t.Del(node.word) //剪枝
		}

		for _, v := range dirs {
			i, j := x+v[0], y+v[1]
			if i >= 0 && i < m && j >= 0 && j < n {
				ch := board[i][j] - 'a'
				dfs(i, j, node.childs[ch])
			}
		}
		// 现场恢复
		used[x][y] = false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, t.childs[board[i][j]-'a'])
		}
	}
	return ret
}

type Trie struct {
	// 基本信息
	childs [26]*Trie
	isEnd  bool
	// 额外信息，用于加速
	count  int    //记录本节点被覆盖了多少次，用于剪枝
	word   string //保存整个路径，便于直接扔到本题的结果数组中。否则就得在dfs过程中额外维护路径信息
}

func (t *Trie) Add(word string) {
	node := t
	for _, c := range word {
		ch := c - 'a'
		if node.childs[ch] == nil {
			node.childs[ch] = &Trie{}
		}
		node = node.childs[ch]
		node.count++
	}
	node.isEnd = true
	node.word = word
}

func (t *Trie) Del(word string) {
	node := t
	for _, c := range word {
		ch := c - 'a'
		node = node.childs[ch]
		node.count--
	}
	node.isEnd = false
}

func main() {
	bytes := [][]byte{{'o','a','a','n'},{'e','t','a','e'},{'i','h','k','r'},{'i','f','l','v'}}
	fmt.Printf("findWords(bytes,\"ABCCED\") = %v\n",findWords(bytes,[]string{ "oath","pea","eat","rain" })) // ["eat","oath"]
	fmt.Printf("findWords([][]byte{ {'a','b'},{'c','d'}},[]string{ \"abcb\" } = %v\n",findWords([][]byte{ {'a','b'},{'c','d'}},[]string{ "abcb" })) // []

	fmt.Printf("findWordsBest(bytes,\"ABCCED\") = %v\n",findWordsBest(bytes,[]string{ "oath","pea","eat","rain" })) // ["eat","oath"]
	fmt.Printf("findWordsBest([][]byte{ {'a','b'},{'c','d'}},[]string{ \"abcb\" } = %v\n",findWordsBest([][]byte{ {'a','b'},{'c','d'}},[]string{ "abcb" })) // []

}