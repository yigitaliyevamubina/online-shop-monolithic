package main

import (
	"fmt"
	"log"
	"net/http"
	"online_shop/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.ListUsers).Methods("GET")            //get list
	router.HandleFunc("/users/get", handlers.GetUser).Methods("GET")          //get single
	router.HandleFunc("/users/create", handlers.CreateUser).Methods("POST")   //create
	router.HandleFunc("/users/update", handlers.UpdateUser).Methods("PUT")    //update
	router.HandleFunc("/users/delete", handlers.DeleteUser).Methods("DELETE") //delete

	router.HandleFunc("/products", handlers.ListProducts).Methods("GET")            //get list
	router.HandleFunc("/product/get", handlers.GetProduct).Methods("GET")           //get single
	router.HandleFunc("/products/create", handlers.CreateProduct).Methods("POST")   //create
	router.HandleFunc("/products/update", handlers.UpdateProduct).Methods("PUT")    //update
	router.HandleFunc("/products/delete", handlers.DeleteProduct).Methods("DELETE") //delete

	router.HandleFunc("/products/add", handlers.AddProductToShoppingCart).Methods("POST")         //add product to the shopping cart or just create one
	router.HandleFunc("/products/remove", handlers.RemoveProductFromShoppingCart).Methods("POST") //remove product from the shopping cart or delete your shopping cart
	router.HandleFunc("/products/user", handlers.ListUsersProducts).Methods("GET")                     //get you shopping cart (provide user id)
	router.HandleFunc("/products/carts", handlers.ListShoppingCarts).Methods("GET")               // list all users' shopping carts

	fmt.Println("Listening on book-service:6060...")
	log.Fatal(http.ListenAndServe(":6060", router))
}
