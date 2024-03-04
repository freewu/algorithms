// 2692. Make Object Immutable
// Write a function that takes an object obj and returns a new immutable version of this object.
// An immutable object is an object that can't be altered and will throw an error if any attempt is made to alter it.

// There are three types of error messages that can be produced from this new object.
//         Attempting to modify a key on the object will result in this error message: `Error Modifying: ${key}`.
//         Attempting to modify an index on an array will result in this error message: `Error Modifying Index: ${index}`.
//         Attempting to call a method that mutates an array will result in this error message: `Error Calling Method: ${methodName}`. You may assume the only methods that can mutate an array are ['pop', 'push', 'shift', 'unshift', 'splice', 'sort', 'reverse'].
//         obj is a valid JSON object or array, meaning it is the output of JSON.parse().

// Note that a string literal should be thrown, not an Error.

// Example 1:
// Input: 
// obj = {
//   "x": 5
// }
// fn = (obj) => { 
//   obj.x = 5;
//   return obj.x;
// }
// Output: {"value": null, "error": "Error Modifying: x"}
// Explanation: Attempting to modify a key on an object resuts in a thrown error. Note that it doesn't matter that the value was set to the same value as it was before.

// Example 2:
// Input: 
// obj = [1, 2, 3]
// fn = (arr) => { 
//   arr[1] = {}; 
//   return arr[2]; 
// }
// Output: {"value": null, "error": "Error Modifying Index: 1"}
// Explanation: Attempting to modify an array results in a thrown error.

// Example 3:
// Input: 
// obj = {
//   "arr": [1, 2, 3]
// }
// fn = (obj) => { 
//   obj.arr.push(4);
//   return 42;
// }
// Output: { "value": null, "error": "Error Calling Method: push"}
// Explanation: Calling a method that can result in a mutation results in a thrown error.

// Example 4:
// Input: 
// obj = {
//   "x": 2,
//   "y": 2
// }
// fn = (obj) => { 
//   return Object.keys(obj);
// }
// Output: {"value": ["x", "y"], "error": null}
// Explanation: No mutations were attempted so the function returns as normal.

// Constraints:
//         obj is a valid JSON object or array
//         2 <= JSON.stringify(obj).length <= 10^5

/**
 * @param {Object | Array} obj
 * @return {Object | Array} immutable obj
 */
// use Proxy
var makeImmutable = function(obj) {
    let newObj
    switch (Object.prototype.toString.call(obj)) {
        case "[object Object]": // 如果是 Object 递归给 value 加上 proxy
            newObj = {}
            for(key of Object.keys(obj)) newObj[key] = makeImmutable(obj[key]);
            break;
        case "[object Array]": // 如果是 Array 递归给 value 加上 proxy
            newObj = []
            for(let i = 0; i < obj.length; i++) newObj[i] = makeImmutable(obj[i]);
            break;
        default:
            return obj;
    }
    const handler = {
        get(target,p) {
            const methods  =  {
                'pop':true,
                'push':true,
                'shift':true,
                'unshift':true,
                'splice':true,
                'sort':true,
                'reverse':true
            };
            // Attempting to call a method that mutates an array will result in this error message: `Error Calling Method: ${methodName}`.
            // You may assume the only methods that can mutate an array are ['pop', 'push', 'shift', 'unshift', 'splice', 'sort', 'reverse'].
            if(methods[p]) {
                return () => {
                    throw `Error Calling Method: ${p}` 
                }
            }
            return target[p]
        },
        set(target,pos) {
            // Attempting to modify an index on an array will result in this error message: `Error Modifying Index: ${index}`.
            if(Array.isArray(target)) throw `Error Modifying Index: ${pos}`;
            // Attempting to modify a key on the object will result in this error message: `Error Modifying: ${key}`.
            throw `Error Modifying: ${pos}`;
        }
    }
    return new Proxy(newObj,handler);
};

// use Proxy + Reflect
var makeImmutable1 = function(obj) {
    const proxify = (obj) => {
        // 如果是数组，给可以改变数据的方法加上 Proxy
        if (Array.isArray(obj)) {
            ['pop', 'push', 'shift', 'unshift', 'splice', 'sort', 'reverse'].forEach(method => {
                obj[method] = new Proxy(obj[method], {
                    apply(target) {
                        throw `Error Calling Method: ${target.name}`
                    }
                })
            })
        }
        return new Proxy(obj, {
            get(target, prop) {
                const val = Reflect.get(target, prop);
                // 不为对象支持返回值 
                if (!(typeof val === 'object' && val !== null)) return val;
                return proxify(val);
            },
            set(target, prop) {
                const preStr = Array.isArray(target) ? 'Error Modifying Index' : 'Error Modifying';
                throw `${preStr}: ${String(prop)}`
            }
        })
    };
    return proxify(obj);
};

/**
 * const obj = makeImmutable({x: 5});
 * obj.x = 6; // throws "Error Modifying x"
 */
try {
    const obj = makeImmutable({x: 5});
    obj.x = 6; // throws "Error Modifying x"
} catch (err) {
    console.log(err) // Error Modifying x
}

try {
    const arr = makeImmutable([1, 2, 3]);
    arr[1] = {}; // throws "Error Modifying Index: 1"
} catch (err) {
    console.log(err) // Error Modifying Index: 1
}

try {
    const obj = makeImmutable({"arr": [1, 2, 3]});
    obj.arr.push(4); // throws "Error Calling Method: push"
} catch (err) {
    console.log(err) // Error Calling Method: push
}

try {
    const obj = makeImmutable({"x": 2,"y": 2});
    console.log(Object.keys(obj)) // ["x", "y"]
} catch (err) {
    console.log(err) // 
}

try {
    const obj = makeImmutable1({x: 5});
    obj.x = 6; // throws "Error Modifying x"
} catch (err) {
    console.log(err) // Error Modifying x
}

try {
    const arr = makeImmutable1([1, 2, 3]);
    arr[1] = {}; // throws "Error Modifying Index: 1"
} catch (err) {
    console.log(err) // Error Modifying Index: 1
}

try {
    const obj = makeImmutable1({"arr": [1, 2, 3]});
    obj.arr.push(4); // throws "Error Calling Method: push"
} catch (err) {
    console.log(err) // Error Calling Method: push
}

try {
    const obj = makeImmutable1({"x": 2,"y": 2});
    console.log(Object.keys(obj)) // ["x", "y"]
} catch (err) {
    console.log(err) // 
}