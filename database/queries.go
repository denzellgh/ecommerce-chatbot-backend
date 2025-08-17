package database

import "strings"

func GetProductsQuery(query string) string {
	var sqlQuery string

	query = strings.ToLower(query)

	if strings.Contains(query, "laptop") || strings.Contains(query, "computer") {
		sqlQuery = `SELECT p.id, p.name, p.brand, p.category_id, p.price, p.stock_quantity, p.description, p.specs 
                   FROM products p JOIN categories c ON p.category_id = c.id 
                   WHERE c.name ILIKE '%laptop%' OR c.name ILIKE '%computer%' 
                   AND p.deleted_at IS NULL ORDER BY p.stock_quantity DESC LIMIT 10`
	} else if strings.Contains(query, "phone") || strings.Contains(query, "mobile") {
		sqlQuery = `SELECT p.id, p.name, p.brand, p.category_id, p.price, p.stock_quantity, p.description, p.specs 
                   FROM products p JOIN categories c ON p.category_id = c.id 
                   WHERE c.name ILIKE '%phone%' 
                   AND p.deleted_at IS NULL ORDER BY p.stock_quantity DESC LIMIT 10`
	} else if strings.Contains(query, "gaming") || strings.Contains(query, "game") {
		sqlQuery = `SELECT p.id, p.name, p.brand, p.category_id, p.price, p.stock_quantity, p.description, p.specs 
                   FROM products p JOIN categories c ON p.category_id = c.id 
                   WHERE c.name ILIKE '%gaming%' OR p.name ILIKE '%gaming%'
                   AND p.deleted_at IS NULL ORDER BY p.stock_quantity DESC LIMIT 10`
	} else if strings.Contains(query, "audio") || strings.Contains(query, "headphone") || strings.Contains(query, "speaker") {
		sqlQuery = `SELECT p.id, p.name, p.brand, p.category_id, p.price, p.stock_quantity, p.description, p.specs 
                   FROM products p JOIN categories c ON p.category_id = c.id 
                   WHERE c.name ILIKE '%audio%' OR p.name ILIKE '%audio%' OR p.name ILIKE '%headphone%'
                   AND p.deleted_at IS NULL ORDER BY p.stock_quantity DESC LIMIT 10`
	} else {
		sqlQuery = `SELECT id, name, brand, category_id, price, stock_quantity, description, specs 
                   FROM products WHERE deleted_at IS NULL 
                   ORDER BY stock_quantity DESC LIMIT 15`
	}

	return sqlQuery
}
