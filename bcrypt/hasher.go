package bcrypt

import (
	"github.com/go-fires/hashing"
	"golang.org/x/crypto/bcrypt"
)

var global hashing.Hasher = &hasher{}

func New() hashing.Hasher {
	return global
}

type hasher struct{}

// Make returns the hashed value.
func (h *hasher) Make(value string) (string, error) {
	return h.MakeWithCost(value, bcrypt.DefaultCost)
}

// MustMake returns the hashed value.
func (h *hasher) MustMake(value string) string {
	hashedValue, err := h.Make(value)

	if err != nil {
		panic(err)
	}

	return hashedValue
}

// MakeWithCost returns the hashed value with the given cost.
func (h *hasher) MakeWithCost(value string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), cost)

	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// Check returns true if the value matches the hashed value.
func (h *hasher) Check(value, hashedValue string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value)) == nil
}
