package main

import (
	"github.com/gin-gonic/gin" // menangani permintaan HTTP dan membangun API RESTful.
	"github.com/joho/godotenv" // membaca variabel lingkungan dari file .env

	"log"
	"os"

	"github.com/salsabilarafifah/API-ToDoList/internal/config"
	"github.com/salsabilarafifah/API-ToDoList/internal/delivery/http"
	"github.com/salsabilarafifah/API-ToDoList/internal/domain"
	"github.com/salsabilarafifah/API-ToDoList/internal/repository"
	"github.com/salsabilarafifah/API-ToDoList/internal/usecase"
)

// dieksekusi sebelum fungsi main untuk memuat variabel lingkungan menggunakan library godontev
func init() {
	err := godotenv.Load(".env"); 
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//menginisialisasi router Gin untuk menangani HTTP requests.
	r := gin.Default()

	//terhubung ke database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	//melakukan migrasi otomatis model Todo ke database
	db.AutoMigrate(&domain.Todo{})

	//menyiapkan rute HTTP
	todoRepository := repository.NewTodoRepository(db)
	todoUseCase := usecase.NewTodoUseCase(todoRepository)
	todoHandler := http.NewTodoHandler(todoUseCase)

	// Mendefinisikan grup rute untuk operasi CRUD pada Todos dengan awalan /todos/api/v1
	apiV1 := r.Group("/todos/api/v1")
	{
		apiV1.POST("/create", todoHandler.Create)
		apiV1.GET("", todoHandler.GetAll)
		apiV1.GET("/:id", todoHandler.GetByID)
		apiV1.PUT("/update/:id", todoHandler.Update)          // Ubah menjadi "/todos/:id"
		apiV1.DELETE("/delete/:id", todoHandler.Delete)       // Ubah menjadi "/todos/:id"
		apiV1.PUT("/:id/completed", todoHandler.MarkAsCompleted)
		apiV1.GET("/completed", todoHandler.GetCompleted)
		apiV1.GET("/uncompleted", todoHandler.GetUnCompleted)
		apiV1.GET("/search/:title", todoHandler.SearchByTitle)
	}

	//menjalankan aplikasi Gin pada port 3000 sebagai default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(r.Run(":" + port))
}
