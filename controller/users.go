package controller

import (
	"Api-Aula1/models"
	"Api-Aula1/persistency"
	"Api-Aula1/repository"
	"Api-Aula1/responses"
	"Api-Aula1/security"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Lê o request.Body
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer r.Body.Close()

	// Descompacta (Unmarshal) o conteúdo JSON em uma Struct
	var newUser models.Users
	if err = json.Unmarshal(bodyRequest, &newUser); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	// Chama os métodos de preparação do User
	if err = newUser.Prepare("create"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	// Conexão com o banco
	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Instanciar Repo
	repo := repository.NewUsersRepo(db)
	newUser.ID, err = repo.Create(newUser)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, newUser)
}

// ==========================
// LOGIN
// ==========================

// Struct específica para o body do login
type loginInput struct {
	Email    string `json:"email_usuario"`
	Password string `json:"senha"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Lê o body da requisição
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer r.Body.Close()

	var input loginInput
	if err = json.Unmarshal(bodyRequest, &input); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	// Validação básica
	input.Email = strings.TrimSpace(strings.ToLower(input.Email))
	if input.Email == "" || input.Password == "" {
		responses.Err(w, http.StatusBadRequest, errors.New("e-mail e senha são obrigatórios"))
		return
	}

	// Conecta no banco
	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUsersRepo(db)

	// Busca usuário pelo e-mail
	userFromDB, err := repo.FindByEmail(input.Email)
	if err != nil {
		// Não expõe se foi e-mail ou senha para segurança
		responses.Err(w, http.StatusUnauthorized, errors.New("usuário ou senha inválidos"))
		return
	}

	// Compara a senha enviada com o hash salvo no banco
	if err := security.Verify(userFromDB.Password, input.Password); err != nil {
		// Se der erro aqui, senha está errada
		responses.Err(w, http.StatusUnauthorized, errors.New("usuário ou senha inválidos"))
		return
	}

	// Login OK – por segurança, não devolve o hash da senha
	userFromDB.Password = ""

	responses.JSON(w, http.StatusOK, userFromDB)
}

func FetchUser(writer http.ResponseWriter, request *http.Request) {

}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {

}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {}
