package chain

import (
	"fmt"
	"testing"
)

func TestTransactionHashing(t *testing.T) {
	txOne := Transaction{
		extraData: []uint8{1, 1, 1, 0, 0},
		nonce:     0,
	}

	txTwo := Transaction{
		extraData: []uint8{1, 1, 1, 0, 0},
		nonce:     1,
	}

	txHexOne := fmt.Sprintf("%02x", txOne.GetHash())
	txHexTwo := fmt.Sprintf("%02x", txTwo.GetHash())

	if txHexOne == txHexTwo {
		t.Errorf("Two transactions have the same hash with different nonces!")
	}

}
