-- 620. Not Boring Movies
-- Table: Cinema
-- +----------------+----------+
-- | Column Name    | Type     |
-- +----------------+----------+
-- | id             | int      |
-- | movie          | varchar  |
-- | description    | varchar  |
-- | rating         | float    |
-- +----------------+----------+
-- id is the primary key for this table.
-- Each row contains information about the name of a movie, its genre, and its rating.
-- rating is a 2 decimal places float in the range [0, 10]
--  
-- Write an SQL query to report the movies with an odd-numbered ID and a description that is not "boring".
-- Return the result table ordered by rating in descending order.
-- The query result format is in the following example.
--
-- Example 1:
-- Input:
-- Cinema table:
-- +----+------------+-------------+--------+
-- | id | movie      | description | rating |
-- +----+------------+-------------+--------+
-- | 1  | War        | great 3D    | 8.9    |
-- | 2  | Science    | fiction     | 8.5    |
-- | 3  | irish      | boring      | 6.2    |
-- | 4  | Ice song   | Fantacy     | 8.6    |
-- | 5  | House card | Interesting | 9.1    |
-- +----+------------+-------------+--------+
-- Output:
-- +----+------------+-------------+--------+
-- | id | movie      | description | rating |
-- +----+------------+-------------+--------+
-- | 5  | House card | Interesting | 9.1    |
-- | 1  | War        | great 3D    | 8.9    |
-- +----+------------+-------------+--------+
-- Explanation:
-- We have three movies with odd-numbered IDs: 1, 3, and 5. 
-- The movie with ID = 3 is boring so we do not include it in the answer.

-- Create table If Not Exists cinema (id int, movie varchar(255), description varchar(255), rating float(2, 1))
-- Truncate table cinema
-- insert into cinema (id, movie, description, rating) values ('1', 'War', 'great 3D', '8.9')
-- insert into cinema (id, movie, description, rating) values ('2', 'Science', 'fiction', '8.5')
-- insert into cinema (id, movie, description, rating) values ('3', 'irish', 'boring', '6.2')
-- insert into cinema (id, movie, description, rating) values ('4', 'Ice song', 'Fantacy', '8.6')
-- insert into cinema (id, movie, description, rating) values ('5', 'House card', 'Interesting', '9.1')

-- Write your MySQL query statement below
SELECT
    *
FROM
    cinema
WHERE
    id NOT IN ( -- 非 boring (不无聊)
        SELECT
            id
        FROM
            cinema
        WHERE
            description = 'boring'
    ) AND
    id % 2 = 1 -- id 为奇数
ORDER BY
    rating DESC -- 按等级 rating 排列