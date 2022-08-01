package utils

import (
	"encoding/base64"
	"reflect"
	"time"
)

var timeFormat = "2006-01-02T15:04:05.999Z07:00"

func DecodeStringToTime(encodedTime string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

func EncodeTimeToString(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}

func GetStructFieldMetadata(s interface{}, field string) *reflect.StructTag {
	m := reflect.ValueOf(s)

	if m.Kind() != reflect.Struct {
		panic("Not a struct")
	}

	for i := 0; i < m.NumField(); i++ {
		k := m.Type().Field(i)

		if k.Name == field {
			return &k.Tag
		}
	}

	return nil
}
