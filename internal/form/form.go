package form

import (
	"github.com/Thor-x86/nullable"
	"github.com/gin-gonic/gin"
	"github.com/nmfzone/privy-cake-store/internal/errors"
	"github.com/nmfzone/privy-cake-store/internal/validator"
	"strconv"
	"strings"
)

// @TODO: Add utils for another nullable type

func GetTrimmedPostForm(ctx *gin.Context, field string) string {
	val := ctx.PostForm(field)

	return strings.TrimRight(strings.TrimLeft(val, " "), " ")
}

func GetNullableStringPostForm(ctx *gin.Context, field string, trim bool) nullable.String {
	val, ok := ctx.GetPostForm(field)

	if trim {
		val = GetTrimmedPostForm(ctx, field)
	}

	res := &val

	if !ok {
		res = nil
	}

	return nullable.NewString(res)
}

func GetNullableInt64PostForm(ctx *gin.Context, field string) (nullable.Int64, errors.ValidationError) {
	val := GetNullableStringPostForm(ctx, field, true)
	var res *int64 = nil

	if val.Get() != nil {
		j, err := toIntOrError(field, *val.Get())
		if err.Errors != nil {
			return nullable.Int64{}, err
		}

		m := int64(j)
		res = &m
	}

	return nullable.NewInt64(res), errors.ValidationError{}
}

func GetNullableIntPostForm(ctx *gin.Context, field string) (nullable.Int, errors.ValidationError) {
	val := GetNullableStringPostForm(ctx, field, true)
	var res *int = nil

	if val.Get() != nil {
		// should be able to convert into integer, if it doesn't, return validation error
		j, err := toIntOrError(field, *val.Get())
		if err.Errors != nil {
			return nullable.Int{}, err
		}

		res = &j
	}

	return nullable.NewInt(res), errors.ValidationError{}
}

func toIntOrError(field string, val string) (int, errors.ValidationError) {
	res, err := strconv.Atoi(val)

	if err == nil {
		return res, errors.ValidationError{}
	}

	vErr := validator.ValidateSingleField(val, "numeric", field)

	return res, vErr
}
