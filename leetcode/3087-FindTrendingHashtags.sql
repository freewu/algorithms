-- 3087. Find Trending Hashtags
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
-- Write a solution to find the top 3 trending hashtags in February 2024.

-- Return the result table orderd by count of hashtag, hastag in descending order.
-- The result format is in the following example.

-- Example 1:
-- Input:
-- Tweets table:
-- +---------+----------+----------------------------------------------+------------+
-- | user_id | tweet_id | tweet                                        | tweet_date |
-- +---------+----------+----------------------------------------------+------------+
-- | 135     | 13       | Enjoying a great start to the day! #HappyDay | 2024-02-01 |
-- | 136     | 14       | Another #HappyDay with good vibes!           | 2024-02-03 |
-- | 137     | 15       | Productivity peaks! #WorkLife                | 2024-02-04 |
-- | 138     | 16       | Exploring new tech frontiers. #TechLife      | 2024-02-04 |
-- | 139     | 17       | Gratitude for today's moments. #HappyDay     | 2024-02-05 |
-- | 140     | 18       | Innovation drives us. #TechLife              | 2024-02-07 |
-- | 141     | 19       | Connecting with nature's serenity. #Nature   | 2024-02-09 |
-- +---------+----------+----------------------------------------------+------------+
-- Output:
-- +-----------+--------------+
-- | hashtag   | hashtag_count|
-- +-----------+--------------+
-- | #HappyDay | 3            |
-- | #TechLife | 2            |
-- | #WorkLife | 1            |
-- +-----------+--------------+
-- Explanation:
-- #HappyDay: Appeared in tweet IDs 13, 14, and 17, with a total count of 3 mentions.
-- #TechLife: Appeared in tweet IDs 16 and 18, with a total count of 2 mentions.
-- #WorkLife: Appeared in tweet ID 15, with a total count of 1 mention.
-- Note: Output table is sorted in descending order by hashtag_count and hashtag respectively.

-- Create table If Not Exists Tweets (user_id int, tweet_id int,  tweet_date date, tweet varchar(100))
-- Truncate table Tweets
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('135', '13', 'Enjoying a great start to the day. #HappyDay', '2024-02-01')
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('136', '14', 'Another #HappyDay with good ', '2024-02-03')
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('137', '15', 'Productivity peaks! #WorkLife', '2024-02-04')
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('138', '16', 'Exploring new tech frontiers. #TechLife', '2024-02-04')
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('139', '17', 'Gratitude for today's moments. #HappyDay', '2024-02-05')
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('140', '18', 'Innovation drives us. #TechLife', '2024-02-07')
-- insert into Tweets (user_id, tweet_id, tweet, tweet_date) values ('141', '19', 'Connecting with nature's serenity. #Nature', '2024-02-09')

-- 正则提取标签
-- SELECT 
--     SUBSTRING(tweet, REGEXP_INSTR(tweet, '#[a-zA-z](.*)')) AS hashtag 
-- FROM Tweets;

SELECT substring_index(REGEXP_SUBSTR(tweet, '#[a-zA-z](.*)'), ' ', -1) AS hashtag FROM Tweets;

-- 使用 SUBSTRING 函数结合正则表达式
-- 在MySQL中，我们还可以使用 SUBSTRING 函数结合正则表达式来提取数据。 SUBSTRING 函数可以用来截取一
-- 个字符串的一部分。
-- 以下是一个使用 SUBSTRING 函数结合正则表达式提取数据的示例：

--     SELECT SUBSTRING(column_name, REGEXP_INSTR(column_name, 'pattern')) AS extracted_data FROM table_name;

-- 这个示例中， table_name 是要查询的表的名称， column_name 是要匹配的字段的名称， pattern 是要匹配的正则
-- 表达式模式。 REGEXP_INSTR 函数可以返回匹配到的子字符串的起始位置。

-- 使用正则提取
SELECT 
    hashtag,
    COUNT(1) AS hashtag_count
FROM
( -- 提取2024年2月份的标签
    SELECT 
        SUBSTRING_INDEX(REGEXP_SUBSTR(tweet, '#[a-zA-z](.*)'), ' ', 1) AS hashtag 
    FROM 
        Tweets
    WHERE
        tweet_date BETWEEN '2024-02-01' AND '2024-02-29'
) AS r 
GROUP BY 
    hashtag
ORDER BY 
    hashtag_count DESC, hashtag DESC
LIMIT
    3


-- 使用 SUBSTRING_INDEX
SELECT
    CONCAT('#', SUBSTRING_INDEX(SUBSTRING_INDEX(tweet, '#', -1), ' ', 1)) AS hashtag,
    COUNT(1) AS hashtag_count
FROM 
    Tweets
WHERE 
    DATE_FORMAT(tweet_date, '%Y%m') = '202402'
GROUP BY 
    hashtag
ORDER BY 
    hashtag_count DESC, hashtag DESC
LIMIT 
    3;