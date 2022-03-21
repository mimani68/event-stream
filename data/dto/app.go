package dto

type Transaction struct {
	Id             string
	TrxHash        string
	Value          string
	BlockNumber    string
	ExpireIn       string
	Confirmed      string
	ConfirmedCount string
}

type Authority struct {
	Id             string
	Amount         string
	Incoming_value string
	Address        string
	Transactions   string
	Wallet         string
	Status         string
}

type Address struct {
	Id      string
	Address string
	Network string
}

type BlockChairBitcoinTransaction struct {
	BalanceChange int    `json:"balance_change"`
	BlockId       int    `json:"block_id"`
	Hash          string `json:"hash"`
	Time          string `json:"time"`
}

type BlockChairEthereumTransaction struct {
	BlockId         int     `json:"block_id"`
	TransactionHash int     `json:"transaction_hash"`
	Index           string  `json:"index"`
	Time            string  `json:"time"`
	Sender          string  `json:"sender"`
	Recipient       string  `json:"recipient"`
	Value           int     `json:"value"`
	ValueUsd        float32 `json:"value_usd"`
	Transferred     bool    `json:"transferred"`
}

type BlockChairGeneralDto struct {
	Context map[string]interface{} `json:"context"`
	// 	Data    struct {
	// 		Address map[string]interface{}          `json:"address"`
	// 		Calls   []BlockChairEthereumTransaction `json:"calls"`
	// 	} `json:"data"`
	Data map[string]interface{} `json:"data"`
}
