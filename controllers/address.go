package controllers

import (
	"context"
	"net/http"

	"github.com/ammar-nousher-ali/ecommerce/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAddress() gin.HandlerFunc {

}

func EditHomeAddress() gin.HandlerFunc {

}

func EditWorkAddress() gin.HandlerFunc {

}

func DeleteAddress() gin.HandlerFunc {

	return func(c *gin.Context) {

		user_id := c.Query("id")

		if user_id == "" {

			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid search index"})
			c.Abort()
			return

		}

		addresses := make([]models.Address, 0)
		user_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal server error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.second)
		defer cancel()

		filter := bson.D{primitive.E{Key:"_id", Value: user_id}}
		update := bson.D{{Key:"$set", Value: bson:D{primitive.E{Key:"address", Value: addresses}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err!=nil {
			
			c.IndentedJSON(404, "wrong command")
			return
		}

		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "successfully deleted")



	}

}
