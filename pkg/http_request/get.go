package http_request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/robertkrimen/otto"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

type Http struct{}

func (h *Http) Get(url string, header map[string]string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log_handler.LoggerF("Send \"GET\" request to %s is unable", url)
		return "", err
	}
	if header["x-api-token"] != "" {
		// req.Header.Add("x-api-token", config.Tatum_token)
		req.Header.Add("x-api-token", utils.ToString(header["x-api-token"]))
	}
	response, err := client.Do(req)
	// responseByte, err := http.Get(url)
	if err != nil {
		log_handler.LoggerF("Send \"GET\" request to %s is unable", url)
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log_handler.LoggerF(err.Error())
		return "", err
	}

	return string(body), err
}

func (h *Http) ParseBlockchairResult(value string, address string, network string) ([]map[string]interface{}, error) {
	response := []map[string]interface{}{}

	// templateCommand := `jq '.data["salam"]'`
	// // templateCommand := `echo '` + value + `' | jq -c '.data["` + address + `"]`
	// cmd := exec.Command("sh", "-c", templateCommand)
	// list, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ctx := v8.NewContext()
	// jsCode := fmt.Sprintf("result=(JSON.parse('%s')).data['%s'].calls[0].block_id", value, address)
	// ctx.RunScript(jsCode, "app.js")
	// val, _ := ctx.RunScript("result", "app.js")
	// fmt.Printf("%s", val)

	switch network {
	case "bitcoin":
		value, _ = h.bitcoinResponseCleaner(value, address)
	case "ethereum":
		value, _ = h.ethereumResponseCleaner(value, address)
	}

	if err := json.Unmarshal([]byte(value), &response); err != nil {
		log_handler.LoggerF(err.Error())
		// 	return dto.BlockChairEthereumTransaction{}, err
		return nil, err
	}

	return response, nil
}

func (h *Http) ethereumResponseCleaner(value string, address string) (string, error) {
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

func (h *Http) bitcoinResponseCleaner(value string, address string) (string, error) {
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
