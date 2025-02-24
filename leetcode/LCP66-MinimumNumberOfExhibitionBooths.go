package main

// LCP 66. 最小展台数量
// 力扣嘉年华将举办一系列展览活动，后勤部将负责为每场展览提供所需要的展台。 
// 已知后勤部得到了一份需求清单，记录了近期展览所需要的展台类型，demand[i][j]表示第i天展览时第j个展台的类型。 
// 在满足每一天展台需求的基础上，请返回后勤部需要准备的最小展台数量。

// 注意：
//     同一展台在不同天中可以重复使用。

// 示例 1：
// 输入：demand = ["acd","bed","accd"]
// 输出：6
// 解释： 第0天需要展台a、c、d； 第1天需要展台b、e、d； 第2天需要展台a、c、c、d； 因此，后勤部准备abccde的展台，可以满足每天的展览需求;

// 示例 2：
// 输入：demand = ["abc","ab","ac","b"]
// 输出：3

// 提示：
//     1 <= demand.length,demand[i].length <= 100
//     demand[i][j]仅为小写字母

import "fmt"
import "strings"

func minNumBooths(demand []string) int {
    mp := make(map[string]int)
    for _, i := range demand {
        for _, v := range i {
            count := strings.Count(i, string(v)) // 判断当前展台类型个数
            if count > mp[string(v)] {
                mp[string(v)]++
            }
        }
    }
    res := 0
    for _, v := range mp {
        res += v
    }
    return res
}

func minNumBooths1(demand []string) int {
    mp := [26]int{}
    for _, s := range demand {
        count := [26]int{}
        for _, c := range s {
            count[c-'a']++
        }
        for i := 0; i < 26; i++ {
            if count[i] > mp[i] {
                mp[i] = count[i]
            }
        }
    }
    res := 0
    for _, v := range mp {
        res += v
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：demand = ["acd","bed","accd"]
    // 输出：6
    // 解释： 第0天需要展台a、c、d； 第1天需要展台b、e、d； 第2天需要展台a、c、c、d； 因此，后勤部准备abccde的展台，可以满足每天的展览需求;
    fmt.Println(minNumBooths([]string{"acd","bed","accd"})) // 6
    // 示例 2：
    // 输入：demand = ["abc","ab","ac","b"]
    // 输出：3
    fmt.Println(minNumBooths([]string{"abc","ab","ac","b"})) // 3

    fmt.Println(minNumBooths([]string{"bluefrog","leetcode"})) // 13

    fmt.Println(minNumBooths1([]string{"acd","bed","accd"})) // 6
    fmt.Println(minNumBooths1([]string{"abc","ab","ac","b"})) // 3
    fmt.Println(minNumBooths1([]string{"bluefrog","leetcode"})) // 13
}