#/usr/bin/env python
# -*- coding:utf-8 -*- 

# Write an algorithm which computes the number of trailing zeros in n factorial.
# Example
# Example 1:
# 	Input:  11
# 	Output: 2
	
# 	Explanation: 
# 	11! = 39916800, so the output should be 2

# Example 2:
# 	Input:  5
# 	Output: 1
	
# 	Explanation: 
# 	5! = 120, so the output should be 1.

# 1 : 1  0
# 2 : 2  0
# 3 : 6  0
# 4 : 24 0
# 5 : 120 1
# 6 : 720 1
# 7 : 5040 1
# 8 : 40320 1
# 9 : 362880 1
# 10 : 3628800 2
# 11 : 39916800 2
# 12 : 479001600 2
# 13 : 6227020800 2
# 14 : 87178291200 2
# 15 : 1307674368000 3
# 16 : 20922789888000 3
# 17 : 355687428096000 3
# 18 : 6402373705728000 3
# 19 : 121645100408832000 3
# 20 : 2432902008176640000 4
# 21 : 51090942171709440000 4
# 22 : 1124000727777607680000 4
# 23 : 25852016738884976640000 4
# 24 : 620448401733239439360000 4 
# 25 : 15511210043330985984000000 6
# 30 : 265252859812191058636308480000000 7
# 35 : 10333147966386144929666651337523200000000 8
# 40 : 815915283247897734345611269596115894272000000000 9
# 45 : 119622220865480194561963161495657715064383733760000000000 10
# 50 : 30414093201713378043612608166064768844377641568960512000000000000 12
# 74 : 330788544151938641225953028221253782145683251820934971170611926835411235700971565459250872320000000000000000 16
# 75 : 24809140811395398091946477116594033660926243886570122837795894512655842677572867409443815424000000000000000000 18
# 99 : 933262154439441526816992388562667004907159682643816214685929638952175999932299156089414639761565182862536979208272237582511852109168640000000000000000000000 22
# 100 : 93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000 24


def factorial(n):
    if n == 1 : return 1
    return n * factorial(n - 1)

class Solution:
    """
    @param: n: An integer
    @return: An integer, denote the number of trailing zeros in n!
    """
    # 计算5的个数 25以后就有问题了
    def trailingZeros1(self, n):
        # write your code here, try to do it without arithmetic operators.
        r = 0
        # 计算5的个数
        while (n - 5) >= 0:
            n = n - 5
            r = r + 1
            # 被5整除就加1
            if 0 == r % 5:
                r = r + 1

        return r

    # 分别算出每个数的因数2的个数之和和因数5的个数之和，取其小值。
    # 10 = 2*5 5的倍数
    def trailingZeros(self, n):
        count = 0
        i = 5
        # 如果大于5
        while n/i>=1:
            count += n/i
            i = i * 5 # 5 25 50 75 100
        return count

if __name__ == "__main__":
    # for i in range(1,121):
    #     print "%d : %d" % (i,factorial(i))
    s = Solution()
    print s.trailingZeros(11)
    print s.trailingZeros(20)
    print s.trailingZeros(4)
    print s.trailingZeros(5)
    print s.trailingZeros(15)
    print s.trailingZeros(16)
    print s.trailingZeros(105) # 25
    print s.trailingZeros(5)
    print s.trailingZeros(25)
    print s.trailingZeros(26)
    print s.trailingZeros(74)