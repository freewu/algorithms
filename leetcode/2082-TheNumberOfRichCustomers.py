# 2082. The Number of Rich Customers
# Table: Store
# +-------------+------+
# | Column Name | Type |
# +-------------+------+
# | bill_id     | int  |
# | customer_id | int  |
# | amount      | int  |
# +-------------+------+
# bill_id is the primary key (column with unique values) for this table.
# Each row contains information about the amount of one bill and the customer associated with it.
# Write a solution to report the number of customers who had at least one bill with an amount strictly greater than 500.
# The result format is in the following example.

# Example 1:
# Input: 
# Store table:
# +---------+-------------+--------+
# | bill_id | customer_id | amount |
# +---------+-------------+--------+
# | 6       | 1           | 549    |
# | 8       | 1           | 834    |
# | 4       | 2           | 394    |
# | 11      | 3           | 657    |
# | 13      | 3           | 257    |
# +---------+-------------+--------+
# Output: 
# +------------+
# | rich_count |
# +------------+
# | 2          |
# +------------+
# Explanation: 
# Customer 1 has two bills with amounts strictly greater than 500.
# Customer 2 does not have any bills with an amount strictly greater than 500.
# Customer 3 has one bill with an amount strictly greater than 500.

import pandas as pd

def count_rich_customers(store: pd.DataFrame) -> pd.DataFrame:
    # least one bill with an amount strictly greater than 500.
    rich_customers = store[store['amount'] > 500]
    # 去重统计数量
    count = rich_customers['customer_id'].nunique()
    return  pd.DataFrame({'rich_count':[count]})


def count_rich_customers1(store: pd.DataFrame) -> pd.DataFrame:
    # 筛选出订单大于500
    df = store[store["amount"] > 500]
    # 去重
    df.drop_duplicates("customer_id",inplace= True)
    return pd.DataFrame({"rich_count":[len(df["customer_id"])]})

if __name__ == "__main__":
    data = [[6, 1, 549], [8, 1, 834], [4, 2, 394], [11, 3, 657], [13, 3, 257]]
    store = pd.DataFrame(data, columns=['bill_id', 'customer_id', 'amount']).astype({'bill_id':'int64', 'customer_id':'int64', 'amount':'int64'})
    print(count_rich_customers(store))
    print(count_rich_customers1(store))