package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"online_shop/helper"
	"online_shop/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//user crud
//product crud
//add products to shopping cart
//remove products from shopping cart
//list shopping carts
//get shopping cart by user id

var database = helper.ConnectToMongoDB()

// Insert user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> CreateUser decoding json")
		return
	}

	user.ID = uuid.NewString()
	_, err = database.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> CreateUser inserting user")
		return
	}

	filter := bson.M{"uuid": user.ID}
	err = database.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> CreateUser finding user")
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Get user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user models.User

	var params = mux.Vars(r)
	fmt.Println(params)

	filter := bson.M{"uuid": params["id"]}
	fmt.Println(params["id"])
	err := database.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> Get user")
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Update user by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var params = mux.Vars(r)

	var user models.User

	filter := bson.M{"uuid": params["id"]}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> Update user decoding json")
		return
	}

	updateReq := bson.M{
		"$set": bson.M{
			"fullName": user.FullName,
			"address":  user.Address,
		},
	}

	err = database.Collection("users").FindOneAndUpdate(context.Background(), filter, updateReq).Decode(&user)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> Update user (finding and updating)")
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Delete user by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var params = mux.Vars(r)

	filter := bson.M{"uuid": params["id"]}

	deleteResponse, err := database.Collection("users").DeleteOne(context.Background(), filter)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> Delete user")
		return
	}

	json.NewEncoder(w).Encode(deleteResponse)
}

// List users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)
	page := cast.ToInt64(params["page"])
	limit := cast.ToInt64(params["limit"])

	offset := (page - 1) * limit
	options := options.Find()

	options.SetSkip(int64(offset))
	options.SetLimit(int64(limit))

	cursor, err := database.Collection("users").Find(context.Background(), bson.D{}, options)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> List users")
		return
	}

	defer cursor.Close(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User

		err = cursor.Decode(&user)
		if err != nil {
			log.Println(err)
			return
		}

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// Insert product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> Create product decoding json")
		return
	}

	product.ID = uuid.NewString()
	_, err = database.Collection("products").InsertOne(context.Background(), product)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> Create product inserting document")
		return
	}

	filter := bson.M{"uuid": product.ID}
	err = database.Collection("products").FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> Create product getting product")
		return
	}

	json.NewEncoder(w).Encode(product)
}

// Get product by id
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var product models.Product

	var params = mux.Vars(r)

	filter := bson.M{"uuid": params["id"]}
	err := database.Collection("products").FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> Get product")
		return
	}

	json.NewEncoder(w).Encode(product)
}

// Update product by id
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var params = mux.Vars(r)

	var product models.Product

	filter := bson.M{"uuid": params["id"]}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> Update product")
		return
	}

	updateReq := bson.M{
		"$set": bson.M{
			"name":  product.Name,
			"price": product.Price,
		},
	}

	err = database.Collection("products").FindOneAndUpdate(context.Background(), filter, updateReq).Decode(&product)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> Update product (finding and updating)")
		return
	}

	json.NewEncoder(w).Encode(product)
}

// Delete product by id
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var params = mux.Vars(r)

	filter := bson.M{"uuid": params["id"]}

	deleteResponse, err := database.Collection("products").DeleteOne(context.Background(), filter)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> Delete product")
		return
	}

	json.NewEncoder(w).Encode(deleteResponse)
}

// List products
func ListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := r.URL.Query()
	page := cast.ToInt64(params.Get("page"))
	limit := cast.ToInt64(params.Get("limit"))

	offset := (page - 1) * limit
	options := options.Find()

	options.SetSkip(int64(offset))
	options.SetLimit(int64(limit))

	cursor, err := database.Collection("products").Find(context.Background(), bson.D{}, options)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> List products")
		return
	}

	defer cursor.Close(context.Background())

	var products []models.Product
	for cursor.Next(context.Background()) {
		var product models.Product

		err = cursor.Decode(&product)
		if err != nil {
			log.Println(err)
			return
		}

		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(products)
}

// Add products to shopping cart
func AddProductToShoppingCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var shoppingCart models.ShoppingCart
	err := json.NewDecoder(r.Body).Decode(&shoppingCart)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> AddProductToShoppingCart -> decoding json")
		return
	}

	shoppingCart.TotalAmount = 0
	filter := bson.M{"userId": shoppingCart.UserId}
	var result models.ShoppingCart

	for _, productId := range shoppingCart.ProductIds {
		var product models.Product

		filter := bson.M{"uuid": productId}
		err := database.Collection("products").FindOne(context.Background(), filter).Decode(&product)
		if err != nil {
			helper.HandleError(err, w, http.StatusInternalServerError, "-> AddProductToShoppingCart -> finding one")
			return
		}
		shoppingCart.TotalAmount += product.Price
	}

	check := database.Collection("shoppingcarts").FindOne(context.Background(), filter)
	if check.Err() == mongo.ErrNoDocuments {
		_, err := database.Collection("shoppingcarts").InsertOne(context.Background(), shoppingCart)
		if err != nil {
			helper.HandleError(err, w, http.StatusInternalServerError, "-> AddProductToShoppingCart -> inserting one")
			return
		}
	} else {
		err = check.Decode(&result)
		if err != nil {
			helper.HandleError(err, w, http.StatusInternalServerError, "-> AddProductToShoppingCart -> decoding")
			return
		}

		result.ProductIds = append(result.ProductIds, shoppingCart.ProductIds...)
		updateReq := bson.M{
			"$set": bson.M{
				"productIds":  result.ProductIds,
				"totalAmount": result.TotalAmount + shoppingCart.TotalAmount,
			},
		}
		_, err := database.Collection("shoppingcarts").UpdateOne(context.Background(), bson.M{"userId": result.UserId}, updateReq)
		if err != nil {
			helper.HandleError(err, w, http.StatusInternalServerError, "-> AddProductToShoppingCart -> updating")
			return
		}
	}

	filter = bson.M{"userId": shoppingCart.UserId}
	err = database.Collection("shoppingcarts").FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> AddProductToShoppingCart -> getting for user")
		return
	}

	json.NewEncoder(w).Encode(result)
}

// Remove products shopping cart
func RemoveProductFromShoppingCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var shoppingCart models.ShoppingCart
	err := json.NewDecoder(r.Body).Decode(&shoppingCart)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> RemoveProductFromShoppingCart -> decoding json")
		return
	}

	shoppingCart.TotalAmount = 0
	filter := bson.M{"userId": shoppingCart.UserId}

	var result models.ShoppingCart
	check := database.Collection("shoppingcarts").FindOne(context.Background(), filter)
	if check.Err() == mongo.ErrNoDocuments {
		helper.HandleError(fmt.Errorf("this user did not buy anything yet"), w, http.StatusOK, "no shopping cart belongs to user")
		return
	} else {
		err = check.Decode(&result)
		if err != nil {
			helper.HandleError(err, w, http.StatusInternalServerError, "-> RemoveProductFromShoppingCart -> decoding")
			return
		}

		for _, idToRemove := range shoppingCart.ProductIds {
			for i, id := range result.ProductIds {
				if id == idToRemove {
					var product models.Product
					filter := bson.M{"uuid": idToRemove}
					err := database.Collection("products").FindOne(context.Background(), filter).Decode(&product)
					if err != nil {
						helper.HandleError(err, w, http.StatusInternalServerError, "-> RemoveProductFromShoppingCart -> finding product")
						return
					}
					result.TotalAmount -= product.Price
					result.ProductIds = append(result.ProductIds[:i], result.ProductIds[i+1:]...)
					break
				}
			}
		}

		if len(result.ProductIds) == 0 {
			_, err := database.Collection("shoppingcarts").DeleteOne(context.Background(), filter)
			if err != nil {
				helper.HandleError(err, w, http.StatusInternalServerError, "-> RemoveProductFromShoppingCart -> deleting one")
				return
			}

			helper.HandleError(fmt.Errorf("ops, you have no products left in shopping cart"), w, http.StatusOK, "no product left")
			return
		}

		updateReq := bson.M{
			"$set": bson.M{
				"productIds":  result.ProductIds,
				"totalAmount": result.TotalAmount,
			},
		}
		_, err := database.Collection("shoppingcarts").UpdateOne(context.Background(), bson.M{"userId": result.UserId}, updateReq)
		if err != nil {
			helper.HandleError(err, w, http.StatusInternalServerError, "-> RemoveProductFromShoppingCart -> updating")
			return
		}
	}

	err = database.Collection("shoppingcarts").FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> RemoveProductFromShoppingCart -> finding")
		return
	}

	json.NewEncoder(w).Encode(result)
}

// List products in shopping cart by user id
func ListUsersProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	userID := params["user_id"]

	filter := bson.M{"userId": userID}
	var respUser models.ShoppingCart
	err := database.Collection("shoppingcarts").FindOne(context.Background(), filter).Decode(&respUser)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> ListUsersProducts")
		return
	}

	json.NewEncoder(w).Encode(respUser)
}

// List shopping carts
func ListShoppingCarts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	page := cast.ToInt64(params.Get("page"))
	limit := cast.ToInt64(params.Get("limit"))

	offset := (page - 1) * limit

	sortBy := bson.D{{Key: "totalAmount", Value: -1}} //from min to max
	options := options.Find().SetSort(sortBy)
	options.SetSkip(int64(offset))
	options.SetLimit(limit)

	cursor, err := database.Collection("shoppingcarts").Find(context.Background(), bson.D{}, options)
	if err != nil {
		helper.HandleError(err, w, http.StatusInternalServerError, "-> ListShoppingCarts")
		return
	}

	fmt.Println(cursor.ID())
	defer cursor.Close(context.Background())

	var shoppingCarts []models.ShoppingCart
	for cursor.Next(context.Background()) {
		var shoppingCart models.ShoppingCart

		err = cursor.Decode(&shoppingCart)
		if err != nil {
			log.Println(err)
			return
		}

		shoppingCarts = append(shoppingCarts, shoppingCart)
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(shoppingCarts)
}