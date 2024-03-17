# 1907. Count Salary Categories
# Table: Accounts
# +-------------+------+
# | Column Name | Type |
# +-------------+------+
# | account_id  | int  |
# | income      | int  |
# +-------------+------+
# account_id is the primary key (column with unique values) for this table.
# Each row contains information about the monthly income for one bank account.
 
# Write a solution to calculate the number of bank accounts for each salary category. The salary categories are:
#      "Low Salary": All the salaries strictly less than $20000.
#      "Average Salary": All the salaries in the inclusive range [$20000, $50000].
#      "High Salary": All the salaries strictly greater than $50000.

# The result table must contain all three categories. If there are no accounts in a category, return 0.
# Return the result table in any order.
# The result format is in the following example.

# Example 1:
# Input: 
# Accounts table:
# +------------+--------+
# | account_id | income |
# +------------+--------+
# | 3          | 108939 |
# | 2          | 12747  |
# | 8          | 87709  |
# | 6          | 91796  |
# +------------+--------+
# Output: 
# +----------------+----------------+
# | category       | accounts_count |
# +----------------+----------------+
# | Low Salary     | 1              |
# | Average Salary | 0              |
# | High Salary    | 3              |
# +----------------+----------------+
# Explanation: 
# Low Salary: Account 2.
# Average Salary: No accounts.
# High Salary: Accounts 3, 6, and 8.

import pandas as pd

# len
def count_salary_categories(accounts: pd.DataFrame) -> pd.DataFrame:
    df = [
        ['Low Salary',len(accounts[accounts['income'] < 20000])], # < $20000
        ['Average Salary',len(accounts[(accounts['income'] <= 50000) & (accounts['income'] >= 20000)])], # [$20000, $50000]
        ['High Salary', len(accounts[accounts['income'] > 50000])] # > 50000
    ]
    return pd.DataFrame(df, columns=['category','accounts_count'])

# shape
def count_salary_categories1(accounts: pd.DataFrame) -> pd.DataFrame:
    df = [
        ['Low Salary', accounts[accounts['income'] < 20000].shape[0]], # < $20000
        ['Average Salary', sum((accounts['income'] >= 20000) & (accounts['income'] <= 50000))], # [$20000, $50000]
        ['High Salary', accounts[accounts['income'] > 50000].shape[0]] # > 50000
    ]
    return pd.DataFrame(df, columns=['category','accounts_count'])

# sum
def count_salary_categories2(accounts: pd.DataFrame) -> pd.DataFrame:
    df = [
        ['Low Salary', (accounts['income'] < 20000).sum()], # < $20000
        ['Average Salary', ((20000 <= accounts['income']) & (accounts['income'] <= 50000)).sum()], # [$20000, $50000]
        ['High Salary', (accounts['income'] > 50000).sum()] # > 50000
    ]
    return pd.DataFrame(df, columns=['category','accounts_count'])


if __name__ == "__main__":
    data = [[3, 108939], [2, 12747], [8, 87709], [6, 91796]]
    accounts = pd.DataFrame(data, columns=['account_id', 'income']).astype({'account_id':'Int64', 'income':'Int64'})
    print(count_salary_categories(accounts))
    print(count_salary_categories1(accounts))
    print(count_salary_categories2(accounts))
