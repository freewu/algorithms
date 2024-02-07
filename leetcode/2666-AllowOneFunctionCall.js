// 2666. Allow One Function Call
// Given a function fn, return a new function that is identical to the original function except that it ensures fn is called at most once.
//         The first time the returned function is called, it should return the same result as fn.
//         Every subsequent time it is called, it should return undefined.
 
// Example 1:
// Input: fn = (a,b,c) => (a + b + c), calls = [[1,2,3],[2,3,6]]
// Output: [{"calls":1,"value":6}]
// Explanation:
// const onceFn = once(fn);
// onceFn(1, 2, 3); // 6
// onceFn(2, 3, 6); // undefined, fn was not called

// Example 2:
// Input: fn = (a,b,c) => (a * b * c), calls = [[5,7,4],[2,3,6],[4,6,8]]
// Output: [{"calls":1,"value":140}]
// Explanation:
// const onceFn = once(fn);
// onceFn(5, 7, 4); // 140
// onceFn(2, 3, 6); // undefined, fn was not called
// onceFn(4, 6, 8); // undefined, fn was not called

// Constraints:
//         calls is a valid JSON array
//         1 <= calls.length <= 10
//         1 <= calls[i].length <= 100
//         2 <= JSON.stringify(calls).length <= 1000

/**
 * @param {Function} fn
 * @return {Function}
 */
var once = function(fn) {
    // 新建一个闭包私有变量singleton，初始无值。
    let singleton = null
    return function (...args) {
        // 当第一次调用时，singleton 赋值同时返回其值
        // 后续 singleton 有值的情况下，只会返回 undefined
        //return !singleton ? singleton = fn(...args) : undefined // 如果第一个返回 0 或 false 会出现误导
        return singleton === null ? singleton = fn(...args) : undefined
    }
};

// best solution
var once1 = function(fn) {
    return function(...args){
        try{
            return fn(...args)
        }
        finally{
            fn=()=>{}
        }
    }
};

/**
 * let fn = (a,b,c) => (a + b + c)
 * let onceFn = once(fn)
 *
 * onceFn(1,2,3); // 6
 * onceFn(2,3,6); // returns undefined without calling fn
 */

let fn = (a,b,c) => (a + b + c)
let onceFn = once(fn)

console.log(onceFn(1,2,3)); // 6
try {
    console.log(onceFn(2,3,6)); // returns undefined without calling fn
} catch (err) {
    console.log(err.message)
}

let onceFn1 = once1(fn)

console.log(onceFn1(1,2,3)); // 6
try {
    console.log(onceFn1(2,3,6)); // returns undefined without calling fn
} catch (err) {
    console.log(err.message)
}