package main

// 49. Group Anagrams
// Given an array of strings strs, group the anagrams together. 
// You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Constraints:
//     1 <= strs.length <= 10^4
//     0 <= strs[i].length <= 100
//     strs[i] consists of lowercase English letters.

// Example 1:
//     Input: strs = ["eat","tea","tan","ate","nat","bat"]
//     Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// Example 2:
//     Input: strs = [""]
//     Output: [[""]]

// Example 3:
//     Input: strs = ["a"]
//     Output: [["a"]]

// 给出一个字符串数组，要求对字符串数组里面有 Anagrams 关系的字符串进行分组。
// Anagrams 关系是指两个字符串的字符完全相同，顺序不同，两者是由排列组合组成。

// 将每个字符串都排序，排序完成以后，
// 相同 Anagrams 的字符串必然排序结果一样。
// 把排序以后的字符串当做 key 存入到 map 中。
// 遍历数组以后，就能得到一个 map，key 是排序以后的字符串，value 对应的是这个排序字符串以后的 Anagrams 字符串集合。
// 最后再将这些 value 对应的字符串数组输出即可。


import "fmt"
import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func groupAnagrams(strs []string) [][]string {
    var res [][]string
    record := map[string][]string{} // key 是排序以后的字符串，value 对应的是这个排序字符串以后的 Anagrams 字符串集合
    for _, str := range strs {
        sByte := []rune(str)
        sort.Sort(sortRunes(sByte)) // 按字符排序

        sstrs := record[string(sByte)]
        sstrs = append(sstrs, str)
        record[string(sByte)] = sstrs
    }
    for _, v := range record {
        res = append(res, v)
    }
    return res
}

// 这边把字符串生成 hash 值
func groupAnagrams1(strs []string) [][]string {
	m := make(map[int][]string)
	for _, s := range strs {
		h := hash(s)
		m[h] = append(m[h], s)
	}
	groups := make([][]string, 0)
	for _, v := range m {
		groups = append(groups, v)
	}
	return groups
}

// a2, b3, c5, d7, e11, f13, g17, h19, i23, g29, k31, l37, m41, n43, o47, p53, q59, r61, s67, t71, u73, v79, w83, x89, y97, z101
var primeNumbers = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

const bigPrime = 276906403

func hash(s string) int {
    //fmt.Printf("s = %v\n",s)
    result := 1
    for _, c := range s {
        result = (result * primeNumbers[c - 'a']) % bigPrime
        // fmt.Printf("primeNumbers[c - 'a'] = %v\n",primeNumbers[c - 'a'])
        // fmt.Printf("result :%v\n",result)
    }
    return result
}

//  best solution 
func groupAnagrams2(strs []string) [][]string {
    // HashMap 对strs 利的元素在bytes 排序后 就相同了，作为key, values 是个原始切片
    // HashMap 用计数作为key, 将计数转化为 [26]int 数组，代表 strs 中每个元素 有几个26字母
    hmap := make(map[[26]int][]string)

    for _, str := range strs {
        // key 表示str 的字符最多26个，每个字符有多少个
        key := [26]int{}
        // 计数
        for _, ch := range str {
            // ch - 'a'代表字符距离小写字母的距离, b 的话是1
            key[ch - 'a'] ++
        }
        hmap[key] = append(hmap[key], str)

        // fmt.Println("key: ",key)
        // fmt.Println("hmap[key]: ",hmap[key])
    }
    result := make([][]string, 0, len(hmap))
    for _, v := range hmap {
        result = append(result, v)
    }   
    return result
}

func main() {
	fmt.Printf("groupAnagrams([]string{\"eat\",\"tea\",\"tan\",\"ate\",\"nat\",\"bat\"}) = %v\n",groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"})) //  [["bat"],["nat","tan"],["ate","eat","tea"]]
	fmt.Printf("groupAnagrams([]string{\"\"}) = %v\n",groupAnagrams([]string{""})) // [[""]]
	fmt.Printf("groupAnagrams([]string{\"a\"}) = %v\n",groupAnagrams([]string{"a"})) // [["a"]]

	fmt.Printf("groupAnagrams1([]string{\"eat\",\"tea\",\"tan\",\"ate\",\"nat\",\"bat\"}) = %v\n",groupAnagrams1([]string{"eat","tea","tan","ate","nat","bat"})) //  [["bat"],["nat","tan"],["ate","eat","tea"]]
	fmt.Printf("groupAnagrams1([]string{\"\"}) = %v\n",groupAnagrams1([]string{""})) // [[""]]
	fmt.Printf("groupAnagrams1([]string{\"a\"}) = %v\n",groupAnagrams1([]string{"a"})) // [["a"]]

	fmt.Printf("groupAnagrams2([]string{\"eat\",\"tea\",\"tan\",\"ate\",\"nat\",\"bat\"}) = %v\n",groupAnagrams2([]string{"eat","tea","tan","ate","nat","bat"})) //  [["bat"],["nat","tan"],["ate","eat","tea"]]
	fmt.Printf("groupAnagrams2([]string{\"\"}) = %v\n",groupAnagrams2([]string{""})) // [[""]]
	fmt.Printf("groupAnagrams2([]string{\"a\"}) = %v\n",groupAnagrams2([]string{"a"})) // [["a"]]
}
