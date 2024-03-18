# 511. Game Play Analysis I
# Table: Activity
# +--------------+---------+
# | Column Name  | Type    |
# +--------------+---------+
# | player_id    | int     |
# | device_id    | int     |
# | event_date   | date    |
# | games_played | int     |
# +--------------+---------+
# (player_id, event_date) is the primary key of this table.
# This table shows the activity of players of some games.
# Each row is a record of a player who logged in and played a number of games (possibly 0) before logging out on someday using some device.

# Write an SQL query to report the first login date for each player.
# Return the result table in any order.
# The query result format is in the following example.

# Example 1:
# Input:
# Activity table:
# +-----------+-----------+------------+--------------+
# | player_id | device_id | event_date | games_played |
# +-----------+-----------+------------+--------------+
# | 1         | 2         | 2016-03-01 | 5            |
# | 1         | 2         | 2016-05-02 | 6            |
# | 2         | 3         | 2017-06-25 | 1            |
# | 3         | 1         | 2016-03-02 | 0            |
# | 3         | 4         | 2018-07-03 | 5            |
# +-----------+-----------+------------+--------------+
# Output:
# +-----------+-------------+
# | player_id | first_login |
# +-----------+-------------+
# | 1         | 2016-03-01  |
# | 2         | 2017-06-25  |
# | 3         | 2016-03-02  |
# +-----------+-------------+

import pandas as pd

def game_analysis(activity: pd.DataFrame) -> pd.DataFrame:
    # 按用户分类取最小的 event_date
    activity = activity.groupby(["player_id"])[['event_date']].min().reset_index()
    # 按要求改列名
    activity = activity.rename(columns= {"event_date": "first_login"})
    return activity.sort_values('player_id')



if __name__ == "__main__":
    data = [[1, 2, '2016-03-01', 5], [1, 2, '2016-05-02', 6], [2, 3, '2017-06-25', 1], [3, 1, '2016-03-02', 0], [3, 4, '2018-07-03', 5]]
    activity = pd.DataFrame(data, columns=['player_id', 'device_id', 'event_date', 'games_played']).astype({'player_id':'Int64', 'device_id':'Int64', 'event_date':'datetime64[ns]', 'games_played':'Int64'})
    print(game_analysis(activity))