package bcrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBcryptHasher(t *testing.T) {
	bcrypt := New()

	value := "123456"

	hashedValue, _ := bcrypt.Make(value)
	assert.True(t, bcrypt.Check(value, hashedValue))

	hashedValue2 := bcrypt.MustMake(value)
	assert.True(t, bcrypt.Check(value, hashedValue2))
}
