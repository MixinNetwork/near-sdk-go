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
	fmt.Println(address)
	fmt.Println(private)
	fmt.Println(err)
}
