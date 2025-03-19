package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const defaultUri = "localhost:8080"

type student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var students = []student{
	{ID: "1", Name: "Ana", Age: 18},
	{ID: "2", Name: "Ben", Age: 20},
	{ID: "3", Name: "Casey", Age: 22},
	{ID: "4", Name: "Denise", Age: 24},
	{ID: "5", Name: "Elmo", Age: 23},
}

func getStudent(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

func getStudentByID(c *gin.Context) {
	id := c.Param("id")
	for _, student := range students {
		if student.ID == id {
			c.JSON(http.StatusOK, student)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/students", getStudent)
	router.GET("/students/:id", getStudentByID)

	router.Run(defaultUri)
}
