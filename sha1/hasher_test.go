package sha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha1Hasher(t *testing.T) {
	sha1 := New()

	value := "hello"
	hashedValue, err := sha1.Make(value)

	assert.Nil(t, err)
	assert.Equal(t, "68656c6c6fda39a3ee5e6b4b0d3255bfef95601890afd80709", hashedValue)
	assert.True(t, sha1.Check(value, hashedValue))

	assert.True(t, sha1.Check(value, sha1.MustMake(value)))
}
