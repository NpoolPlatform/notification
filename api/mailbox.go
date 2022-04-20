// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notification"
	crud "github.com/NpoolPlatform/notification/pkg/crud/mailbox"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateMail(ctx context.Context, in *npool.CreateMailRequest) (*npool.CreateMailResponse, error) {
	resp, err := crud.CreateMail(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create mailbox error: %v", err)
		return &npool.CreateMailResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateMailForOtherApp(ctx context.Context, in *npool.CreateMailForOtherAppRequest) (*npool.CreateMailForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := crud.CreateMail(ctx, &npool.CreateMailRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("create mailbox error: %v", err)
		return &npool.CreateMailForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateMailForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) UpdateMail(ctx context.Context, in *npool.UpdateMailRequest) (*npool.UpdateMailResponse, error) {
	resp, err := crud.UpdateMail(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("check mailbox error: %v", err)
		return &npool.UpdateMailResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetMails(ctx context.Context, in *npool.GetMailsRequest) (*npool.GetMailsResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get mails error %v", err)
		return &npool.GetMailsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetMailsByApp(ctx context.Context, in *npool.GetMailsByAppRequest) (*npool.GetMailsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get mails error %v", err)
		return &npool.GetMailsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetMailsByOtherApp(ctx context.Context, in *npool.GetMailsByOtherAppRequest) (*npool.GetMailsByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetMailsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get mails error %v", err)
		return &npool.GetMailsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetMailsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}
