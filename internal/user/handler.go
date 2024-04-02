package user

import (
	"chatapp/utl"
	"fmt"
	"github.com/gin-gonic/gin"
)
type Handler struct {
	repo *Repository
}
func NewHandler(r *Repository)*Handler{
	return &Handler{repo : r}

}
func (h *Handler) Signup(c *gin.Context) {
	//get new user infromation
	userReq := new(CreateUserReq)
	err := c.ShouldBindJSON(userReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := userReq.Validate(); err != nil {
		c.JSON(402, gin.H{})
		return
	}
	//check if user exists
	fmt.Println("2.1")

	exists,err := h.repo.UserExists(userReq.Username)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if exists {
		fmt.Println(exists)
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}
	fmt.Println("4")
	//hashpassword
	bPass, err := utl.HashPassword([]byte(userReq.Password))
	if err != nil {
		c.Status(500)
		return
	}
	userReq.Password = string(bPass)
	//create user
	if err := h.repo.CreateUser(userReq); err != nil {
		c.Abort()
		return
	}
	//genrate and send token
	token, err := utl.GenerateJWT(userReq.Username,"1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("7")
	c.SetCookie("jwt-token", token, 60, "/", "localhost:8080", false, true)
	c.JSON(201, gin.H{"msg": "user created successfully!"})
}


func(h *Handler) Login(c *gin.Context){
	
	userLogin := new(Login)
	if err := c.ShouldBindJSON(userLogin); err != nil {
		c.Abort()
	}
	//check if user  exists
	userExists, err := h.repo.UserExists(userLogin.Username)
	if err != nil {
		c.Abort()
	}
	if !userExists{
		c.JSON(500,gin.H{"msg":"username or password mismatch"})
		return
	}
	//get user password from database
	user, err := h.repo.GetUser(userLogin.Username)
	if err != nil {
		c.Abort()
	}
	fmt.Println(user.Password)
	fmt.Println(userLogin.Password)
	fmt.Println("SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS!!!!!!!!!!!!!!!!!!!!!!!!!!")
	//compare user password and hashed password
	if err := utl.ComparePasswords(userLogin.Password,user.Password); err != nil {
		c.JSON(500,gin.H{"msg":"username or password mismatch"})
		return
		
	}
	//send cookie based jwt
	token, err := utl.GenerateJWT(userLogin.Username,"1")
	if err != nil {
		fmt.Println(err)
		c.JSON(501,gin.H{"msg":"something bad happnede!"})
		return
	}
	c.JSON(200,user)
	c.SetCookie("jwt-token", token, 60, "/", "localhost:8080", false, true)
} 