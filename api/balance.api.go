package api

import (
	"encoding/json"
	"fmt"
	"go-firebond-assignment/api/dto"
	"go-firebond-assignment/common"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetWalletBalance struct {
	Address string `uri:"address" binding:"required"`
}

func (server *Server) getWalletBalance(ctx *gin.Context) {
	var req GetWalletBalance
	if err := ctx.ShouldBindUri(&req); err != nil {
		dto.BadRequestResponse(ctx, common.ErrBindingRequest.Error(), err)
		return
	}
	requestUrl := fmt.Sprintf("%v?module=account&action=balance&address=%v&tag=latest&apikey=%v", server.config.EthScanUrl, req.Address, server.config.EthScanApiKey)

	fmt.Println(requestUrl)
	client := &http.Client{}
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Printf("%v \n %v\n", common.ErrProcessingRequest.Error(), err)
		dto.InternalServerErrorResponse(ctx, common.ErrProcessingRequest.Error(), err)
		return
	}
	response, err := client.Do(request)
	if err != nil {
		log.Printf(" GET BALANCE :: %v \n", err)
		dto.InternalServerErrorResponse(ctx, common.ErrProcessingRequest.Error(), err)
		return
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("error processing request ... \n statusCode - %v \n", response.StatusCode)
		dto.CustomResponse(ctx, "something went wrong", common.ErrProcessingRequest, 500, false, nil)
		return
	}

	var result map[string]any
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Printf("error parsing response :: %v \n", err)
		dto.InternalServerErrorResponse(ctx, common.ErrProcessingRequest.Error(), err)

		return
	}

	if result["message"] != "OK" {
		dto.CustomResponse(ctx, "something went wrong", common.ErrProcessingRequest, 500, false, nil)
		return
	}
	dto.OkResponse(ctx, fmt.Sprintf("wallet address (%v) balance", req.Address), result)
}
