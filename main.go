package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Eat food", Completed: false},
}

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newtodo todo
	if err := context.BindJSON(&newtodo); err != nil {
		return
	}

	todos = append(todos, newtodo)
	context.IndentedJSON(http.StatusCreated, newtodo)
}
func GetTodosbyid(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func toggltodostates(context *gin.Context) {
	id := context.Param("id")
	todo, err := GetTodosbyid(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func gettodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := GetTodosbyid(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", GetTodos)
	router.GET("/todos/:id", gettodo)
	router.PATCH("/todos/:id", toggltodostates)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
