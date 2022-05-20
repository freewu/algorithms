# 2142 · Replace the elements in a string
# Description
# String is a common data type in Python, and when we want to convert it to a list, we can use the character in the split() as the delimiter.

# We can iterate through the elements of any sequence with a for loop, 
# and use if to determine execute what kind of operations when the condition is met

# for i in range(len(var_list)): 
#     if condition: 
#         do sth 
# There is a string var_str as "Zhuge_Dan#Susan_Molina#Zhuge_Dan#Jennifer_Lee#Jerry_Davis#", 
# please parse out each element in var_str and replace the element "Zhuge_Dan" with the element "Zhuge Liang".(Replace underscores with spaces)
# Please write the relevant python code in the solution.py 
# file to implement the above functionality, and we will import your module in main.py to check if your code implements the above functionality. 
# Your code needs to output two lists var_list, the first with the parsed list and the second with the replaced elements

# Some code is provided, do not modify uncommented code
# Learn to use hints in comments
# Example
# We will pass in the string var_str by calling the solution function

# When var_str = 'Zhuge_Dan#Susan_Molina#Zhuge_Dan#Jennifer_Lee#Jerry_Davis#', your code should output

# ['Zhuge Dan', 'Susan Molina', 'Zhuge Dan', 'Jennifer Lee', 'Jerry Davis']
# ['Zhuge Liang', 'Susan Molina', 'Zhuge Liang', 'Jennifer Lee', 'Jerry Davis']

def solution(var_str):
    print([i for i in var_str.replace('_',' ').split('#') if i != ''])
    print([i for i in var_str.replace('Zhuge_Dan','Zhuge Liang').replace('_',' ').split('#') if i != ''])