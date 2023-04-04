package bank

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
	Charge(ctx context.Context, request TransferRequest) (TransferResponse, error)
}

type DefaultBankService struct {
	configuration.Flutter
	utils.FlutterHttpClient
}

func (d DefaultBankService) Charge(ctx context.Context, request TransferRequest) (TransferResponse, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return TransferResponse{}, err
	}

	endpoint := fmt.Sprintf("%s/charges?type=bank_transfer", d.GetBaseUrl())
	response, err := d.Post(ctx, endpoint, d.GetSecretKey(), data, nil)

	if err != nil {
		return TransferResponse{}, err
	}

	if response.StatusCode != 200 {
		return TransferResponse{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return TransferResponse{}, err
	}

	var transferResponse TransferResponse
	err = json.Unmarshal(body, &transferResponse)

	if err != nil {
		return TransferResponse{}, err
	}

	return transferResponse, nil
}
