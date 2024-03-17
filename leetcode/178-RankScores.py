# 178. Rank Scores
# Table: Scores
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | id          | int     |
# | score       | decimal |
# +-------------+---------+
# id is the primary key for this table.
# Each row of this table contains the score of a game. Score is a floating point value with two decimal places.
#  
# Write an SQL query to rank the scores. The ranking should be calculated according to the following rules:

# The scores should be ranked from the highest to the lowest.
# If there is a tie between two scores, both should have the same ranking.
# After a tie, the next ranking number should be the next consecutive integer value. In other words, there should be no holes between ranks.
# Return the result table ordered by score in descending order.

# The query result format is in the following example.

# Example 1:
# Input:
# Scores table:
# +----+-------+
# | id | score |
# +----+-------+
# | 1  | 3.50  |
# | 2  | 3.65  |
# | 3  | 4.00  |
# | 4  | 3.85  |
# | 5  | 4.00  |
# | 6  | 3.65  |
# +----+-------+
# Output:
# +-------+------+
# | score | rank |
# +-------+------+
# | 4.00  | 1    |
# | 4.00  | 1    |
# | 3.85  | 2    |
# | 3.65  | 3    |
# | 3.65  | 3    |
# | 3.50  | 4    |
# +-------+------+

import pandas as pd

def order_scores(scores: pd.DataFrame) -> pd.DataFrame:
    # 按 score 排序
    scores = scores.sort_values(by = 'score', ascending = False)
    # 给数据排 rank 
    scores['rank'] = scores['score'].rank(method = 'dense', ascending = False)
    # 将 rank 设置为 int型
    scores = scores.astype({'rank':'Int64'})
    # 输出 score & rank
    return scores[['score','rank']]


def order_scores1(scores: pd.DataFrame) -> pd.DataFrame:
    # 按 score DESC rank 排名 并取 ['score','rank']
    scores = scores.assign(rank = scores['score'].rank(ascending=False, method='dense'))[['score','rank']]
    # 按 rank 排序输出
    return scores.sort_values(by='rank')
                

if __name__ == "__main__":
    data = [[1, 3.5], [2, 3.65], [3, 4.0], [4, 3.85], [5, 4.0], [6, 3.65]]
    scores = pd.DataFrame(data, columns=['id', 'score']).astype({'id':'Int64', 'score':'Float64'})
    print(order_scores(scores))
    print(order_scores1(scores))
    