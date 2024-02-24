// 2796. Repeat String
// Write code that enhances all strings such that you can call the string.replicate(x) method on any string and it will return repeated string x times.
// Try to implement it without using the built-in method string.repeat.

// Example 1:
// Input: str = "hello", times = 2
// Output: "hellohello"
// Explanation: "hello" is repeated 2 times

// Example 2:
// Input: str = "code", times = 3
// Output: "codecodecode"
// Explanation: "code" is repeated 3 times

// Example 3:
// Input: str = "js", times = 1
// Output: "js"
// Explanation: "js" is repeated 1 time

// Constraints:
//         1 <= str.length, times <= 10^5

/**
 * @param {number} times
 * @return {string}
 */
String.prototype.replicate = function(times) {
    let res = "";
    for(let i = 0; i <= times - 1; i++) {
        res += this;
    }
    return res;
}

String.prototype.replicate1 = function(times) {
    let s = this.toString()
    
    function dfs(n) {
        if (n == 1) {
            return s
        }
        var ns = dfs(n >> 1)
        if (n%2) return ns+ns+s
        return ns + ns
    }
    return dfs(times)
}

console.log("hello".replicate(2)) // hellohello
console.log("code".replicate(3)) // codecodecode
console.log("js".replicate(1)) // js

console.log("hello".replicate1(2)) // hellohello
console.log("code".replicate1(3)) // codecodecode
console.log("js".replicate1(1)) // js