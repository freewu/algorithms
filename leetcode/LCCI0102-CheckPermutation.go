package main

// 面试题 01.02. Check Permutation LCCI
// Given two strings,write a method to decide if one is a permutation of the other.

// Example 1:
// Input: s1 = "abc", s2 = "bca"
// Output: true

// Example 2:
// Input: s1 = "abc", s2 = "bad"
// Output: false

// Note:
//         0 <= len(s1) <= 100 
//         0 <= len(s2) <= 100

import "fmt"

func CheckPermutation(s1 string, s2 string) bool {
    if s1 == s2  {
        return true
    }
    if len(s1) != len(s2)  {
        return false
    }
    // 有问题 遇到 (aa,bb) 这种 就判断不正确 
    // 计算异或值 
    // b1, b2 := byte(0), byte(0)
    // for i := 0; i < len(s1); i += 1 {
    //     b1 ^= s1[i]
    //     b2 ^= s2[i]
    // }
    // return b1 == b2

    // 使用 map
    m1 := make(map[byte]int)
    m2 := make(map[byte]int)
    for i := 0; i < len(s1); i += 1 {
        m1[s1[i]] += 1
        m2[s2[i]] += 1
    }
    // 比对判断
    for k,v := range(m1) {
        if m2[k] != v {
            return false
        }
    }
    return true
}

// 只需要一个 map 
func CheckPermutation1(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    hash := make([]int, 26)
    for i, c := range s1 {
        hash[c - 'a']++; // s1 有就 + 
        hash[s2[i] - 'a']--; // s2 有就 -
    }
    for _, v := range hash {
        // 如果字符出现不同肯定不为 0
        if v != 0 {
            return false
        }
    }
    return true
}


func main() {
    fmt.Println(CheckPermutation("abc","bca")) // true
    fmt.Println(CheckPermutation("abc","bad")) // false
    fmt.Println(CheckPermutation("aa","bb")) // false

    fmt.Println(CheckPermutation1("abc","bca")) // true
    fmt.Println(CheckPermutation1("abc","bad")) // false
    fmt.Println(CheckPermutation1("aa","bb")) // false
}