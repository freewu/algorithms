package leetcode;

// 1116. Print Zero Even Odd
// You have a function printNumber that can be called with an integer parameter and prints it to the console.

// For example, calling printNumber(7) prints 7 to the console.
// You are given an instance of the class ZeroEvenOdd that has three functions: zero, even, and odd. The same instance of ZeroEvenOdd will be passed to three different threads:

//      Thread A: calls zero() that should only output 0's.
//      Thread B: calls even() that should only output even numbers.
//      Thread C: calls odd() that should only output odd numbers.
//      Modify the given class to output the series "010203040506..." where the length of the series must be 2n.

// Implement the ZeroEvenOdd class:

//      ZeroEvenOdd(int n) Initializes the object with the number n that represents the numbers that should be printed.
//      void zero(printNumber) Calls printNumber to output one zero.
//      void even(printNumber) Calls printNumber to output one even number.
//      void odd(printNumber) Calls printNumber to output one odd number.

// Example 1:
//      Input: n = 2
//      Output: "0102"
//      Explanation: There are three threads being fired asynchronously.
//      One of them calls zero(), the other calls even(), and the last one calls odd().
//      "0102" is the correct output.

// Example 2:
//      Input: n = 5
//      Output: "0102030405"

// Constraints:

//      1 <= n <= 1000

// use synchronized
class ZeroEvenOdd {
    private int n;
    private volatile int flag = 1;
    
    public ZeroEvenOdd(int n) {
        this.n = n;
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i++) {
            synchronized (this) {
                while (flag % 2 == 0) wait();
                printNumber.accept(0);
                flag ++;
                notifyAll();
            }
        }
    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        for (int i=2; i<=n; i+=2) {
            synchronized (this) {
                while (flag % 4 != 0) wait();
                printNumber.accept(i);
                flag ++;
                notifyAll();
            }
        }
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i+=2) {
            synchronized(this) {
                while (flag % 4 != 2) wait();
                printNumber.accept(i);
                flag ++;
                notifyAll();
            }
        }
    }
}

// use LockSupport 

import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.locks.LockSupport;

class ZeroEvenOdd {
    private int n;
    volatile int flag = 1;
    ConcurrentHashMap<Integer, Thread> map = new ConcurrentHashMap<>();

    public ZeroEvenOdd(int n) {
        this.n = n;
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        map.putIfAbsent(1, Thread.currentThread());

        for (int i=1; i<=n; i++) {
            while (flag % 2 == 0) LockSupport.park();
            printNumber.accept(0);
            flag ++;
            if (flag % 4 == 0) LockSupport.unpark(map.get(0));
            if (flag % 4 == 2) LockSupport.unpark(map.get(1));
        }
    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        map.putIfAbsent(0, Thread.currentThread());
        for (int i=2; i<=n; i+=2) {
            while (flag % 4 != 0) LockSupport.park();
            printNumber.accept(i);
            flag ++;
            LockSupport.unpark(map.get(1));
        }
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i+=2) {
            while (flag % 4 != 2) LockSupport.park();
            printNumber.accept(i);
            flag ++;
            LockSupport.unpark(map.get(1));
        }
    }
}

//  yield + 原子类
import java.util.concurrent.atomic.AtomicInteger;

class ZeroEvenOdd {
    private int n;
    volatile AtomicInteger dice = new AtomicInteger(1);


    public ZeroEvenOdd(int n) {
        this.n = n;
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i++) {
            while (dice.get() % 2 == 0) Thread.yield();
            printNumber.accept(0);
            dice.getAndIncrement();
        }
    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        for (int i=2; i<=n; i+=2) {
            while (dice.get() % 4 != 0) Thread.yield();
            printNumber.accept(i);
            dice.getAndIncrement();
        }
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i+=2) {
            while (dice.get() % 4 != 2) Thread.yield();
            printNumber.accept(i);
            dice.getAndIncrement();
        }
    }
}

// ReentrantLock + Condition 
class ZeroEvenOdd {
    private int n;
    private volatile int flag = 1;
    private ReentrantLock lock = new ReentrantLock();
    private Condition condition = lock.newCondition();
    
    public ZeroEvenOdd(int n) {
        this.n = n;
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i++) {
            lock.lock();
            printNumber.accept(0);
            flag ++;
            condition.signalAll();
            while (flag %2 == 0) condition.await();
            lock.unlock();
        }
    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        for (int i=2; i<=n; i+=2) {
            lock.lock();
            while (flag % 4 != 0) condition.await();
            printNumber.accept(i);
            flag ++;
            condition.signalAll();
            lock.unlock();
        }
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i+=2) {
            lock.lock();
            while (flag % 4 != 2) condition.await();
            printNumber.accept(i);
            flag ++;
            condition.signalAll();
            lock.unlock();
        }       
    }
}

 Semaphore 的解法

import java.util.concurrent.Semaphore;

class ZeroEvenOdd {
    private int n;
    Semaphore[] semaphores = new Semaphore[3];

    public ZeroEvenOdd(int n) {
        this.n = n;
        semaphores[0] = new Semaphore(1);
        semaphores[1] = new Semaphore(0);
        semaphores[2] = new Semaphore(0);
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i++) {
            semaphores[0].acquire();
            printNumber.accept(0);
            if (i % 2 == 0) semaphores[1].release();
            else semaphores[2].release();
        }
    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        for (int i=2; i<=n; i=i+2) {
            semaphores[1].acquire();
            printNumber.accept(i);
            semaphores[0].release();
        }
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i=i+2) {
            semaphores[2].acquire();
            printNumber.accept(i);
            semaphores[0].release();
        }
    }
}

//  SynchronousQueue 的解法
import java.util.concurrent.SynchronousQueue;

class ZeroEvenOdd {
    private int n;
    private SynchronousQueue<Integer>[] queues = new SynchronousQueue[3];

    public ZeroEvenOdd(int n) {
        this.n = n;
        for (int i=0; i<queues.length; i++) queues[i] = new SynchronousQueue<>();
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i++) {
            printNumber.accept(0);
            if (i % 2 == 1) {
                queues[1].put(1);
            } else {
                queues[2].put(1);
            }
            queues[0].take();
        }
    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        for (int i=2; i<=n; i+=2) {
            queues[2].take();
            printNumber.accept(i);
            queues[0].put(1);
        }
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i+=2) {
            queues[1].take();
            printNumber.accept(i);
            queues[0].put(1);
        }
    }
}

// use LinkedBlockingQueue 
import java.util.concurrent.LinkedBlockingQueue;

class ZeroEvenOdd {
    private int n;
    private LinkedBlockingQueue<Integer>[] queues = new LinkedBlockingQueue[3];

    public ZeroEvenOdd(int n) {
        this.n = n;
        for (int i=0; i<queues.length; i++) {
            queues[i] = new LinkedBlockingQueue<>(1);
        }
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        queues[0].put(1);
        for (int i=1; i<=n; i++) {
            queues[0].take();
            printNumber.accept(0);
            if (i % 2 == 0) {
                queues[2].put(1);
            } else {
                queues[1].put(1);
            }
        }
    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        for (int i=2; i<=n; i+=2) {
            queues[2].take();
            printNumber.accept(i);
            if (i < n) queues[0].put(1);
        }
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i+=2) {
            queues[1].take();
            printNumber.accept(i);
            if (i < n) queues[0].put(1);
        }
    }
}


// use CyclicBarrier 
import java.util.concurrent.BrokenBarrierException;
import java.util.concurrent.CyclicBarrier;

class ZeroEvenOdd {
    private int n;
    CyclicBarrier[] cbs = new CyclicBarrier[3];

    public ZeroEvenOdd(int n) {
        this.n = n;
        for (int i=0; i<cbs.length; i++) cbs[i] = new CyclicBarrier(2);
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i++) {
            printNumber.accept(0);
            try {
                if (i % 2 == 1) {
                    cbs[1].await();
                } else {
                    cbs[2].await();
                }
                cbs[0].await();
            } catch (BrokenBarrierException e) {
                e.printStackTrace();
            }
        }
    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        for (int i=2; i<=n; i+=2) {
            try {
                cbs[2].await();
                printNumber.accept(i);
                cbs[0].await();
            } catch (BrokenBarrierException e) {
                e.printStackTrace();
            }
        }
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i+=2) {
            try {
                cbs[1].await();
                printNumber.accept(i);
                cbs[0].await();
            } catch (BrokenBarrierException e) {
                e.printStackTrace();
            }
        }
    }
}

// use Exchanger 
import java.util.concurrent.Exchanger;

class ZeroEvenOdd {
    private int n;
    Exchanger<Integer>[] exchanges = new Exchanger[3];

    public ZeroEvenOdd(int n) {
        this.n = n;
        for (int i=0; i< exchanges.length; i++) exchanges[i] = new Exchanger<>();
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i++) {
            printNumber.accept(0);
            if (i % 2 == 1) {
                exchanges[1].exchange(1);
            } else {
                exchanges[2].exchange(1);
            }
            exchanges[0].exchange(1);
        }
    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        for (int i=2; i<=n; i+=2) {
            exchanges[2].exchange(1);
            printNumber.accept(i);
            exchanges[0].exchange(1);
        }
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i+=2) {
            exchanges[1].exchange(1);
            printNumber.accept(i);
            exchanges[0].exchange(1);
        }
    }
}