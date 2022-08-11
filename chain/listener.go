package chain

import "math/big"

type Listener interface {
	run()
	handlePastBlock(fromBlock, toBlock *big.Int) error
}
