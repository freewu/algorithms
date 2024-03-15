-- 1341. Movie Rating
-- Table: Movies
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | movie_id      | int     |
-- | title         | varchar |
-- +---------------+---------+
-- movie_id is the primary key (column with unique values) for this table.
-- title is the name of the movie.
 
-- Table: Users
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | user_id       | int     |
-- | name          | varchar |
-- +---------------+---------+
-- user_id is the primary key (column with unique values) for this table.
 
-- Table: MovieRating
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | movie_id      | int     |
-- | user_id       | int     |
-- | rating        | int     |
-- | created_at    | date    |
-- +---------------+---------+
-- (movie_id, user_id) is the primary key (column with unique values) for this table.
-- This table contains the rating of a movie by a user in their review.
-- created_at is the user's review date. 
 
-- Write a solution to:
--      Find the name of the user who has rated the greatest number of movies. In case of a tie, return the lexicographically smaller user name.
--      Find the movie name with the highest average rating in February 2020. In case of a tie, return the lexicographically smaller movie name.

-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Movies table:
-- +-------------+--------------+
-- | movie_id    |  title       |
-- +-------------+--------------+
-- | 1           | Avengers     |
-- | 2           | Frozen 2     |
-- | 3           | Joker        |
-- +-------------+--------------+
-- Users table:
-- +-------------+--------------+
-- | user_id     |  name        |
-- +-------------+--------------+
-- | 1           | Daniel       |
-- | 2           | Monica       |
-- | 3           | Maria        |
-- | 4           | James        |
-- +-------------+--------------+
-- MovieRating table:
-- +-------------+--------------+--------------+-------------+
-- | movie_id    | user_id      | rating       | created_at  |
-- +-------------+--------------+--------------+-------------+
-- | 1           | 1            | 3            | 2020-01-12  |
-- | 1           | 2            | 4            | 2020-02-11  |
-- | 1           | 3            | 2            | 2020-02-12  |
-- | 1           | 4            | 1            | 2020-01-01  |
-- | 2           | 1            | 5            | 2020-02-17  | 
-- | 2           | 2            | 2            | 2020-02-01  | 
-- | 2           | 3            | 2            | 2020-03-01  |
-- | 3           | 1            | 3            | 2020-02-22  | 
-- | 3           | 2            | 4            | 2020-02-25  | 
-- +-------------+--------------+--------------+-------------+
-- Output: 
-- +--------------+
-- | results      |
-- +--------------+
-- | Daniel       |
-- | Frozen 2     |
-- +--------------+
-- Explanation: 
-- Daniel and Monica have rated 3 movies ("Avengers", "Frozen 2" and "Joker") but Daniel is smaller lexicographically.
-- Frozen 2 and Joker have a rating average of 3.5 in February but Frozen 2 is smaller lexicographically.

-- Create table If Not Exists Movies (movie_id int, title varchar(30))
-- Create table If Not Exists Users (user_id int, name varchar(30))
-- Create table If Not Exists MovieRating (movie_id int, user_id int, rating int, created_at date)
-- Truncate table Movies
-- insert into Movies (movie_id, title) values ('1', 'Avengers')
-- insert into Movies (movie_id, title) values ('2', 'Frozen 2')
-- insert into Movies (movie_id, title) values ('3', 'Joker')
-- Truncate table Users
-- insert into Users (user_id, name) values ('1', 'Daniel')
-- insert into Users (user_id, name) values ('2', 'Monica')
-- insert into Users (user_id, name) values ('3', 'Maria')
-- insert into Users (user_id, name) values ('4', 'James')
-- Truncate table MovieRating
-- insert into MovieRating (movie_id, user_id, rating, created_at) values ('1', '1', '3', '2020-01-12')
-- insert into MovieRating (movie_id, user_id, rating, created_at) values ('1', '2', '4', '2020-02-11')
-- insert into MovieRating (movie_id, user_id, rating, created_at) values ('1', '3', '2', '2020-02-12')
-- insert into MovieRating (movie_id, user_id, rating, created_at) values ('1', '4', '1', '2020-01-01')
-- insert into MovieRating (movie_id, user_id, rating, created_at) values ('2', '1', '5', '2020-02-17')
-- insert into MovieRating (movie_id, user_id, rating, created_at) values ('2', '2', '2', '2020-02-01')
-- insert into MovieRating (movie_id, user_id, rating, created_at) values ('2', '3', '2', '2020-03-01')
-- insert into MovieRating (movie_id, user_id, rating, created_at) values ('3', '1', '3', '2020-02-22')
-- insert into MovieRating (movie_id, user_id, rating, created_at) values ('3', '2', '4', '2020-02-25')

-- Write your MySQL query statement below
(-- 查找评论电影数量最多的用户名。如果出现平局，返回字典序较小的用户名
    SELECT
        u.name AS results      
    FROM
        Users AS u, 
        MovieRating AS r
    WHERE
        u.user_id = r.user_id
    GROUP BY
        r.user_id    
    ORDER BY 
        COUNT(r.movie_id) DESC,u.name ASC
    LIMIT 1 
) 
UNION ALL 
(-- 查找在 February 2020 平均评分最高 的电影名称。如果出现平局，返回字典序较小的电影名称
    SELECT
        m.title AS results      
    FROM
        Movies AS m, 
        MovieRating AS r
    WHERE
        m.movie_id = r.movie_id  AND
        r.created_at  between '2020-02-01' AND '2020-02-29'  
    GROUP BY
        r.movie_id    
    ORDER BY 
        AVG(r.rating) DESC,m.title ASC
    LIMIT 1  
)

-- best solution
( -- 查找评论电影数量最多的用户名。如果出现平局，返回字典序较小的用户名
    SELECT 
        name AS results
    FROM 
        movieRating 
    INNER JOIN users USING(user_id)
    GROUP BY 
        user_id
    ORDER BY 
        COUNT(movie_id) DESC, name ASC
    LIMIT 1
) 
UNION ALL 
( -- 查找在 February 2020 平均评分最高 的电影名称。如果出现平局，返回字典序较小的电影名称
    SELECT
        m.title AS results      
    FROM
        Movies AS m, 
        MovieRating AS r
    WHERE
        m.movie_id = r.movie_id  AND
        r.created_at  between '2020-02-01' AND '2020-02-29'  
    GROUP BY
        r.movie_id    
    ORDER BY 
        AVG(r.rating) DESC,m.title ASC
    LIMIT 1  
)