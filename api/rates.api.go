package api

import (
	"database/sql"
	"fmt"
	"go-firebond-assignment/api/dto"
	"go-firebond-assignment/common"
	db "go-firebond-assignment/db/sqlc"

	"github.com/gin-gonic/gin"
)

type CurrentCryptoToFiatRateRequest struct {
	CryptoCurrencyId    string `uri:"cryptocurrency" binding:"required"`
	FiatCurrencyAcronym string `uri:"fiat" binding:"required"`
}

func (server *Server) getCurrentCryptoToFiatRate(ctx *gin.Context) {
	var req CurrentCryptoToFiatRateRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		dto.BadRequestResponse(ctx, common.ErrBindingRequest.Error(), err)
	}

	if len(common.SupportedCrypto[req.CryptoCurrencyId]) < 1 {
		dto.CustomResponse(ctx, fmt.Sprintf("supported currencies :: %v ", common.SupportedCrypto), common.ErrCurrencyNotSupported, 400, false, nil)
		return
	}
	if len(common.SupportedFiat[req.FiatCurrencyAcronym]) < 1 {
		dto.CustomResponse(ctx, fmt.Sprintf("supported currencies :: %v ", common.SupportedFiat), common.ErrFiatNotSupported, 400, false, nil)
		return
	}

	arg := db.GetCurrentCryptoFiatRateParams{
		FiatCurrencyAcronym: sql.NullString{
			String: req.FiatCurrencyAcronym,
			Valid:  true,
		},
		CryptoCurrencyId: sql.NullString{
			String: req.CryptoCurrencyId,
			Valid:  true,
		},
	}
	rates, err := server.store.GetCurrentCryptoFiatRate(ctx, arg)
	if err != nil {
		dto.InternalServerErrorResponse(ctx, common.ErrProcessingRequest.Error(), err)
		return
	}

	dto.OkResponse(ctx, "current exchange", rates)

}

type GetCryptoCurrencyRateRequest struct {
	CryptoCurrencyId string `uri:"cryptocurrency" binding:"required"`
}

func (server *Server) getCryptoCurrencyRates(ctx *gin.Context) {
	var req GetCryptoCurrencyRateRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		dto.BadRequestResponse(ctx, common.ErrBindingRequest.Error(), err)
	}

	if len(common.SupportedCrypto[req.CryptoCurrencyId]) < 1 {
		dto.CustomResponse(ctx, fmt.Sprintf("supported currencies :: %v ", common.SupportedCrypto), common.ErrCurrencyNotSupported, 400, false, nil)
		return
	}

	arg := sql.NullString{
		String: req.CryptoCurrencyId,
		Valid:  true,
	}

	rates, err := server.store.GetCurrentCryptoExchangeRate(ctx, arg)
	if err != nil {
		dto.InternalServerErrorResponse(ctx, common.ErrProcessingRequest.Error(), err)
		return
	}

	dto.OkResponse(ctx, fmt.Sprintf("%v rates ", req.CryptoCurrencyId), rates)
}

func (server *Server) getRates(ctx *gin.Context) {
	rates, err := server.store.GetCurrentExchangeRates(ctx)
	if err != nil {
		dto.InternalServerErrorResponse(ctx, common.ErrProcessingRequest.Error(), err)
		return
	}

	dto.OkResponse(ctx, "success", rates)
}

func (server *Server) getExchangeRateHistory(ctx *gin.Context) {
	var req CurrentCryptoToFiatRateRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		dto.BadRequestResponse(ctx, common.ErrBindingRequest.Error(), err)
	}

	if len(common.SupportedCrypto[req.CryptoCurrencyId]) < 1 {
		dto.CustomResponse(ctx, fmt.Sprintf("supported currencies :: %v ", common.SupportedCrypto), common.ErrCurrencyNotSupported, 400, false, nil)
		return
	}
	if len(common.SupportedFiat[req.FiatCurrencyAcronym]) < 1 {
		dto.CustomResponse(ctx, fmt.Sprintf("supported currencies :: %v ", common.SupportedFiat), common.ErrFiatNotSupported, 400, false, nil)
		return
	}

	arg := db.GetExchangeRateHistoryParams{
		FiatCurrencyAcronym: sql.NullString{
			String: req.FiatCurrencyAcronym,
			Valid:  true,
		},
		CryptoCurrencyId: sql.NullString{
			String: req.CryptoCurrencyId,
			Valid:  true,
		},
	}
	rates, err := server.store.GetExchangeRateHistory(ctx, arg)
	if err != nil {
		dto.InternalServerErrorResponse(ctx, common.ErrProcessingRequest.Error(), err)
		return
	}

	dto.OkResponse(ctx, fmt.Sprintf("%v - %v , Rates :: ", req.CryptoCurrencyId, req.FiatCurrencyAcronym), rates)
}
