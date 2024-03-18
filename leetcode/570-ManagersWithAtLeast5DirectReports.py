# 570. Managers with at Least 5 Direct Reports
# Table: Employee
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | id          | int     |
# | name        | varchar |
# | department  | varchar |
# | managerId   | int     |
# +-------------+---------+
# id is the primary key column for this table.
# Each row of this table indicates the name of an employee, their department, and the id of their manager.
# If managerId is null, then the employee does not have a manager.
# No employee will be the manager of themself.
#  
# Write an SQL query to report the managers with at least five direct reports.
# Return the result table in any order.
# The query result format is in the following example.

# Example 1:
# Input:
# Employee table:
# +-----+-------+------------+-----------+
# | id  | name  | department | managerId |
# +-----+-------+------------+-----------+
# | 101 | John  | A          | None      |
# | 102 | Dan   | A          | 101       |
# | 103 | James | A          | 101       |
# | 104 | Amy   | A          | 101       |
# | 105 | Anne  | A          | 101       |
# | 106 | Ron   | B          | 101       |
# +-----+-------+------------+-----------+
# Output:
# +------+
# | name |
# +------+
# | John |
# +------+

import pandas as pd

def find_managers(employee: pd.DataFrame) -> pd.DataFrame:
    #manager = employee.groupby('managerId').count().reset_index(name = "count")
    manager = employee.groupby('managerId').size().reset_index(name="count")
    filter = manager["count"] >= 5
    df = manager[filter][["managerId"]]
    print(df)
    print(df.index)
    return employee[employee["id"].isin(df.index)][["name"]]

# merge
def find_managers1(employee: pd.DataFrame) -> pd.DataFrame:
    manager = employee.groupby('managerId').size().reset_index(name='count')
    manager = manager[manager['count'] >= 5 ]
    df = manager.merge(employee, left_on='managerId', right_on='id', how='inner')
    return df[['name']]

def find_managers2(employee: pd.DataFrame) -> pd.DataFrame:
    employee['managerId'] = employee['managerId'].fillna(employee['id'])
    cnt = employee.groupby('managerId').size()
    inds = cnt[cnt >= 5].index
    return employee[employee['id'].isin(inds)][['name']]

if __name__ == "__main__":
    data = [[101, 'John', 'A', None], [102, 'Dan', 'A', 101], [103, 'James', 'A', 101], [104, 'Amy', 'A', 101], [105, 'Anne', 'A', 101], [106, 'Ron', 'B', 101]]
    employee = pd.DataFrame(data, columns=['id', 'name', 'department', 'managerId']).astype({'id':'Int64', 'name':'object', 'department':'object', 'managerId':'Int64'})
    
    print(find_managers(employee))
    print(find_managers1(employee))
    print(find_managers2(employee))