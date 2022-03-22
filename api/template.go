package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notification"
	templatecrud "github.com/NpoolPlatform/notification/pkg/crud/template"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateTemplate(ctx context.Context, in *npool.CreateTemplateRequest) (*npool.CreateTemplateResponse, error) {
	resp, err := templatecrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create template: %v", err)
		return &npool.CreateTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateTemplateForOtherApp(ctx context.Context, in *npool.CreateTemplateForOtherAppRequest) (*npool.CreateTemplateForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := templatecrud.Create(ctx, &npool.CreateTemplateRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("fail create template: %v", err)
		return &npool.CreateTemplateForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateTemplateForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) GetTemplate(ctx context.Context, in *npool.GetTemplateRequest) (*npool.GetTemplateResponse, error) {
	resp, err := templatecrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get template: %v", err)
		return &npool.GetTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateTemplate(ctx context.Context, in *npool.UpdateTemplateRequest) (*npool.UpdateTemplateResponse, error) {
	resp, err := templatecrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update template: %v", err)
		return &npool.UpdateTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetTemplatesByApp(ctx context.Context, in *npool.GetTemplatesByAppRequest) (*npool.GetTemplatesByAppResponse, error) {
	resp, err := templatecrud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get  templates by app: %v", err)
		return &npool.GetTemplatesByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetTemplatesByOtherApp(ctx context.Context, in *npool.GetTemplatesByOtherAppRequest) (*npool.GetTemplatesByOtherAppResponse, error) {
	resp, err := templatecrud.GetByApp(ctx, &npool.GetTemplatesByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail get templates by other app: %v", err)
		return &npool.GetTemplatesByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetTemplatesByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetTemplateByAppLangUsedFor(ctx context.Context, in *npool.GetTemplateByAppLangUsedForRequest) (*npool.GetTemplateByAppLangUsedForResponse, error) {
	resp, err := templatecrud.GetByAppLangUsedFor(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get template by app lang used for: %v", err)
		return &npool.GetTemplateByAppLangUsedForResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
