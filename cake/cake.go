package cake

import (
	"context"
	"github.com/nmfzone/privy-cake-store/cake/dto"
	"github.com/nmfzone/privy-cake-store/domain"
)

type Usecase interface {
	FetchCakes(ctx context.Context, cursor string, limit int) (res []domain.Cake, nextCursor string, err error)
	StoreCake(ctx context.Context, payload dto.CreateCakeDto) (domain.Cake, error)
	ShowCake(ctx context.Context, cakeId uint64) (domain.Cake, error)
	UpdateCake(ctx context.Context, cakeId uint64, payload dto.UpdateCakeDto) (domain.Cake, error)
	DestroyCake(ctx context.Context, id uint64) error
}
