package models

import (
	"errors"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	ID       uint64 `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	CPF      string `json:"cpf,omitempty"`
	Password string `json:"password,omitempty"`
}

// step: "create", "update" etc.
func (u *Users) Prepare(step string) error {
	if err := u.format(step); err != nil {
		return err
	}

	if err := u.validate(step); err != nil {
		return err
	}

	return nil
}

func (u *Users) validate(step string) error {
	if step == "update" && u.ID == 0 {
		return errors.New("id é obrigatório para atualização")
	}

	if strings.TrimSpace(u.Name) == "" && step == "create" {
		return errors.New("nome é obrigatório")
	}

	if strings.TrimSpace(u.Email) == "" {
		return errors.New("email é obrigatório")
	}

	if strings.TrimSpace(u.CPF) == "" {
		return errors.New("cpf é obrigatório")
	}

	if !validarCPF(u.CPF) {
		return errors.New("cpf inválido")
	}

	if step == "create" && strings.TrimSpace(u.Password) == "" {
		return errors.New("senha é obrigatória")
	}

	return nil
}

func (u *Users) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.CPF = strings.TrimSpace(u.CPF)
	u.Password = strings.TrimSpace(u.Password)

	// na criação, gera hash da senha
	if step == "create" && u.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hash)
	}

	return nil
}

// ------------------ Validador de CPF ------------------

func validarCPF(cpf string) bool {
	// mantém só dígitos
	digits := make([]rune, 0, 11)
	for _, r := range cpf {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		}
	}

	if len(digits) != 11 {
		return false
	}

	// rejeita todos iguais (11111111111 etc)
	iguais := true
	for i := 1; i < 11; i++ {
		if digits[i] != digits[0] {
			iguais = false
			break
		}
	}
	if iguais {
		return false
	}

	toInt := func(r rune) int { return int(r - '0') }

	// 1º dígito
	sum := 0
	for i := 0; i < 9; i++ {
		sum += toInt(digits[i]) * (10 - i)
	}
	d1 := (sum * 10) % 11
	if d1 == 10 {
		d1 = 0
	}
	if d1 != toInt(digits[9]) {
		return false
	}

	// 2º dígito
	sum = 0
	for i := 0; i < 10; i++ {
		sum += toInt(digits[i]) * (11 - i)
	}
	d2 := (sum * 10) % 11
	if d2 == 10 {
		d2 = 0
	}
	if d2 != toInt(digits[10]) {
		return false
	}

	return true
}
