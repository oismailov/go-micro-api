package dtu

import (
	"context"
	"encoding/json"
	"fmt"
	proto "github.com/oismailov/go-micro-api/proto"
	"io/ioutil"
	"log"
	"net/http"
)

type Account struct {
	Items []Item
}

type Item struct {
	ProviderCode string
	RegionCode   string
}

func (g *Dtu) AccountLookup(ctx context.Context, req *proto.AccountRequest, rsp *proto.AccountResponse) error {
	url := fmt.Sprintf(
		"https://www.dingconnect.com/api/V1/GetAccountLookup?AccountNumber=%s",
		req.Number,
	)

	r, err := http.NewRequest("GET", url, nil)

	r.Header.Set("api_key", "4H1hsF5ZIWh6ENQ6GmXGqX")

	client := &http.Client{}

	response, err := client.Do(r)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	var data Account

	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal("Error parsing response. ", err)
	}

	if len(data.Items) == 0 {
		return fmt.Errorf("account does not exist")
	}

	rsp.ProviderCode = data.Items[0].ProviderCode
	rsp.RegionCode = data.Items[0].RegionCode

	return nil
}
