package card

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
	InitiateCharge(ctx context.Context, request InitiateRequest) (InitiateResponse, error)
	AuthorizeCharge(ctx context.Context, request AuthorizeRequest) (AuthorizeResponse, error)
	ValidateCharge(ctx context.Context, request ValidateChargeRequest) (ValidateChargeResponse, error)

	TokenizedCharge(ctx context.Context, request TokenizedChargeRequest) (TokenizedChargeResponse, error)
}

type DefaultCardService struct {
	configuration.Flutter
	utils.FlutterHttpClient
}

type EncryptedRequest struct {
	Client string `json:"client"`
}

func (d DefaultCardService) InitiateCharge(ctx context.Context, request InitiateRequest) (InitiateResponse, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return InitiateResponse{}, err
	}

	encrypted, err := d.Encrypt3Des(data)
	if err != nil {
		return InitiateResponse{}, err
	}

	endpoint := fmt.Sprintf("%s/charges?type=card", d.GetBaseUrl())
	requestBody := EncryptedRequest{Client: encrypted}

	encryptedRequest, err := json.Marshal(requestBody)
	if err != nil {
		return InitiateResponse{}, err
	}

	response, err := d.Post(ctx, endpoint, d.GetSecretKey(), encryptedRequest, nil)

	if err != nil {
		return InitiateResponse{}, err
	}

	if response.StatusCode != 200 {
		return InitiateResponse{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return InitiateResponse{}, err
	}

	var initiateResponse InitiateResponse
	err = json.Unmarshal(body, &initiateResponse)

	if err != nil {
		return InitiateResponse{}, err
	}

	return initiateResponse, nil
}

func (d DefaultCardService) AuthorizeCharge(ctx context.Context, request AuthorizeRequest) (AuthorizeResponse, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return AuthorizeResponse{}, err
	}

	encrypted, err := d.Encrypt3Des(data)
	if err != nil {
		return AuthorizeResponse{}, err
	}

	endpoint := fmt.Sprintf("%s/charges?type=card", d.GetBaseUrl())
	requestBody := EncryptedRequest{Client: encrypted}

	encryptedRequest, err := json.Marshal(requestBody)
	if err != nil {
		return AuthorizeResponse{}, err
	}

	response, err := d.Post(ctx, endpoint, d.GetSecretKey(), encryptedRequest, nil)

	if err != nil {
		return AuthorizeResponse{}, err
	}

	if response.StatusCode != 200 {
		return AuthorizeResponse{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return AuthorizeResponse{}, err
	}

	var authorizeResponse AuthorizeResponse
	err = json.Unmarshal(body, &authorizeResponse)

	if err != nil {
		return AuthorizeResponse{}, err
	}

	return authorizeResponse, nil
}

func (d DefaultCardService) ValidateCharge(ctx context.Context, request ValidateChargeRequest) (ValidateChargeResponse, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return ValidateChargeResponse{}, err
	}

	endpoint := fmt.Sprintf("%s/validate-charge", d.GetBaseUrl())
	response, err := d.Post(ctx, endpoint, d.GetSecretKey(), data, nil)

	if err != nil {
		return ValidateChargeResponse{}, err
	}

	if response.StatusCode != 200 {
		return ValidateChargeResponse{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return ValidateChargeResponse{}, err
	}

	var validateChargeResponse ValidateChargeResponse
	err = json.Unmarshal(body, &validateChargeResponse)

	if err != nil {
		return ValidateChargeResponse{}, err
	}

	return validateChargeResponse, nil
}

func (d DefaultCardService) TokenizedCharge(ctx context.Context, request TokenizedChargeRequest) (TokenizedChargeResponse, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return TokenizedChargeResponse{}, err
	}

	endpoint := fmt.Sprintf("%s/tokenized-charges", d.GetBaseUrl())
	response, err := d.Post(ctx, endpoint, d.GetSecretKey(), data, nil)

	if err != nil {
		return TokenizedChargeResponse{}, err
	}

	if response.StatusCode != 200 {
		return TokenizedChargeResponse{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return TokenizedChargeResponse{}, err
	}

	var tokenizedChargeResponse TokenizedChargeResponse
	err = json.Unmarshal(body, &tokenizedChargeResponse)

	if err != nil {
		return TokenizedChargeResponse{}, err
	}

	return tokenizedChargeResponse, nil
}
