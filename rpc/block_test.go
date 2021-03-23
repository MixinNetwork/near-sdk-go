package rpc

import (
	"fmt"
	"testing"
)

func TestGetBlock(t *testing.T) {
	c := NewClient(MainnetRPCEndpoint)
	height, err := c.GetHeight()
	fmt.Println(height)
	fmt.Println(err)
	b, err := c.GetBlock(height)
	fmt.Println(err)
	fmt.Println(b)
}
