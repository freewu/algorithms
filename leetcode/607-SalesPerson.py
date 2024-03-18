# 607. Sales Person
# Table: SalesPerson
# +-----------------+---------+
# | Column Name     | Type    |
# +-----------------+---------+
# | sales_id        | int     |
# | name            | varchar |
# | salary          | int     |
# | commission_rate | int     |
# | hire_date       | date    |
# +-----------------+---------+
# sales_id is the primary key column for this table.
# Each row of this table indicates the name and the ID of a salesperson alongside their salary, commission rate, and hire date.

# Table: Company
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | com_id      | int     |
# | name        | varchar |
# | city        | varchar |
# +-------------+---------+
# com_id is the primary key column for this table.
# Each row of this table indicates the name and the ID of a company and the city in which the company is located.

# Table: Orders
# +-------------+------+
# | Column Name | Type |
# +-------------+------+
# | order_id    | int  |
# | order_date  | date |
# | com_id      | int  |
# | sales_id    | int  |
# | amount      | int  |
# +-------------+------+
# order_id is the primary key column for this table.
# com_id is a foreign key to com_id from the Company table.
# sales_id is a foreign key to com_id from the SalesPerson table.
# Each row of this table contains information about one order. 
# This includes the ID of the company, the ID of the salesperson, the date of the order, and the amount paid.

# Write an SQL query to report the names of all the salespersons who did not have any orders related to the company with the name "RED".
# Return the result table in any order.
# The query result format is in the following example.

# Example 1:
# Input:
# SalesPerson table:
# +----------+------+--------+-----------------+------------+
# | sales_id | name | salary | commission_rate | hire_date  |
# +----------+------+--------+-----------------+------------+
# | 1        | John | 100000 | 6               | 4/1/2006   |
# | 2        | Amy  | 12000  | 5               | 5/1/2010   |
# | 3        | Mark | 65000  | 12              | 12/25/2008 |
# | 4        | Pam  | 25000  | 25              | 1/1/2005   |
# | 5        | Alex | 5000   | 10              | 2/3/2007   |
# +----------+------+--------+-----------------+------------+
# Company table:
# +--------+--------+----------+
# | com_id | name   | city     |
# +--------+--------+----------+
# | 1      | RED    | Boston   |
# | 2      | ORANGE | New York |
# | 3      | YELLOW | Boston   |
# | 4      | GREEN  | Austin   |
# +--------+--------+----------+
# Orders table:
# +----------+------------+--------+----------+--------+
# | order_id | order_date | com_id | sales_id | amount |
# +----------+------------+--------+----------+--------+
# | 1        | 1/1/2014   | 3      | 4        | 10000  |
# | 2        | 2/1/2014   | 4      | 5        | 5000   |
# | 3        | 3/1/2014   | 1      | 1        | 50000  |
# | 4        | 4/1/2014   | 1      | 4        | 25000  |
# +----------+------------+--------+----------+--------+
# Output:
# +------+
# | name |
# +------+
# | Amy  |
# | Mark |
# | Alex |
# +------+
# Explanation:
# According to orders 3 and 4 in the Orders table, it is easy to tell that only salesperson John and Pam have sales to company RED,
# so we report all the other names in the table salesperson.

import pandas as pd

def sales_person(sales_person: pd.DataFrame, company: pd.DataFrame, orders: pd.DataFrame) -> pd.DataFrame:
    # 取公司名为 RED 的数据
    company = company[company['name'] == 'RED'][['com_id']]
    # 取出 RED 公司下的所有 销售记录
    co = pd.merge(company,orders,how='left',on='com_id')[['sales_id']]
    # 取出不是 上面的 sales_person 名称  ~sales_person['sales_id'].isin(co['sales_id']
    return sales_person[~sales_person['sales_id'].isin(co['sales_id'])][['name']]

def sales_person1(sales_person: pd.DataFrame, company: pd.DataFrame, orders: pd.DataFrame) -> pd.DataFrame:
    df = orders[orders["com_id"].isin(company[company["name"] =="RED"]["com_id"])]["sales_id"]
    return sales_person[~sales_person["sales_id"].isin(df)][["name"]]
 

if __name__ == "__main__":
    data = [[1, 'John', 100000, 6, '4/1/2006'], [2, 'Amy', 12000, 5, '5/1/2010'], [3, 'Mark', 65000, 12, '12/25/2008'], [4, 'Pam', 25000, 25, '1/1/2005'], [5, 'Alex', 5000, 10, '2/3/2007']]
    sp = pd.DataFrame(data, columns=['sales_id', 'name', 'salary', 'commission_rate', 'hire_date']).astype({'sales_id':'Int64', 'name':'object', 'salary':'Int64', 'commission_rate':'Int64', 'hire_date':'datetime64[ns]'})
    data = [[1, 'RED', 'Boston'], [2, 'ORANGE', 'New York'], [3, 'YELLOW', 'Boston'], [4, 'GREEN', 'Austin']]
    company = pd.DataFrame(data, columns=['com_id', 'name', 'city']).astype({'com_id':'Int64', 'name':'object', 'city':'object'})
    data = [[1, '1/1/2014', 3, 4, 10000], [2, '2/1/2014', 4, 5, 5000], [3, '3/1/2014', 1, 1, 50000], [4, '4/1/2014', 1, 4, 25000]]
    orders = pd.DataFrame(data, columns=['order_id', 'order_date', 'com_id', 'sales_id', 'amount']).astype({'order_id':'Int64', 'order_date':'datetime64[ns]', 'com_id':'Int64', 'sales_id':'Int64', 'amount':'Int64'})

    print(sales_person(sp,company,orders))
    print(sales_person1(sp,company,orders))