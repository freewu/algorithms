# 2143 · Implementing a shopping cart program
# Description
# In this problem, we try to design a shopping cart system using the learned dictionary and the
# for loop and if else statements.
# We will give the information of all the items in the code box as follows.

# goods = [
#     {"name": "Computer", "price": 1999},
#     {"name": "Mouse", "price": 10},
#     {"name": "Yachts", "price": 20},
#     {"name": "Airplane", "price": 998}
# ]
# Please write your code in the code box.
# where asset in main.py represents the user's current asset (a positive integer) and
# input_list is the list of inputs that the user will enter in the next process.
# The input list is represented by a positive integer including 0, e.g. input_list = [0, 2, 4, 5, 1, 0], and the purchase is executed according to the following rules.

# if the input is 0 then the current purchase is ended.
# If the price of the product is less than or equal to the user's current asset, the output is
# "Purchase successful!" and the price of the product is deducted from the user's current asset and re-read into the input; if the price of the product exceeds the user's current asset, the output is output "The balance is low, so go ahead and top up!" and end the current purchase.
# 3. if the input is greater than 4, the user is prompted "Please re-enter" and the user input is re-read.

# Pay attention to the indentation of the code.

# Pay attention to code capitalization and punctuation in English and Chinese.

# Example
# The evaluator will execute your code by executing the command python main.py, and your code should output different results for different asset and input_list.

# When the input information is as follows.

# asset = 3165 
# input_list = [4,6,10,5,4,4,0] 
# We should have the following output.

# You want to buy: Airplane
# It is priced at: 998
# Purchase successful!
# Please re-enter
# Please re-enter
# Please re-enter
# You want to buy: Airplane
# It is priced at: 998
# Purchase successful!
# You want to buy: Airplane
# It is priced at: 998
# Purchase successful!
# When the input information is as follows.

# asset = 2436 
# input_list = [8,4,3,1,5,5,8,0]
# We should have the following output.

# Please re-enter
# You want to buy: Airplane
# It is priced at: 998
# Purchase successful!
# You want to buy: Yachts
# It is priced at: 20
# Purchase successful!
# You want to buy: Computer
# It is priced at: 1999
# The balance is low, so go ahead and top up!

import sys

# keep the code below
goods = [
    {"name": "Computer", "price": 1999},
    {"name": "Mouse", "price": 10},
    {"name": "Yachts", "price": 20},
    {"name": "Airplane", "price": 998}
]


asset = int(sys.argv[1])
input_list = eval(sys.argv[2])

# write your code here
for want in input_list:
    if want == 0:
        break
    if want > 4:
        print('Please re-enter')
        continue
    good_name, good_price = goods[want-1]['name'], goods[want-1]['price']
    print(f'You want to buy: {good_name}')
    print(f'It is priced at: {good_price}')
    if  good_price > asset:
        print('The balance is low, so go ahead and top up!')
        break
    else:
        print('Purchase successful!')
        asset -= good_price