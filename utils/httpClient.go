package utils

import (
	"bytes"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type FlutterHttpClient interface {
	Get(ctx context.Context, endpoint string, token string, headers map[string]string) (*http.Response, error)
	Options(ctx context.Context, endpoint string, token string, headers map[string]string) (*http.Response, error)
	Post(ctx context.Context, endpoint string, token string, body []byte, headers map[string]string) (*http.Response, error)
	Delete(ctx context.Context, endpoint string, token string, body []byte, headers map[string]string) (*http.Response, error)
	Put(ctx context.Context, endpoint string, token string, body []byte, headers map[string]string) (*http.Response, error)
}

type DefaultFlutterHttpClient struct{}

func (d DefaultFlutterHttpClient) Get(ctx context.Context, endpoint string, token string, headers map[string]string) (*http.Response, error) {
	return d.do(ctx, endpoint, token, nil, headers, http.MethodGet)
}

func (d DefaultFlutterHttpClient) Options(ctx context.Context, endpoint string, token string, headers map[string]string) (*http.Response, error) {
	return d.do(ctx, endpoint, token, nil, headers, http.MethodOptions)
}

func (d DefaultFlutterHttpClient) Post(ctx context.Context, endpoint string, token string, body []byte, headers map[string]string) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	return d.do(ctx, endpoint, token, bodyReader, headers, http.MethodPost)
}

func (d DefaultFlutterHttpClient) Delete(ctx context.Context, endpoint string, token string, body []byte, headers map[string]string) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	return d.do(ctx, endpoint, token, bodyReader, headers, http.MethodDelete)
}

func (d DefaultFlutterHttpClient) Put(ctx context.Context, endpoint string, token string, body []byte, headers map[string]string) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	return d.do(ctx, endpoint, token, bodyReader, headers, http.MethodPut)
}

func (d DefaultFlutterHttpClient) do(ctx context.Context, endpoint string, token string, bodyReader io.Reader, headers map[string]string, method string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, endpoint, bodyReader)

	if err != nil {
		log.WithError(err).Error(fmt.Sprintf("could not complete %s request to endpoint %s", method, endpoint))
		return nil, err
	}

	req = EnrichRequestWithAuth(req, token, headers)
	return client.Do(req)
}
