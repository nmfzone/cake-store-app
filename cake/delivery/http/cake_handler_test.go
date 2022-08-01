package http

import (
	"encoding/json"
	"github.com/bxcodec/faker"
	"github.com/gin-gonic/gin"
	"github.com/nmfzone/privy-cake-store/cake/dto"
	"github.com/nmfzone/privy-cake-store/domain"
	"github.com/nmfzone/privy-cake-store/internal/errors"
	cakeMocks "github.com/nmfzone/privy-cake-store/mocks/cake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestFetchHandlerIsSuccessful(t *testing.T) {
	var cakeMock domain.Cake
	// @TODO: We need to add provider for the nullable type
	// See: https://github.com/bxcodec/faker/blob/master/example_custom_faker_test.go
	err := faker.FakeData(&cakeMock)
	assert.NoError(t, err)

	usecaseMocks := new(cakeMocks.Usecase)
	cakes := make([]domain.Cake, 0)
	cakes = append(cakes, cakeMock)

	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)

	limit := 1
	cursor := ""

	cursorExpect := "10"

	usecaseMocks.
		On("FetchCakes", mock.Anything, cursor, limit).
		Return(cakes, cursorExpect, nil)

	req, err := http.NewRequest(http.MethodGet, "/public/api/cakes", strings.NewReader(""))
	assert.NoError(t, err)

	q := req.URL.Query()
	q.Add("limit", strconv.Itoa(limit))

	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	handler := &CakeHandler{
		usecase: usecaseMocks,
	}
	handler.Fetch(ctx)

	assert.Equal(t, cursorExpect, rec.Header().Get("X-Cursor"))
	assert.Equal(t, http.StatusOK, rec.Code)

	usecaseMocks.AssertExpectations(t)
}

func TestStoreHandlerIsSuccessful(t *testing.T) {
	usecaseMocks := new(cakeMocks.Usecase)

	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)

	form := url.Values{
		"title": {"Test"},
	}

	req, err := http.NewRequest(http.MethodPost, "/public/api/cakes", strings.NewReader(form.Encode()))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	ctx.Request = req

	cakeMock := domain.Cake{
		Title: "Test",
	}

	usecaseMocks.
		On("StoreCake", mock.Anything, dto.CreateCakeDto{
			Title: "Test",
		}).
		Return(cakeMock, nil)

	handler := &CakeHandler{
		usecase: usecaseMocks,
	}
	handler.Store(ctx)

	var response struct {
		Message string              `json:"message"`
		Data    dto.CakeResponseDto `json:"data"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "Cake created successfully.", response.Message)
	assert.Equal(t, cakeMock.Title, response.Data.Title)

	usecaseMocks.AssertExpectations(t)
}

func TestStoreHandlerNotPassValidationWhenDataIsInvalid(t *testing.T) {
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)

	form := url.Values{
		"title":  {""},
		"rating": {"-1"},
	}

	req, err := http.NewRequest(http.MethodPost, "/public/api/cakes", strings.NewReader(form.Encode()))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	ctx.Request = req

	handler := &CakeHandler{}
	handler.Store(ctx)

	var response struct {
		errors.ValidationError
		Message string `json:"message"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	_, ok := response.Errors["title"]
	assert.True(t, ok)
	_, ok = response.Errors["rating"]
	assert.True(t, ok)
}

func TestShowHandlerIsSuccessful(t *testing.T) {
	usecaseMocks := new(cakeMocks.Usecase)

	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)

	req, err := http.NewRequest(http.MethodGet, "/public/api/cakes", strings.NewReader(""))
	assert.NoError(t, err)

	pId := "1"

	ctx.Request = req
	// We set param here, not in URL
	ctx.Params = []gin.Param{{Key: "id", Value: pId}}

	cakeMock := domain.Cake{
		Title: "Test",
	}

	id, _ := strconv.Atoi(pId)
	usecaseMocks.
		On("ShowCake", mock.Anything, uint64(id)).
		Return(cakeMock, nil)

	handler := &CakeHandler{
		usecase: usecaseMocks,
	}
	handler.Show(ctx)

	var response struct {
		Data dto.CakeResponseDto `json:"data"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, cakeMock.Title, response.Data.Title)

	usecaseMocks.AssertExpectations(t)
}

func TestUpdateHandlerIsSuccessful(t *testing.T) {
	usecaseMocks := new(cakeMocks.Usecase)

	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)

	form := url.Values{
		"title": {"Test"},
	}

	req, err := http.NewRequest(http.MethodPut, "/public/api/cakes", strings.NewReader(form.Encode()))
	assert.NoError(t, err)

	pId := "1"

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	ctx.Request = req
	// We set param here, not in URL
	ctx.Params = []gin.Param{{Key: "id", Value: pId}}

	cakeMock := domain.Cake{
		Title: "Test",
	}

	id, _ := strconv.Atoi(pId)
	usecaseMocks.
		On("UpdateCake", mock.Anything, uint64(id), dto.UpdateCakeDto{
			Title: "Test",
		}).
		Return(cakeMock, nil)

	handler := &CakeHandler{
		usecase: usecaseMocks,
	}
	handler.Update(ctx)

	var response struct {
		Message string              `json:"message"`
		Data    dto.CakeResponseDto `json:"data"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Cake updated successfully.", response.Message)
	assert.Equal(t, cakeMock.Title, response.Data.Title)

	usecaseMocks.AssertExpectations(t)
}

func TestDestroyHandlerIsSuccessful(t *testing.T) {
	usecaseMocks := new(cakeMocks.Usecase)

	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)

	req, err := http.NewRequest(http.MethodDelete, "/public/api/cakes", strings.NewReader(""))
	assert.NoError(t, err)

	pId := "1"

	ctx.Request = req
	// We set param here, not in URL
	ctx.Params = []gin.Param{{Key: "id", Value: pId}}

	id, _ := strconv.Atoi(pId)
	usecaseMocks.
		On("DestroyCake", mock.Anything, uint64(id)).
		Return(nil)

	handler := &CakeHandler{
		usecase: usecaseMocks,
	}
	handler.Destroy(ctx)

	var response struct {
		Message string `json:"message"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Cake deleted successfully.", response.Message)

	usecaseMocks.AssertExpectations(t)
}
