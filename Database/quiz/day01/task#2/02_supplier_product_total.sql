SELECT 
    s.supplier_id,
    s.company_name,
    COUNT(p.product_id) AS total_product
FROM oe.suppliers s
LEFT JOIN oe.products p ON p.supplier_id = s.supplier_id
GROUP BY 
    s.supplier_id,
    s.company_name
ORDER BY s.supplier_id;