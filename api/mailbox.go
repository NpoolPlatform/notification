// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/notification/message/npool"
	crud "github.com/NpoolPlatform/notification/pkg/crud/mailbox"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateMail(ctx context.Context, in *npool.CreateMailRequest) (*npool.CreateMailResponse, error) {
	resp, err := crud.CreateMail(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create mailbox error: %v", err)
		return &npool.CreateMailResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateMail(ctx context.Context, in *npool.UpdateMailRequest) (*npool.UpdateMailResponse, error) {
	resp, err := crud.UpdateMail(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("check mailbox error: %v", err)
		return &npool.UpdateMailResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
