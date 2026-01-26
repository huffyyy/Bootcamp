SELECT
    s.shipper_id,
    s.company_name,
    p.product_id,
    p.product_name,
    SUM(od.quantity) AS total_qty_ordered
FROM oe.shippers s
JOIN oe.orders o
    ON o.ship_via = s.shipper_id
JOIN oe.order_detail od
    ON od.order_id = o.order_id
JOIN oe.products p
    ON p.product_id = od.product_id
GROUP BY
    s.shipper_id,
    s.company_name,
    p.product_id,
    p.product_name
ORDER BY
    s.shipper_id,
    total_qty_ordered DESC;