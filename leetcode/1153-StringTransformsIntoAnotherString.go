package main

// 1153. String Transforms Into Another String
// Given two strings str1 and str2 of the same length, determine whether you can transform str1 into str2 by doing zero or more conversions.
// In one conversion you can convert all occurrences of one character in str1 to any other lowercase English character.
// Return true if and only if you can transform str1 into str2.

// Example 1:
// Input: str1 = "aabcc", str2 = "ccdee"
// Output: true
// Explanation: Convert 'c' to 'e' then 'b' to 'd' then 'a' to 'c'. Note that the order of conversions matter.

// Example 2:
// Input: str1 = "leetcode", str2 = "codeleet"
// Output: false
// Explanation: There is no way to transform str1 to str2.

// Constraints:
//     1 <= str1.length == str2.length <= 10^4
//     str1 and str2 contain only lowercase English letters.

import "fmt"

// 参考模式匹配，只是这个是没有“.”也没有“*”的模式匹配
// 不是所有连续的数字，是所有相同的数字
// 为什么说转化的顺序也很重要？？懂了，假如说abc可以转成aaa；自己试着转一下leetcode -> codeleet，o转成e，d转成e；或者e先转成t；都不行。aabcc - > ccdee，如果先把a变成c就搞不定
// 就是连续的e永远都连续leet中间的ee永远也变不成code的od；但是od可以变成ee，先转o，再转e
// 就是熵增，字母在变化过程中趋同，连续的字母只会维持或者变多，而不会变少；那就是str2中连续字母的位置必须覆盖str1中连续字母的位置；仅仅这样够了么？不够，下面这个例子就是非连续，只要是相同字母的非连续也必须覆盖。
// str1:cbbc   str2: abbc
// 只需要返回是或者否，递推?dp?
// **如果str1和str2的长度不超过32位，可以用位运算，a & b + a ^ b = a证明在b出现1的地方a都出现1。这个方法可以快速匹配看str2中某个字母s出现的地点是否覆盖str1中某个相同字母t出现的地点（a表示str2中的s出现的地点，b表示str1中t出现的地点）
// 否则，就是检查t每个出现的地方，是否都含在str2中的某个字母t的位置集中，每个地方都是同一个位置集
// 可以先构造str2的位置集，然后遍历str1，看是否str1的位置集能被str2覆盖
// str1: "abcdefghijklmnopqrstuvwxyz"
// str2: "bcdefghijklmnopqrstuvwxyza" 这个错了，这个情况下，str1变一次，那么就会有超过str2位置集的字母出现了，如何保证不会有这种情况出现，如果26个字母都有的情况下，str1必须和str2完全匹配？简化法假设只能有a和b两个字母互相变，那么str1是ab，str2是ba，是变不过去的；或者str1是abab，str2是baba。如果位置和数量不完全一样的话，那么str1的变化必然导致熵增，也不存在str2中某些位置集比str1大的情况，假设str2是aaab，str1是aabb，str2中a的位置集确实覆盖了str1中a的位置集，但也导致str1中b的位置集大于str2中b的位置集，因为我们限定了只能是a和b，相同长度的前提下，a位置多了，b位置就少了，26个字母是相同的道理。因此断定，在str2用满字符集的前提下，str1和str2必须完全一样。
func canConvert(str1 string, str2 string) bool {
    findIPosSet := func(i int, m2 map[byte]map[int]bool) map[int]bool {
        for _, m := range m2 { // for key, val := range myMap
            if _, e := m[i]; e {
                return m
            }
        }
        mp := make(map[int]bool, 0)
        return mp
    }
    buildMap := func (s string) map[byte]map[int]bool {
        mp := make(map[byte]map[int]bool)
        for i, r := range s {
            c := byte(r)
            tm, e := mp[c]
            if !e {
                tm = make(map[int]bool)
                mp[c] = tm
            }
            tm[i] = true
        }
        return mp
    }
    map2 := buildMap(str2)
    if len(map2) == 26 {
        return str1 == str2
    }
    map1 := make(map[byte]map[int]bool)
    for i, r := range str1 {
        c := byte(r)
        curM, e := map1[c]
        if !e { // 之前没出现过相同的字母
            map1[c] = findIPosSet(i, map2) // 找到涵盖i位置的位置集
        } else {
            _, e1 := curM[i]
            if !e1 {
                return false
            } 
        }
    }
    return true
}

func canConvert1(str1 string, str2 string) bool {
    if str1 == str2 {
        return true
    }
    mp1, mp2 := make(map[byte]byte), make(map[byte]bool)
    for i := range str1 {
        if curr, ok := mp1[str1[i]]; ok && curr != str2[i] {
            return false
        }
        mp1[str1[i]] = str2[i]
        mp2[str2[i]] = true
    }
    if len(mp1) == 0 {
        return true
    }
    return len(mp2) < 26
}

func main() {
    // Example 1:
    // Input: str1 = "aabcc", str2 = "ccdee"
    // Output: true
    // Explanation: Convert 'c' to 'e' then 'b' to 'd' then 'a' to 'c'. Note that the order of conversions matter.
    fmt.Println(canConvert("aabcc", "ccdee")) // true
    // Example 2:
    // Input: str1 = "leetcode", str2 = "codeleet"
    // Output: false
    // Explanation: There is no way to transform str1 to str2.
    fmt.Println(canConvert("leetcode", "codeleet")) // false

    fmt.Println(canConvert("bluefrog", "leetcode")) // true
    fmt.Println(canConvert("bluefrog", "frogblue")) // true

    fmt.Println(canConvert1("aabcc", "ccdee")) // true
    fmt.Println(canConvert1("leetcode", "codeleet")) // false
    fmt.Println(canConvert1("bluefrog", "leetcode")) // true
    fmt.Println(canConvert1("bluefrog", "frogblue")) // true
}