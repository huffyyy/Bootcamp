SELECT 
    c.category_id,
    c.category_name,
    COUNT(p.product_id) AS total_product
FROM oe.categories c
LEFT JOIN oe.products p ON p.category_id = c.category_id
GROUP BY 
    c.category_id,
    c.category_name
ORDER BY c.category_id;