package repository

import (
	"Api-Aula1/models"
	"errors"
	"sync"
)

type UsersRepository struct {
	mu     sync.Mutex
	data   map[uint64]models.Users
	lastID uint64
}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{
		data: make(map[uint64]models.Users),
	}
}

// CREATE
func (r *UsersRepository) Create(user models.Users) (models.Users, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastID++
	user.ID = r.lastID
	r.data[user.ID] = user

	return user, nil
}

// READ ALL
func (r *UsersRepository) FindAll() ([]models.Users, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	users := make([]models.Users, 0, len(r.data))
	for _, u := range r.data {
		users = append(users, u)
	}
	return users, nil
}

// READ BY ID (opcional, mas é útil)
func (r *UsersRepository) FindByID(id uint64) (models.Users, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.data[id]
	if !ok {
		return models.Users{}, errors.New("usuário não encontrado")
	}
	return user, nil
}

// UPDATE
func (r *UsersRepository) Update(id uint64, user models.Users) (models.Users, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return models.Users{}, errors.New("usuário não encontrado")
	}

	user.ID = id
	r.data[id] = user
	return user, nil
}

// DELETE
func (r *UsersRepository) Delete(id uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return errors.New("usuário não encontrado")
	}

	delete(r.data, id)
	return nil
}
