SELECT
    t.category_id,
    t.category_name,
    t.total_qty_ordered
FROM (
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
) t
WHERE t.total_qty_ordered = (
        SELECT MAX(total_qty_ordered)
        FROM (
            SELECT SUM(od.quantity) AS total_qty_ordered
            FROM oe.order_detail od
            JOIN oe.products p ON p.product_id = od.product_id
            GROUP BY p.category_id
        ) x
    )
   OR t.total_qty_ordered = (
        SELECT MIN(total_qty_ordered)
        FROM (
            SELECT SUM(od.quantity) AS total_qty_ordered
            FROM oe.order_detail od
            JOIN oe.products p ON p.product_id = od.product_id
            GROUP BY p.category_id
        ) y
    )
ORDER BY t.total_qty_ordered DESC;