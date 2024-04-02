package main

import (
	"chatapp/db"
	"chatapp/internal/http"
	"chatapp/internal/shop"
	"chatapp/internal/user"
	"log"
)
func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
		return
	}
	//make two isntanse for each userRepo and shopRepo
	userRepository := user.NewUserRepository(db.GetDB())
	shopRepository := shop.NewShopRepository(db.GetDB())
	userHandler := user.NewHandler(userRepository)
	shopHandler := shop.NewHandler(shopRepository)
	http.InitRouter(userHandler,shopHandler)
	http.Start("127.0.0.1:8080")
}

