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
