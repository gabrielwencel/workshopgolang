package controller

import (
	"Api-Aula1/models"
	"Api-Aula1/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// repositório global em memória
var usersRepo = repository.NewUsersRepository()

// --------- CREATE (POST /users) ---------

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Print("CreateUser chamado")

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var user models.Users
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err := user.Prepare("create"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := usersRepo.Create(user)
	if err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

// --------- READ (GET /users) ---------

func FetchUser(w http.ResponseWriter, r *http.Request) {
	log.Print("FetchUser chamado")

	users, err := usersRepo.FindAll()
	if err != nil {
		http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// --------- UPDATE (PUT /users/{userID}) ---------

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Print("UpdateUser chamado")

	vars := mux.Vars(r)
	idStr := vars["userID"]

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var user models.Users
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	user.ID = id
	if err := user.Prepare("update"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := usersRepo.Update(id, user)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

// --------- DELETE (DELETE /users/{userID}) ---------

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteUser chamado")

	vars := mux.Vars(r)
	idStr := vars["userID"]

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := usersRepo.Delete(id); err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
