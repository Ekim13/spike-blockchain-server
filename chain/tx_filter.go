package chain

const blockConfirmHeight = 15

const (
	governanceToken TokenType = iota
	gameToken
	usdc
	bnb
	gameVault
	gameNft
)

var shortNames = map[TokenType]string{
	governanceToken: "governanceToken",
	gameToken:       "gameToken",
	bnb:             "bnb",
	gameVault:       "gameVault",
	gameNft:         "gameNft",
	usdc:            "usdc",
}

func (t TokenType) String() string {
	n, ok := shortNames[t]
	if !ok {
		return "unknown"
	}
	return n
}

const (
	SKK_RECHARGE = iota + 1
	SKS_RECHARGE
	USDC_RECHARGE
	BNB_RECHARGE
	SKK_WITHDRAW
	SKS_WITHDRAW
	USDC_WITHDRAW
	BNB_WITHDRAW
	AUNFT_TRANSFER
	AUNFT_IMPORT
	NOT_EXIST
)

var recharge = map[int]struct{}{
	SKK_RECHARGE:  {},
	SKS_RECHARGE:  {},
	USDC_RECHARGE: {},
	BNB_RECHARGE:  {},
}

var nftImport = map[int]struct{}{
	AUNFT_IMPORT: {},
}

type TokenType int

type TxFilter interface {
	Accept(fromAddr, toAddr string) (bool, uint64)
}

func checkRecharge(txType int) bool {
	_, ok := recharge[txType]
	return ok
}

func checkImport(txType int) bool {
	_, ok := nftImport[txType]
	return ok
}
