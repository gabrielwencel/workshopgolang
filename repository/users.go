package repository

import (
	"Api-Aula1/models"
	"database/sql"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

func (u UsersRepo) Create(user models.Users) (int8, error) {
	query := `INSERT INTO treehousedb.users(
                    name,
                    email,
                    password,
                    cpf
                ) VALUES (?,?,?,?)`

	statement, err := u.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Password, user.CPF)
	if err != nil {
		return 0, err
	}

	lastid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int8(uint64(lastid)), nil
}

// 🔹 NOVO: buscar usuário pelo e-mail (para o login)
func (u UsersRepo) FindByEmail(email string) (models.Users, error) {
	var user models.Users

	query := `SELECT 
                    id,
                    name,
                    cpf,
                    email,
                    password
              FROM treehousedb.users
              WHERE email = ?
              LIMIT 1`

	row := u.db.QueryRow(query, email)

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.CPF,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}
