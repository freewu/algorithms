# 183. Customers Who Never Order
# Table: Customers
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | id          | int     |
# | name        | varchar |
# +-------------+---------+
# id is the primary key (column with unique values) for this table.
# Each row of this table indicates the ID and name of a customer.
 
# Table: Orders
# +-------------+------+
# | Column Name | Type |
# +-------------+------+
# | id          | int  |
# | customerId  | int  |
# +-------------+------+
# id is the primary key (column with unique values) for this table.
# customerId is a foreign key (reference columns) of the ID from the Customers table.
# Each row of this table indicates the ID of an order and the ID of the customer who ordered it.
 
# Write a solution to find all customers who never order anything.
# Return the result table in any order.
# The result format is in the following example.

# Example 1:
# Input: 
# Customers table:
# +----+-------+
# | id | name  |
# +----+-------+
# | 1  | Joe   |
# | 2  | Henry |
# | 3  | Sam   |
# | 4  | Max   |
# +----+-------+
# Orders table:
# +----+------------+
# | id | customerId |
# +----+------------+
# | 1  | 3          |
# | 2  | 1          |
# +----+------------+
# Output: 
# +-----------+
# | Customers |
# +-----------+
# | Henry     |
# | Max       |
# +-----------+

import pandas as pd

def find_customers(customers: pd.DataFrame, orders: pd.DataFrame) -> pd.DataFrame:
    # Select the customers whose 'id' is not present in the orders DataFrame's 'customerId' column.
    df = customers[~customers['id'].isin(orders['customerId'])] # not in    ~xxx.isin
    # Build a DataFrame that only contains the 'name' column and rename it as 'Customers'.
    df = df[['name']].rename(columns={'name': 'Customers'})
    return df

def find_customers1(customers: pd.DataFrame, orders: pd.DataFrame) -> pd.DataFrame:
    # Merge the customers DataFrame with the orders DataFrame using a left join on 'id' and 'customerId'
    merged_df = customers.merge(orders, how='left', left_on='id', right_on='customerId')
    # Use the 'customerId' column to create a boolean mask for customers who never placed any orders
    mask = merged_df['customerId'].isna()
    # Filter the rows using the boolean mask
    result_df = merged_df[mask]
    # Select only the 'name' column from the result DataFrame and rename it as 'Customers'
    result_df = result_df[['name']].rename(columns={'name': 'Customers'})
    return result_df

if __name__ == "__main__":
    data = [[1, 'Joe'], [2, 'Henry'], [3, 'Sam'], [4, 'Max']]
    customers = pd.DataFrame(data, columns=['id', 'name']).astype({'id':'Int64', 'name':'object'})
    data = [[1, 3], [2, 1]]
    orders = pd.DataFrame(data, columns=['id', 'customerId']).astype({'id':'Int64', 'customerId':'Int64'})
    print(find_customers(customers,orders))
    print(find_customers1(customers,orders))