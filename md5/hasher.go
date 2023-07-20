package md5

import (
	"crypto/md5"
	"fmt"

	"github.com/go-fires/hashing"
)

var global hashing.Hasher = &hasher{}

func New() hashing.Hasher {
	return global
}

type hasher struct{}

// Make generates a new hashed value.
func (h *hasher) Make(value string) (string, error) {
	hashedValue := md5.Sum([]byte(value))

	return fmt.Sprintf("%x", hashedValue), nil
}

// MustMake generates a new hashed value.
func (h *hasher) MustMake(value string) string {
	hashedValue, err := h.Make(value)

	if err != nil {
		panic(err)
	}

	return hashedValue
}

// Check checks the given value and hashed value.
func (h *hasher) Check(value, hashedValue string) bool {
	hv, err := h.Make(value)

	if err != nil {
		return false
	}

	return hv == hashedValue
}
