package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/notification"
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
