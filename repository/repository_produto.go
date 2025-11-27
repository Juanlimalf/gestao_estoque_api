package repository

import (
	"database/sql"
	"fmt"
	"main/model"
)

type ProductRepository struct {
	db *sql.DB
}

func CreateProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{db: db}
}
func (r *ProductRepository) FetchAllProducts() []model.Product {

	result, error := r.db.Query("SELECT * from product")

	if error != nil {
		fmt.Println("Error fetching products:", error)
		return []model.Product{}
	}

	products := []model.Product{}
	valueTotal := 0.0

	for result.Next() {
		var product model.Product
		error := result.Scan(&product.ID, &product.Description, &product.Price)

		if error != nil {
			fmt.Println("Error scanning product:", error)
			continue
		}
		fmt.Println("Product:", product)
		products = append(products, product)
		valueTotal += product.Price
	}
	fmt.Println("Total value of products:", fmt.Sprintf("%.2f", valueTotal))

	result.Close()

	return products
}

func (r *ProductRepository) InsertProduct(product model.Product) error {
	query := "INSERT INTO product (description, price) VALUES (?, ?)"
	_, err := r.db.Exec(query, product.Description, product.Price)
	return err
}
