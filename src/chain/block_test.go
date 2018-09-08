package chain

import (
	"fmt"
	"testing"
)

func TestBlockHashing(t *testing.T) {
	txOne := Transaction{
		extraData: []uint8{1, 1, 1, 0, 0},
		nonce:     0,
	}

	txTwo := Transaction{
		extraData: []uint8{1, 1, 1, 0, 0},
		nonce:     1,
	}

	blockOne := Block{
		transactions: []*Transaction{&txOne, &txTwo},
		nonce:        0,
	}

	blockTwo := Block{
		transactions: []*Transaction{&txOne, &txTwo},
		nonce:        1,
	}

	blockHexOne := fmt.Sprintf("%02x", blockOne)
	blockHexTwo := fmt.Sprintf("%02x", blockTwo)

	if blockHexOne == blockHexTwo {
		t.Errorf("Two blocks have the same hash with different nonces!")
	}

}
