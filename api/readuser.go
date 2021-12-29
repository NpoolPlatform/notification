// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/notification/message/npool"
	crud "github.com/NpoolPlatform/notification/pkg/crud/readuser"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateReadUser(ctx context.Context, in *npool.CreateReadUserRequest) (*npool.CreateReadUserResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create read user error: %v", err)
		return &npool.CreateReadUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) CheckReadUser(ctx context.Context, in *npool.CheckReadUserRequest) (*npool.CheckReadUserResponse, error) {
	resp, err := crud.Check(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("check read user error: %v", err)
		return &npool.CheckReadUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
