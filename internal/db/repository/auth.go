package repository

import (
	"context"
	"gorm.io/gorm"
	"ticket-booking-app/internal/db/models"
)

var _ AuthRepository = (*Auth)(nil)

func NewAuthRepository(db *gorm.DB) *Auth {
	return &Auth{
		dbClient: db,
	}
}

type Auth struct {
	dbClient *gorm.DB
}

func (a Auth) RegisterUser(ctx context.Context, registerData *models.AuthCredentials) (*models.User, error) {
	user := &models.User{
		Email:    registerData.Email,
		Password: registerData.Password,
	}

	res := a.dbClient.Model(&models.User{}).WithContext(ctx).Create(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (a Auth) GetUser(ctx context.Context, query interface{}, args ...interface{}) (*models.User, error) {
	user := &models.User{}

	if res := a.dbClient.Model(user).Where(query, args...).WithContext(ctx).First(user); res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
