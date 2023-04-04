package transaction

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hub1989/flutter-api-wrapper/configuration"
	"github.com/hub1989/flutter-api-wrapper/utils"
	"io"
)

type Service interface {
	VerifyChargeByTransactionReference(ctx context.Context, transactionReference string) (VerifyChargeResponse, error)
	VerifyChargeByTransactionId(ctx context.Context, transactionId int) (VerifyChargeResponse, error)
	RefundFull(ctx context.Context, transactionId int) (RefundResponse, error)
	RefundPartial(ctx context.Context, transactionId int, amount float64) (RefundResponse, error)
}

type DefaultTransactionService struct {
	configuration.Flutter
	utils.FlutterHttpClient
}

func (d DefaultTransactionService) VerifyChargeByTransactionReference(ctx context.Context, transactionReference string) (VerifyChargeResponse, error) {
	endpoint := fmt.Sprintf("%s/transactions/verify_by_reference?tx_ref=%s", d.GetBaseUrl(), transactionReference)
	response, err := d.Get(ctx, endpoint, d.GetSecretKey(), nil)

	if err != nil {
		return VerifyChargeResponse{}, err
	}

	if response.StatusCode != 200 {
		return VerifyChargeResponse{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return VerifyChargeResponse{}, err
	}

	var verifyChargeResponse VerifyChargeResponse
	err = json.Unmarshal(body, &verifyChargeResponse)

	if err != nil {
		return VerifyChargeResponse{}, err
	}

	return verifyChargeResponse, nil
}

func (d DefaultTransactionService) VerifyChargeByTransactionId(ctx context.Context, transactionId int) (VerifyChargeResponse, error) {
	endpoint := fmt.Sprintf("%s/transactions/%d/verify", d.GetBaseUrl(), transactionId)
	response, err := d.Get(ctx, endpoint, d.GetSecretKey(), nil)

	if err != nil {
		return VerifyChargeResponse{}, err
	}

	if response.StatusCode != 200 {
		return VerifyChargeResponse{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return VerifyChargeResponse{}, err
	}

	var verifyChargeResponse VerifyChargeResponse
	err = json.Unmarshal(body, &verifyChargeResponse)

	if err != nil {
		return VerifyChargeResponse{}, err
	}

	return verifyChargeResponse, nil
}

func (d DefaultTransactionService) RefundFull(ctx context.Context, transactionId int) (RefundResponse, error) {
	endpoint := fmt.Sprintf("%s/transactions/%d/refund", d.GetBaseUrl(), transactionId)
	response, err := d.Post(ctx, endpoint, d.GetSecretKey(), nil, nil)

	if err != nil {
		return RefundResponse{}, err
	}

	if response.StatusCode != 200 {
		return RefundResponse{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return RefundResponse{}, err
	}

	var refundResponse RefundResponse
	err = json.Unmarshal(body, &refundResponse)

	if err != nil {
		return RefundResponse{}, err
	}

	return refundResponse, nil
}

func (d DefaultTransactionService) RefundPartial(ctx context.Context, transactionId int, amount float64) (RefundResponse, error) {
	endpoint := fmt.Sprintf("%s/transactions/%d/refund", d.GetBaseUrl(), transactionId)
	req := struct {
		Amount float64 `json:"amount"`
	}{
		Amount: amount,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return RefundResponse{}, err
	}

	response, err := d.Post(ctx, endpoint, d.GetSecretKey(), data, nil)

	if err != nil {
		return RefundResponse{}, err
	}

	if response.StatusCode != 200 {
		return RefundResponse{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return RefundResponse{}, err
	}

	var refundResponse RefundResponse
	err = json.Unmarshal(body, &refundResponse)

	if err != nil {
		return RefundResponse{}, err
	}

	return refundResponse, nil
}
