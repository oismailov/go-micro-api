package dtu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	proto "github.com/oismailov/go-micro-api/proto"
	"io/ioutil"
	"net/http"
)

func (g *Dtu) SendTransfer(ctx context.Context, req *proto.TransferRequest, rsp *proto.TransferResponse) error {
	url := "https://www.dingconnect.com/api/V1/SendTransfer"

	requestBody, err := json.Marshal(req)

	if err != nil {
		return err
	}

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	r.Header.Set("api_key", "4H1hsF5ZIWh6ENQ6GmXGqX")
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(r)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, rsp)

	fmt.Println(string(body))

	if err != nil {
		return err
	}

	return nil
}
