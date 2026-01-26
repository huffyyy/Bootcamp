SELECT
    o.order_id,
    o.customer_id,
    o.order_date,
    o.required_date,
    o.shipped_date,
    (o.shipped_date - o.order_date) AS delivery_time
FROM oe.orders o
WHERE o.order_id IN (
    SELECT order_id
    FROM oe.orders
    WHERE shipped_date IS NOT NULL
      AND (shipped_date - order_date) > 7
)
ORDER BY delivery_time DESC, o.order_id;