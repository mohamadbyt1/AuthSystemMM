package shop

import "github.com/gin-gonic/gin"

type Handler struct {
	db *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{
		db: repo,
	}
}
func (h *Handler) AddProduct(c *gin.Context) {
	addProduct := new(AddProduct)
	err := c.ShouldBindJSON(&addProduct)
	if err != nil {
		c.Abort()
		return
	}
	err = h.db.AddProduct(addProduct)
	if err != nil {
		c.JSON(400,gin.H{"msg":"product not added!"})
	
	}
	c.JSON(201,gin.H{"msg":"product added"})
	

	

}
func (h *Handler) GetAllProducts(c *gin.Context) {

}
func (h *Handler) GetProductById(c *gin.Context) {

}
func (h *Handler) DeleateProduct(c *gin.Context) {

}