package main

// LCR 066.  键值映射
// 实现一个 MapSum 类，支持两个方法，insert 和 sum：
//     MapSum() 
//         初始化 MapSum 对象
//     void insert(String key, int val) 
//         插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。
//         如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
//     int sum(string prefix) 
//         返回所有以该前缀 prefix 开头的键 key 的值的总和。

// 示例：
// 输入：
// inputs = ["MapSum", "insert", "sum", "insert", "sum"]
// inputs = [[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
// 输出：
// [null, null, 3, null, 5]
// 解释：
// MapSum mapSum = new MapSum();
// mapSum.insert("apple", 3);  
// mapSum.sum("ap");           // return 3 (apple = 3)
// mapSum.insert("app", 2);    
// mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)

// 提示：
//     1 <= key.length, prefix.length <= 50
//     key 和 prefix 仅由小写英文字母组成
//     1 <= val <= 1000
//     最多调用 50 次 insert 和 sum

import "fmt"

type MapSum struct {
    dict map[string]int
    set map[string]int
}

func Constructor() MapSum {
    return MapSum{ make(map[string]int), make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int)  {
    diff := val
    if v, ok := this.set[key]; ok {
        diff -= v
    }
    this.set[key] = val
    for i := 1; i <= len(key); i++ {
        this.dict[key[:i]] += diff
    }
}

func (this *MapSum) Sum(prefix string) int {
    return this.dict[prefix]
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

func main() {
    // MapSum mapSum = new MapSum();
    obj := Constructor()
    // mapSum.insert("apple", 3);  
    obj.Insert("apple", 3)
    fmt.Println(obj)
    // mapSum.sum("ap");           // return 3 (apple = 3)
    fmt.Println(obj.Sum("ap")) // 3
    // mapSum.insert("app", 2);    
    obj.Insert("app", 2)
    fmt.Println(obj)
    // mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)
    fmt.Println(obj.Sum("ap")) // 5
}