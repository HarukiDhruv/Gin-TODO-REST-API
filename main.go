package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Walking on Golang", Completed: false},
	{ID: "2", Item: "Drink Some Water", Completed: true},
	{ID: "3", Item: "Read a book", Completed: true},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func newTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(context *gin.Context) {
	id := context.Param("id")

	for _, todo := range todos {
		if todo.ID == id {
			context.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func updateTodoByID(context *gin.Context) {
	id := context.Param("id")
	var updatedTodo todo

	if err := context.BindJSON(&updatedTodo); err != nil {
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = updatedTodo
			context.IndentedJSON(http.StatusOK, updatedTodo)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func deleteTodoByID(context *gin.Context) {

	id := context.Param("id")

	for i, t := range todos {

		if t.ID == id {

			todos = append(todos[:i], todos[i+1:]...)

			context.JSON(http.StatusOK, gin.H{
				"message": "Todo deleted",
			})
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{
		"error": "Todo not found",
	})
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", newTodo)
	router.GET("/todos/:id", getTodoByID)
	router.DELETE("/todos/:id", deleteTodoByID)
	router.PUT("/todos/:id", updateTodoByID)
	router.Run("localhost:8080")
}
