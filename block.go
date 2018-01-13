package main

type Block struct {
	Index        int
	PreviousHash int
	Timestamp    int
	Data         string
	ThisHash     int
}

func NewBlock(index int, timestamp int, data string, previousHash int) Block {
	block := new(Block)
	block.Index = index
	block.Timestamp = timestamp
	block.Data = data
	block.PreviousHash = previousHash
	block.ThisHash = block.CalculateHash()

	return *block
}

func (block Block) CalculateHash() int {
	return block.Index + block.PreviousHash + block.Timestamp + len(block.Data)
}
