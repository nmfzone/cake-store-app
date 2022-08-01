package http

import (
	"github.com/gin-gonic/gin"
	module "github.com/nmfzone/privy-cake-store/cake"
	"github.com/nmfzone/privy-cake-store/cake/dto"
	"github.com/nmfzone/privy-cake-store/internal/response"
	"github.com/nmfzone/privy-cake-store/internal/validator"
	"strconv"
)

type CakeHandler struct {
	usecase module.Usecase
}

func NewCakeHandler(e *gin.Engine, u module.Usecase) {
	handler := &CakeHandler{
		usecase: u,
	}

	router := e.Group("/public/api/")
	router.GET("/cakes", handler.Fetch)
	router.POST("/cakes", handler.Store)
	router.GET("/cakes/:id", handler.Show)
	router.PUT("/cakes/:id", handler.Update)
	router.DELETE("/cakes/:id", handler.Destroy)
}

func (h *CakeHandler) Fetch(ctx *gin.Context) {
	qLimit, ok := ctx.GetQuery("limit")
	if !ok {
		qLimit = "10"
	}

	limit, _ := strconv.Atoi(qLimit)
	cursor := ctx.Query("cursor")

	cakes, nextCursor, err := h.usecase.FetchCakes(ctx.Request.Context(), cursor, limit)
	if err != nil {
		response.NewResponseError(ctx, err)
		return
	}

	ctx.Header("X-Cursor", nextCursor)

	responseDto := dto.CakeResponseDto{}

	response.OkResponseData(ctx, responseDto.Collection(cakes))
}

func (h *CakeHandler) Store(ctx *gin.Context) {
	var createCakeDto dto.CreateCakeDto
	vErr := createCakeDto.Bind(ctx)
	if vErr.Errors != nil {
		response.ValidationErrorResponse(ctx, vErr)
		return
	}

	v := validator.NewRequestValidator()

	ok, vErr := v.Validate(createCakeDto)
	if !ok {
		response.ValidationErrorResponse(ctx, vErr)
		return
	}

	cake, err := h.usecase.StoreCake(ctx.Request.Context(), createCakeDto)

	if err != nil {
		response.NewResponseError(ctx, err)
		return
	}

	responseDto := dto.CakeResponseDto{}

	response.CreatedResponse(
		ctx,
		"Cake created successfully.",
		responseDto.Make(cake),
	)
}

func (h *CakeHandler) Show(ctx *gin.Context) {
	pId, _ := strconv.Atoi(ctx.Param("id"))

	id := uint64(pId)
	cake, err := h.usecase.ShowCake(ctx.Request.Context(), id)

	if err != nil {
		response.NewResponseError(ctx, err)
		return
	}

	responseDto := dto.CakeResponseDto{}

	response.OkResponseData(ctx, responseDto.Make(cake))
}

func (h *CakeHandler) Update(ctx *gin.Context) {
	pId, _ := strconv.Atoi(ctx.Param("id"))

	var updateCakeDto dto.UpdateCakeDto
	vErr := updateCakeDto.Bind(ctx)
	if vErr.Errors != nil {
		response.ValidationErrorResponse(ctx, vErr)
		return
	}

	v := validator.NewRequestValidator()

	ok, vErr := v.Validate(updateCakeDto)
	if !ok {
		response.ValidationErrorResponse(ctx, vErr)
		return
	}

	id := uint64(pId)
	cake, err := h.usecase.UpdateCake(ctx.Request.Context(), id, updateCakeDto)

	if err != nil {
		response.NewResponseError(ctx, err)
		return
	}

	responseDto := dto.CakeResponseDto{}

	response.OkResponseWith(
		ctx,
		"Cake updated successfully.",
		responseDto.Make(cake),
	)
}

func (h *CakeHandler) Destroy(ctx *gin.Context) {
	pId, _ := strconv.Atoi(ctx.Param("id"))

	id := uint64(pId)
	err := h.usecase.DestroyCake(ctx.Request.Context(), id)

	if err != nil {
		response.NewResponseError(ctx, err)
		return
	}

	response.OkResponse(
		ctx,
		"Cake deleted successfully.",
	)
}
