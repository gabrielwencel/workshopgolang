package security

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma senha em texto puro e devolve o hash para salvar no banco
func Hash(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// Verify compara a senha em texto puro com o hash vindo do banco
// Se a senha estiver errada, ele retorna um erro (bcrypt.ErrMismatchedHashAndPassword)
func Verify(hashFromDB, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashFromDB), []byte(password))
}
