package dto

import (
	"github.com/Thor-x86/nullable"
	"github.com/gin-gonic/gin"
	"github.com/nmfzone/privy-cake-store/domain"
	"github.com/nmfzone/privy-cake-store/internal/errors"
	"github.com/nmfzone/privy-cake-store/internal/form"
	"time"
)

type CreateCakeDto struct {
	Title       string `binding:"required"`
	Description nullable.String
	Rating      nullable.Int    `binding:"omitempty,numeric,min=0,max=100"`
	Image       nullable.String `binding:"omitempty,min=2,max=10"`
}

func (c *CreateCakeDto) Bind(ctx *gin.Context) errors.ValidationError {
	// use manual bind, since gin context bind not support `nullable`
	err := errors.ValidationError{}

	c.Title = form.GetTrimmedPostForm(ctx, "title")
	c.Description = form.GetNullableStringPostForm(ctx, "description", true)
	c.Rating, err = form.GetNullableIntPostForm(ctx, "rating")
	c.Image = form.GetNullableStringPostForm(ctx, "image", true)

	return err
}

type UpdateCakeDto struct {
	Title       string `binding:"required"`
	Description nullable.String
	Rating      nullable.Int    `binding:"omitempty,numeric,min=0,max=100"`
	Image       nullable.String `binding:"omitempty,min=0,max=1000"`
}

func (c *UpdateCakeDto) Bind(ctx *gin.Context) errors.ValidationError {
	err := errors.ValidationError{}

	c.Title = form.GetTrimmedPostForm(ctx, "title")
	c.Description = form.GetNullableStringPostForm(ctx, "description", true)
	c.Rating, err = form.GetNullableIntPostForm(ctx, "rating")
	c.Image = form.GetNullableStringPostForm(ctx, "image", true)

	return err
}

type CakeResponseDto struct {
	ID          nullable.Uint64 `json:"id"`
	Title       string          `json:"title"`
	Description nullable.String `json:"description"`
	Rating      nullable.Int    `json:"rating"`
	Image       nullable.String `json:"image"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}

func (c *CakeResponseDto) Make(cake domain.Cake) CakeResponseDto {
	data := CakeResponseDto{
		ID:          cake.ID,
		Title:       cake.Title,
		Description: cake.Description,
		Rating:      cake.Rating,
		Image:       cake.Image,
		CreatedAt:   cake.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   cake.UpdatedAt.Format(time.RFC3339),
	}

	return data
}

func (c *CakeResponseDto) Collection(cakes []domain.Cake) []CakeResponseDto {
	response := make([]CakeResponseDto, 0)

	for _, cake := range cakes {
		response = append(response, c.Make(cake))
	}

	return response
}
