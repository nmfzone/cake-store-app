package mysql

import (
	"context"
	"database/sql"
	"github.com/Thor-x86/nullable"
	"github.com/godruoyi/go-snowflake"
	"github.com/nmfzone/privy-cake-store/domain"
	"github.com/nmfzone/privy-cake-store/internal/errors"
	"github.com/nmfzone/privy-cake-store/internal/utils"
	"github.com/rs/zerolog/log"
	"time"
)

type mysqlCakeRepository struct {
	Conn *sql.DB
}

func NewMysqlCakeRepository(Conn *sql.DB) domain.CakeRepository {
	return &mysqlCakeRepository{Conn}
}

func (m *mysqlCakeRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Cake, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Err(errRow).Send()
		}
	}()

	result = make([]domain.Cake, 0)
	for rows.Next() {
		t := domain.Cake{}

		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Rating,
			&t.Image,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		t = t.New(t)

		if err != nil {
			log.Err(err).Send()
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlCakeRepository) FindAll(ctx context.Context, cursor string, limit int) (res []domain.Cake, nextCursor string, err error) {
	query := `SELECT * FROM cakes WHERE created_at > ? ORDER BY created_at LIMIT ?`

	decodedCursor, err := utils.DecodeStringToTime(cursor)
	if err != nil && cursor != "" {
		return nil, "", errors.ErrBadParamInput
	}

	if cursor == "" {
		decodedCursor = time.Time{}.AddDate(1000, 1, 1)
	}

	res, err = m.fetch(ctx, query, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	count := len(res)
	if count == limit {
		nextCursor = utils.EncodeTimeToString(res[count-1].CreatedAt)
	}

	return
}

func (m *mysqlCakeRepository) FindById(ctx context.Context, id uint64) (res domain.Cake, err error) {
	query := `SELECT * FROM cakes WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Cake{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, errors.ErrNotFound
	}

	return
}

func (m *mysqlCakeRepository) FindByTitle(ctx context.Context, title string) (res domain.Cake, err error) {
	query := `SELECT * FROM cakes WHERE title = ?`

	list, err := m.fetch(ctx, query, title)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, errors.ErrNotFound
	}

	return
}

func (m *mysqlCakeRepository) Save(ctx context.Context, cake *domain.Cake) (err error) {
	if cake.ID.Get() == nil {
		id := snowflake.ID()
		cake.ID = nullable.NewUint64(&id)
		err = m.insert(ctx, cake)
	} else {
		err = m.update(ctx, cake)
	}

	if err != nil {
		log.Err(err).Send()
		return err
	}

	return
}

func (m *mysqlCakeRepository) insert(ctx context.Context, cake *domain.Cake) (err error) {
	query := `INSERT cakes SET id=?, title=?, description=?, rating=?, image=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(
		ctx,
		cake.ID,
		cake.Title,
		cake.Description,
		cake.Rating,
		cake.Image,
		cake.CreatedAt,
		cake.UpdatedAt)

	if err != nil {
		return err
	}

	return
}

func (m *mysqlCakeRepository) update(ctx context.Context, cake *domain.Cake) (err error) {
	query := `UPDATE cakes set title=?`

	args := make([]any, 0)

	args = append(args, cake.Title)

	if cake.Description.Get() != nil {
		query += `, description=?`
		args = append(args, cake.Description)
	}
	if cake.Rating.Get() != nil {
		query += `, rating=?`
		args = append(args, cake.Rating)
	}
	if cake.Image.Get() != nil {
		query += `, image=?`
		args = append(args, cake.Image)
	}

	query += `, updated_at=? WHERE ID=?`
	args = append(args, cake.UpdatedAt, cake.ID)

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(
		ctx,
		args...)

	if err != nil {
		return
	}

	return
}

func (m *mysqlCakeRepository) Remove(ctx context.Context, cake *domain.Cake) (err error) {
	query := "DELETE FROM cakes WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Send()
		return
	}

	_, err = stmt.ExecContext(ctx, cake.ID)
	if err != nil {
		log.Err(err).Send()
		return
	}

	return
}
