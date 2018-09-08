package chain

import (
	"crypto/sha256"
)

// Transaction is the data that will be put into blocks in the blockchain
type Transaction struct {
	extraData []byte
	nonce     uint8
}

// GetHash returns the hash of the transaction
func (t *Transaction) GetHash() [sha256.Size]uint8 {
	hash := sha256.Sum256(t.hashableFields())
	return hash
}

// TODO: Implement once hashable fields are decided
// hashableFields returns the hashable fields of the transaction
func (t *Transaction) hashableFields() []uint8 {
	data := t.extraData
	data = append(data, t.nonce)
	return data
}
