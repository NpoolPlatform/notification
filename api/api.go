package api

import (
	"context"

	"github.com/NpoolPlatform/notification/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedNotificationServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterNotificationServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterNotificationHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
