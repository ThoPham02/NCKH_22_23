package logic

import (
	"database/sql"
	db "github/ThoPham02/research_management/api/db/sqlc"
	"github/ThoPham02/research_management/api/types"
	"time"
)

func (l *Logic) UpdateUserInfo(userID int64, req *types.UserInfoResponse) error {
	l.logHelper.Infof("Start processing update user info, input: %d, %v", userID, req)

	description := sql.NullString{
		Valid:  req.Description != nil,
		String: getString(req.Description),
	}
	avataUrl := sql.NullString{
		Valid:  req.AvatarUrl != nil,
		String: getString(req.AvatarUrl),
	}
	birthday := sql.NullTime{
		Valid: req.Birthday != nil,
		Time:  getTime(req.Birthday),
	}
	bankAccount := sql.NullString{
		Valid:  req.BankAccount != nil,
		String: getString(req.BankAccount),
	}
	phone := sql.NullString{
		Valid:  req.Phone != nil,
		String: getString(req.Phone),
	}
	sex := sql.NullInt64{
		Valid: req.Sex != nil,
		Int64: getInt64(req.Sex),
	}

	_, err := l.svcCtx.Store.GetUserInfo(l.ctx, userID)
	if err == sql.ErrNoRows {
		err = l.svcCtx.Store.CreateUserInfo(l.ctx, db.CreateUserInfoParams{
			UserID:      userID,
			Name:        req.Name,
			Description: description,
			AvataUrl:    avataUrl,
			FaculityID:  req.FaculityID,
			YearStart:   req.YearStart,
			Birthday:    birthday,
			BankAccount: bankAccount,
			Phone:       phone,
			Sex:         sex,
		})
	}
	if err != nil {
		l.logHelper.Errorf("Failed while creating user info, error: %v", err)
		return err
	}

	if err = l.svcCtx.Store.UpdateUserInfo(l.ctx, db.UpdateUserInfoParams{
		UserID:      userID,
		Name:        req.Name,
		Description: description,
		AvataUrl:    avataUrl,
		FaculityID:  req.FaculityID,
		YearStart:   req.YearStart,
		Birthday:    birthday,
		BankAccount: bankAccount,
		Phone:       phone,
		Sex:         sex,
	}); err != nil {
		l.logHelper.Errorf("Failed while update user info: %v", err)
		return err
	}

	return nil
}

func getString(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}

func getInt64(input *int64) int64 {
	if input == nil {
		return 0
	}
	return *input
}

func getTime(input *time.Time) time.Time {
	if input == nil {
		return time.Now()
	}
	return *input
}
