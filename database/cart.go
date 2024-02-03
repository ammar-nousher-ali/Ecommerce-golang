package database

import (
	"context"
	"errors"
	"log"

	"github.com/ammar-nousher-ali/go-ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrCantFindProduct   = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't find the product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErroCantUpdateUser    = errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item form cart")
	ErrCantGetItem        = errors.New("was unable to get item from the cart")
	ErrCanBuyCartItem     = errors.New("cannot update the purchase")
)

func AddProductToCart(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {

	searchfromdb, err := prodCollection.Find(ctx, bson.M{"_id": productID})
	if err!=nil {
		log.Println(err)
		return ErrCantFindProduct
		
	}

	var productCart []models.ProductUser
	err = searchfromdb.All(ctx, &productCart)
	if err!=nil {
		log.Println(err)
		return ErrCantDecodeProducts
		
	}


	id, err := primitive.ObjectIDFromHex(userID)
	if err!=nil {

		log.Println(err)
		return ErrUserIdIsNotValid
		
	}


	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "usercart", Value: bson.D{{Key: "$each", Value: productCart}}}}}}


	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err!=nil {

		return ErroCantUpdateUser
		
	}

	return nil




}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
