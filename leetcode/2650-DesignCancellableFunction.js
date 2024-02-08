// 2650. Design Cancellable Function
// Sometimes you have a long running task, and you may wish to cancel it before it completes. 
// To help with this goal, write a function cancellable that accepts a generator object and returns an array of two values: a cancel function and a promise.

// You may assume the generator function will only yield promises. 
// It is your function's responsibility to pass the values resolved by the promise back to the generator. 
// If the promise rejects, your function should throw that error back to the generator.

// If the cancel callback is called before the generator is done, your function should throw an error back to the generator. 
// That error should be the string "Cancelled" (Not an Error object). If the error was caught, the returned promise should resolve with the next value that was yielded or returned. 
// Otherwise, the promise should reject with the thrown error. No more code should be executed.

// When the generator is done, the promise your function returned should resolve the value the generator returned. 
// If, however, the generator throws an error, the returned promise should reject with the error.

// An example of how your code would be used:

//         function* tasks() {
//         const val = yield new Promise(resolve => resolve(2 + 2));
//         yield new Promise(resolve => setTimeout(resolve, 100));
//         return val + 1; // calculation shouldn't be done.
//         }
//         const [cancel, promise] = cancellable(tasks());
//         setTimeout(cancel, 50);
//         promise.catch(console.log); // logs "Cancelled" at t=50ms

// If instead cancel() was not called or was called after t=100ms, the promise would have resolved 5.

// Example 1:
// Input: 
// generatorFunction = function*() { 
//   return 42; 
// }
// cancelledAt = 100
// Output: {"resolved": 42}
// Explanation:
// const generator = generatorFunction();
// const [cancel, promise] = cancellable(generator);
// setTimeout(cancel, 100);
// promise.then(console.log); // resolves 42 at t=0ms
// The generator immediately yields 42 and finishes. Because of that, the returned promise immediately resolves 42. Note that cancelling a finished generator does nothing.

// Example 2:
// Input:
// generatorFunction = function*() { 
//   const msg = yield new Promise(res => res("Hello")); 
//   throw `Error: ${msg}`; 
// }
// cancelledAt = null
// Output: {"rejected": "Error: Hello"}
// Explanation:
// A promise is yielded. The function handles this by waiting for it to resolve and then passes the resolved value back to the generator. Then an error is thrown which has the effect of causing the promise to reject with the same thrown error.

// Example 3:
// Input: 
// generatorFunction = function*() { 
//   yield new Promise(res => setTimeout(res, 200)); 
//   return "Success"; 
// }
// cancelledAt = 100
// Output: {"rejected": "Cancelled"}
// Explanation:
// While the function is waiting for the yielded promise to resolve, cancel() is called. This causes an error message to be sent back to the generator. Since this error is uncaught, the returned promise rejected with this error.

// Example 4:
// Input:
// generatorFunction = function*() { 
//   let result = 0; 
//   yield new Promise(res => setTimeout(res, 100));
//   result += yield new Promise(res => res(1)); 
//   yield new Promise(res => setTimeout(res, 100)); 
//   result += yield new Promise(res => res(1)); 
//   return result;
// }
// cancelledAt = null
// Output: {"resolved": 2}
// Explanation:
// 4 promises are yielded. Two of those promises have their values added to the result. After 200ms, the generator finishes with a value of 2, and that value is resolved by the returned promise.

// Example 5:
// Input: 
// generatorFunction = function*() { 
//   let result = 0; 
//   try { 
//     yield new Promise(res => setTimeout(res, 100)); 
//     result += yield new Promise(res => res(1)); 
//     yield new Promise(res => setTimeout(res, 100)); 
//     result += yield new Promise(res => res(1)); 
//   } catch(e) { 
//     return result; 
//   } 
//   return result; 
// }
// cancelledAt = 150
// Output: {"resolved": 1}
// Explanation:
// The first two yielded promises resolve and cause the result to increment. However, at t=150ms, the generator is cancelled. The error sent to the generator is caught and the result is returned and finally resolved by the returned promise.

// Example 6:
// Input: 
// generatorFunction = function*() { 
//   try { 
//     yield new Promise((resolve, reject) => reject("Promise Rejected")); 
//   } catch(e) { 
//     let a = yield new Promise(resolve => resolve(2));
//     let b = yield new Promise(resolve => resolve(2)); 
//     return a + b; 
//   }; 
// }
// cancelledAt = null
// Output: {"resolved": 4}
// Explanation:
// The first yielded promise immediately rejects. This error is caught. Because the generator hasn't been cancelled, execution continues as usual. It ends up resolving 2 + 2 = 4.
 
// Constraints:
//         cancelledAt == null or 0 <= cancelledAt <= 1000
//         generatorFunction returns a generator object

// # 什么是 generator ？
//     简单的说就是可以中断的函数，该函数里的代码可以用 yield 关键字来中断。如何继续执行？看下面

// # 怎么定义一个 generator ：
//     一般来说都用 function* 标识符来定义，例如
//     function* tasks() {
//         const val = yield new Promise(resolve => resolve(2 + 2));
//         yield new Promise(resolve => setTimeout(resolve, 100));
//         return val + 1;
//     }

// # generator 的实例，或者说返回值，有哪些方法？
//     function* tasks() {
//         const val = yield new Promise(resolve => resolve(2 + 2));
//         yield new Promise(resolve => setTimeout(resolve, 100));
//         return val + 1;
//     }
//     const _t = tasks();
//     总的来说有三种方法，next、throw, return。这三种方法都会返回一个 IteratorResult，也就是形如：
//     interface IteratorResult {
//         value: any;
//         done: boolean
//     }
//     done 则表示该 generator 是否完成，true 完成，false 未完成
//     value 不同方法调用代表的值不同，下面再说。

// # 那么 generator 啥时候完成？就是走到这个 generator 本身 return 的时候。 
//     下面开始介绍三个方法
//     ## next 方法：
//         在 next 方法的返回值中，IteratorResult 中的 value 是 yield 关键词右侧代码的执行结果
//         generator 的执行必需由 next 方法开启，并走到 yield 关键词停止
//         继续调用next() 可以在 yield 后继续往下走，走到下一个 yield 为止
//         该方法同时支持传参，传入的参数会成为上一个 yield 的返回值，例如：
//             function* tasks() {
//                 const val = yield new Promise(resolve => resolve(2 + 2)); // 1
//                 console.log({val});
//                 yield new Promise(resolve => setTimeout(resolve, 100)); // 2
//                 return val + 1;
//             }
//             const _t = tasks();
//             _t.next(); // 走到 1 处，返回的 {value， done} 中的 value 是 1 处的 Promise
//             _t.next(100); // 走到 2 处，打印 val 为 100
//     ## throw 方法：
//         在 throw 方法的返回值中，IteratorResult 中的 value 同样是 yield 关键词右侧代码的执行结果
//         该方法通过传参抛出一个错误，如果该错误没有被 generator 本身 catch 住，则会往外暴露给外层，也就是 generator 的调用方。如果调用方也没有 catch 住，则正常抛错。
//         例子1，被 generator 本身 catch：
//         function* tasks() {
//             try {
//                 const val = yield new Promise(resolve => resolve(2 + 2));
//             } catch (err) {
//                 console.log('catched by generator *tasks', err);
//             }
//             yield new Promise(resolve => setTimeout(resolve, 100));
//             return val + 1;
//         }
//         const _t = tasks();
//         _t.next();
//         _t.throw('err 1');
//         例子2，被 generator 的调用方 catch：
//         function* tasks() {
//             const val = yield new Promise(resolve => resolve(2 + 2));
//             yield new Promise(resolve => setTimeout(resolve, 100));
//             return val + 1;
//         }
//         const _t = tasks();
//         _t.next();
//         try {
//             _t.throw('err 2');
//         } catch(err) {
//             console.log('catched by generator *tasks caller', err);
//         }
//     ## return 方法
//         在 return 方法的返回值中，IteratorResult 中的 value 是 return 方法传递的参数
//         该方法强制 generator 函数完成，其返回 IteratorResult 中的 done 将为 true
//         举例
//         function* tasks() {
//             const val = yield new Promise(resolve => resolve(2 + 2));
//             yield new Promise(resolve => setTimeout(resolve, 100));
//             return val + 1;
//         }
//         const _t = tasks();
//         _t.next();
//         const obj = _t.return('Return by Generator.return method');
//         console.log(obj); // {value: 'Return by Generator.return method', done: true}

/**
 * @param {Generator} generator
 * @return {[Function, Promise]}
 */
var cancellable = function(generator) {
    let cancel = () => {};
    const p = new Promise((resolve, reject) => {
        cancel = (msg = 'Cancelled') => {
            run(msg, 'throw');
        };
        const run = (ret, fnName = 'next') => {
            try {
                const {value, done} = generator[fnName](ret);
                if (done) {
                    resolve(value);
                    return;
                }
                value.then((val) => {
                    run(val);
                }).catch((err) => {
                    run(err, 'throw');
                });
            } catch (errorByGenerator) {
                reject(errorByGenerator);
            }
        };
        run(null);
    });
    return [cancel, p];
};

/**
 * @param {Generator} generator
 * @return {[Function, Promise]}
 */
var cancellable1 = function (generator) {
    let hasCancel = false;
    let cancel = () => {},
      promise;
  
    promise = new Promise((resolve, reject) => {
      function run(preValue, method) {
        try {
          const { value, done } = method
            ? generator[method](preValue)
            : generator.next(preValue);
  
          if (done) {
            return resolve(value);
          }
  
          Promise.resolve(value)
            .then((data) => {
              !hasCancel && run(data);
            })
            .catch((err) => {
              !hasCancel && run(err, "throw");
            });
        } catch (err) {
          reject(err);
        }
      }
  
      run();
  
      cancel = () => {
        hasCancel = true;
        run("Cancelled", "throw");
      };
    });
  
    return [cancel, promise];
  };

/**
 * function* tasks() {
 *   const val = yield new Promise(resolve => resolve(2 + 2));
 *   yield new Promise(resolve => setTimeout(resolve, 100));
 *   return val + 1;
 * }
 * const [cancel, promise] = cancellable(tasks());
 * setTimeout(cancel, 50);
 * promise.catch(console.log); // logs "Cancelled" at t=50ms
 */

function* tasks() {
    const val = yield new Promise(resolve => resolve(2 + 2));
    yield new Promise(resolve => setTimeout(resolve, 100));
    return val + 1;
}
let [cancel, promise] = cancellable(tasks());
setTimeout(cancel, 50);
promise.catch(console.log); // logs "Cancelled" at t=50ms


// Example 1:
let generatorFunction = function*() { 
  return 42; 
}
let generator = generatorFunction();
[cancel, promise] = cancellable(generator);
setTimeout(cancel, 100);
promise.then(console.log); // resolves 42 at t=0ms

let generator11 = generatorFunction();
let [cancel11, promise11] = cancellable1(generator11);
setTimeout(cancel11, 100);
promise11.then(console.log); // resolves 42 at t=0ms

// Example 2:
// try {
//     generatorFunction = function*() { 
//     const msg = yield new Promise(res => res("Hello")); 
//     throw `Error: ${msg}`; 
//     }
//     generator = generatorFunction();
//     [cancel, promise] = cancellable(generator);
//     //setTimeout(cancel, 100);
//     promise.then(console.log); // {"rejected": "Error: Hello"} at t=0ms
// } catch (e) {
//     console.log(e.message)
// }
try {
    let generatorFunction22 = function*() { 
        const msg = yield new Promise(res => res("Hello")); 
        throw `Error: ${msg}`; 
    }
    let generator22 = generatorFunction22();
    let [cancel22, promise22] = cancellable(generator22);
    //setTimeout(cancel22, 100);
    promise22.then(console.log); // {"rejected": "Error: Hello"} at t=0ms
} catch (e) {
    console.log(e.message)
}


// Example 3:
// const generatorFunction3 = function*() { 
//     yield new Promise(res => setTimeout(res, 200)); 
//     return "Success"; 
// }
// const generator3 = generatorFunction3();
// const [cancel3, promise3] = cancellable(generator3);
// setTimeout(cancel3, 100);
// promise3.then(console.log); // {"rejected": "Cancelled"} at t=100ms

// Input: 
// generatorFunction = function*() { 
//   yield new Promise(res => setTimeout(res, 200)); 
//   return "Success"; 
// }
// cancelledAt = 100
// Output: {"rejected": "Cancelled"}
// Explanation:
// While the function is waiting for the yielded promise to resolve, cancel() is called. This causes an error message to be sent back to the generator. Since this error is uncaught, the returned promise rejected with this error.

// Example 4:
// Input:
// generatorFunction = function*() { 
//   let result = 0; 
//   yield new Promise(res => setTimeout(res, 100));
//   result += yield new Promise(res => res(1)); 
//   yield new Promise(res => setTimeout(res, 100)); 
//   result += yield new Promise(res => res(1)); 
//   return result;
// }
// cancelledAt = null
// Output: {"resolved": 2}
// Explanation:
// 4 promises are yielded. Two of those promises have their values added to the result. After 200ms, the generator finishes with a value of 2, and that value is resolved by the returned promise.

// Example 5:
// Input: 
// generatorFunction = function*() { 
//   let result = 0; 
//   try { 
//     yield new Promise(res => setTimeout(res, 100)); 
//     result += yield new Promise(res => res(1)); 
//     yield new Promise(res => setTimeout(res, 100)); 
//     result += yield new Promise(res => res(1)); 
//   } catch(e) { 
//     return result; 
//   } 
//   return result; 
// }
// cancelledAt = 150
// Output: {"resolved": 1}
// Explanation:
// The first two yielded promises resolve and cause the result to increment. However, at t=150ms, the generator is cancelled. The error sent to the generator is caught and the result is returned and finally resolved by the returned promise.

// Example 6:
// Input: 
// generatorFunction = function*() { 
//   try { 
//     yield new Promise((resolve, reject) => reject("Promise Rejected")); 
//   } catch(e) { 
//     let a = yield new Promise(resolve => resolve(2));
//     let b = yield new Promise(resolve => resolve(2)); 
//     return a + b; 
//   }; 
// }
// cancelledAt = null
// Output: {"resolved": 4}
// Explanation:
// The first yielded promise immediately rejects. This error is caught. Because the generator hasn't been cancelled, execution continues as usual. It ends up resolving 2 + 2 = 4.
 