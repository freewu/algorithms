package leetcode;

// 1115. Print FooBar Alternately
//  Suppose you are given the following code:
//     class FooBar {
//         public void foo() {
//             for (int i = 0; i < n; i++) {
//                 print("foo");
//             }
//         }
//        
//         public void bar() {
//             for (int i = 0; i < n; i++) {
//                 print("bar");
//             }
//         }
//     }

// The same instance of FooBar will be passed to two different threads:
// thread A will call foo(), while
// thread B will call bar().
// Modify the given program to output "foobar" n times.

// Example 1:

// Input: n = 1
// Output: "foobar"
// Explanation: There are two threads being fired asynchronously. One of them calls foo(), while the other calls bar().
// "foobar" is being output 1 time.


// Example 2:
// Input: n = 2
// Output: "foobarfoobar"
// Explanation: "foobar" is being output 2 times.
 

// Constraints:
//      1 <= n <= 1000


// 信号量 适合控制顺序
// Semaphore 的核心方法是 acquire 获取信号量和 release 释放信号量
// https://blog.csdn.net/admans/article/details/125957120
class FooBar {

    private int n;
    private Semaphore foo = new Semaphore(1);
    private Semaphore bar = new Semaphore(0);

    public FooBar(int n) {
        this.n = n;
    }

    public void foo(Runnable printFoo) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            foo.acquire();
        	printFoo.run();
            bar.release();
        }
    }
    public void bar(Runnable printBar) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            bar.acquire();
        	printBar.run();
            foo.release();
        }
    }
}

// BlockingQueue
// https://blog.csdn.net/qq_37774171/article/details/122742494
public class FooBar {
    private int n;
    private BlockingQueue<Integer> bar = new LinkedBlockingQueue<>(1);
    private BlockingQueue<Integer> foo = new LinkedBlockingQueue<>(1);

    public FooBar(int n) {
        this.n = n;
    }

    public void foo(Runnable printFoo) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            foo.put(i); // 先放 foo 队列, 后面还要放入需要先被取出
            printFoo.run();
            bar.put(i); // 放入 bar 队列
        }
    }

    public void bar(Runnable printBar) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            bar.take(); // 从 bar 队列中取出,如果队列中没有则等待
            printBar.run();
            foo.take(); // 从 foo 对列表中取出
        }
    }
}


// CyclicBarrier 控制先后
// https://blog.csdn.net/BASK2311/article/details/128145305
class FooBar {
    private int n;

    public FooBar(int n) {
        this.n = n;
    }

    CyclicBarrier cb = new CyclicBarrier(2);
    volatile boolean fin = true;

    public void foo(Runnable printFoo) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            while(!fin);
            printFoo.run();
            fin = false;
            try {
                cb.await();
            } catch (BrokenBarrierException e) {}
        }
    }

    public void bar(Runnable printBar) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            try {
                cb.await();
            } catch (BrokenBarrierException e) {}
            printBar.run();
            fin = true;
        }
    }
}

// 自旋 + 让出CPU
class FooBar {
    private int n;

    public FooBar(int n) {
        this.n = n;
    }

    volatile boolean permitFoo = true;

    public void foo(Runnable printFoo) throws InterruptedException {     
        for (int i = 0; i < n; ) {
            if(permitFoo) {
        	    printFoo.run();
            	i++;
            	permitFoo = false;
            } else {
                // Thread.yield() 方法，使当前线程由执行状态，变成为就绪状态，让出CPU，在下一个线程执行时候，此线程有可能被执行，也有可能没有被执行
                Thread.yield();
            }
        }
    }

    public void bar(Runnable printBar) throws InterruptedException {       
        for (int i = 0; i < n; ) {
            if(!permitFoo) {
                printBar.run();
                i++;
                permitFoo = true;
            } else{
                Thread.yield();
            }
        }
    }
}

// ReentrantLock 可重入锁 + Condition
class FooBar {
    private int n;

    public FooBar(int n) {
        this.n = n;
    }

    Lock lock = new ReentrantLock(true);
    private final Condition foo = lock.newCondition();
    volatile boolean flag = true;

    public void foo(Runnable printFoo) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            lock.lock();
            try {
            	while(!flag) {
                    foo.await();
                }
                printFoo.run();
                flag = false;
                foo.signal();
            } finally {
            	lock.unlock();
            }
        }
    }

    public void bar(Runnable printBar) throws InterruptedException {
        for (int i = 0; i < n;i++) {
            lock.lock();
            try {
            	while(flag) {
                    foo.await();
            	}
                printBar.run();
                flag = true;
                foo.signal();
            }finally {
            	lock.unlock();
            }
        }
    }
}

//  synchronized + 标志位 + 唤醒
class FooBar {
    private int n;
    // 标志位，控制执行顺序，true执行printFoo，false执行printBar
    private volatile boolean type = true;
    private final Object foo=  new Object(); // 锁标志

    public FooBar(int n) {
        this.n = n;
    }
    public void foo(Runnable printFoo) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            synchronized (foo) {
                while(!type) {
                    foo.wait();
                }
                printFoo.run();
                type = false;
                foo.notifyAll();
            }
        }
    }

    public void bar(Runnable printBar) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            synchronized (foo) {
                while(type) {
                    foo.wait();
                }
                printBar.run();
                type = true;
                foo.notifyAll();
            }
        }
    }
}