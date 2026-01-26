SELECT
    c.customer_id,
    c.company_name,
    COUNT(o.order_id) AS total_order
FROM oe.customers c
LEFT JOIN oe.orders o
    ON o.customer_id = c.customer_id
GROUP BY
    c.customer_id,
    c.company_name
ORDER BY
    total_order DESC,
    c.customer_id;