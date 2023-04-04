package paypal

import "context"

type Service interface {
	Charge(ctx context.Context, request Request) (Response, error)
}
