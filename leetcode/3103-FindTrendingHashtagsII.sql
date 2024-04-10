-- 3103. Find Trending Hashtags II 
-- Table: Tweets
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | user_id     | int     |
-- | tweet_id    | int     |
-- | tweet_date  | date    |
-- | tweet       | varchar |
-- +-------------+---------+
-- tweet_id is the primary key (column with unique values) for this table.
-- Each row of this table contains user_id, tweet_id, tweet_date and tweet.
-- Write a solution to find the top 3 trending hashtags in February 2024. Every tweet may contain several hashtags.

-- Return the result table orderd by count of hashtag, hashtag in descending order.
-- The result format is in the following example.

-- Example 1:
-- Input:
-- Tweets table:
-- +---------+----------+------------------------------------------------------------+------------+
-- | user_id | tweet_id | tweet                                                      | tweet_date |
-- +---------+----------+------------------------------------------------------------+------------+
-- | 135     | 13       | Enjoying a great start to the day. #HappyDay #MorningVibes | 2024-02-01 |
-- | 136     | 14       | Another #HappyDay with good vibes! #FeelGood               | 2024-02-03 |
-- | 137     | 15       | Productivity peaks! #WorkLife #ProductiveDay               | 2024-02-04 |
-- | 138     | 16       | Exploring new tech frontiers. #TechLife #Innovation        | 2024-02-04 |
-- | 139     | 17       | Gratitude for today's moments. #HappyDay #Thankful         | 2024-02-05 |
-- | 140     | 18       | Innovation drives us. #TechLife #FutureTech                | 2024-02-07 |
-- | 141     | 19       | Connecting with nature's serenity. #Nature #Peaceful       | 2024-02-09 |
-- +---------+----------+------------------------------------------------------------+------------+
-- Output:
-- +-----------+---------------+
-- | hashtag   | hashtag_count |
-- +-----------+---------------+
-- | #HappyDay | 3             |
-- | #TechLife | 2             |
-- | #WorkLife | 1             |
-- +-----------+---------------+
-- Explanation:
-- #HappyDay: Appeared in tweet IDs 13, 14, and 17, with a total count of 3 mentions.
-- #TechLife: Appeared in tweet IDs 16 and 18, with a total count of 2 mentions.
-- #WorkLife: Appeared in tweet ID 15, with a total count of 1 mention.
-- Note: Output table is sorted in descending order by hashtag_count and hashtag respectively.

-- Create table If Not Exists Tweets (user_id int, tweet_id int,  tweet_date date, tweet varchar(100))
-- Truncate table Tweets
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('135', '13', 'Enjoying a great start to the day. #HappyDay #MorningVibes', '2024-02-01')
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('136', '14', 'Another #HappyDay with good vibes! #FeelGood', '2024-02-03')
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('137', '15', 'Productivity peaks! #WorkLife #ProductiveDay', '2024-02-04')
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('138', '16', 'Exploring new tech frontiers. #TechLife #Innovation', '2024-02-04')
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('139', '17', 'Gratitude for today's moments. #HappyDay #Thankful', '2024-02-05')
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('140', '18', 'Innovation drives us. #TechLife #FutureTech', '2024-02-07')
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('141', '19', 'Connecting with nature's serenity. #Nature #Peaceful', '2024-02-09')


-- SELECT 
--     substring_index(REGEXP_SUBSTR(tweet, '#[a-zA-z](.*)'), ' ', -1) AS hashtag1, -- 第一个标签
--     substring_index(REGEXP_SUBSTR(tweet, '#[a-zA-z](.*)'), ' ', 1) AS hashtag2  -- 第二个标签
-- FROM 
--     Tweets;

-- | hashtag1       | hashtag2  |
-- | -------------- | --------- |
-- | #MorningVibes  | #HappyDay |
-- | #FeelGood      | #HappyDay |
-- | #ProductiveDay | #WorkLife |
-- | #Innovation    | #TechLife |
-- | #Thankful      | #HappyDay |
-- | #FutureTech    | #TechLife |
-- | #Peaceful      | #Nature   |

-- 不止两个
-- SELECT 
--     REGEXP_SUBSTR(tweet, '#[a-zA-z](.*)',1, 1, 'i') AS hashtag1, -- 第一个标签
--     REGEXP_SUBSTR(tweet, '#[a-zA-z](.*)',1, 2, 'i') AS hashtag2  -- 第二个标签
-- FROM 
--     Tweets;

-- Write your MySQL query statement below
WITH RECURSIVE t AS ( -- 将 tweet 按空格拆分
    ( 
        SELECT 
            SUBSTRING_INDEX(SUBSTRING_INDEX(tweet,' ', 1), ' ',-1) AS w, 
            1 AS n, 
            tweet_id AS tid  
        FROM 
            tweets 
        WHERE DATE_FORMAT(tweet_date , '%Y%m') = '202402'
    )
    UNION ALL
    (
        SELECT 
            SUBSTRING_INDEX(SUBSTRING_INDEX(tweet,' ', n + 1), ' ',-1) AS w, 
            n + 1 AS n, 
            tweet_id 
        FROM 
            t, tweets 
        WHERE 
            n < LENGTH(tweet) - LENGTH(REPLACE(tweet,' ','')) + 1 AND 
            DATE_FORMAT(tweet_date , '%Y%m') = '202402' AND 
            tid = tweets.tweet_id 
    )
)
-- SELECT * FROM t;

SELECT 
    w AS hashtag, 
    COUNT(1) AS count 
FROM 
    t 
WHERE
    w REGEXP '#.+?\\b' -- 只取 # 开头的单词
GROUP BY 
    w 
ORDER BY 
    count DESC, hashtag DESC -- the result table orderd by count of hashtag, hashtag in descending order.
LIMIT 3


