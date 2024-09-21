package main

// 997. Find the Town Judge
// In a town, there are n people labeled from 1 to n. 
// There is a rumor that one of these people is secretly the town judge.
// If the town judge exists, then:
//     The town judge trusts nobody.
//     Everybody (except for the town judge) trusts the town judge.
//     There is exactly one person that satisfies properties 1 and 2.

// You are given an array trust where trust[i] = [ai, bi] representing that the person labeled ai trusts the person labeled bi. 
// If a trust relationship does not exist in trust array, then such a trust relationship does not exist.
// Return the label of the town judge if the town judge exists and can be identified, or return -1 otherwise.

// Example 1:
// Input: n = 2, trust = [[1,2]]
// Output: 2

// Example 2:
// Input: n = 3, trust = [[1,3],[2,3]]
// Output: 3

// Example 3:
// Input: n = 3, trust = [[1,3],[2,3],[3,1]]
// Output: -1

// Constraints:
//     1 <= n <= 1000
//     0 <= trust.length <= 10^4
//     trust[i].length == 2
//     All the pairs of trust are unique.
//     ai != bi
//     1 <= ai, bi <= n

// 小镇里有 n 个人，按从 1 到 n 的顺序编号。传言称，这些人中有一个暗地里是小镇法官。
// 如果小镇法官真的存在，那么：
//      小镇法官不会信任任何人。
//      每个人（除了小镇法官）都信任这位小镇法官。
//      只有一个人同时满足属性 1 和属性 2 。

// 给你一个数组 trust ，其中 trust[i] = [ai, bi] 表示编号为 ai 的人信任编号为 bi 的人。
// 如果小镇法官存在并且可以确定他的身份，请返回该法官的编号；否则，返回 -1 
import "fmt"

func findJudge(n int, trust [][]int) int {
    // 只一一个人且 小镇法官不会信任任何人
    if n == 1 && len(trust) == 0 {
        return 1
    }
    judges := make(map[int]int)
    //  [ai, bi] 表示编号为 ai 的人信任编号为 bi 的人
    for _, v := range trust {
        judges[v[1]] += 1
    }
    for _, v := range trust {
        // 小镇法官不会信任任何人,剔除掉有相信过人的人
        if _, ok := judges[v[0]]; ok {
            delete(judges, v[0])
        }
    }
    for k, v := range judges {
        // 每个人（除了小镇法官）都信任这位小镇法官。
        if v == n-1 {
            return k
        }
    }
    return -1
}

func findJudge1(n int, trust [][]int) int {
    l, r := make([]int, n+1), make([]int, n+1)
    for _, t := range trust {
        l[t[0]]++ // 投票人 + 1
        r[t[1]]++ // 信任人 + 1
    }
    for i := 1; i <= n; i++ {
        // 小镇法官不会信任任何人 l[i] == 0
        // 每个人（除了小镇法官）都信任这位小镇法官 r[i] == n - 1(除了小镇法官)
        if l[i] == 0 && r[i] == n - 1 {
            return i
        }
    }
    return -1
}

func findJudge2(n int, trust [][]int) int {
    indegree := make([]int, n + 1)
    for _,t := range trust {
        indegree[t[1]]++
        indegree[t[0]] = -1
    }
    count, j := 0, -1
    for i := 1; i< len(indegree); i++ {
        if indegree[i] == n-1 {
            count++
            j = i
        }
    }
    if count == 1 {
        return j
    }
    return -1
}

func main() {
    fmt.Println(findJudge(2,[][]int{[]int{1,2}})) // 2
    fmt.Println(findJudge(3,[][]int{[]int{1,3},[]int{2,3}})) // 3
    fmt.Println(findJudge(3,[][]int{[]int{1,3},[]int{2,3},[]int{3,1}})) // -1

    fmt.Println(findJudge1(2,[][]int{[]int{1,2}})) // 2
    fmt.Println(findJudge1(3,[][]int{[]int{1,3},[]int{2,3}})) // 3
    fmt.Println(findJudge1(3,[][]int{[]int{1,3},[]int{2,3},[]int{3,1}})) // -1

    fmt.Println(findJudge2(2,[][]int{[]int{1,2}})) // 2
    fmt.Println(findJudge2(3,[][]int{[]int{1,3},[]int{2,3}})) // 3
    fmt.Println(findJudge2(3,[][]int{[]int{1,3},[]int{2,3},[]int{3,1}})) // -1
}