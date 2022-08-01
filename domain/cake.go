package domain

import (
	"context"
	"github.com/Thor-x86/nullable"
	"time"
)

type Cake struct {
	ID          nullable.Uint64 `json:"id"`
	Title       string          `json:"title"`
	Description nullable.String `json:"description"`
	Rating      nullable.Int    `json:"rating"`
	Image       nullable.String `json:"image"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

func (c *Cake) New(n Cake) Cake {
	dm := Cake{
		ID:          n.ID,
		Title:       n.Title,
		Description: n.Description,
		Rating:      n.Rating,
		Image:       n.Image,
		CreatedAt:   n.CreatedAt,
		UpdatedAt:   n.UpdatedAt,
	}

	return dm
}

type CakeRepository interface {
	FindById(ctx context.Context, id uint64) (Cake, error)
	FindAll(ctx context.Context, cursor string, limit int) ([]Cake, string, error)
	FindByTitle(ctx context.Context, title string) (res Cake, err error)
	Save(ctx context.Context, cake *Cake) error
	Remove(ctx context.Context, cake *Cake) (err error)
}
