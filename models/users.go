package models

import (
	"Api-Aula1/security"
	"Api-Aula1/utils"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type Users struct {
	ID       int8   `json:"id"`
	Name     string `json:"nome_usuario"`
	CPF      string `json:"cpf"`
	Email    string `json:"email_usuario"`
	Password string `json:"senha"`
}

// Prepare executa em duas etapas:
// 1. validate()
// 2. format()
func (u *Users) Prepare(step string) error {
	if err := u.validate(step); err != nil {
		return err
	}

	if err := u.format(step); err != nil {
		return err
	}

	return nil
}

// Validar campos obrigatórios e formatos
func (u *Users) validate(step string) error {

	if u.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if u.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if err := utils.CPFValidator(u.CPF); err != nil {
		return err
	}

	if step == "create" && u.Password == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

// Formatar dados e aplicar hash na senha quando for criação de usuário
func (u *Users) format(step string) error {

	// Remover espaços
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.CPF = strings.TrimSpace(u.CPF)

	// Padronizar para minúsculo
	u.Name = strings.ToLower(u.Name)
	u.Email = strings.ToLower(u.Email)

	// Criptografar senha apenas no CREATE
	if step == "create" {
		hashedPassword, err := security.Hash(u.Password)
		if err != nil {
			return err
		}
		u.Password = hashedPassword
	}

	return nil
}
