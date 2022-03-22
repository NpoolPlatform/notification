// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notification"
	crud "github.com/NpoolPlatform/notification/pkg/crud/announcement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAnnouncement(ctx context.Context, in *npool.CreateAnnouncementRequest) (*npool.CreateAnnouncementResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create announcement error: %v", err)
		return &npool.CreateAnnouncementResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateAnnouncement(ctx context.Context, in *npool.UpdateAnnouncementRequest) (*npool.UpdateAnnouncementResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update announcement error: %v", err)
		return &npool.UpdateAnnouncementResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetAnnouncementsByApp(ctx context.Context, in *npool.GetAnnouncementsByAppRequest) (*npool.GetAnnouncementsByAppResponse, error) {
	resp, err := crud.GetAnnouncementsByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get announcements by app error: %v", err)
		return &npool.GetAnnouncementsByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
