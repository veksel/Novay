package chain

import "crypto/sha256"

// Block is a list of transactions
type Block struct {
	transactions []*Transaction
	nonce        uint8
}

// GetHash returns the hash of the block
func (b *Block) GetHash() [sha256.Size]uint8 {
	hash := sha256.Sum256(b.hashableFields())
	return hash
}

func (b *Block) hashableFields() []uint8 {
	var data []uint8
	for _, tx := range b.transactions {
		hash := tx.GetHash()
		data = append(data, hash[:]...)
	}
	data = append(data, b.nonce)
	return data
}
