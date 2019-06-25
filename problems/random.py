#/usr/bin/env python
# -*- coding:utf-8 -*- 

import time

NEXT = 1L

# 线性同余法
# y＝(ax＋b)(mod  n)
# 其中n一般是一个很大的素数（几万）。a也是大素数，而且a，b，n都是常数。
# 所以rand的产生决定于x，他被称seed。每一个seed都是上一次产生的y的函数。
# 这样，如果直接取seed＝y的话，虽然产生的rand之间相关性甚小，但只要知道某个y，就能推知以后的rand。   
# 为避免这种情况，一般取seed为y和当时计算机的时间的函数，如seed＝y＋t系统里的随机数是利用初等数论中的同余定理来实现的
# 1 b,n 互质
# 2 n的所有质因子的积能整除 a-1
# 3 若m是4的倍数,a-1也是
# 4 a b y 都比 x 小
# 5 a,b都是正整数

def random():
    global NEXT
    # NEXT = (NEXT * 1103515245 + 12345)  % 32768
    # return int (NEXT / 65536)
    NEXT = (NEXT * 9301 + 49297) % 233280
    return NEXT / (233280.0)

def srand(seed):
    global NEXT
    NEXT = seed

# 把“线性同余”，“移位轮转”和“带记忆乘法”这3种基本的随机数发生法一起用
# unsigned int x = 123456789,
#                    y = 362436000,
#                    z = 521288629,
#                    c = 7654321; /* Seed variables */ 
 
# unsigned int KISS()
# {  
#     unsigned long long t, A = 698769069ULL;  
 
#     x = 69069*x+12345;  
#     y ^= (y<<13); 
#     y ^= (y>>17); 
#     y ^= (y<<5);  
    
#     t = (A*z + c);
#     c = (t >> 32);
#     z = t;
     
#     return x+y+z;  
# }

X = 123456789
Y = 362436000
Z = 521288629
C = 7654321



def random2():
    global X,Y,Z,C
    t = 698769069
    a = 698769069

    X = 69069 * X + 12345
    Y ^= (Y << 13)
    Y ^= (Y >> 17)
    Y ^= (Y << 5)

    t = (a * Z + C)
    C = (t >> 32)
    Z = t

    return X + Y + Z


if __name__ == "__main__":
    srand(time.time())
    # print(random())
    # print(random())
    # print(random())
    # 测试分布
    a = [0]*10
    start = time.time()
    len = 1000000
    for num in range(0,len):
        d = (int)((random() * 100 ) % 10)
        a[d] = a[d] + 1

    for key in range(0,10):
        print key," : ",(float(a[key]) / float(len))

    print "耗时: %f "  % ( time.time() - start)

    # print(random2())
    # print(random2())