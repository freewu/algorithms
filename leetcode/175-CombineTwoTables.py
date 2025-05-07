# 175. Combine Two Tables
# Table: Person
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | personId    | int     |
# | lastName    | varchar |
# | firstName   | varchar |
# +-------------+---------+
# personId is the primary key (column with unique values) for this table.
# This table contains information about the ID of some persons and their first and last names.

# Table: Address
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | addressId   | int     |
# | personId    | int     |
# | city        | varchar |
# | state       | varchar |
# +-------------+---------+
# addressId is the primary key (column with unique values) for this table.
# Each row of this table contains information about the city and state of one person with ID = PersonId.
 
# Write a solution to report the first name, last name, city, and state of each person in the Person table. If the address of a personId is not present in the Address table, report null instead.

# Return the result table in any order.

# The result format is in the following example.

# Example 1:
# Input: 
# Person table:
# +----------+----------+-----------+
# | personId | lastName | firstName |
# +----------+----------+-----------+
# | 1        | Wang     | Allen     |
# | 2        | Alice    | Bob       |
# +----------+----------+-----------+
# Address table:
# +-----------+----------+---------------+------------+
# | addressId | personId | city          | state      |
# +-----------+----------+---------------+------------+
# | 1         | 2        | New York City | New York   |
# | 2         | 3        | Leetcode      | California |
# +-----------+----------+---------------+------------+
# Output: 
# +-----------+----------+---------------+----------+
# | firstName | lastName | city          | state    |
# +-----------+----------+---------------+----------+
# | Allen     | Wang     | Null          | Null     |
# | Bob       | Alice    | New York City | New York |
# +-----------+----------+---------------+----------+
# Explanation: 
# There is no address in the address table for the personId = 1 so we return null in their city and state.
# addressId = 1 contains information about the address of personId = 2.

# data = [[1, 'Wang', 'Allen'], [2, 'Alice', 'Bob']]
# person = pd.DataFrame(data, columns=['personId', 'firstName', 'lastName']).astype({'personId':'Int64', 'firstName':'object', 'lastName':'object'})
# data = [[1, 2, 'New York City', 'New York'], [2, 3, 'Leetcode', 'California']]
# address = pd.DataFrame(data, columns=['addressId', 'personId', 'city', 'state']).astype({'addressId':'Int64', 'personId':'Int64', 'city':'object', 'state':'object'})

import pandas as pd

# merge
def combine_two_tables(person: pd.DataFrame, address: pd.DataFrame) -> pd.DataFrame:
    merged_df = person.merge(address, on='personId', how='left') # 通过 personId 关联两个 DataFrame
    return merged_df[['firstName', 'lastName', 'city', 'state']]

# join
def combine_two_tables1(person: pd.DataFrame, address: pd.DataFrame) -> pd.DataFrame:
    merged_df = person.join(
        address.set_index('personId'), 
        on='personId', # '对齐的列'
        how='left', # 'left'-默认'/inner'/'outer'/'right'
        lsuffix='_left', # 相同列出现，左表后缀（‘_left’）
        rsuffix='_right', # 相同列出现，右表后缀
    ) # 通过 personId 关联两个 DataFrame
    return merged_df[['firstName', 'lastName', 'city', 'state']]

if __name__ == "__main__":
    data = [[1, 'Wang', 'Allen'], [2, 'Alice', 'Bob']]
    person = pd.DataFrame(data, columns=['personId', 'firstName', 'lastName']).astype({'personId':'Int64', 'firstName':'object', 'lastName':'object'})
    data = [[1, 2, 'New York City', 'New York'], [2, 3, 'Leetcode', 'California']]
    address = pd.DataFrame(data, columns=['addressId', 'personId', 'city', 'state']).astype({'addressId':'Int64', 'personId':'Int64', 'city':'object', 'state':'object'}) 
    print(combine_two_tables(person, address))
    print(combine_two_tables1(person, address))