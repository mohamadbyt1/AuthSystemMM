package shop

import "database/sql"

type Repository struct {
	db *sql.DB
}
func NewShopRepository(db *sql.DB) *Repository {
	return &Repository{db:db}
}

func (db *Repository)AddProduct(product *AddProduct)error{
	queryString := "INSERT INTO products(name,category,price,stock) VALUES ($1, $2, $3,$4)"
	_,err := db.db.Exec(queryString, product.Name,product.Category,product.Price,product.Stock)
	if err != nil {
		return err
	}
	return nil
}