package main

import (
	"fmt"
)

func main()  {
	bc := NewBlockchain()
	bc.AddBlock("send 1 btc to ivan")
	bc.AddBlock("send 2 more btc to ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}