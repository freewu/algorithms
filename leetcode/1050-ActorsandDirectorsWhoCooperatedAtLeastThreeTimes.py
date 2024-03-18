# 1050. Actors and Directors Who Cooperated At Least Three Times
# Table: ActorDirector
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | actor_id    | int     |
# | director_id | int     |
# | timestamp   | int     |
# +-------------+---------+
# timestamp is the primary key column for this table.
# Write a SQL query for a report that provides the pairs (actor_id, director_id) where the actor has cooperated with the director at least three times.
# Return the result table in any order.
# The query result format is in the following example.

# Example 1:
# Input:
# ActorDirector table:
# +-------------+-------------+-------------+
# | actor_id    | director_id | timestamp   |
# +-------------+-------------+-------------+
# | 1           | 1           | 0           |
# | 1           | 1           | 1           |
# | 1           | 1           | 2           |
# | 1           | 2           | 3           |
# | 1           | 2           | 4           |
# | 2           | 1           | 5           |
# | 2           | 1           | 6           |
# +-------------+-------------+-------------+
# Output:
# +-------------+-------------+
# | actor_id    | director_id |
# +-------------+-------------+
# | 1           | 1           |
# +-------------+-------------+
# Explanation: The only pair is (1, 1) where they cooperated exactly 3 times.

import pandas as pd


# size
def actors_and_directors(actor_director: pd.DataFrame) -> pd.DataFrame:
    # 按 director_id & actor_id 分组统计数量
    actor_director = actor_director.groupby(["director_id","actor_id"])["timestamp"].size().reset_index(name="count")
    # 过滤出 >= 3 的数据
    filter = actor_director["count"] >= 3
    return actor_director[filter][["actor_id","director_id"]]

# count
def actors_and_directors1(actor_director: pd.DataFrame) -> pd.DataFrame:
    df = actor_director.groupby(['actor_id','director_id'], as_index = False).count()
    # 过滤出 >= 3 的数据
    filter = df["timestamp"] >= 3
    return df[filter][['actor_id','director_id']]


if __name__ == "__main__":
    data = [[1, 1, 0], [1, 1, 1], [1, 1, 2], [1, 2, 3], [1, 2, 4], [2, 1, 5], [2, 1, 6]]
    actor_director = pd.DataFrame(data, columns=['actor_id', 'director_id', 'timestamp']).astype({'actor_id':'int64', 'director_id':'int64', 'timestamp':'int64'})
    print(actors_and_directors(actor_director))
    print(actors_and_directors1(actor_director))