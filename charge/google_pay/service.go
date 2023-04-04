package google_pay

import "context"

type Service interface {
	Charge(ctx context.Context, request Request) (Response, error)
}
