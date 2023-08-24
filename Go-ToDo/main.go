package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID          string `json:"id"`
	Task        string `json:"title"`
	Iscompleted bool   `json:"iscompleted"`
}

var todos = []todo{
	{ID: "1", Task: "Read Book", Iscompleted: false},
	{ID: "2", Task: "Go to School", Iscompleted: true},
	{ID: "1", Task: "Painting", Iscompleted: true},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodos)
	router.Run("localhost:8080")
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodos(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {

			return &todos[i], nil
		}
	}
	return nil, errors.New("todos not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	todo.Iscompleted = !todo.Iscompleted

	context.IndentedJSON(http.StatusOK, todo)
}
