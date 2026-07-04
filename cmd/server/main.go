package main

import (
	"fmt"
	"go-agenda/internal/handler"
	"go-agenda/internal/repository"
	"net/http"
)

func main() {
	repo := repository.NewTaskRepository()
	taskHandler := handler.NewTaskHandler(repo)

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodPost:
				taskHandler.HandleCreate(w, r)
			case http.MethodGet:
				taskHandler.HandleList(w, r)
			default:
				http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Servidor rodando com sucesso na porta 8080...")
	err := http.ListenAndServe(":8080", nil);
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)	
	}
}
