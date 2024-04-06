package main 

// 1600. Throne Inheritance
// A kingdom consists of a king, his children, his grandchildren, and so on. 
// Every once in a while, someone in the family dies or a child is born.

// The kingdom has a well-defined order of inheritance that consists of the king as the first member. 
// Let's define the recursive function Successor(x, curOrder), which given a person x and the inheritance order so far, 
// returns who should be the next person after x in the order of inheritance.

// Successor(x, curOrder):
//     if x has no children or all of x's children are in curOrder:
//         if x is the king return null
//         else return Successor(x's parent, curOrder)
//     else return x's oldest child who's not in curOrder

// For example, assume we have a kingdom that consists of the king, his children Alice and Bob (Alice is older than Bob), and finally Alice's son Jack.
//     In the beginning, curOrder will be ["king"].
//     Calling Successor(king, curOrder) will return Alice, so we append to curOrder to get ["king", "Alice"].
//     Calling Successor(Alice, curOrder) will return Jack, so we append to curOrder to get ["king", "Alice", "Jack"].
//     Calling Successor(Jack, curOrder) will return Bob, so we append to curOrder to get ["king", "Alice", "Jack", "Bob"].
//     Calling Successor(Bob, curOrder) will return null. Thus the order of inheritance will be ["king", "Alice", "Jack", "Bob"].

// Using the above function, we can always obtain a unique order of inheritance.
// Implement the ThroneInheritance class:
//     ThroneInheritance(string kingName) 
//         Initializes an object of the ThroneInheritance class. 
//         The name of the king is given as part of the constructor.
//     void birth(string parentName, string childName) 
//         Indicates that parentName gave birth to childName.
//     void death(string name) 
//         Indicates the death of name. 
//         The death of the person doesn't affect the Successor function nor the current inheritance order. 
//         You can treat it as just marking the person as dead.
//     string[] getInheritanceOrder() 
//         Returns a list representing the current order of inheritance excluding dead people.
    
// Example 1:
// Input
// ["ThroneInheritance", "birth", "birth", "birth", "birth", "birth", "birth", "getInheritanceOrder", "death", "getInheritanceOrder"]
// [["king"], ["king", "andy"], ["king", "bob"], ["king", "catherine"], ["andy", "matthew"], ["bob", "alex"], ["bob", "asha"], [null], ["bob"], [null]]
// Output
// [null, null, null, null, null, null, null, ["king", "andy", "matthew", "bob", "alex", "asha", "catherine"], null, ["king", "andy", "matthew", "alex", "asha", "catherine"]]
// Explanation
// ThroneInheritance t= new ThroneInheritance("king"); // order: king
// t.birth("king", "andy"); // order: king > andy
// t.birth("king", "bob"); // order: king > andy > bob
// t.birth("king", "catherine"); // order: king > andy > bob > catherine
// t.birth("andy", "matthew"); // order: king > andy > matthew > bob > catherine
// t.birth("bob", "alex"); // order: king > andy > matthew > bob > alex > catherine
// t.birth("bob", "asha"); // order: king > andy > matthew > bob > alex > asha > catherine
// t.getInheritanceOrder(); // return ["king", "andy", "matthew", "bob", "alex", "asha", "catherine"]
// t.death("bob"); // order: king > andy > matthew > bob > alex > asha > catherine
// t.getInheritanceOrder(); // return ["king", "andy", "matthew", "alex", "asha", "catherine"]
 
// Constraints:
//     1 <= kingName.length, parentName.length, childName.length, name.length <= 15
//     kingName, parentName, childName, and name consist of lowercase English letters only.
//     All arguments childName and kingName are distinct.
//     All name arguments of death will be passed to either the constructor or as childName to birth first.
//     For each call to birth(parentName, childName), it is guaranteed that parentName is alive.
//     At most 105 calls will be made to birth and death.
//     At most 10 calls will be made to getInheritanceOrder.

import "fmt"

// // DFS + Adjoint List + Map
// type ThroneInheritance struct {
//     Family map[string][]string
//     Order []string
    
//     DeathList map[string]bool
//     Used map[string]bool
//     Result []string
// }

// func Constructor(kingName string) ThroneInheritance {
//     t := ThroneInheritance{}
//     t.Family = map[string][]string{}
//     t.DeathList = map[string]bool{}
//     t.Order = append(t.Order, kingName)
//     t.Family[kingName] = make([]string, 0)
//     return t
// }

// func (t *ThroneInheritance) Birth(parentName string, childName string)  {
//     t.Family[parentName] = append(t.Family[parentName], childName)
//     t.Order = append(t.Order, childName)
// }

// func (t *ThroneInheritance) Death(name string)  {
//     t.DeathList[name] = true
// }

// func (t *ThroneInheritance) GetInheritanceOrder() []string {
//     t.Used = map[string]bool{}
//     t.Result = make([]string,0)
//     for _, n := range t.Order {
//         t.dfs(n)
//     }
//     return t.Result
// }

// func (t *ThroneInheritance) dfs(name string) {
//     _, isDeath := t.DeathList[name]
//     _, isPlaced := t.Used[name]
//     if !isDeath && !isPlaced {
//         t.Result = append(t.Result, name)
//     }
//     t.Used[name] = true
//     for _, val := range t.Family[name] {
//         _, isDeath := t.DeathList[val]
//         _, isPlaced := t.Used[val]
//         if !isDeath && !isPlaced {
//             t.Result = append(t.Result, val)
//             t.Used[val] = true
//         }
//         if _, found := t.Family[val]; found && !isPlaced {
//             t.dfs(val)
//         }
//     }
// }

type ThroneInheritance struct {
    king  string // 王的名字
    edges map[string][]string // 记录每个人下面的孩子名字
    dead  map[string]bool // 记录是否挂了 持了需要从 GetInheritanceOrder 移除
}

func Constructor(kingName string) (t ThroneInheritance) {
    return ThroneInheritance{kingName, map[string][]string{}, map[string]bool{}}
}

func (t *ThroneInheritance) Birth(parentName, childName string) {
    t.edges[parentName] = append(t.edges[parentName], childName)
}

func (t *ThroneInheritance) Death(name string) {
    t.dead[name] = true
}

func (t *ThroneInheritance) GetInheritanceOrder() []string {
    res := []string{}
    var preorder func(string)
    preorder = func(name string) {
        // 如果没有挂就加入到列表中
        if !t.dead[name] {
            res = append(res, name)
        }
        // 如果有孩子递归出 孩子了名字
        for _, childName := range t.edges[name] {
            preorder(childName)
        }
    }
    preorder(t.king)
    return res
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */

func main() {
    // ThroneInheritance t= new ThroneInheritance("king"); // order: king
    obj := Constructor("king")
    fmt.Println(obj)
    // t.birth("king", "andy"); // order: king > andy
    obj.Birth("king","andy")
    fmt.Println(obj)
    // t.birth("king", "bob"); // order: king > andy > bob
    obj.Birth("king","bob")
    fmt.Println(obj)
    // t.birth("king", "catherine"); // order: king > andy > bob > catherine
    obj.Birth("king","catherine")
    fmt.Println(obj)
    // t.birth("andy", "matthew"); // order: king > andy > matthew > bob > catherine
    obj.Birth("andy", "matthew")
    fmt.Println(obj)
    // t.birth("bob", "alex"); // order: king > andy > matthew > bob > alex > catherine
    obj.Birth("bob", "alex")
    fmt.Println(obj)
    // t.birth("bob", "asha"); // order: king > andy > matthew > bob > alex > asha > catherine
    obj.Birth("bob", "asha")
    fmt.Println(obj)
    // t.getInheritanceOrder(); // return ["king", "andy", "matthew", "bob", "alex", "asha", "catherine"]
    obj.Birth("bob", "asha")
    fmt.Println(obj.GetInheritanceOrder())
    // t.death("bob"); // order: king > andy > matthew > bob > alex > asha > catherine
    obj.Death("bob")
    fmt.Println(obj)
    // t.getInheritanceOrder(); // return ["king", "andy", "matthew", "alex", "asha", "catherine"]
    fmt.Println(obj.GetInheritanceOrder())
}