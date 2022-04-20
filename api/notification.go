//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notification"
	crud "github.com/NpoolPlatform/notification/pkg/crud/notification"
	mv "github.com/NpoolPlatform/notification/pkg/middleware/notification"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateNotification(ctx context.Context, in *npool.CreateNotificationRequest) (*npool.CreateNotificationResponse, error) {
	resp, err := mv.CreateNotification(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create notification error: %v", err)
		return &npool.CreateNotificationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) CreateNotificationForAppOtherUser(ctx context.Context, in *npool.CreateNotificationForAppOtherUserRequest) (*npool.CreateNotificationForAppOtherUserResponse, error) {
	info := in.GetInfo()
	info.UserID = in.GetTargetUserID()

	resp, err := mv.CreateNotification(ctx, &npool.CreateNotificationRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorw("create notification error: %v", err)
		return &npool.CreateNotificationForAppOtherUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return &npool.CreateNotificationForAppOtherUserResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) CreateNotificationForOtherAppUser(ctx context.Context, in *npool.CreateNotificationForOtherAppUserRequest) (*npool.CreateNotificationForOtherAppUserResponse, error) {
	info := in.GetInfo()
	info.UserID = in.GetTargetUserID()

	resp, err := mv.CreateNotification(ctx, &npool.CreateNotificationRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorw("create notification error: %v", err)
		return &npool.CreateNotificationForOtherAppUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return &npool.CreateNotificationForOtherAppUserResponse{
		Info: resp.Info,
	}, nil
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

func (s *Server) GetNotificationsByApp(ctx context.Context, in *npool.GetNotificationsByAppRequest) (*npool.GetNotificationsByAppResponse, error) {
	resp, err := crud.GetNotificationsByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get notifications by app user error: %v", err)
		return &npool.GetNotificationsByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetNotificationsByOtherApp(ctx context.Context, in *npool.GetNotificationsByOtherAppRequest) (*npool.GetNotificationsByOtherAppResponse, error) {
	resp, err := crud.GetNotificationsByApp(ctx, &npool.GetNotificationsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorw("get notifications by app user error: %v", err)
		return &npool.GetNotificationsByOtherAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return &npool.GetNotificationsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}
