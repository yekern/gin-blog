package utils

import "golang.org/x/crypto/bcrypt"

type Password struct {
}

func NewPassword() *Password {
	return &Password{}
}

// EncodePassword 加密密码
func (p *Password) EncodePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// DeCodePassword 验证密码
func (p *Password) DeCodePassword(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
