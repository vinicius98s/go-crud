package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User represents an instance of a user
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

var collection *mongo.Collection

// UsersCollection handle the users collection
func UsersCollection(c *mongo.Database) {
	collection = c.Collection("users")
}

func parseID(c *gin.Context) (primitive.ObjectID, error) {
	id := c.Param("id")
	parsedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "The user ID is not valid",
		})
		return parsedID, err
	}
	return parsedID, nil
}

// ListUsers list all users available
func ListUsers(c *gin.Context) {
	var users []*User
	cursor, err := collection.Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		fmt.Printf("Error getting users: %s \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	for cursor.Next(context.TODO()) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Printf("Failed to decode user: %s \n", err)
		}

		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		fmt.Printf("Cursor error: %s \n", err)
	}

	cursor.Close(context.TODO())
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

// CreateUser creates a user in database
func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	name, email, password := user.Name, user.Email, user.Password
	newUser := User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := collection.InsertOne(context.TODO(), newUser)

	if err != nil {
		fmt.Printf("Failed to create user: %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	data := map[string]interface{}{
		"_id":      createdUser.InsertedID,
		"name":     name,
		"email":    email,
		"password": password,
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data":   data,
	})
}

// UpdateUser updates the user by Id
func UpdateUser(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	var user User
	c.BindJSON(&user)
	updateUser := bson.M{}
	if user.Name != "" {
		updateUser["name"] = user.Name
	}
	if user.Email != "" {
		updateUser["email"] = user.Email
	}
	if user.Name != "" {
		updateUser["password"] = user.Password
	}

	filter := bson.M{"_id": id}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	update := bson.D{{Key: "$set", Value: updateUser}}

	var updatedDocument bson.M
	e := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)
	if e != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if e == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Could not find user",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   updatedDocument,
	})
}

// DeleteUser delete a user
func DeleteUser(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	filter := bson.M{"_id": id}
	opts := options.FindOneAndDelete()
	var deletedDocument bson.M
	e := collection.FindOneAndDelete(context.TODO(), filter, opts).Decode(&deletedDocument)
	if e != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if e == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Could not find user",
			})
			return
		}
	}

	c.JSON(http.StatusNoContent, nil)
}
