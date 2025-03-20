// 2691. Immutability Helper
// Creating clones of immutable objects with minor alterations can be a tedious process. Write a class ImmutableHelper that serves as a tool to help with this requirement. 
// The constructor accepts an immutable object obj which will be a JSON object or array.
// The class has a single method produce which accepts a function mutator. The function returns a new object which is similar to the original except it has those mutations applied.
// mutator accepts a proxied version of obj. A user of this function can (appear to) mutate this object, but the original object obj should not actually be effected.
// For example, a user could write code like this
//         const originalObj = {"x": 5};
//         const helper = new ImmutableHelper(originalObj);
//         const newObj = helper.produce((proxy) => {
//         proxy.x = proxy.x + 1;
//         });
//         console.log(originalObj); // {"x": 5}
//         console.log(newObj); // {"x": 6}

// Properties of the mutator function:
//         It will always return undefined.
//         It will never access keys that don't exist.
//         It will never delete keys (delete obj.key)
//         It will never call methods on a proxied object (push, shift, etc).
//         It will never set keys to objects (proxy.x = {})

// Note on how the solution will be tested: 
//         the solution validator will only analyze differences between what was returned and the original obj. 
//         Doing a full comparison would be too computationally expensive. Also, any mutations to the original object will result in a wrong answer.


// Example 1:
// Input: 
// obj = {"val": 10}, 
// mutators = [
//   proxy => { proxy.val += 1; },
//   proxy => { proxy.val -= 1; }
// ]
// Output: 
// [
//   {"val": 11},
//   {"val": 9}
// ]
// Explanation:
// const helper = new ImmutableHelper({val: 10});
// helper.produce(proxy => { proxy.val += 1; }); // { "val": 11 }
// helper.produce(proxy => { proxy.val -= 1; }); // { "val": 9 }

// Example 2:
// Input: 
// obj = {"arr": [1, 2, 3]} 
// mutators = [
//  proxy => { 
//    proxy.arr[0] = 5; 
//    proxy.newVal = proxy.arr[0] + proxy.arr[1];
//  }
// ]
// Output: 
// [
//   {"arr": [5, 2, 3], "newVal": 7 } 
// ]
// Explanation: 
//     Two edits were made to the original array. 
//     The first element in the array was to set 5. 
//     Then a new key was added with a value of 7.

// Example 3:
// Input: 
// obj = {"obj": {"val": {"x": 10, "y": 20}}}
// mutators = [
//   proxy => { 
//     let data = proxy.obj.val; 
//     let temp = data.x; 
//     data.x = data.y; 
//     data.y = temp; 
//   }
// ]
// Output: 
// [
//   {"obj": {"val": {"x": 20, "y": 10}}}
// ]
// Explanation: The values of "x" and "y" were swapped.
 
// Constraints:
//         2 <= JSON.stringify(obj).length <= 4 * 10^5
//         mutators is an array of functions
//         total calls to produce() < 10^5

const isObj = o => typeof o === 'object' && o !== null;

var ImmutableHelper = function(obj) {
    this.obj = obj;
};

const setMapConstructor = () => {
    return {
        hasValue: false,
        value: null,
        map: new Map(),
    }
}


ImmutableHelper.prototype.produce = function(mutator) {
    function createProxy(obj, setMap) {
        return new Proxy(obj, {
              set(_, prop, value) {
                  if (!setMap.map.has(prop)) {
                      setMap.map.set(prop, setMapConstructor());
                  }
                  setMap.map.get(prop).hasValue = true;
                  setMap.map.get(prop).value = value;
              },
              get(_, prop, value) {
                  if (setMap.hasValue) {
                      return setMap.value;
                  }

                  if (!setMap.map.has(prop)) {
                      setMap.map.set(prop, setMapConstructor());
                  }
                  if (setMap.map.get(prop).hasValue) {
                      return setMap.map.get(prop).value;
                  }
                  if (isObj(obj[prop])) {
                      return createProxy(obj[prop], setMap.map.get(prop));
                  }
                  return obj[prop];
              },
        });
    }
    
    function simplify(setMap) {
        if (setMap.hasValue) return true
        let hasMut = false;
        for (const [key, value] of [...setMap.map]) {
            if (!simplify(value)) {
                setMap.map.delete(key)
            } else {
                hasMut = true
            }
        }
        return hasMut;
    }
    
    function transform(obj, setMap) {
        if (setMap.hasValue) {
            return setMap.value
        }
        if (setMap.map.size === 0) return obj;
        let clone;
        if (isObj(obj)) {
            if (Array.isArray(obj)) {
                clone = [...obj]
            } else {
                clone = {...obj};
            }
        } else {
            return obj;
        }
        for (const [key, value] of [...setMap.map]) {
            clone[key] = transform(obj[key], value);
        }
        return clone
    }
    
    const setMap = setMapConstructor();
    const proxy = createProxy(this.obj, setMap);
   
    mutator(proxy);
    simplify(setMap)
    return transform(this.obj, setMap);
};

/**
 * const originalObj = {"x": 5};
 * const mutator = new ImmutableHelper(originalObj);
 * const newObj = mutator.produce((proxy) => {
 *   proxy.x = proxy.x + 1;
 * });
 * console.log(originalObj); // {"x: 5"}
 * console.log(newObj); // {"x": 6}
 */
const originalObj = {"x": 5};
const mutator = new ImmutableHelper(originalObj);
const newObj = mutator.produce((proxy) => {
    proxy.x = proxy.x + 1;
});
console.log(originalObj); // {"x: 5"}
console.log(newObj); // {"x": 6}

const helper = new ImmutableHelper({val: 10});
console.log(helper.produce(proxy => { proxy.val += 1; })); // { "val": 11 }
console.log(helper.produce(proxy => { proxy.val -= 1; })); // { "val": 9 }
console.log(helper) // ImmutableHelper { obj: { val: 10 } }

const helper2 = new ImmutableHelper({"arr": [1, 2, 3]});
console.log(helper2.produce(proxy => { 
    proxy.arr[0] = 5; 
    proxy.newVal = proxy.arr[0] + proxy.arr[1];
})); // [{"arr": [5, 2, 3], "newVal": 7 }]
console.log(helper2)

const helper3 = new ImmutableHelper({"obj": {"val": {"x": 10, "y": 20}}});
console.log(helper3.produce(proxy => { 
    let data = proxy.obj.val; 
    let temp = data.x; 
    data.x = data.y; 
    data.y = temp; 
})); // {"obj": {"val": {"x": 20, "y": 10}}}
console.log(helper3)

// {"arr":[1,2,3]}
// [proxy => { proxy.arr[0] = 5; proxy.newVal = proxy.arr[0] + proxy.arr[1]; }]

// Proxy是ES6中新增的一个功能，它可以在某个对象前架设一个“拦截器”，从而可以对该对象的访问进行拦截和控制。
// 可以理解为是对对象访问的一个代理，通过代理可以改变对象的默认行为。

// Proxy的作用主要有以下几个方面：
//     1. 对象的拦截和控制：可以对对象的属性访问、赋值、函数调用等操作进行拦截和控制，从而实现对对象行为的定制。
//     2. 数据劫持：可以通过Proxy实现数据双向绑定、深度监听、表单校验等数据劫持操作。
//     3. 权限控制：可以使用Proxy实现对象属性的访问权限控制，限制一些敏感属性的访问。
//     4. 性能优化：可以使用Proxy进行缓存、懒加载和单例模式等性能优化操作。

// Proxy和Object.defineProperty都可以用于拦截和控制对象的属性访问，但是它们之间有以下几个区别：
//     1. Proxy支持拦截更多的操作：Proxy可以拦截更多的对象操作，包括对象属性的读取和设置、函数调用、in操作符、for...in循环等等，而Object.defineProperty只能拦截属性的访问和设置。
//     2. Proxy是基于对象的拦截：Proxy是基于对象的拦截，即一个Proxy实例对应一个被拦截的对象，通过代理可以改变整个对象的行为。而Object.defineProperty是基于属性的拦截，可以对单个属性进行拦截。
//     3. Proxy具有别名效应：在不需要拦截的情况下，可以直接使用对象的别名和引用来访问对象。而Object.defineProperty修改对象行为之后，不可以直接使用对象的别名和引用来访问属性。
//     4. Proxy可以使用Reflect对象：Proxy通过Reflect对象来执行默认操作，而Object.defineProperty则不能
