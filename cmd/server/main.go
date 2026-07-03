package main

import (
	"fmt"
	"go-agenda-agil/internal/handler"
	"go-agenda-agil/internal/repository"
	"net/http"
)

func main() {
	repo := repository.NewTaskRepository()
	taskHandler := handler.NewTaskHandler(repo)

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.method {
			case http.MethodPost:
				taskHandler.HandleCreate(w, r)
			case httpMethodGet:
				taskHandler.HandleList(w, r)
			default:
				http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	})

	fmt.Println("Servidor rodando com sucesso na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)	
	}
}
