SELECT
    c.category_id,
    c.category_name,
    SUM(od.quantity) AS total_qty_ordered
FROM oe.order_detail od
JOIN oe.products p
    ON p.product_id = od.product_id
JOIN oe.categories c
    ON c.category_id = p.category_id
GROUP BY
    c.category_id,
    c.category_name
ORDER BY
    total_qty_ordered DESC;