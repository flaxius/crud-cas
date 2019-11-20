package cmd

import (
	"context"
	"github.com/flaxius/crud-cas/pkg/protocol/grpc"
	"github.com/flaxius/crud-cas/pkg/protocol/rest"
	"github.com/flaxius/crud-cas/pkg/services"
	"github.com/flaxius/crud-cas/ports"
	"time"
)

const (
	contextTimeOut = 5 * time.Second
)

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	v1API := oauth.NewTokenService()
	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, ports.PortiaGrpCDefaultPort, ports.RestPortiaDefaultPort)
	}()
	return grpc.RunServer(ctx, v1API, ports.PortiaGrpCDefaultPort)
}
