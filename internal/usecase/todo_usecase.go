//Use case bertanggung jawab untuk mengimplementasikan logika bisnis aplikasi dan menjadi perantara antara lapisan domain dan lapisan infrastruktur.

package usecase

import (
	"github.com/salsabilarafifah/API-ToDoList/internal/domain"
	"github.com/salsabilarafifah/API-ToDoList/internal/repository"
)

// Interface lebih berguna ketika ada logika kompleks atau kebutuhan untuk bergantung pada berbagai implementasi.
// interface yang mendefinisikan kontrak untuk operasi bisnis pada entitas Todo.
// sebagai abstraksi antarmuka bagi use case yang akan diimplementasikan.
type TodoUseCase interface {
	Create(todo *domain.Todo) error
	GetAll() ([]domain.Todo, error)
	GetByID(id uint) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id uint) error
	MarkAsCompleted(id uint) error
	GetCompleted() ([]domain.Todo, error)
	GetUnCompleted() ([]domain.Todo, error)
	SearchByTitle(title string) ([]*domain.Todo, error)
}

// implementasi konkret dari TodoUseCase.
// Struct menyimpan instance dari TodoRepository sebagai properti.
type todoUseCase struct {
	todoRepository repository.TodoRepository
}

// Parameter todoRepository adalah instance dari interface TodoRepository digunakan sebagai dependensi untuk menghubungkan 'todoUseCase dengan lapisan repository
// Function konstruktor untuk membuat dan mengembalikan instance baru dari todoUseCase dengan tipe data TodoUseCase yang merupakan interface operasi bisnis pada Todo
func NewTodoUseCase(todoRepository repository.TodoRepository) TodoUseCase {
	return &todoUseCase{todoRepository}
}

func (uc *todoUseCase) Create(todo *domain.Todo) error {
	return uc.todoRepository.Create(todo)
}

func (uc *todoUseCase) GetAll() ([]domain.Todo, error) {
	return uc.todoRepository.GetAll()
}

func (uc *todoUseCase) GetByID(id uint) (*domain.Todo, error) {
	return uc.todoRepository.GetByID(id)
}

func (uc *todoUseCase) Update(todo *domain.Todo) error {
	return uc.todoRepository.Update(todo)
}

func (uc *todoUseCase) Delete(id uint) error {
	return uc.todoRepository.Delete(id)
}

func (uc *todoUseCase) MarkAsCompleted(id uint) error {
	return uc.todoRepository.MarkAsCompleted(id)
}

func (uc *todoUseCase) GetCompleted() ([]domain.Todo, error) {
	return uc.todoRepository.GetCompleted()
}

func (uc *todoUseCase) GetUnCompleted() ([]domain.Todo, error) {
	return uc.todoRepository.GetUnCompleted()
}

func (uc *todoUseCase) SearchByTitle(title string) ([]*domain.Todo, error) {
	return uc.todoRepository.SearchByTitle(title)
}
