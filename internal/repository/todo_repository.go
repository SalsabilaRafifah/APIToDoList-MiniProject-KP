//repository yang bertanggung jawab untuk mengakses dan memanipulasi data entitas Todo pada database menggunakan ORM GORM.

package repository

import (
	"time"
	"github.com/salsabilarafifah/API-ToDoList/internal/domain"
	"gorm.io/gorm"
)

//Interface lebih berguna ketika ada logika kompleks atau kebutuhan untuk bergantung pada berbagai implementasi.
//interface yang mendefinisikan kontrak untuk operasi CRUD pada entitas Todo
type TodoRepository interface {
	Create(todo *domain.Todo) error
	GetAll() ([]domain.Todo, error)
	GetByID(id uint) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id uint) error
	MarkAsCompleted(id uint) error
	GetCompleted() ([]domain.Todo, error)
	GetUnCompleted() ([]domain.Todo, error)
}

//implementasi konkret dari TodoRepository.
//Struct menyimpan instance dari database GORM sebagai properti.
//menyimpan koneksi database GORM.
type todoRepository struct {
	db *gorm.DB
}

//Membuat instance baru dari todoRepository dengan menyediakan koneksi database GORM.
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) Create(todo *domain.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) GetAll() ([]domain.Todo, error) {
	var todos []domain.Todo //menggunakan tipe data slice
	//&todos adalah alamat memori dari slice todos, sehingga data hasil query dapat langsung dimasukkan ke dalam slice tersebut.
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *todoRepository) GetByID(id uint) (*domain.Todo, error) {
	//mengambil satu objek Todo berdasarkan ID
	var todo domain.Todo
	//untuk mengambil satu entitas pertama yang memiliki ID sesuai dengan nilai id yang diberikan
	//Data entitas tersebut kemudian dimasukkan ke variabel todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *todoRepository) Update(todo *domain.Todo) error {
	//menyimpan perubahan pada objek todo ke dalam database.
	return r.db.Save(todo).Error
}

//
func (r *todoRepository) Delete(id uint) error {
	//menentukan tabel atau model data yang akan dihapus
	//hanya entitas dengan ID yang sesuai yang akan dihapus
	return r.db.Delete(&domain.Todo{}, id).Error
}

//
func (r *todoRepository) MarkAsCompleted(id uint) error {
	//&domain.Todo{} menunjukkan model data atau tabel yang akan dimodelkan
	//memperbarui kolom completed dan completed_at dari record yang memiliki ID sesuai dengan parameter id dari tabel todos.
	return r.db.Model(&domain.Todo{}).Where("id = ?", id).Updates(map[string]interface{}{
		"completed":     true,
		"completed_at":  time.Now(),
	}).Error
}

func (r *todoRepository) GetCompleted() ([]domain.Todo, error) {
	var completedTodos []domain.Todo
	//mengekstrak semua record dari tabel todos yang memiliki completed = true ke dalam slice completedTodos.
	err := r.db.Where("completed = ?", true).Find(&completedTodos).Error
	return completedTodos, err
}

func (r *todoRepository) GetUnCompleted() ([]domain.Todo, error) {
    var unCompletedTodos []domain.Todo
    err := r.db.Where("completed = ?", false).Find(&unCompletedTodos).Error
    return unCompletedTodos, err
}