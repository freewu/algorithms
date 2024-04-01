-- 3055. Top Percentile Fraud
-- Table: Fraud
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | policy_id   | int     |
-- | state       | varchar |
-- | fraud_score | int     |
-- +-------------+---------+
-- policy_id is column of unique values for this table.
-- This table contains policy id, state, and fraud score.
-- The Leetcode Insurance Corp has developed an ML-driven predictive model to detect the likelihood of fraudulent claims. Consequently, they allocate their most seasoned claim adjusters to address the top 5% of claims flagged by this model.

-- Write a solution to find the top 5 percentile of claims from each state.
-- Return the result table ordered by state in ascending order, fraud_score in descending order, and policy_id in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Fraud table:
-- +-----------+------------+-------------+
-- | policy_id | state      | fraud_score | 
-- +-----------+------------+-------------+
-- | 1         | California | 0.92        | 
-- | 2         | California | 0.68        |   
-- | 3         | California | 0.17        | 
-- | 4         | New York   | 0.94        | 
-- | 5         | New York   | 0.81        | 
-- | 6         | New York   | 0.77        |  
-- | 7         | Texas      | 0.98        |  
-- | 8         | Texas      | 0.97        | 
-- | 9         | Texas      | 0.96        | 
-- | 10        | Florida    | 0.97        |  
-- | 11        | Florida    | 0.98        | 
-- | 12        | Florida    | 0.78        | 
-- | 13        | Florida    | 0.88        | 
-- | 14        | Florida    | 0.66        | 
-- +-----------+------------+-------------+
-- Output: 
-- +-----------+------------+-------------+
-- | policy_id | state      | fraud_score |
-- +-----------+------------+-------------+
-- | 1         | California | 0.92        | 
-- | 11        | Florida    | 0.98        | 
-- | 4         | New York   | 0.94        | 
-- | 7         | Texas      | 0.98        |  
-- +-----------+------------+-------------+
-- Explanation
-- - For the state of California, only policy ID 1, with a fraud score of 0.92, falls within the top 5 percentile for this state.
-- - For the state of Florida, only policy ID 11, with a fraud score of 0.98, falls within the top 5 percentile for this state. 
-- - For the state of New York, only policy ID 4, with a fraud score of 0.94, falls within the top 5 percentile for this state. 
-- - For the state of Texas, only policy ID 7, with a fraud score of 0.98, falls within the top 5 percentile for this state. 
-- Output table is ordered by state in ascending order, fraud score in descending order, and policy ID in ascending order.

-- Create table If Not Exists Fraud (policy_id int, state varchar(50), fraud_score decimal(5,2))
-- Truncate table Fraud
-- insert into Fraud (policy_id, state, fraud_score) values ('1', 'California', '0.92')
-- insert into Fraud (policy_id, state, fraud_score) values ('2', 'California', '0.68')
-- insert into Fraud (policy_id, state, fraud_score) values ('3', 'California', '0.17')
-- insert into Fraud (policy_id, state, fraud_score) values ('4', 'New York', '0.94')
-- insert into Fraud (policy_id, state, fraud_score) values ('5', 'New York', '0.81')
-- insert into Fraud (policy_id, state, fraud_score) values ('6', 'New York', '0.77')
-- insert into Fraud (policy_id, state, fraud_score) values ('7', 'Texas', '0.98')
-- insert into Fraud (policy_id, state, fraud_score) values ('8', 'Texas', '0.97')
-- insert into Fraud (policy_id, state, fraud_score) values ('9', 'Texas', '0.96')
-- insert into Fraud (policy_id, state, fraud_score) values ('10', 'Florida', '0.97')
-- insert into Fraud (policy_id, state, fraud_score) values ('11', 'Florida', '0.98')
-- insert into Fraud (policy_id, state, fraud_score) values ('12', 'Florida', '0.78')
-- insert into Fraud (policy_id, state, fraud_score) values ('13', 'Florida', '0.88')
-- insert into Fraud (policy_id, state, fraud_score) values ('14', 'Florida', '0.66')

-- # PERCENT_RANK()函数
--      PERCENT_RANK() 函数用于将每行按照(rank - 1) / (rows - 1) 进行计算,用以求MySQL中前百分之N问题。
--      其中，rank为RANK()函数产生的序号，rows为当前窗口的记录总行数
--      PERCENT_RANK()函数返回介于 0 和 1 之间的小数值

-- PERCENT_RANK
SELECT
    policy_id,
    state,
    fraud_score
FROM
    (
        SELECT
            policy_id,
            state,
            fraud_score,
            PERCENT_RANK() OVER(PARTITION BY state ORDER BY fraud_score DESC) AS percent -- 按每州欺诈分由高到低计算出占比 
        FROM Fraud
    ) AS t 
WHERE
    t.percent <= 0.05 -- 取前 5 百分位数
ORDER BY 
    -- Return the result table ordered by state in ascending order, fraud_score in descending order, and policy_id in ascending order.
    state ASC,fraud_score DESC,policy_id ASC

-- rank
SELECT
    policy_id,
    state,
    fraud_score
FROM
    (
        SELECT 
            policy_id,
            state,
            fraud_score,
            (
                (RANK() OVER(PARTITION BY state ORDER BY fraud_score DESC) - 1)  -- 欺诈分在州的排名
                / 
                COUNT(*) OVER(partition by state) -- 每个州的数量
            ) AS percent
        FROM 
            fraud
    ) AS t
WHERE 
    percent <= 0.05
ORDER BY 
    -- Return the result table ordered by state in ascending order, fraud_score in descending order, and policy_id in ascending order.
    state ASC,fraud_score DESC,policy_id ASC

