package errors

import (
	"net/http"

	"github.com/joomcode/errorx"
)

var (
	invalidInput = errorx.NewNamespace("validation error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	dbError      = errorx.NewNamespace("db error")
	duplicate    = errorx.NewNamespace("duplicate").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	dataNotFound = errorx.NewNamespace("data not found").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
)

var (
	ErrInvalidUserInput = errorx.NewType(invalidInput, "invalid user input")
	ErrWriteError       = errorx.NewType(dbError, "could not write to db")
	ErrReadError        = errorx.NewType(dbError, "could not read data from db")
	ErrDataExists       = errorx.NewType(duplicate, "data already exists")
	ErrNoRecordFound    = errorx.NewType(dataNotFound, "no record found")
)

var ErrorMap = map[*errorx.Type]int{
	ErrInvalidUserInput: http.StatusBadRequest,
	ErrWriteError:       http.StatusInternalServerError,
	ErrReadError:        http.StatusInternalServerError,
	ErrDataExists:       http.StatusBadRequest,
	ErrNoRecordFound:    http.StatusNotFound,
}
