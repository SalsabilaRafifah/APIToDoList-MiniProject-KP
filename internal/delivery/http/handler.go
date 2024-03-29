package http

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/salsabilarafifah/API-ToDoList/internal/domain"
	"github.com/salsabilarafifah/API-ToDoList/internal/usecase"
)

//untuk menangani operasi CRUD pada TODO.
type TodoHandler struct {
	todoUseCase usecase.TodoUseCase
}

//konstruktor untuk membuat instance baru dari TodoHandler dengan dependensi TodoUsecase
func NewTodoHandler(todoUseCase usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{todoUseCase}
}

func (h *TodoHandler) Create(c *gin.Context) {
	// Struktur untuk mengambil data dari body request
	type TodoCreateInput struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
	}

	// Mengikat data JSON dari body request ke dalam struktur TodoCreateInput.
	var input TodoCreateInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": "error parsing request body"})
		return
	}

	// Menghapus spasi di awal dan di akhir input title dan description
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	// Menambahkan validasi bahwa title harus diisi
	if input.Title == "" {
		c.JSON(400, gin.H{"error": "title is required"})
		return
	}

	// Membuat objek Todo dengan nilai default Completed: false
	todo := &domain.Todo{
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
	}

	// Memanggil Create dari todoUseCase untuk menyimpan Todo ke dalam penyimpanan data.
	if err := h.todoUseCase.Create(todo); err != nil {
		c.JSON(500, gin.H{"error": "error creating todo"})
		return
	}

	c.JSON(201, gin.H{"message": "todo created successfully", "todo": todo})
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	//untuk mendapatkan semua entitas Todo dari penyimpanan data.
	todos, err := h.todoUseCase.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": "error getting todos"})
		return
	}

	c.JSON(200, gin.H{"todos": todos})
}

func (h *TodoHandler) GetByID(c *gin.Context) {
	//mendapatkan entitas Todo berdasarkan ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid todo ID"})
		return
	}

	todo, err := h.todoUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "todo not found"})
		return
	}

	c.JSON(200, gin.H{"todo": todo})
}

func (h *TodoHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid todo ID"})
		return
	}

	// Mendapatkan todo dari database berdasarkan ID
	existingTodo, err := h.todoUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "todo not found"})
		return
	}

	// mengambil data JSON dari body permintaan HTTP dan mengubahnya menjadi objek Todo
	var updateData domain.Todo
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "error parsing request body"})
		return
	}

	// Update atribut sesuai data yang diterima
	if updateData.Title != "" {
		existingTodo.Title = updateData.Title
	}
	if updateData.Description != "" {
		existingTodo.Description = updateData.Description
	}

	// Update todo
	if err := h.todoUseCase.Update(existingTodo); err != nil {
		c.JSON(500, gin.H{"error": "error updating todo"})
		return
	}

	c.JSON(200, gin.H{"message": "todo updated successfully", "todo": existingTodo})
}

func (h *TodoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid todo ID"})
		return
	}

	if err := h.todoUseCase.Delete(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": "error deleting todo"})
		return
	}

	c.JSON(200, gin.H{"message": "todo deleted successfully"})
}

func (h *TodoHandler) MarkAsCompleted(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid todo ID"})
		return
	}

	// Dapatkan todo dari database berdasarkan ID
	existingTodo, err := h.todoUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "todo not found"})
		return
	}

	// Menandai tugas sebagai selesai hanya dengan memperbarui atribut Completed
	existingTodo.Completed = !existingTodo.Completed
	existingTodo.CompletedAt = time.Now()

	// Update todo untuk menyimpan perubahan
	if err := h.todoUseCase.Update(existingTodo); err != nil {
		c.JSON(500, gin.H{"error": "error updating todo"})
		return
	}

	c.JSON(200, gin.H{"message": "todo marked as completed successfully", "todo": existingTodo})
}

//Memanggil GetCompleted dari todoUseCase untuk mendapatkan daftar Todo yang sudah selesai.
func (h *TodoHandler) GetCompleted(c *gin.Context) {
	completedTodos, err := h.todoUseCase.GetCompleted()
	if err != nil {
		c.JSON(500, gin.H{"error": "error getting completed todos"})
		return
	}

	if len(completedTodos) == 0 {
		c.JSON(200, gin.H{"message": "No completed todos found"})
		return
	}

	c.JSON(200, gin.H{"completedTodos": completedTodos})
}

func (h *TodoHandler) GetUnCompleted(c *gin.Context) {
    unCompletedTodos, err := h.todoUseCase.GetUnCompleted()
    if err != nil {
        c.JSON(500, gin.H{"error": "error getting uncompleted todos"})
        return
    }

    if len(unCompletedTodos) == 0 {
        c.JSON(200, gin.H{"message": "No uncompleted todos found"})
        return
    }

    c.JSON(200, gin.H{"unCompletedTodos": unCompletedTodos})
}

func (h *TodoHandler) SearchByTitle(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		c.JSON(400, gin.H{"error": "title parameter is required"})
		return
	}

	// Memanggil metode SearchByTitle dari todoUseCase
	todos, err := h.todoUseCase.SearchByTitle(title)
	if err != nil {
		c.JSON(500, gin.H{"error": "error searching todos by title"})
		return
	}

	if len(todos) == 0 {
		c.JSON(200, gin.H{"message": "No todos found with the specified title"})
		return
	}

	c.JSON(200, gin.H{"todos": todos})
}