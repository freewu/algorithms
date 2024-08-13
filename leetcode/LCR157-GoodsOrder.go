package main

// LCR 157. 套餐内商品的排列顺序
// 某店铺将用于组成套餐的商品记作字符串 goods，其中 goods[i] 表示对应商品。请返回该套餐内所含商品的 全部排列方式 。
// 返回结果 无顺序要求，但不能含有重复的元素。

// 示例 1:
// 输入：goods = "agew"
// 输出：["aegw","aewg","agew","agwe","aweg","awge","eagw","eawg","egaw","egwa","ewag","ewga","gaew","gawe","geaw","gewa","gwae","gwea","waeg","wage","weag","wega","wgae","wgea"]

// 提示：
//     1 <= goods.length <= 8

import "fmt"
import "sort"

func goodsOrder(goods string) []string {
    goodsList := []byte(goods)
    res := make([]string, 0)
    sort.Slice(goodsList, func(i, j int) bool {
        return goodsList[i] < goodsList[j]
    })
    reverse := func(goods []byte) {
        n := len(goods)
        for i := 0; i < n/2; i++ {
            goods[i], goods[n-i-1] = goods[n-i-1], goods[i]
        }
    }
    nextOrder := func(goods []byte) bool {
        n := len(goods)
        i := n-2
        for i >= 0 && goods[i] >= goods[i+1] { i-- } 
        if i < 0 {  return false }
        j := n-1
        for j >= 0 && goods[i] >= goods[j] { j-- }
        goods[i], goods[j] = goods[j], goods[i]    
        reverse(goods[i+1:]) 
        return true
    }
    for {
        res = append(res, string(goodsList))
        if nextOrder(goodsList) == false {
            return res
        }
    }
    return res
}

func goodsOrder1(s string) []string {
    res := []string{}
    var dfs func(s string, path string)
    dfs = func(s string, path string) {
        if len(s) == 0 {
            res = append(res, path)
        }
        encounterd := make(map[byte]struct{})
        for i := range s {
            if _, ok := encounterd[s[i]]; ok { continue }
            encounterd[s[i]] = struct{}{}
            dfs(s[:i]+s[i+1:], path+string(s[i]))
        }
    }
    dfs(s, "")
    return res 
}

func main() {
    // 示例 1:
    // 输入：goods = "agew"
    // 输出：["aegw","aewg","agew","agwe","aweg","awge","eagw","eawg","egaw","egwa","ewag","ewga","gaew","gawe","geaw","gewa","gwae","gwea","waeg","wage","weag","wega","wgae","wgea"]
    fmt.Println(goodsOrder("agew")) // ["aegw","aewg","agew","agwe","aweg","awge","eagw","eawg","egaw","egwa","ewag","ewga","gaew","gawe","geaw","gewa","gwae","gwea","waeg","wage","weag","wega","wgae","wgea"]
    
    fmt.Println(goodsOrder1("agew")) // ["aegw","aewg","agew","agwe","aweg","awge","eagw","eawg","egaw","egwa","ewag","ewga","gaew","gawe","geaw","gewa","gwae","gwea","waeg","wage","weag","wega","wgae","wgea"]
}