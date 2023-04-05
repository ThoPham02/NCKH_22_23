package logic

import (
	"database/sql"
	"github/ThoPham02/research_management/api/types"
	"time"
)

func (l *Logic) GetUserInfo(userID int64) (*types.UserInfoResponse, error) {
	l.logHelper.Infof("Start processing get user info, input: %d", userID)

	userInfo, err := l.svcCtx.Store.GetUserInfo(l.ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.UserInfoResponse{}, nil
		}
		l.logHelper.Errorf("Failed to get user info: %s", err.Error())
		return nil, err
	}

	var (
		description *string
		avataUrl    *string
		birthday    *time.Time
		bankAccount *string
		phone       *string
		sex         *int64
	)
	if userInfo.Description.Valid {
		description = &userInfo.Description.String
	}
	if userInfo.Birthday.Valid {
		birthday = &userInfo.Birthday.Time
	}
	if userInfo.AvataUrl.Valid {
		avataUrl = &userInfo.AvataUrl.String
	}
	if userInfo.BankAccount.Valid {
		bankAccount = &userInfo.BankAccount.String
	}
	if userInfo.Phone.Valid {
		phone = &userInfo.Phone.String
	}
	if userInfo.Sex.Valid {
		sex = &userInfo.Sex.Int64
	}

	return &types.UserInfoResponse{
		Name:        userInfo.Name,
		Description: description,
		AvatarUrl:   avataUrl,
		Birthday:    birthday,
		FaculityID:  userInfo.FaculityID,
		YearStart:   userInfo.YearStart,
		BankAccount: bankAccount,
		Phone:       phone,
		Sex:         sex,
	}, nil
}
