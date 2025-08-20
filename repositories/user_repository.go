package repositories

import (
	"context"
	"learn_golang/database/myquery"
	"learn_golang/models"
)

func GetAllUsers(ctx context.Context) ([]*models.User, error) {
	userQuery := myquery.User
	users, err := userQuery.WithContext(ctx).Find()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(ctx context.Context, id uint64) (*models.User, error) {
	userQuery := myquery.User
	user, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func SaveUser(ctx context.Context, user *models.User) (*models.User, error) {
	userQuery := myquery.User
	err := userQuery.WithContext(ctx).Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil

}
