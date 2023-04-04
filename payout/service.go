package payout

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
	CreateSubAccount(ctx context.Context, request CreateSubAccountRequest) (SubAccountResponse, error)
	ListSubAccounts(ctx context.Context) (SubAccountsResponse, error)
	FetchSubAccountBalance(ctx context.Context, accountId string, currency string) (SubAccountBalance, error)
	FetchSubAccountDetails(ctx context.Context, accountId string, currency string) (StaticAccountResponse, error)
	FetchSubAccountTransactions(ctx context.Context, accountId string, currency string) (SubAccountTransactionsResponse, error)
}

type DefaultPayoutService struct {
	configuration.Flutter
	utils.FlutterHttpClient
}

func (d DefaultPayoutService) CreateSubAccount(ctx context.Context, request CreateSubAccountRequest) (SubAccountResponse, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return SubAccountResponse{}, err
	}

	endpoint := fmt.Sprintf("%s/payout-subaccounts", d.GetBaseUrl())
	response, err := d.Post(ctx, endpoint, d.GetSecretKey(), data, nil)

	if err != nil {
		return SubAccountResponse{}, err
	}

	if response.StatusCode != 200 {
		return SubAccountResponse{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return SubAccountResponse{}, err
	}

	var jsonResponse SubAccountResponse
	err = json.Unmarshal(body, &jsonResponse)

	if err != nil {
		return SubAccountResponse{}, err
	}

	return jsonResponse, nil
}

func (d DefaultPayoutService) ListSubAccounts(ctx context.Context) (SubAccountsResponse, error) {
	endpoint := fmt.Sprintf("%s/payout-subaccounts", d.GetBaseUrl())
	response, err := d.Get(ctx, endpoint, d.GetSecretKey(), nil)

	if err != nil {
		return SubAccountsResponse{}, err
	}

	if response.StatusCode != 200 {
		return SubAccountsResponse{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return SubAccountsResponse{}, err
	}

	var jsonResponse SubAccountsResponse
	err = json.Unmarshal(body, &jsonResponse)

	if err != nil {
		return SubAccountsResponse{}, err
	}

	return jsonResponse, nil
}

func (d DefaultPayoutService) FetchSubAccountBalance(ctx context.Context, accountId string, currency string) (SubAccountBalance, error) {
	endpoint := fmt.Sprintf("%s/payout-subaccounts/%s/balances?currency=%s", d.GetBaseUrl(), accountId, currency)
	response, err := d.Get(ctx, endpoint, d.GetSecretKey(), nil)

	if err != nil {
		return SubAccountBalance{}, err
	}

	if response.StatusCode != 200 {
		return SubAccountBalance{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return SubAccountBalance{}, err
	}

	var jsonResponse SubAccountBalance
	err = json.Unmarshal(body, &jsonResponse)

	if err != nil {
		return SubAccountBalance{}, err
	}

	return jsonResponse, nil
}

func (d DefaultPayoutService) FetchSubAccountDetails(ctx context.Context, accountId string, currency string) (StaticAccountResponse, error) {
	endpoint := fmt.Sprintf("%s/payout-subaccounts/%s/balances?currency=%s", d.GetBaseUrl(), accountId, currency)
	response, err := d.Get(ctx, endpoint, d.GetSecretKey(), nil)

	if err != nil {
		return StaticAccountResponse{}, err
	}

	if response.StatusCode != 200 {
		return StaticAccountResponse{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return StaticAccountResponse{}, err
	}

	var jsonResponse StaticAccountResponse
	err = json.Unmarshal(body, &jsonResponse)

	if err != nil {
		return StaticAccountResponse{}, err
	}

	return jsonResponse, nil
}

func (d DefaultPayoutService) FetchSubAccountTransactions(ctx context.Context, accountId string, currency string) (SubAccountTransactionsResponse, error) {
	endpoint := fmt.Sprintf("%s/payout-subaccounts/%s/transactions", d.GetBaseUrl(), accountId)
	response, err := d.Get(ctx, endpoint, d.GetSecretKey(), nil)

	if err != nil {
		return SubAccountTransactionsResponse{}, err
	}

	if response.StatusCode != 200 {
		return SubAccountTransactionsResponse{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return SubAccountTransactionsResponse{}, err
	}

	var jsonResponse SubAccountTransactionsResponse
	err = json.Unmarshal(body, &jsonResponse)

	if err != nil {
		return SubAccountTransactionsResponse{}, err
	}

	return jsonResponse, nil
}
