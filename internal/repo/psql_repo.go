package repo

import (
	"card-service/internal/errmsg"
	"card-service/internal/model"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type psqlRepo struct {
	pool *pgxpool.Pool
}

func NewPsqlRepo(ctx context.Context, dsn string) (IRepository, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return &psqlRepo{pool: pool}, nil
}

func (repo *psqlRepo) Close() error {
	repo.pool.Close()
	return nil
}

func (repo *psqlRepo) CreateCard(ctx context.Context, card *model.Card) error {
	sql := `
		insert into cards
		values ($1, $2, $3, $4, $5, NOW())
	`
	_, err := repo.pool.Exec(ctx, sql, card.ID, card.UserID, card.Debit, card.Credit, card.Status)

	return err
}

func (repo *psqlRepo) CountCardByUserID(ctx context.Context, userID string) (int32, error) {
	sql := `
		select count(*)
		from cards
		where user_id = $1
	`

	var count int32
	if err := repo.pool.QueryRow(ctx, sql, userID).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (repo *psqlRepo) GetCardByID(ctx context.Context, id string) (*model.Card, error) {
	sql := `
		select *
		from cards
		where id = $1
	`

	var card model.Card
	if err := repo.pool.QueryRow(ctx, sql, id).Scan(
		&card.ID,
		&card.UserID,
		&card.Debit,
		&card.Credit,
		&card.Status,
		&card.UpdatedAt,
	); err != nil {
		return nil, errmsg.CardNotFound
	}

	return &card, nil
}

func (repo *psqlRepo) GetCardByUserID(ctx context.Context, userID string) (*model.Card, error) {
	sql := `
		select *
		from cards
		where user_id = $1
	`

	var card model.Card
	if err := repo.pool.QueryRow(ctx, sql, userID).Scan(
		&card.ID,
		&card.UserID,
		&card.Debit,
		&card.Credit,
		&card.Status,
		&card.UpdatedAt,
	); err != nil {
		return nil, errmsg.CardNotFound
	}

	return &card, nil
}

func (repo *psqlRepo) UpdateCardStatus(ctx context.Context, id string, status model.Status) error {
	sql := `
		update cards
		set status = $2, updated_at = NOW()
		where id = $1
	`

	_, err := repo.pool.Exec(ctx, sql, id, status)

	return err
}
