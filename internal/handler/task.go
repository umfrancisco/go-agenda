package handler

import (
	"encoding/json"
	"go-agenda-agil/internal/repository"
	"net/http"
)

type TaskHandler struct {
	Repo *repository.TaskRepository
}

func NewTaskHandler(repo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{Repo: repo}
}

func (h *TaskHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return	
	}

	tasks := h.Repo.GetAll()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.newEncoder(w).Encode(tasks)
}
