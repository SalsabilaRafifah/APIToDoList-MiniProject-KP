//Use case bertanggung jawab untuk mengimplementasikan logika bisnis aplikasi dan menjadi perantara antara lapisan domain dan lapisan infrastruktur.

package usecase

import "github.com/salsabilarafifah/API-ToDoList/internal/domain"

//interface yang mendefinisikan kontrak untuk operasi CRUD pada entitas Todo.
//sebagai dependensi untuk use case.
type TodoRepository interface {
	Create(todo *domain.Todo) error
	GetAll() ([]domain.Todo, error)
	GetByID(id uint) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id uint) error
	MarkAsCompleted(id uint) error
	GetCompleted() ([]domain.Todo, error)
}

//interface yang mendefinisikan kontrak untuk operasi bisnis pada entitas Todo.
//sebagai abstraksi antarmuka bagi use case yang akan diimplementasikan.
type TodoUseCase interface {
	Create(todo *domain.Todo) error
	GetAll() ([]domain.Todo, error)
	GetByID(id uint) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id uint) error
	MarkAsCompleted(id uint) error
	GetCompleted() ([]domain.Todo, error)
}

//implementasi konkret dari TodoUseCase.
//Struct menyimpan instance dari TodoRepository sebagai properti.
type todoUseCase struct {
	todoRepository TodoRepository
}

//membuat instance baru dari todoUseCase dengan dependensi TodoRepository.
func NewTodoUseCase(todoRepository TodoRepository) TodoUseCase {
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