package utils

import "golang.org/x/crypto/bcrypt"

type PwHash struct {
}

func NewPwHash() PwHash {
	return PwHash{}
}

func (p PwHash) Hash(password string) string {
	bcrypt, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bcrypt)
}

func (p PwHash) Compare(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
