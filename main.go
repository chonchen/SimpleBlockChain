package main

import (
	"fmt"
)

func main() {
	chain := NewChain()

	fmt.Printf("%v\n", chain.IsChainValid())

	chain.AddBlock(NewBlock(1, 456, "{data:new data}", chain.GetLatestBlock().ThisHash))

	fmt.Printf("%v\n", chain.IsChainValid())

	chain.AddBlock(NewBlock(2, 456, "{data:fake block}", 4589234985435))

	fmt.Printf("%v\n", chain.IsChainValid())

}
