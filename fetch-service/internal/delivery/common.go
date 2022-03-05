package delivery

import "context"

type Server interface {
	Start(port string) error
	Stop(ctx context.Context) error
}
