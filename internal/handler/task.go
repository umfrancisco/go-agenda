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
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return	
	}

	var input struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Title == "" {
		http.Error(w, "Requisição inválida. O campo 'title' é obrigatório", http.StatusBadRequest)
		return
	}

	newTask := h.Repo.Create(input.Title)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func (h *TaskHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	tasks := h.Repo.GetAll()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}


