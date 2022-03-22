// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notification"
	crud "github.com/NpoolPlatform/notification/pkg/crud/notification"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateNotification(ctx context.Context, in *npool.CreateNotificationRequest) (*npool.CreateNotificationResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create notification error: %v", err)
		return &npool.CreateNotificationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateNotification(ctx context.Context, in *npool.UpdateNotificationRequest) (*npool.UpdateNotificationResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update notification error: %v", err)
		return &npool.UpdateNotificationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetNotificationsByAppUser(ctx context.Context, in *npool.GetNotificationsByAppUserRequest) (*npool.GetNotificationsByAppUserResponse, error) {
	resp, err := crud.GetNotificationsByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get notifications by app user error: %v", err)
		return &npool.GetNotificationsByAppUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
