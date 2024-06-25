-- 3198. Find Cities in Each State
-- Table: cities
-- +-------------+---------+
-- | Column Name | Type    | 
-- +-------------+---------+
-- | state       | varchar |
-- | city        | varchar |
-- +-------------+---------+
-- (state, city) is the primary key (combination of columns with unique values) for this table.
-- Each row of this table contains the state name and the city name within that state.
-- Write a solution to find all the cities in each state and combine them into a single comma-separated string.

-- Return the result table ordered by state in ascending order.
-- The result format is in the following example.

-- Example:
-- Input:
-- cities table:
-- +-------------+---------------+
-- | state       | city          |
-- +-------------+---------------+
-- | California  | Los Angeles   |
-- | California  | San Francisco |
-- | California  | San Diego     |
-- | Texas       | Houston       |
-- | Texas       | Austin        |
-- | Texas       | Dallas        |
-- | New York    | New York City |
-- | New York    | Buffalo       |
-- | New York    | Rochester     |
-- +-------------+---------------+
-- Output:
-- +-------------+---------------------------------------+
-- | state       | cities                                |
-- +-------------+---------------------------------------+
-- | California  | Los Angeles, San Diego, San Francisco |
-- | New York    | Buffalo, New York City, Rochester     |
-- | Texas       | Austin, Dallas, Houston               |
-- +-------------+---------------------------------------+
-- Explanation:
-- California: All cities ("Los Angeles", "San Diego", "San Francisco") are listed in a comma-separated string.
-- New York: All cities ("Buffalo", "New York City", "Rochester") are listed in a comma-separated string.
-- Texas: All cities ("Austin", "Dallas", "Houston") are listed in a comma-separated string.
-- Note: The output table is ordered by the state name in ascending order.

-- Create table if not exists cities( state varchar(100),city varchar(100))
-- Truncate table cities
-- insert into cities (state, city) values ('California', 'Los Angeles')
-- insert into cities (state, city) values ('California', 'San Francisco')
-- insert into cities (state, city) values ('California', 'San Diego')
-- insert into cities (state, city) values ('Texas', 'Houston')
-- insert into cities (state, city) values ('Texas', 'Austin')
-- insert into cities (state, city) values ('Texas', 'Dallas')
-- insert into cities (state, city) values ('New York', 'New York City')
-- insert into cities (state, city) values ('New York', 'Buffalo')
-- insert into cities (state, city) values ('New York', 'Rochester')

-- Write your MySQL query statement below
SELECT 
    state,
    GROUP_CONCAT(DISTINCT city ORDER BY city SEPARATOR ', ') AS cities
FROM
    cities  
GROUP BY
    state
ORDER BY
    state