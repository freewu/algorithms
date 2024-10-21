-- 3328. Find Cities in Each State II
-- Table: cities
-- +-------------+---------+
-- | Column Name | Type    | 
-- +-------------+---------+
-- | state       | varchar |
-- | city        | varchar |
-- +-------------+---------+
-- (state, city) is the combination of columns with unique values for this table.
-- Each row of this table contains the state name and the city name within that state.
-- Write a solution to find all the cities in each state and analyze them based on the following requirements:

-- Combine all cities into a comma-separated string for each state.
-- Only include states that have at least 3 cities.
-- Only include states where at least one city starts with the same letter as the state name.
-- Return the result table ordered by the count of matching-letter cities in descending order and then by state name in ascending order.

-- The result format is in the following example.

-- Example:
-- Input:
-- cities table:
-- +--------------+---------------+
-- | state        | city          |
-- +--------------+---------------+
-- | New York     | New York City |
-- | New York     | Newark        |
-- | New York     | Buffalo       |
-- | New York     | Rochester     |
-- | California   | San Francisco |
-- | California   | Sacramento    |
-- | California   | San Diego     |
-- | California   | Los Angeles   |
-- | Texas        | Tyler         |
-- | Texas        | Temple        |
-- | Texas        | Taylor        |
-- | Texas        | Dallas        |
-- | Pennsylvania | Philadelphia  |
-- | Pennsylvania | Pittsburgh    |
-- | Pennsylvania | Pottstown     |
-- +--------------+---------------+
-- Output:
-- +-------------+-------------------------------------------+-----------------------+
-- | state       | cities                                    | matching_letter_count |
-- +-------------+-------------------------------------------+-----------------------+
-- | Pennsylvania| Philadelphia, Pittsburgh, Pottstown       | 3                     |
-- | Texas       | Dallas, Taylor, Temple, Tyler             | 2                     |
-- | New York    | Buffalo, Newark, New York City, Rochester | 2                     |
-- +-------------+-------------------------------------------+-----------------------+
-- Explanation:

-- Pennsylvania:
-- Has 3 cities (meets minimum requirement)
-- All 3 cities start with 'P' (same as state)
-- matching_letter_count = 3
-- Texas:
-- Has 4 cities (meets minimum requirement)
-- 2 cities (Temple, Tyler) start with 'T' (same as state)
-- matching_letter_count = 2
-- New York:
-- Has 4 cities (meets minimum requirement)
-- 2 cities (Newark, New York City) start with 'N' (same as state)
-- matching_letter_count = 2
-- California is not included in the output because:
-- Although it has 4 cities (meets minimum requirement)
-- No cities start with 'C' (doesn't meet the matching letter requirement)
-- Note:
-- Results are ordered by matching_letter_count in descending order
-- When matching_letter_count is the same (Texas and New York both have 2), they are ordered by state name alphabetically
-- Cities in each row are ordered alphabetically

-- Create table if not exists cities( state varchar(100),city varchar(100))
-- Truncate table cities
-- insert into cities (state, city) values ('New York', 'New York City')
-- insert into cities (state, city) values ('New York', 'Newark')
-- insert into cities (state, city) values ('New York', 'Buffalo')
-- insert into cities (state, city) values ('New York', 'Rochester')
-- insert into cities (state, city) values ('California', 'San Francisco')
-- insert into cities (state, city) values ('California', 'Sacramento')
-- insert into cities (state, city) values ('California', 'San Diego')
-- insert into cities (state, city) values ('California', 'Los Angeles')
-- insert into cities (state, city) values ('Texas', 'Tyler')
-- insert into cities (state, city) values ('Texas', 'Temple')
-- insert into cities (state, city) values ('Texas', 'Taylor')
-- insert into cities (state, city) values ('Texas', 'Dallas')
-- insert into cities (state, city) values ('Pennsylvania', 'Philadelphia')
-- insert into cities (state, city) values ('Pennsylvania', 'Pittsburgh')
-- insert into cities (state, city) values ('Pennsylvania', 'Pottstown')

-- Write your MySQL query statement below
WITH t AS (
    SELECT
        state,
        LEFT(state, 1) AS s_f,
        city,
        LEFT(city, 1) AS c_f
    FROM 
       cities 
)
-- SELECT * FROM t;
-- | state        | s_f | city          | c_f |
-- | ------------ | --- | ------------- | --- |
-- | New York     | N   | New York City | N   |
-- | New York     | N   | Newark        | N   |
-- | New York     | N   | Buffalo       | B   |
-- | New York     | N   | Rochester     | R   |
-- | California   | C   | San Francisco | S   |
-- | California   | C   | Sacramento    | S   |
-- | California   | C   | San Diego     | S   |
-- | California   | C   | Los Angeles   | L   |
-- | Texas        | T   | Tyler         | T   |
-- | Texas        | T   | Temple        | T   |
-- | Texas        | T   | Taylor        | T   |
-- | Texas        | T   | Dallas        | D   |
-- | Pennsylvania | P   | Philadelphia  | P   |
-- | Pennsylvania | P   | Pittsburgh    | P   |
-- | Pennsylvania | P   | Pottstown     | P   |
, t1 AS (
    SELECT 
        state,
        GROUP_CONCAT(city ORDER BY city ASC SEPARATOR ", ") AS cities,
        SUM(IF(s_f = c_f, 1, 0)) AS matching_letter_count,
        COUNT(*) AS num
    FROM 
        t
    GROUP BY
        state
)
-- | state        | cities                                         | matching_letter_count |
-- | ------------ | ---------------------------------------------- | --------------------- |
-- | Pennsylvania | Philadelphia,Pittsburgh,Pottstown              | 3                     |
-- | Texas        | Tyler,Temple,Taylor,Dallas                     | 3                     |
-- | New York     | New York City,Newark,Buffalo,Rochester         | 2                     |
-- | California   | San Francisco,Sacramento,San Diego,Los Angeles | 0                     |

SELECT 
    state,
    cities,
    matching_letter_count
FROM 
    t1 
WHERE 
    matching_letter_count > 0 AND -- Only include states where at least one city starts with the same letter as the state name.
    num >= 3 -- Only include states that have at least 3 cities
ORDER BY
    matching_letter_count DESC, state -- matching-letter cities in descending order and then by state name in ascending order.