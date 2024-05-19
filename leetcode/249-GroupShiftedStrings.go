package main

// 249. Group Shifted Strings
// We can shift a string by shifting each of its letters to its successive letter.
//     For example, "abc" can be shifted to be "bcd".

// We can keep shifting the string to form a sequence.
//     For example, we can keep shifting "abc" to form the sequence: "abc" -> "bcd" -> ... -> "xyz".

// Given an array of strings strings, group all strings[i] that belong to the same shifting sequence. 
// You may return the answer in any order.

// Example 1:
// Input: strings = ["abc","bcd","acef","xyz","az","ba","a","z"]
// Output: [["acef"],["a","z"],["abc","bcd","xyz"],["az","ba"]]

// Example 2:
// Input: strings = ["a"]
// Output: [["a"]]

// Constraints:
//     1 <= strings.length <= 200
//     1 <= strings[i].length <= 50
//     strings[i] consists of lowercase English letters.

import "fmt"
import "sort"

// 暴力
func groupStrings(strings []string) [][]string {
    sort.Strings(strings)
    res := [][]string{}
    for i := 0; i < len(strings); i++ {
        str, find := strings[i], false
        for j := 0; j < len(res); j++ {
            if len(res[j][0]) != len(str) {
                continue
            }
            add, same := int(str[0] - res[j][0][0]), 1
            for k := 1; k < len(str); k++ {
                if int((str[k] - res[j][0][k]) + 26) % 26 == add { // 计算差值是否相等
                    same++
                } else {
                    break
                }
            }
            if same == len(str) { // 需要所有差值都相等
                res[j] = append(res[j], str)
                find = true
            }
        }
        if !find {
            res = append(res, []string{str}) // 没有同组的，新增一个组
        }
    }
    return res
}

// 哈希表记录分组
func groupStrings1(strings []string) [][]string {
    res := make(map[string][]string)
    change := func (inp string) (ret string){
        offset, tmpByte := inp[0] - 'a', []byte(inp)
        for i := 0; i <  len(tmpByte); i++ {
            tmpByte[i] = tmpByte[i] - offset
            if tmpByte[i] < 'a' {
                tmpByte[i] = tmpByte[i] + 26
            }
            if inp[i] > 'z' {
                tmpByte[i] = tmpByte[i] - 26
            }
        }
        return string(tmpByte) 
    }
    for _, cur := range(strings) {
        changeStr := change(cur)
        if val, ok := res[changeStr]; ok {
            res[changeStr] = append(val, cur)
        } else {
            res[changeStr] = make([]string, 0)
            res[changeStr] = append(res[changeStr], cur)
        }
    }
    result := make([][]string, 0)
    for _, v := range(res){
        result = append(result, v)
    }
    return result
}

func main() {
    // Example 1:
    // Input: strings = ["abc","bcd","acef","xyz","az","ba","a","z"]
    // Output: [["acef"],["a","z"],["abc","bcd","xyz"],["az","ba"]]
    fmt.Println(groupStrings([]string{"abc","bcd","acef","xyz","az","ba","a","z"})) // [["acef"],["a","z"],["abc","bcd","xyz"],["az","ba"]]
    // Example 2:
    // Input: strings = ["a"]
    // Output: [["a"]]
    fmt.Println(groupStrings([]string{"a"})) // [["a"]]

    fmt.Println(groupStrings1([]string{"abc","bcd","acef","xyz","az","ba","a","z"})) // [["acef"],["a","z"],["abc","bcd","xyz"],["az","ba"]]
    fmt.Println(groupStrings1([]string{"a"})) // [["a"]]
}