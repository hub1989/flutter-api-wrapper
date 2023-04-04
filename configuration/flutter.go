package configuration

import (
	"github.com/hub1989/flutter-api-wrapper/utils"
)

type Flutter interface {
	GetBaseUrl() string
	GetSecretKey() string
	GetPublicKey() string
	Encrypt3Des(payload []byte) (string, error)
	GetTransactionPrefix() string
}

type FlutterConfig struct {
	BaseUrl           string
	SecretKey         string
	EncryptionKey     string
	PublicKey         string
	TransactionPrefix string
}

func (f FlutterConfig) GetBaseUrl() string {
	return f.BaseUrl
}

func (f FlutterConfig) GetSecretKey() string {
	return f.SecretKey
}

func (f FlutterConfig) GetPublicKey() string {
	return f.PublicKey
}

func (f FlutterConfig) Encrypt3Des(payload []byte) (string, error) {
	return utils.Encrypt3Des(f.EncryptionKey, payload)
}

func (f FlutterConfig) GetTransactionPrefix() string {
	return f.TransactionPrefix
}
