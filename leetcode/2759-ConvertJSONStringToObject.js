// 2759. Convert JSON String to Object
// Given a string str, return parsed JSON parsedStr. 
// You may assume the str is a valid JSON string hence it only includes strings, numbers, arrays, objects, booleans, and null. 
// str will not include invisible characters and escape characters. 

// Please solve it without using the built-in JSON.parse method.

// Example 1:
// Input: str = '{"a":2,"b":[1,2,3]}'
// Output: {"a":2,"b":[1,2,3]}
// Explanation: Returns the object represented by the JSON string.

// Example 2:
// Input: str = 'true'
// Output: true
// Explanation: Primitive types are valid JSON.

// Example 3:
// Input: str = '[1,5,"false",{"a":2}]'
// Output: [1,5,"false",{"a":2}]
// Explanation: Returns the array represented by the JSON string.

// Constraints:
//         str is a valid JSON string
//         1 <= str.length <= 10^5

/**
 * @param {string} str
 * @return {null|boolean|number|string|Array|Object}
 */
var jsonParse = function(str) {
    return JSON.parse(str);
};

var jsonParse1 = function(str) {
    // 当前读取字符串的下标位置变量
    let index = 0
    
    // 用于判断数字的特征表
    const book = new Set("0123456789-.".split(""))

    // 递归函数
    const dfs = () => {
        // 如果当前碰到 n 必然就是 null ，返回 null 并移动 4 下指针
        if (str[index] === "n") {
            index += 4
            return null
        }
        // 如果当前碰到 t 必然就是 true ，返回 null 并移动 4 下指针
        if (str[index] === "t") {
            index += 4
            return true
        }
        // 如果当前碰到 f 必然就是 false ，返回 null 并移动 4 下指针
        if (str[index] === "f") {
            index += 5
            return false
        }
        // 如果当前碰到 " 必然就是 字符串 ，在没有碰到下一个 " 号之前都记录下来即可，最后碰到结束符后移动 1 下指针
        if (str[index] === "\"") {
            let res = ""
            while (str[++index] !== "\"") res += str[index]
            index++
            return res
        }
        // 如果当前碰到 [ 必然就是 数组 ，在没有碰到下一个 ] 号之前都递归下去进行每个元素判别，最后碰到结束符 ] 后移动 1 下指针
        if (str[index] === "[") {
            const res = []
            index++ // 跳过第一个 [ 号
            while (str[index] !== "]") {
                res.push(dfs())
                if (str[index] === ",") index++ // 跳过 , 号
            }
            index++
            return res
        }
        // 如果当前碰到 { 必然就是 对象 ，在没有碰到下一个 } 号之前都递归下去进行每个元素判别，由于必然键值是字符串可以直接当做独立元素进行判别，直接递归即可，最后碰到结束符 } 后移动 1 下指针
        if (str[index] === "{") {
            const res = {}
            index++ // 跳过第一个 { 号
            while (str[index] !== "}") {
                const key = dfs()
                index++ // 跳过 : 号
                const value = dfs()
                res[key] = value
                if (str[index] === ",") index++ // 跳过 , 号
            }
            index++
            return res
        }
        // 如果都不是，只能是数字，把具体字符记录下来直接转即可
        let res = ""
        while (index < str.length) {
            if (book.has(str[index])) res = res + str[index]
            else break
            index++
        }
        return Number(res)
    }

    return dfs()
};

var jsonParse2 = function(str) {
    eval('var a = ' + str);
    return a;
};

console.log(jsonParse('{"a":2,"b":[1,2,3]}')) // {"a":2,"b":[1,2,3]}
console.log(jsonParse('true')) // true
console.log(jsonParse('[1,5,"false",{"a":2}]')) // [1,5,"false",{"a":2}]

console.log(jsonParse1('{"a":2,"b":[1,2,3]}')) // {"a":2,"b":[1,2,3]}
console.log(jsonParse1('true')) // true
console.log(jsonParse1('[1,5,"false",{"a":2}]')) // [1,5,"false",{"a":2}]

console.log(jsonParse2('{"a":2,"b":[1,2,3]}')) // {"a":2,"b":[1,2,3]}
console.log(jsonParse2('true')) // true
console.log(jsonParse2('[1,5,"false",{"a":2}]')) // [1,5,"false",{"a":2}]