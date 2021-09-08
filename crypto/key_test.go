package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyGenerateKey(t *testing.T) {
	assert := assert.New(t)
	address, private, err := GenerateKey()
	assert.Nil(err)
	fmt.Println(private)
	a, err := NormalizePublicKey(address)
	assert.Nil(err)
	fmt.Println(a)
}
