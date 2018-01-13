package main

type Chain struct {
	blockChain []Block
}

func NewChain() Chain {
	chain := new(Chain)
	chain.blockChain = make([]Block, 1)
	chain.blockChain[0] = createGenesisBlock()

	return *chain
}

func (chain *Chain) GetLatestBlock() Block {
	return chain.blockChain[len(chain.blockChain)-1]
}

func (chain *Chain) AddBlock(block Block) {
	chain.blockChain = append(chain.blockChain, block)
}

func (chain *Chain) IsChainValid() bool {

	if len(chain.blockChain) < 2 {
		return true
	}

	for i := 1; i < len(chain.blockChain); i++ {
		currentBlock := chain.blockChain[i]
		previousBlock := chain.blockChain[i-1]

		if currentBlock.ThisHash != currentBlock.CalculateHash() {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.ThisHash {
			return false
		}
	}

	return true
}

func createGenesisBlock() Block {
	return NewBlock(0, 123, "{mydata: ok}", 0)
}
