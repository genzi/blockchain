package blockchain

import "testing"

func TestCreateSimpleBlockchain(t *testing.T) {
	simplechain := NewBlockchain()

	if simplechain == nil {
		t.Error("Error creating Blockchain object")
	}
}

func TestCreateNextBlock(t *testing.T) {
	simplechain := NewBlockchain()

	block := simplechain.NextBlock()

	if block == nil {
		t.Error("Error creating next block in Blockchain")
	}
}

func TestCreateAndPrintBlockchain(t *testing.T) {
	simplechain := NewBlockchain()

	for i := 0; i < 10; i++ {
		simplechain.NextBlock()
	}
	simplechain.Print()
}
