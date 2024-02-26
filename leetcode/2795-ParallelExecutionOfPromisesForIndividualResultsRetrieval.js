// 2795. Parallel Execution of Promises for Individual Results Retrieval
// Given an array functions, return a promise promise. functions is an array of functions that return promises fnPromise. 
// Each fnPromise can be resolved or rejected.  

// If fnPromise is resolved:
//         obj = { status: "fulfilled", value: resolved value}
// If fnPromise is rejected:
//         obj = { status: "rejected", reason: reason of rejection (catched error message)}

// The promise should resolve with an array of these objects obj. 
// Each obj in the array should correspond to the promises in the original array function, maintaining the same order.
// Try to implement it without using the built-in method Promise.allSettled().


// Example 1:
// Input: functions = [
//     () => new Promise(resolve => setTimeout(() => resolve(15), 100))
// ]
// Output: {"t":100,"values":[{"status":"fulfilled","value":15}]}
// Explanation: 
// const time = performance.now()
// const promise = promiseAllSettled(functions);           
// promise.then(res => {
//     const out = {t: Math.floor(performance.now() - time), values: res}
//     console.log(out) // {"t":100,"values":[{"status":"fulfilled","value":15}]}
// })
// The returned promise resolves within 100 milliseconds. Since promise from the array functions is fulfilled, the resolved value of the returned promise is set to [{"status":"fulfilled","value":15}].

// Example 2:
// Input: functions = [
//     () => new Promise(resolve => setTimeout(() => resolve(20), 100)), 
//     () => new Promise(resolve => setTimeout(() => resolve(15), 100))
// ]
// Output: 
// {
//     "t":100,
//     "values": [
//         {"status":"fulfilled","value":20},
//         {"status":"fulfilled","value":15}
//     ]
// }
// Explanation: The returned promise resolves within 100 milliseconds, because the resolution time is determined by the promise that takes the longest time to fulfill. Since promises from the array functions are fulfilled, the resolved value of the returned promise is set to [{"status":"fulfilled","value":20},{"status":"fulfilled","value":15}].

// Example 3:
// Input: functions = [
//     () => new Promise(resolve => setTimeout(() => resolve(30), 200)), 
//     () => new Promise((resolve, reject) => setTimeout(() => reject("Error"), 100))
// ]
// Output:
// {
//     "t":200,
//     "values": [
//         {"status":"fulfilled","value":30},
//         {"status":"rejected","reason":"Error"}
//     ]
// }
// Explanation: The returned promise resolves within 200 milliseconds, as its resolution time is determined by the promise that takes the longest time to fulfill. Since one promise from the array function is fulfilled and another is rejected, the resolved value of the returned promise is set to an array containing objects in the following order: [{"status":"fulfilled","value":30}, {"status":"rejected","reason":"Error"}]. Each object in the array corresponds to the promises in the original array function, maintaining the same order.

// Constraints:
//         1 <= functions.length <= 10

/**
 * @param {Array<Function>} functions
 * @return {Promise<Array>}
 */
var promiseAllSettled = function(functions) {
    // const statusesPromise = Promise.allSettled(functions);
    // return new Promise(() => {
    //     try {
    //         return statusesPromise;
    //     } catch (error) {
    //         console.log(error);
    //         throw error;
    //     }
    // })
    

    //return Promise.allSettled(functions)
    return new Promise((resolve) => {
        let ret = []
        let settledNums = 0
        for (let i in functions) {
            functions[i]()
            .then(value => ({ status: 'fulfilled', value }))
            .catch(err => ({ status: 'rejected', reason: err }))
            .then(result => {
                ret[i] = result
                // 中断方式：需要记录完成 Promise 的数量，该值与 functions 数量一致时中断
                if (++settledNums === functions.length) resolve(ret)
            })
        }
    })
};


/**
 * const functions = [
 *    () => new Promise(resolve => setTimeout(() => resolve(15), 100))
 * ]
 * const time = performance.now()
 *
 * const promise = promiseAllSettled(functions);
 *              
 * promise.then(res => {
 *     const out = {t: Math.floor(performance.now() - time), values: res}
 *     console.log(out) // {"t":100,"values":[{"status":"fulfilled","value":15}]}
 * })
 */

let functions = [
    () => new Promise(resolve => setTimeout(() => resolve(15), 100))
]
//const time = performance.now()
let time = process.hrtime();
let promise = promiseAllSettled(functions);              
promise.then(res => {
    //const out = {t: Math.floor(performance.now() - time), values: res}
    const out = {t: Math.floor(process.hrtime() - time), values: res}
    console.log(out) // {"t":100,"values":[{"status":"fulfilled","value":15}]}
})


functions = [
    () => new Promise(resolve => setTimeout(() => resolve(20), 100)), 
    () => new Promise(resolve => setTimeout(() => resolve(15), 100))
]
//const time = performance.now()
time = process.hrtime();
promise = promiseAllSettled(functions);              
promise.then(res => {
    //const out = {t: Math.floor(performance.now() - time), values: res}
    const out = {t: Math.floor(process.hrtime() - time), values: res}
    console.log(out) // {"t":100,"values":[{"status":"fulfilled","value":15}]}
})


functions = [
    () => new Promise(resolve => setTimeout(() => resolve(30), 200)), 
    () => new Promise((resolve, reject) => setTimeout(() => reject("Error"), 100))
]
//const time = performance.now()
time = process.hrtime();
promise = promiseAllSettled(functions);              
promise.then(res => {
    //const out = {t: Math.floor(performance.now() - time), values: res}
    const out = {t: Math.floor(process.hrtime() - time), values: res}
    console.log(out) // {"t":100,"values":[{"status":"fulfilled","value":15}]}
})


// # promise.allSettled()详解
// 当需要执行并行和独立的异步操作并收集所有结果时，Promise.allSettled() 就是不错的选择，即使一些异步操作可能失败。
// 和promise.all()的区别：
//     返回的Promise总是以一系列状态实现，无论是否有一些（或者全部）Promise被拒绝。

// Promise.allSettled() 永远不会被 rejected