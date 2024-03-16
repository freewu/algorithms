# 1667. Fix Names in a Table
# Table: Users
# +----------------+---------+
# | Column Name    | Type    |
# +----------------+---------+
# | user_id        | int     |
# | name           | varchar |
# +----------------+---------+
# user_id is the primary key for this table.
# This table contains the ID and the name of the user. The name consists of only lowercase and uppercase characters.

# Write an SQL query to fix the names so that only the first character is uppercase and the rest are lowercase.
# Return the result table ordered by user_id.
# The query result format is in the following example.

# Example 1:
# Input:
# Users table:
# +---------+-------+
# | user_id | name  |
# +---------+-------+
# | 1       | aLice |
# | 2       | bOB   |
# +---------+-------+
# Output:
# +---------+-------+
# | user_id | name  |
# +---------+-------+
# | 1       | Alice |
# | 2       | Bob   |
# +---------+-------+

import pandas as pd

def fix_names(users: pd.DataFrame) -> pd.DataFrame:
    users['name'] = users.apply(
        # fix the names so that only the first character is uppercase and the rest are lowercase
        lambda x: x['name'].capitalize(),
        axis = 1
    )
    # Return the result table ordered by user_id.
    return users.sort_values('user_id')

def fix_names1(users: pd.DataFrame) -> pd.DataFrame:
    users['name'] = users['name'].str.capitalize()
    return users.sort_values(by='user_id')

if __name__ == "__main__":
    data = [[1, 'aLice'], [2, 'bOB']]
    users = pd.DataFrame(data, columns=['user_id', 'name']).astype({'user_id':'Int64', 'name':'object'})
    
    print(fix_names(users))
    print(fix_names1(users))
