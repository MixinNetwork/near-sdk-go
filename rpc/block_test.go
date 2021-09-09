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
	b, err := c.GetBlockByNumber(47315560)
	fmt.Println(err)
	fmt.Println(b)

	fmt.Println(b.Chunks)
	for _, a := range b.Chunks {
		fmt.Println(a.ChunkHash)
		chunk, err := c.GetChunk(a.ChunkHash)
		fmt.Println(err)
		fmt.Println(chunk)
	}
}
