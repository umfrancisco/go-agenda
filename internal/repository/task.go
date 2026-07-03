package repository

import (
	"errors"
	"go-agenda-agil/internal/model"
	"sync"
	"time"
)

type TaskRepository struct {
	mu sync.Mutex
	tasks []model.Task
	nextID int
}
