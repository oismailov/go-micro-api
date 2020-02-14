package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/oismailov/go-micro-api/pkg/util"
	proto "github.com/oismailov/go-micro-api/proto"
	dtuService "github.com/oismailov/go-micro-api/services/dtu"
	"net/http"
)

func GetAccount(c *gin.Context) {
	number := c.Param("number")

	dtu := dtuService.Dtu{}

	var account proto.AccountResponse
	accountRequest := proto.AccountRequest{Number: number}

	err := dtu.AccountLookup(context.Background(), &accountRequest, &account)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.Message{Message: err.Error()})

		return
	}

	var product proto.ProductsResponse

	err = dtu.GetProducts(context.Background(), &account, &product)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.Message{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, product)
}
