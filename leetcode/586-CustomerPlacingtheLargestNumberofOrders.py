# 586. Customer Placing the Largest Number of Orders
# Table: Orders
# +-----------------+----------+
# | Column Name     | Type     |
# +-----------------+----------+
# | order_number    | int      |
# | customer_number | int      |
# +-----------------+----------+
# order_number is the primary key for this table.
# This table contains information about the order ID and the customer ID.

# Write an SQL query to find the customer_number for the customer who has placed the largest number of orders.
# The test cases are generated so that exactly one customer will have placed more orders than any other customer.
# The query result format is in the following example.

# Example 1:
# Input:
# Orders table:
# +--------------+-----------------+
# | order_number | customer_number |
# +--------------+-----------------+
# | 1            | 1               |
# | 2            | 2               |
# | 3            | 3               |
# | 4            | 3               |
# +--------------+-----------------+
# Output:
# +-----------------+
# | customer_number |
# +-----------------+
# | 3               |
# +-----------------+
# Explanation:
# The customer with number 3 has two orders, which is greater than either customer 1 or 2 because each of them only has one order.
# So the result is customer_number 3.

# Follow up: What if more than one customer has the largest number of orders, can you find all the customer_number in this case?

import pandas as pd

def largest_orders(orders: pd.DataFrame) -> pd.DataFrame:
    # 按 customer_number 分组统计数据
    orders = orders.groupby("customer_number").size().reset_index(name="cnt")
    # 取最大的 
    orders = orders[orders["cnt"] >= orders["cnt"].max()]
    return orders[["customer_number"]]

# count
def largest_orders1(orders: pd.DataFrame) -> pd.DataFrame:
    orders = orders.groupby(by='customer_number', as_index=False).count()
    filter = orders['order_number'] == orders['order_number'].max()
    return orders[['customer_number']][filter]

if __name__ == "__main__":
    data = [[1, 1], [2, 2], [3, 3], [4, 3]]
    orders = pd.DataFrame(data, columns=['order_number', 'customer_number']).astype({'order_number':'Int64', 'customer_number':'Int64'})
    print(largest_orders(orders))
    print(largest_orders1(orders))