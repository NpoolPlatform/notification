package announcement

import (
	"context"
	"time"

	"github.com/NpoolPlatform/notification/message/npool"
	"github.com/NpoolPlatform/notification/pkg/db"
	"github.com/NpoolPlatform/notification/pkg/db/ent"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

const (
	dbTimeout = 5 * time.Second
)

func validateAnnouncement(info *npool.Announcement) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if info.GetTitle() == "" {
		return xerrors.Errorf("invalid title")
	}
	if info.GetContent() == "" {
		return xerrors.Errorf("invalid content")
	}
	return nil
}

func dbRowToAnnouncement(row *ent.Announcement) *npool.Announcement {
	return &npool.Announcement{
		ID:       row.ID.String(),
		AppID:    row.AppID.String(),
		Title:    row.Title,
		Content:  row.Content,
		CreateAt: row.CreateAt,
	}
}

func Create(ctx context.Context, in *npool.CreateAnnouncementRequest) (*npool.CreateAnnouncementResponse, error) {
	if err := validateAnnouncement(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		Announcement.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetTitle(in.GetInfo().GetTitle()).
		SetContent(in.GetInfo().GetContent()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create announcement: %v", err)
	}

	return &npool.CreateAnnouncementResponse{
		Info: dbRowToAnnouncement(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAnnouncementRequest) (*npool.UpdateAnnouncementResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	if err = validateAnnouncement(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invlaid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		Announcement.
		UpdateOneID(id).
		SetTitle(in.GetInfo().GetTitle()).
		SetContent(in.GetInfo().GetContent()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update announcement: %v", err)
	}

	return &npool.UpdateAnnouncementResponse{
		Info: dbRowToAnnouncement(info),
	}, nil
}

func GetAnnouncementsByApp(ctx context.Context, in *npool.GetAnnouncementsByAppRequest) (*npool.GetAnnouncementsByAppResponse, error) {
	return nil, nil
}
