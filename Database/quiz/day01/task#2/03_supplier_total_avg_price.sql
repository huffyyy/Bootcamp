SELECT
    s.supplier_id,
    s.company_name,
    COUNT(p.product_id) AS total_product,
    to_char(AVG(p.unit_price), 'FM999999990.00') AS avg_unit_price
FROM oe.suppliers s
LEFT JOIN oe.products p
    ON p.supplier_id = s.supplier_id
GROUP BY
    s.supplier_id,
    s.company_name
ORDER BY
    total_product DESC,
    s.supplier_id;