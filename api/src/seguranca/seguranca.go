package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe ia, string e coloca hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha Compara hash e senha e retorna se são iguais
func VerificarSenha(senhaString, senhaHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}
