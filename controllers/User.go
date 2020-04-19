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

// User represents a isntance of a user
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

var collection *mongo.Collection

// UsersCollection : Makes the users collection
func UsersCollection(c *mongo.Database) {
	collection = c.Collection("users")
}

// ListUsers all users available
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

// CreateUser : creates a user in database
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

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data":   createdUser,
	})
}
