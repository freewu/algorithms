-- 1149. Article Views II
-- Table: Views
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | article_id    | int     |
-- | author_id     | int     |
-- | viewer_id     | int     |
-- | view_date     | date    |
-- +---------------+---------+
-- This table may have duplicate rows.
-- Each row of this table indicates that some viewer viewed an article (written by some author) on some date. 
-- Note that equal author_id and viewer_id indicate the same person.
 
-- Write a solution to find all the people who viewed more than one article on the same date.
-- Return the result table sorted by id in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Views table:
-- +------------+-----------+-----------+------------+
-- | article_id | author_id | viewer_id | view_date  |
-- +------------+-----------+-----------+------------+
-- | 1          | 3         | 5         | 2019-08-01 |
-- | 3          | 4         | 5         | 2019-08-01 |
-- | 1          | 3         | 6         | 2019-08-02 |
-- | 2          | 7         | 7         | 2019-08-01 |
-- | 2          | 7         | 6         | 2019-08-02 |
-- | 4          | 7         | 1         | 2019-07-22 |
-- | 3          | 4         | 4         | 2019-07-21 |
-- | 3          | 4         | 4         | 2019-07-21 |
-- +------------+-----------+-----------+------------+
-- Output: 
-- +------+
-- | id   |
-- +------+
-- | 5    |
-- | 6    |
-- +------+

-- Create table If Not Exists Views (article_id int, author_id int, viewer_id int, view_date date)
-- Truncate table Views
-- insert into Views (article_id, author_id, viewer_id, view_date) values ('1', '3', '5', '2019-08-01')
-- insert into Views (article_id, author_id, viewer_id, view_date) values ('3', '4', '5', '2019-08-01')
-- insert into Views (article_id, author_id, viewer_id, view_date) values ('1', '3', '6', '2019-08-02')
-- insert into Views (article_id, author_id, viewer_id, view_date) values ('2', '7', '7', '2019-08-01')
-- insert into Views (article_id, author_id, viewer_id, view_date) values ('2', '7', '6', '2019-08-02')
-- insert into Views (article_id, author_id, viewer_id, view_date) values ('4', '7', '1', '2019-07-22')
-- insert into Views (article_id, author_id, viewer_id, view_date) values ('3', '4', '4', '2019-07-21')
-- insert into Views (article_id, author_id, viewer_id, view_date) values ('3', '4', '4', '2019-07-21')

SELECT
    DISTINCT viewer_id AS id
FROM
    Views 
GROUP BY
    view_date,viewer_id 
HAVING 
    COUNT(DISTINCT article_id) >= 2 -- 同一天阅读至少两篇文章
ORDER BY 
    id -- 结果按照 id 升序排序


SELECT
    DISTINCT viewer_id AS id
FROM
( -- 每天每人阅读的文章数
    SELECT 
        viewer_id, 
        view_date,
        COUNT(DISTINCT article_id) AS cnt
    FROM 
        views
    GROUP BY
        viewer_id, view_date
) AS t
WHERE 
    cnt >= 2
ORDER BY 
    id -- 结果按照 id 升序排序