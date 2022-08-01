package usecase

import (
	"context"
	module "github.com/nmfzone/privy-cake-store/cake"
	"github.com/nmfzone/privy-cake-store/cake/dto"
	"github.com/nmfzone/privy-cake-store/domain"
	"time"
)

type cakeUsecase struct {
	cakeRepo       domain.CakeRepository
	contextTimeout time.Duration
}

func NewCakeUsecase(r domain.CakeRepository, timeout time.Duration) module.Usecase {
	return &cakeUsecase{
		cakeRepo:       r,
		contextTimeout: timeout,
	}
}

func (usecase *cakeUsecase) FetchCakes(ctx context.Context, cursor string, limit int) (res []domain.Cake, nextCursor string, err error) {
	if limit == 0 {
		limit = 10
	}

	//ctx, cancel := context.WithTimeout(ctx, usecase.contextTimeout)
	//defer cancel()

	res, nextCursor, err = usecase.cakeRepo.FindAll(ctx, cursor, limit)

	if err != nil {
		return nil, "", err
	}

	return
}

func (usecase *cakeUsecase) StoreCake(ctx context.Context, payload dto.CreateCakeDto) (domain.Cake, error) {
	cake := domain.Cake{
		Title:       payload.Title,
		Description: payload.Description,
		Rating:      payload.Rating,
		Image:       payload.Image,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := usecase.cakeRepo.Save(ctx, &cake)

	return cake, err
}

func (usecase *cakeUsecase) ShowCake(ctx context.Context, cakeId uint64) (domain.Cake, error) {
	cake, err := usecase.cakeRepo.FindById(ctx, cakeId)

	if err != nil {
		return domain.Cake{}, err
	}

	return cake, err
}

func (usecase *cakeUsecase) UpdateCake(ctx context.Context, cakeId uint64, payload dto.UpdateCakeDto) (domain.Cake, error) {
	cake, err := usecase.cakeRepo.FindById(ctx, cakeId)

	if err != nil {
		return domain.Cake{}, err
	}

	cake = domain.Cake{
		ID:          cake.ID,
		Title:       payload.Title,
		Description: payload.Description,
		Rating:      payload.Rating,
		Image:       payload.Image,
		CreatedAt:   cake.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	err = usecase.cakeRepo.Save(ctx, &cake)

	if err != nil {
		return domain.Cake{}, err
	}

	return cake, err
}

func (usecase *cakeUsecase) DestroyCake(ctx context.Context, cakeId uint64) error {
	cake, err := usecase.cakeRepo.FindById(ctx, cakeId)

	if err != nil {
		return err
	}

	err = usecase.cakeRepo.Remove(ctx, &cake)

	if err != nil {
		return err
	}

	return nil
}
