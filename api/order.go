package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/oismailov/go-micro-api/pkg/util"
	proto "github.com/oismailov/go-micro-api/proto"
	dtuService "github.com/oismailov/go-micro-api/services/dtu"
	"net/http"
)

func Order(c *gin.Context) {
	var transferRequest proto.TransferRequest

	err := c.BindJSON(&transferRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.Message{
			Message: string(err.Error()),
		})

		return
	}

	dtu := dtuService.Dtu{}

	transfer := proto.TransferResponse{}

	err = dtu.SendTransfer(context.Background(), &transferRequest, &transfer)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.Message{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, transfer)
}
