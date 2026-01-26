SELECT
    p.product_id,
    p.product_name,
    SUM(od.quantity) AS total_qty
FROM oe.order_detail od
JOIN oe.products p
    ON p.product_id = od.product_id
GROUP BY
    p.product_id,
    p.product_name
ORDER BY
    total_qty DESC;