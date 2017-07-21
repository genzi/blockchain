package blockchain

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

/////////////////////////////////begin Block////////////////////////////////////

type Block struct {
	index         int
	timestamp     string
	data          string
	previous_hash string
	hash          string
}

func newBlock(index int, timestamp string, data string, previous_hash string) *Block {
	block := new(Block)
	block.index = index
	block.timestamp = timestamp
	block.data = data
	block.previous_hash = previous_hash
	block.hash = block.hashBlock()
	return block
}

func (self Block) hashBlock() string {
	hasher := md5.New()
	hasher.Write([]byte(strconv.Itoa(self.index)))
	hasher.Write([]byte(self.timestamp))
	hasher.Write([]byte(self.data))
	hasher.Write([]byte(self.previous_hash))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (self Block) print() {
	fmt.Println(self.index)
	fmt.Println(self.timestamp)
	fmt.Println(self.data)
	fmt.Println(self.previous_hash)
	fmt.Println(self.hash)
	fmt.Println("-----------------------")
}

///////////////////////////////////end Block////////////////////////////////////

/////////////////////////////////begin Blockchain///////////////////////////////

type Blockchain struct {
	chain []*Block
}

func NewBlockchain() *Blockchain {
	blockchain := new(Blockchain)
	block := blockchain.createGenesisBlock()
	blockchain.chain = append(blockchain.chain, block)

	return blockchain
}

func (self *Blockchain) createGenesisBlock() *Block {
	return newBlock(0, time.Now().Format(time.RFC850), "Genesis Block", "0")
}

func (self *Blockchain) NextBlock() *Block {
	last_block := self.chain[len(self.chain)-1]
	index := last_block.index + 1
	timestamp := time.Now().Format(time.RFC850)
	data := "Hey! I'm block " + strconv.Itoa(index)
	hash := last_block.hash
	currentblock := newBlock(index, timestamp, data, hash)
	self.chain = append(self.chain, currentblock)

	return currentblock
}

func (self *Blockchain) Print() {
	for _, block := range self.chain {
		block.print()
	}
}

///////////////////////////////////end Blockchain///////////////////////////////
