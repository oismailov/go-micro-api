package dtu

import (
	"context"
	"encoding/json"
	"fmt"
	proto "github.com/oismailov/go-micro-api/proto"
	"io/ioutil"
	"net/http"
)

func (g *Dtu) GetProducts(ctx context.Context, req *proto.AccountResponse, rsp *proto.ProductsResponse) error {
	url := fmt.Sprintf(
		"https://www.dingconnect.com/api/V1/GetProducts?ProviderCodes[]=%s&RegionCodes[]=%s",
		req.ProviderCode,
		req.RegionCode,
	)

	r, err := http.NewRequest("GET", url, nil)

	r.Header.Set("api_key", "4H1hsF5ZIWh6ENQ6GmXGqX")

	client := &http.Client{}

	response, err := client.Do(r)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	var data proto.ProductsResponse

	err = json.Unmarshal(body, &data)

	if err != nil {
		return err
	}

	for i := 0; i < len(data.Items); i++ {
		rsp.Items = append(rsp.Items, data.Items[i])
	}

	return nil
}
