package blockchair

import (
	"encoding/json"
	"fmt"

	"github.com/robertkrimen/otto"
	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/log_handler"
)

type BlockchairHttpValidation struct{}

func (h *BlockchairHttpValidation) ParseBlockchairResult(value string, address string, network string) (map[string]interface{}, error) {
	response := map[string]interface{}{}

	switch network {
	case config.BITCOIN:
		value, _ = h.bitcoinResponseCleaner(value, address)
	case config.ETHEREUM:
		value, _ = h.ethereumResponseCleaner(value, address)
	}

	if err := json.Unmarshal([]byte(value), &response); err != nil {
		log_handler.LoggerF(err.Error())
		// 	return dto.BlockChairEthereumTransaction{}, err
		return nil, err
	}

	return response, nil
}

func (h *BlockchairHttpValidation) ParseBlockchairListResult(value string, address string, network string) ([]map[string]interface{}, error) {
	response := []map[string]interface{}{}

	switch network {
	case config.BITCOIN:
		value, _ = h.bitcoinResponseCleaner(value, address)
	case config.ETHEREUM:
		value, _ = h.ethereumResponseCleaner(value, address)
	}

	if err := json.Unmarshal([]byte(value), &response); err != nil {
		log_handler.LoggerF(err.Error())
		// 	return dto.BlockChairEthereumTransaction{}, err
		return nil, err
	}

	return response, nil
}

func (h *BlockchairHttpValidation) ethereumResponseCleaner(value string, address string) (string, error) {
	vm := otto.New()
	jsCode := fmt.Sprintf(`
		result=JSON.stringify((JSON.parse('%s')).data['%s'].calls)
	`, value, address)
	vm.Run(jsCode)
	cleanedText, err := vm.Get("result")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", cleanedText), nil
}

func (h *BlockchairHttpValidation) bitcoinResponseCleaner(value string, address string) (string, error) {
	vm := otto.New()
	jsCode := fmt.Sprintf(`
		result=JSON.stringify((JSON.parse('%s')).data['%s'].transactions)
	`, value, address)
	vm.Run(jsCode)
	cleanedText, err := vm.Get("result")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", cleanedText), nil
}

func (h *BlockchairHttpValidation) blockchairOkResponse(response string) (bool, error) {
	a := blockchairDto{}
	if err := json.Unmarshal([]byte(response), &a); err != nil {
		log_handler.LoggerF(err.Error())
		return false, err
	}
	if a.Context.Code != 200 {
		return false, nil
	} else {
		return true, nil
	}
}

type blockchairDto struct {
	Context struct {
		Code int `json:"code"`
	} `json:"context"`
}
