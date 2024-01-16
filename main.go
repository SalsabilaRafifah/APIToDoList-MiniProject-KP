package main

import (
	"github.com/gin-gonic/gin" // menangani permintaan HTTP dan membangun API RESTful.
	"github.com/joho/godotenv"

	"log"
	"os"

	"github.com/salsabilarafifah/API-ToDoList/internal/config"
	"github.com/salsabilarafifah/API-ToDoList/internal/delivery/http"
	"github.com/salsabilarafifah/API-ToDoList/internal/domain"
	"github.com/salsabilarafifah/API-ToDoList/internal/repository"
	"github.com/salsabilarafifah/API-ToDoList/internal/usecase"
)

// dieksekusi sebelum fungsi main memuat variabel lingkungan
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//menginisialisasi router Gin
	r := gin.Default()

	//terhubung ke database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	//melakukan migrasi model Todo ke database
	db.AutoMigrate(&domain.Todo{})

	//menyiapkan rute HTTP
	todoRepository := repository.NewTodoRepository(db)
	todoUseCase := usecase.NewTodoUseCase(todoRepository)
	todoHandler := http.NewTodoHandler(todoUseCase)

	//mendefinisikan rute untuk operasi CRUD pada Todos
	r.POST("/todos", todoHandler.Create)
	r.GET("/todos", todoHandler.GetAll)
	r.GET("/todos/:id", todoHandler.GetByID)
	r.PUT("/todos/:id", todoHandler.Update)
	r.DELETE("/todos/:id", todoHandler.Delete)
	r.PUT("/todos/:id/complete", todoHandler.MarkAsCompleted)
	r.GET("/todos/completed", todoHandler.GetCompleted)
	r.GET("/todos/uncompleted", todoHandler.GetUnCompleted)

	//menjalankan aplikasi Gin pada port 3000 sebagai default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(r.Run(":" + port))
}
