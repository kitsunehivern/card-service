package repo

import (
	"card-service/internal/model"
	"context"

	"gorm.io/gorm"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, user *model.User) error
	CountUser(ctx context.Context, params model.UserParams) (int64, error)
	GetUserPasswordHash(ctx context.Context, params model.UserParams) (string, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) buildUserQuery(ctx context.Context, params model.UserParams) model.UserQuerySet {
	query := model.NewUserQuerySet(repo.db.WithContext(ctx))
	if params.ID > 0 {
		query = query.IDEq(params.ID)
	}

	if params.PhoneNumber != "" {
		query = query.PhoneNumberEq(params.PhoneNumber)
	}

	return query
}

func (repo *UserRepo) CreateUser(ctx context.Context, user *model.User) error {
	return repo.db.WithContext(ctx).Create(user).Error
}

func (repo *UserRepo) CountUser(ctx context.Context, params model.UserParams) (int64, error) {
	count, err := repo.buildUserQuery(ctx, params).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *UserRepo) GetUserPasswordHash(ctx context.Context, params model.UserParams) (string, error) {
	var user model.User
	if err := repo.buildUserQuery(ctx, params).Select(model.UserDBSchema.PasswordHash).One(&user); err != nil {
		return "", err
	}
	return user.PasswordHash, nil
}
