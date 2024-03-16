# 196. Delete Duplicate Emails
# Table: Person
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | id          | int     |
# | email       | varchar |
# +-------------+---------+
# id is the primary key column for this table.
# Each row of this table contains an email. The emails will not contain uppercase letters.
# Write an SQL query to delete all the duplicate emails, keeping only one unique email with the smallest id.
# Note that you are supposed to write a DELETE statement and not a SELECT one.
# After running your script, the answer shown is the Person table.
# The driver will first compile and run your piece of code and then show the Person table.
# The final order of the Person table does not matter.

# The query result format is in the following example.

# Example 1:
# Input:
# Person table:
# +----+------------------+
# | id | email            |
# +----+------------------+
# | 1  | john@example.com |
# | 2  | bob@example.com  |
# | 3  | john@example.com |
# +----+------------------+
# Output:
# +----+------------------+
# | id | email            |
# +----+------------------+
# | 1  | john@example.com |
# | 2  | bob@example.com  |
# +----+------------------+
# Explanation: john@example.com is repeated two times. We keep the row with the smallest Id = 1.

import pandas as pd

def delete_duplicate_emails(person: pd.DataFrame) -> None:
    # person[["email"]].drop_duplicates(inplace=True)
    # 先排序 去重需要保留小的 id
    person.sort_values(by = 'id', ascending = True, inplace = True)
    person.drop_duplicates(subset='email', inplace=True)

def delete_duplicate_emails1(person: pd.DataFrame) -> None:
    person.sort_values(by='id',inplace=True)
    person.drop_duplicates(subset='email',keep='first',inplace=True)

if __name__ == "__main__":
    data = [[1, 'john@example.com'], [2, 'bob@example.com'], [3, 'john@example.com']]
    person = pd.DataFrame(data, columns=['id', 'email']).astype({'id':'int64', 'email':'object'})
    delete_duplicate_emails(person)
    print(person)

# drop_duplicates() 函数的具体参数
#
# 用法：
#
#     DataFrame.drop_duplicates(subset=None, keep=‘first’, inplace=False)
#
# 参数说明
#
#     参数	说明
#     subset	根据指定的列名进行去重，默认整个数据集
#     keep	可选{‘first’, ‘last’, False}，默认first，即默认保留第一次出现的重复值，并删去其他重复的数据，False是指删去所有重复数据。
#     inplace	是否对数据集本身进行修改，默认False

# sort_values()函数的具体参数
#
# 用法：
#
#     DataFrame.sort_values(by=‘##’,axis=0,ascending=True, inplace=False, na_position=‘last’)
#
# 参数说明
#
#     参数	    说明
#     by	        指定列名(axis=0或’index’)或索引值(axis=1或’columns’)
#     axis	    若axis=0或’index’，则按照指定列中数据大小排序；若axis=1或’columns’，则按照指定索引中数据大小排序，默认axis=0
#     ascending	是否按指定列的数组升序排列，默认为True，即升序排列
#     inplace	    是否用排序后的数据集替换原来的数据，默认为False，即不替换
#     na_position	{‘first’,‘last’}，设定缺失值的显示位置