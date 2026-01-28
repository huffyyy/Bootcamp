SELECT r.region_id, region_name, c.country_id, country_name,
       l.location_id, street_address, postal_code, state_province
FROM hr.regions r
JOIN hr.countries c ON r.region_id = c.region_id
JOIN hr.locations l ON l.country_id = c.country_id;