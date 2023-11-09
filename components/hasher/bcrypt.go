package hasher

import "golang.org/x/crypto/bcrypt"

type bcryptHash struct {
	cost int
}

func NewBcryptHash(cost int) *bcryptHash {
	return &bcryptHash{cost: cost}
}

func (h *bcryptHash) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(bytes), err
}

func (h *bcryptHash) Compare(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
