# 1173. Immediate Food Delivery I
# Table: Delivery
# +-----------------------------+---------+
# | Column Name                 | Type    |
# +-----------------------------+---------+
# | delivery_id                 | int     |
# | customer_id                 | int     |
# | order_date                  | date    |
# | customer_pref_delivery_date | date    |
# +-----------------------------+---------+
# delivery_id is the primary key (column with unique values) of this table.
# The table holds information about food delivery to customers that make orders at some date and specify a preferred delivery date (on the same order date or after it).
# If the customer's preferred delivery date is the same as the order date, then the order is called immediate; otherwise, it is called scheduled.
# Write a solution to find the percentage of immediate orders in the table, rounded to 2 decimal places.
# The result format is in the following example.

# Example 1:
# Input: 
# Delivery table:
# +-------------+-------------+------------+-----------------------------+
# | delivery_id | customer_id | order_date | customer_pref_delivery_date |
# +-------------+-------------+------------+-----------------------------+
# | 1           | 1           | 2019-08-01 | 2019-08-02                  |
# | 2           | 5           | 2019-08-02 | 2019-08-02                  |
# | 3           | 1           | 2019-08-11 | 2019-08-11                  |
# | 4           | 3           | 2019-08-24 | 2019-08-26                  |
# | 5           | 4           | 2019-08-21 | 2019-08-22                  |
# | 6           | 2           | 2019-08-11 | 2019-08-13                  |
# +-------------+-------------+------------+-----------------------------+
# Output: 
# +----------------------+
# | immediate_percentage |
# +----------------------+
# | 33.33                |
# +----------------------+
# Explanation: The orders with delivery id 2 and 3 are immediate while the others are scheduled.

import pandas as pd

def food_delivery(delivery: pd.DataFrame) -> pd.DataFrame:
    # 筛选出订单日期与顾客首选配送日期相同的订单(即时订单)
    immediate_delivery = delivery[delivery['order_date'] == delivery['customer_pref_delivery_date']]
    # 计算总订单数量和即时订单数量
    total_deliveries = len(delivery)
    immediate_deliveries = len(immediate_delivery)
    # 计算即时配送的百分比
    immediate_percentage = round((immediate_deliveries / total_deliveries) * 100, 2)
    # 将结果转化为 DataFrame 格式
    return  pd.DataFrame({'immediate_percentage': [immediate_percentage]})

# mean
def food_delivery1(delivery: pd.DataFrame) -> pd.DataFrame:
    rate = round((delivery.customer_pref_delivery_date == delivery.order_date).mean() * 100, 2)
    return pd.DataFrame({"immediate_percentage": [rate ]})


if __name__ == "__main__":
    data = [[1, 1, '2019-08-01', '2019-08-02'], [2, 5, '2019-08-02', '2019-08-02'], [3, 1, '2019-08-11', '2019-08-11'], [4, 3, '2019-08-24', '2019-08-26'], [5, 4, '2019-08-21', '2019-08-22'], [6, 2, '2019-08-11', '2019-08-13']]
    delivery = pd.DataFrame(data, columns=['delivery_id', 'customer_id', 'order_date', 'customer_pref_delivery_date']).astype({'delivery_id':'Int64', 'customer_id':'Int64', 'order_date':'datetime64[ns]', 'customer_pref_delivery_date':'datetime64[ns]'})
    print(food_delivery(delivery))

    print(food_delivery1(delivery))
