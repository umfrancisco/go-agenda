package repository

import (
	"go-agenda-agil/internal/model"
	"sync"
	"time"
)

type TaskRepository struct {
	mu sync.Mutex
	tasks []model.Task
	nextID int
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository {
		tasks: []model.Task{},
		nextID: 1,
	}
}

func (r *TaskRepository) Create(title string) model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	task := model.Task {
		ID: r.nextID,
		Title: title,
		Done: false,
		CreatedAt: time.Now(),
	}

	r.tasks = append(r.tasks, task)
	r.nextID++
	return task
}

func (r *TaskRepository) GetAll() []model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.tasks
}
