-- 3401. Find Circular Gift Exchange Chains
-- Table: SecretSanta
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | giver_id    | int  |
-- | receiver_id | int  |
-- | gift_value  | int  |
-- +-------------+------+
-- (giver_id, receiver_id) is the unique key for this table.   
-- Each row represents a record of a gift exchange between two employees, giver_id represents the employee who gives a gift, receiver_id represents the employee who receives the gift and gift_value represents the value of the gift given.  
-- Write a solution to find the total gift value and length of circular chains of Secret Santa gift exchanges:

-- A circular chain is defined as a series of exchanges where:
--     1. Each employee gives a gift to exactly one other employee.
--     2. Each employee receives a gift from exactly one other employee.
--     3. The exchanges form a continuous loop (e.g., employee A gives a gift to B, B gives to C, and C gives back to A).

-- Return the result ordered by the chain length and total gift value of the chain in descending order. 

-- The result format is in the following example.
-- Example:
-- Input:
-- SecretSanta table:
-- +----------+-------------+------------+
-- | giver_id | receiver_id | gift_value |
-- +----------+-------------+------------+
-- | 1        | 2           | 20         |
-- | 2        | 3           | 30         |
-- | 3        | 1           | 40         |
-- | 4        | 5           | 25         |
-- | 5        | 4           | 35         |
-- +----------+-------------+------------+
-- Output:
-- +----------+--------------+------------------+
-- | chain_id | chain_length | total_gift_value |
-- +----------+--------------+------------------+
-- | 1        | 3            | 90               |
-- | 2        | 2            | 60               |
-- +----------+--------------+------------------+
-- Explanation:
-- Chain 1 involves employees 1, 2, and 3:
-- Employee 1 gives a gift to 2, employee 2 gives a gift to 3, and employee 3 gives a gift to 1.
-- Total gift value for this chain = 20 + 30 + 40 = 90.
-- Chain 2 involves employees 4 and 5:
-- Employee 4 gives a gift to 5, and employee 5 gives a gift to 4.
-- Total gift value for this chain = 25 + 35 = 60.
-- The result table is ordered by the chain length and total gift value of the chain in descending order.

-- CREATE TABLE SecretSanta (
--     giver_id INT,
--     receiver_id INT,
--     gift_value INT
-- )
-- Truncate table SecretSanta
-- insert into SecretSanta (giver_id, receiver_id, gift_value) values ('1', '2', '20')
-- insert into SecretSanta (giver_id, receiver_id, gift_value) values ('2', '3', '30')
-- insert into SecretSanta (giver_id, receiver_id, gift_value) values ('3', '1', '40')
-- insert into SecretSanta (giver_id, receiver_id, gift_value) values ('4', '5', '25')
-- insert into SecretSanta (giver_id, receiver_id, gift_value) values ('5', '4', '35')

WITH RECURSIVE cycles_cte AS (
    SELECT 
        giver_id, 
        receiver_id, 
        gift_value, 
        giver_id AS start
    FROM 
        SecretSanta
    UNION
    SELECT 
        s.giver_id, 
        s.receiver_id, 
        s.gift_value, 
        c.start
    FROM 
        SecretSanta s
    INNER JOIN 
        cycles_cte c ON s.giver_id = c.receiver_id AND s.giver_id != c.start
),
cycle_start_length_value_triples_cte AS (
    SELECT 
        start, 
        COUNT(*) AS chain_length, 
        SUM(gift_value) AS total_gift_value
    FROM 
        cycles_cte
    GROUP BY 
        start
    ORDER BY NULL
),
unique_cycle_length_value_pairs_cte AS (
    SELECT 
        chain_length, 
        total_gift_value
    FROM cycle_start_length_value_triples_cte
    GROUP BY 
        chain_length, total_gift_value
    ORDER BY NULL
)
SELECT 
    RANK() OVER (ORDER BY chain_length DESC, total_gift_value DESC) AS chain_id,
    chain_length,
    total_gift_value
FROM 
    unique_cycle_length_value_pairs_cte