package shop
type Product struct {
	Id int
	Name string
	Category string
	Price float64
	Stock int 
}
type AddProduct struct {
	Name string `json:"name"`
	Category string `json:"category"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
}
type ProductStore struct {
	Products map[int]*Product
}
func NewProductStore() *ProductStore {
	return &ProductStore{
		Products: make(map[int]*Product),
	}
}
