package main

// LCP 73. 探险营地
// 探险家小扣的行动轨迹，都将保存在记录仪中。expeditions[i] 表示小扣第 i 次探险记录，用一个字符串数组表示。
// 其中的每个「营地」由大小写字母组成，通过子串 -> 连接。
//     例："Leet->code->Campsite"，表示到访了 "Leet"、"code"、"Campsite" 三个营地。

// expeditions[0] 包含了初始小扣已知的所有营地；
// 对于之后的第 i 次探险(即 expeditions[i] 且 i > 0)，如果记录中包含了之前均没出现的营地，则表示小扣 新发现 的营地。

// 请你找出小扣发现新营地最多且索引最小的那次探险，并返回对应的记录索引。
// 如果所有探险记录都没有发现新的营地，返回 -1

// 注意：
//     大小写不同的营地视为不同的营地；
//     营地的名称长度均大于 0。

// 示例 1：
// 输入：expeditions = ["leet->code","leet->code->Campsite->Leet","leet->code->leet->courier"]
// 输出：1
// 解释： 初始已知的所有营地为 "leet" 和 "code" 第 1 次，到访了 "leet"、"code"、"Campsite"、"Leet"，新发现营地 2 处："Campsite"、"Leet" 第 2 次，到访了 "leet"、"code"、"courier"，新发现营地 1 处："courier" 第 1 次探险发现的新营地数量最多，因此返回 1

// 示例 2：
// 输入：expeditions = ["Alice->Dex","","Dex"]
// 输出：-1
// 解释： 初始已知的所有营地为 "Alice" 和 "Dex" 第 1 次，未到访任何营地； 第 2 次，到访了 "Dex"，未新发现营地； 因为两次探险均未发现新的营地，返回 -1

// 示例 3：
// 输入：expeditions = ["","Gryffindor->Slytherin->Gryffindor","Hogwarts->Hufflepuff->Ravenclaw"]
// 输出：2
// 解释： 初始未发现任何营地； 第 1 次，到访 "Gryffindor"、"Slytherin" 营地，其中重复到访 "Gryffindor" 两次， 因此新发现营地为 2 处："Gryffindor"、"Slytherin" 第 2 次，到访 "Hogwarts"、"Hufflepuff"、"Ravenclaw" 营地； 新发现营地 3 处："Hogwarts"、"Hufflepuff"、"Ravenclaw"； 第 2 次探险发现的新营地数量最多，因此返回 2

// 提示：
//     1 <= expeditions.length <= 1000
//     0 <= expeditions[i].length <= 1000
//     探险记录中只包含大小写字母和子串"->"

import "fmt"
import "strings"

func adventureCamp(expeditions []string) int {
    mp := make(map[string]bool) // 存储已知营地的地图
    for _, campsite := range strings.Split(expeditions[0], "->") {
        mp[campsite] = true // 增加从第一次探险到已知营地的营地
    }
    mx, res := 0, -1 // 存储已发现的最大数量的新营地, 存储具有最大新营地的探险队索引
    for i, expedition := range expeditions {
        if expedition == "" { continue }
        count := 0
        for _, campsite := range strings.Split(expedition, "->") { // 把探险队分成营地
            if !mp[campsite] && campsite != "" { // 检查营地是否是新的
                mp[campsite] = true 
                count++
            }
        }
        if count > mx { 
            mx, res = count, i
        }
    }
    if mx == 0 { return -1 }
    return res
}

func main() {
    // 示例 1：
    // 输入：expeditions = ["leet->code","leet->code->Campsite->Leet","leet->code->leet->courier"]
    // 输出：1
    // 解释： 初始已知的所有营地为 "leet" 和 "code" 第 1 次，到访了 "leet"、"code"、"Campsite"、"Leet"，新发现营地 2 处："Campsite"、"Leet" 第 2 次，到访了 "leet"、"code"、"courier"，新发现营地 1 处："courier" 第 1 次探险发现的新营地数量最多，因此返回 1
    fmt.Println(adventureCamp([]string{"leet->code","leet->code->Campsite->Leet","leet->code->leet->courier"})) // 1
    // 示例 2：
    // 输入：expeditions = ["Alice->Dex","","Dex"]
    // 输出：-1
    // 解释： 初始已知的所有营地为 "Alice" 和 "Dex" 第 1 次，未到访任何营地； 第 2 次，到访了 "Dex"，未新发现营地； 因为两次探险均未发现新的营地，返回 -1
    fmt.Println(adventureCamp([]string{"Alice->Dex","","Dex"})) // -1
    // 示例 3：
    // 输入：expeditions = ["","Gryffindor->Slytherin->Gryffindor","Hogwarts->Hufflepuff->Ravenclaw"]
    // 输出：2
    // 解释： 初始未发现任何营地； 第 1 次，到访 "Gryffindor"、"Slytherin" 营地，其中重复到访 "Gryffindor" 两次， 因此新发现营地为 2 处："Gryffindor"、"Slytherin" 第 2 次，到访 "Hogwarts"、"Hufflepuff"、"Ravenclaw" 营地； 新发现营地 3 处："Hogwarts"、"Hufflepuff"、"Ravenclaw"； 第 2 次探险发现的新营地数量最多，因此返回 2
    fmt.Println(adventureCamp([]string{"","Gryffindor->Slytherin->Gryffindor","Hogwarts->Hufflepuff->Ravenclaw"})) // 2
}