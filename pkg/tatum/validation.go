package tatum

import (
	"encoding/json"
	"fmt"

	"github.com/robertkrimen/otto"
	"zarinworld.ir/event/pkg/log_handler"
)

type TatumHttpValidation struct{}

func (h *TatumHttpValidation) ParseTatumResult(value string, network string) (map[string]interface{}, error) {
	response := map[string]interface{}{}

	// switch network {
	// case "bitcoin":
	// 	value, _ = h.bitcoinResponseCleaner(value, address)
	// case "ethereum":
	// 	value, _ = h.ethereumResponseCleaner(value, address)
	// }

	if err := json.Unmarshal([]byte(value), &response); err != nil {
		log_handler.LoggerF(err.Error())
		// 	return dto.BlockChairEthereumTransaction{}, err
		return nil, err
	}

	return response, nil
}

func (h *TatumHttpValidation) ethereumResponseCleaner(value string, address string) (string, error) {
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

func (h *TatumHttpValidation) bitcoinResponseCleaner(value string, address string) (string, error) {
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
