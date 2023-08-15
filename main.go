package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{Name: "Ananda", Email: "ananda@gmail.com"},
	{Name: "Fineta", Email: "fineta@gmail.com"},
}

func GetUsers(c *gin.Context) {
	response := users
	c.JSON(http.StatusOK, response)
}

func GetDetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	response := users[id]
	c.JSON(http.StatusOK, response)
}

func AddUser(c *gin.Context) {
	var user User
	c.ShouldBindJSON(&user)
	users = append(users, user)

	response := users
	c.JSON(http.StatusOK, response)
}

func DeleteUser(c *gin.Context) {
	removeId, _ := strconv.Atoi(c.Param("id"))

	if removeId >= len(users) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Index User Not Found"})
		return
	}

	users = append(users[:removeId], users[removeId+1:]...) // Swap this index value with next index value
	response := users
	c.JSON(http.StatusOK, response)
}

func DeleteUserByName(c *gin.Context) {
	var filteredUsers []User
	removeName := c.Param("name")

	for _, user := range users {
		if user.Name != removeName {
			filteredUsers = append(filteredUsers, user)
		}
	}

	users = filteredUsers
	response := users
	c.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())
	api := router.Group("/api")

	api.GET("/", GetUsers)
	api.GET("/user/:id", GetDetailUser)
	api.POST("/user", AddUser)
	api.DELETE("/user/:id", DeleteUser)
	api.DELETE("/user/name/:name", DeleteUserByName)

	router.Run("localhost:8080")
}
