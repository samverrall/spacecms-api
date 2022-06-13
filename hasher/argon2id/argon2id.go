package argon2id

import (
	"github.com/alexedwards/argon2id"
	argon "github.com/alexedwards/argon2id"
)

type argonHasher struct{}

func New() *argonHasher {
	return &argonHasher{}
}

func (ah *argonHasher) Hash(password string) (string, error) {
	hash, err := argon.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil

}

func (ah *argonHasher) Compare(hashed, password string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hashed)
}
