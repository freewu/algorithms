# 2258 · Find the largest prime number backwards
# Description
# The main purpose of this question is that we will provide a natural number greater than 1,
# and then determine whether the number is a prime number, if it is a prime number, 
# then print it directly, if it is not a prime number, look for the previous digit, then the first prime number found is the largest Prime number.
# Please write the relevant Python code in solution.py and print out the largest prime number it generates.s.

# Example
# The evaluation opportunity executes your code by executing the command python main.py {n}, 
# and passing in n as a command line parameter, you can learn how the code runs in main.py.

# Example 1
# When the input data is:
#   5
# The output data is:
#   5

# Example 2
# When the input data is:
#   14
# The output data is:
#   13

def prime(num: int) -> int:
    """
    :param num: a random integer
    :return: determine if the result is a prime number and return 1 otherwise other values
    """
	# write your code here
    if num == 0 or num == 1:
        return 0
    else:
        for i in range(2, int(num / 2)): # 取一半数据就行了
            if num % i == 0:
                return 0
        return 1

def Find_prime_numbers_backwards(n:int) -> int:
    '''
    :param n: Natural number greater than 1
    :return: Maximum prime number
    '''
    # write your code here
    while(n > 0):
        if prime(n): return n
        n = n - 1
    return -1

if __name__ == "__main__":
    print(Find_prime_numbers_backwards(5)) # 5
    print(Find_prime_numbers_backwards(14)) # 13