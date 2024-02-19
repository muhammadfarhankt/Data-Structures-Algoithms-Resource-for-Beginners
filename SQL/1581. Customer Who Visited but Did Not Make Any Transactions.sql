-- Write your PostgreSQL query statement below
SELECT customer_id, COUNT(*) AS count_no_trans
FROM Visits
WHERE visit_id NOT IN (SELECT visit_id FROM Transactions)
GROUP BY customer_id;
--SELECT V.customer_id, COUNT(*) AS count_no_trans
--FROM Visits AS V
--LEFT JOIN Transactions AS T ON V.visit_id = T.visit_id
--WHERE T.transaction_id IS NULL
--GROUP BY V.customer_id;
