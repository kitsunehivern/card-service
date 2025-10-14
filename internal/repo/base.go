package repo

import "gorm.io/gorm"

type Repo struct {
	DB       *gorm.DB
	UserRepo IUserRepo
	CardRepo ICardRepo
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		DB:       db,
		UserRepo: NewUserRepo(db),
		CardRepo: NewCardRepo(db),
	}
}
