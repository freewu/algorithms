# 2211 · Judge the number of palindrome
# Description
# In this question, we will provide a string str_1. 
# We have already written the is_palindrome function for you in solution.py. 
# The function str_1 represents the initial string. You need to determine whether the input string is It is a palindrome, 
# if it is, it returns True, if it is not, it returns False.

# Example
# The evaluation machine will execute your code by executing python main.py {input_path}, 
# and the test data will be placed in the file corresponding to input_path.
# You can see how the code works in main.py.

# Example 1
# When the input string is:
#   '123454321'
# The print result is:
#   True


# Example 2
# When the input string is:
#   'abcdcbad'
# The print result is:
#   False

def is_palindrome(str_1: str) -> bool:
    '''
    :param str_1: Input a string
    :return: Whether it is a palindrome number
    '''
    # -- write your code here --
    # while 循环
    low, high = 0, len(str_1) - 1
    while low < high:
        # 如果 不相等 说明不是回文
        if str_1[low] != str_1[high]:
            return False
        # 向里逼近
        low = low + 1 
        high = high - 1
    return True


def is_palindrome1(str_1: str) -> bool:
    '''
    :param str_1: Input a string
    :return: Whether it is a palindrome number
    '''
    m = len(str_1) # 取字符串长度
    for i in range(m // 2): # for 循环遍历，只需要遍历一半的长度即可
        if str_1[i] != str_1[m - i - 1]: # 与下标 i 对应的下标是 m - i - 1
            return False
    return True

# 
def is_palindrome2(str_1: str) -> bool:
    '''
    :param str_1: Input a string
    :return: Whether it is a palindrome number
    '''
    return str_1 == str_1[::-1]


if __name__ == "__main__":
    print(is_palindrome("'123454321'")) # True
    print(is_palindrome("'abcdcbad'")) # False
    print(is_palindrome1("'123454321'")) # True
    print(is_palindrome1("'abcdcbad'")) # False
    print(is_palindrome2("'123454321'")) # True
    print(is_palindrome2("'abcdcbad'")) # False