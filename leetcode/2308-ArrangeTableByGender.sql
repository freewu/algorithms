-- 2308. Arrange Table by Gender
-- Table: Genders
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | user_id     | int     |
-- | gender      | varchar |
-- +-------------+---------+
-- user_id is the primary key (column with unique values) for this table.
-- gender is ENUM (category) of type 'female', 'male', or 'other'.
-- Each row in this table contains the ID of a user and their gender.
-- The table has an equal number of 'female', 'male', and 'other'.
 
-- Write a solution to rearrange the Genders table such that the rows alternate between 'female', 'other', and 'male' in order. 
-- The table should be rearranged such that the IDs of each gender are sorted in ascending order.

-- Return the result table in the mentioned order.
-- The result format is shown in the following example.

-- Example 1:
-- Input: 
-- Genders table:
-- +---------+--------+
-- | user_id | gender |
-- +---------+--------+
-- | 4       | male   |
-- | 7       | female |
-- | 2       | other  |
-- | 5       | male   |
-- | 3       | female |
-- | 8       | male   |
-- | 6       | other  |
-- | 1       | other  |
-- | 9       | female |
-- +---------+--------+
-- Output: 
-- +---------+--------+
-- | user_id | gender |
-- +---------+--------+
-- | 3       | female |
-- | 1       | other  |
-- | 4       | male   |
-- | 7       | female |
-- | 2       | other  |
-- | 5       | male   |
-- | 9       | female |
-- | 6       | other  |
-- | 8       | male   |
-- +---------+--------+
-- Explanation: 
-- Female gender: IDs 3, 7, and 9.
-- Other gender: IDs 1, 2, and 6.
-- Male gender: IDs 4, 5, and 8.
-- We arrange the table alternating between 'female', 'other', and 'male'.
-- Note that the IDs of each gender are sorted in ascending order.

-- Create table If Not Exists Genders (user_id int, gender ENUM('female', 'other', 'male'))
-- Truncate table Genders
-- insert into Genders (user_id, gender) values ('4', 'male')
-- insert into Genders (user_id, gender) values ('7', 'female')
-- insert into Genders (user_id, gender) values ('2', 'other')
-- insert into Genders (user_id, gender) values ('5', 'male')
-- insert into Genders (user_id, gender) values ('3', 'female')
-- insert into Genders (user_id, gender) values ('8', 'male')
-- insert into Genders (user_id, gender) values ('6', 'other')
-- insert into Genders (user_id, gender) values ('1', 'other')
-- insert into Genders (user_id, gender) values ('9', 'female')

WITH f AS
( -- 女性数据
    SELECT
        user_id,
        gender,
        RANK() OVER (ORDER BY user_id) AS rk
    FROM
        Genders 
    WHERE
        gender = 'female'
),
o AS
( -- 其它数据
    SELECT
        user_id,
        gender,
        RANK() OVER (ORDER BY user_id) AS rk
    FROM
        Genders 
    WHERE
        gender = 'other'
),
m AS
( -- 男性数据
    SELECT
        user_id,
        gender,
        RANK() OVER (ORDER BY user_id) AS rk
    FROM
        Genders 
    WHERE
        gender = 'male'
)

SELECT
    user_id,
    gender
FROM
(
    (
        SELECT user_id,gender, rk * 3 AS rk FROM f
    )
    UNION ALL 
    (
        SELECT user_id,gender, rk * 3 + 1 AS rk FROM o
    )
    UNION ALL 
    (
        SELECT user_id,gender, rk  * 3  + 2 AS rk FROM m
    )
) AS r 
ORDER BY 
    rk

-- best solution
SELECT 
    *
FROM 
    genders
ORDER BY 
    ROW_NUMBER() OVER(PARTITION BY gender ORDER BY user_id),
    length(gender) DESC