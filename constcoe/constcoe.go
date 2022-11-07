//constcoe.go

package constcoe

const (
	Difficulty          = 20
	InitCoin            = 1000
	TransactionPoolFile = "./tmp/transaction_pool.data" //缓冲池
	BCPath              = "./tmp/blocks"
	BCFile              = "./tmp/blocks/MANIFEST"
	ChecksumLength      = 4
	NetworkVersion      = byte(0x00)
	Wallets             = "./tmp/wallets/"
	WalletsRefList      = "./tmp/ref_list/"
)
